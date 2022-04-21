package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"maoim/api/comet"
	"maoim/internal/pkg/encrypt"
	"maoim/internal/pkg/utils"
	"maoim/pkg/merror"
)

type UserVo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

var _ Service = &service{}

type Service interface {
	Register(username, password string) (*User, error)
	Login(username, password string) (string, error)
	Logout(userId string) error
	GetUser(userId string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	Auth(token string) (*User, error)

	GetFriends(userId string) ([]*UserVo, error)
	ApplyFriend(userId, otherUsername, remark string) error
	AddFriend(userId, otherUserId string) error
	RemoveFriend(userId, otherUserId string) error
	IsFriend(userId, friendId string) (bool, error)

	ListApplyRecord(userId string, applying bool) ([]*FriendShipApply, error)
	ListOffsetApplyRecord(userId, recordId string, applying bool) ([]*FriendShipApply, error)
	AgreeFriendShipApply(userId, recordId string) error

	Connect(userId string) error
	Disconnect(userId string) error
	IsOnline(userId string) (bool, error)
}

type service struct {
	dao         Dao
	cometClient comet.CometClient
}

func (s *service) GetUserByUsername(username string) (*User, error) {
	return s.dao.GetUserByUsername(username)
}

func NewService(d Dao, cometClient comet.CometClient) Service {
	return &service{
		dao:         d,
		cometClient: cometClient,
	}
}

func (s *service) Register(username, password string) (u *User, err error) {
	_, err = s.dao.GetUserByUsername(username)
	// err = nil 时，说明没有报找不到err或其它err，说明存在用户
	if err == nil {
		err = merror.WithMessage(err, HasRegisterErr.Error())
		return
	}
	// 查询 db 报其它err
	if err != gorm.ErrRecordNotFound {
		err = merror.WithMessage(err, RegisterErr.Error())
		return
	}

	u = &User{
		Username: username,
		Password: encrypt.Encrypt([]byte(password)),
	}
	err = s.dao.SaveUser(u)
	if err != nil {
		return
	}
	return
}

func (s *service) Login(username, password string) (str string, err error) {
	u, err := s.dao.GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = UserNotFound
		} else {
			err = LoginFailErr
		}
		return
	}
	if encrypt.Encrypt([]byte(password)) != u.Password {
		err = PasswordErrorErr
		return
	}
	return utils.GenToken(u.ID, u.Username)
}

func (s *service) Logout(userId string) error {
	return s.dao.DeleteUser(userId)
}

func (s *service) GetUser(userId string) (*User, error) {
	return s.dao.GetUser(userId)
}

func (s *service) Auth(token string) (u *User, err error) {
	if token == "" {
		err = TokenEmptyErr
		return
	}
	userId, _, err := utils.ValidToken(token)
	if err != nil {
		err = TokenCheckFailErr
		return
	}
	return s.GetUser(userId)
}

func (s *service) ApplyFriend(userId, otherUsername, remark string) (err error) {
	other, err := s.dao.GetUserByUsername(otherUsername)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = UserNotFound
		}
		return
	}
	isFriend, err := s.IsFriend(userId, other.ID)
	if err != nil {
		return err
	}
	if isFriend {
		return AlreadyFriendErr
	}

	record, err := s.dao.GetApplyRecordByUserId(userId, other.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if record != nil && record.ID != "" {
		return errors.New("重复申请")
	}

	err = s.dao.SaveApplyRecord(userId, other.ID, remark)
	if err != nil {
		return err
	}

	_, err = s.cometClient.NewFriendShipApplyNotice(context.Background(), &comet.NewFriendShipApplyNoticeReq{
		UserId: other.ID,
	})
	return err
}

func (s *service) AddFriend(userId, otherUserId string) error {
	user, err := s.GetUser(otherUserId)
	if err != nil {
		return err
	}
	if user.ID == "" {
		return fmt.Errorf("不存在该用户")
	}

	ok, err := s.IsFriend(userId, otherUserId)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("已添加该好友")
	}

	err1 := s.dao.AddFriend(userId, otherUserId)
	err2 := s.dao.AddFriend(otherUserId, userId)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *service) RemoveFriend(userId, otherUserId string) error {
	user, err := s.GetUser(otherUserId)
	if err != nil {
		return err
	}
	if user.ID == "" {
		return fmt.Errorf("不存在该用户")
	}

	ok, err := s.IsFriend(userId, otherUserId)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("无此好友")
	}
	return s.dao.RemoveFriend(userId, otherUserId)
}

func (s *service) GetFriends(userId string) (vos []*UserVo, err error) {
	friends, err := s.dao.GetFriendList(userId)
	if err != nil {
		return
	}

	fuserIds := make([]string, 0)
	for _, f := range friends {
		fuserIds = append(fuserIds, f.FUserId)
	}
	users, err := s.dao.BatchGetUser(fuserIds)
	if err != nil {
		return
	}
	for _, u := range users {
		vos = append(vos, &UserVo{
			ID:       u.ID,
			Username: u.Username,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
		})
	}
	return
}

func (s *service) IsFriend(userId, friendId string) (bool, error) {
	return s.dao.IsFriend(userId, friendId)
}

func (s *service) ListApplyRecord(userId string, applying bool) ([]*FriendShipApply, error) {
	return s.dao.ListApplyRecord(userId, applying)
}

func (s *service) ListOffsetApplyRecord(userId, recordId string, applying bool) ([]*FriendShipApply, error) {
	return s.dao.ListOffsetApplyRecord(userId, recordId, applying)
}

func (s service) AgreeFriendShipApply(userId, recordId string) error {
	record, err := s.dao.GetApplyRecord(recordId)
	if err != nil {
		return err
	}
	if record.OtherUserId != userId {
		return errors.New("操作异常")
	}
	record.Agree = true
	err = s.dao.UpdateApplyRecord(record)
	if err != nil {
		return err
	}

	err = s.AddFriend(record.UserId, record.OtherUserId)
	if err != nil {
		return err
	}

	_, err = s.cometClient.FriendShipApplyPassNotice(context.Background(), &comet.FriendShipApplyPassReq{
		UserId: record.UserId,
	})
	return err
}

func (s *service) Disconnect(userId string) error {
	return s.dao.SetOffline(userId)
}

func (s *service) Connect(userId string) error {
	return s.dao.SetOnline(userId)
}

func (s *service) IsOnline(userId string) (bool, error) {
	return s.dao.IsOnline(userId)
}
