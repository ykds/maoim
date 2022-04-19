package user

import (
	"fmt"
	"maoim/pkg/mysql"
	"maoim/pkg/redis"
	"time"
)

var _ Dao = (*dao)(nil)

type Dao interface {
	// User
	SaveUser(u *User) error
	GetUser(userId string) (*User, error)
	BatchGetUser(userIds []string) ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	DeleteUser(userId string) error

	// Friend
	AddFriend(userId, otherUserId string) error
	RemoveFriend(userId, otherUserId string) error
	GetFriendList(userId string) ([]*FriendShip, error)
	IsFriend(userId, otherUserId string) (bool, error)

	// Apply Record
	GetApplyRecord(recordId string) (*FriendShipApply, error)
	GetApplyRecordByUserId(userId, otherUserId string) (*FriendShipApply, error)
	SaveApplyRecord(userId, otherUserId, remark string) error
	ListApplyRecord(userId string, applying bool) ([]*FriendShipApply, error)
	ListOffsetApplyRecord(userId, recordId string, applying bool) ([]*FriendShipApply, error)
	UpdateApplyRecord(record *FriendShipApply) error

	// Online status
	SetOnline(userId string) error
	SetOffline(userId string) error
	IsOnline(userId string) (bool, error)
}

type dao struct {
	rdb *redis.Redis
	db  *mysql.Mysql
}

func NewDao(rdb *redis.Redis, db *mysql.Mysql) Dao {
	_ = db.GetDB().AutoMigrate(&User{}, &FriendShipApply{}, &FriendShip{})
	return &dao{
		rdb: rdb,
		db:  db,
	}
}

// User Module
// SaveUser 新增用户
func (d *dao) SaveUser(u *User) error {
	return d.db.GetDB().Create(u).Error
}

// GetUser 获取用户
func (d *dao) GetUser(userId string) (u *User, err error) {
	err = d.db.Query().Where("user_id = ?", userId).First(&u).Error
	return
}

// BatchGetUser 批量获取用户
func (d *dao) BatchGetUser(userIds []string) (u []*User, err error) {
	err = d.db.Query().Where("user_id in ?", userIds).Find(&u).Error
	return
}

func (d *dao) GetUserByUsername(username string) (u *User, err error) {
	err = d.db.Query().Where("username = ?", username).First(&u).Error
	return
}

// DeleteUser 删除用户
func (d *dao) DeleteUser(userId string) error {
	return d.db.GetDB().Delete("user_id = ?", userId).Error
}

// Friend Module
// AddFriend 添加好友
func (d *dao) AddFriend(userId, otherUserId string) error {
	return d.db.GetDB().Create(&FriendShip{UserId: userId, FUserId: otherUserId}).Error
}

// RemoveFriend 删除好友
func (d *dao) RemoveFriend(userId, otherUserId string) error {
	return d.db.GetDB().Delete("user_id = ? AND f_user_id = ?", userId, otherUserId).Error
}

// GetFriendList 获取好友列表
func (d *dao) GetFriendList(userId string) (fs []*FriendShip, err error) {
	err = d.db.Query().Where("user_id = ?", userId).Find(&fs).Error
	return
}

// IsFriend 查询是否好友关系
func (d *dao) IsFriend(userId, friendId string) (bool, error) {
	var count int64
	err := d.db.Query().Model(&FriendShip{}).Where("user_id = ? AND f_user_id = ?", userId, friendId).Count(&count).Error
	return count > 0, err
}


// Apply Record Module
func (d *dao) GetApplyRecord(recordId string) (f *FriendShipApply, err error) {
	err = d.db.GetDB().Where("id = ? AND agree = 0", recordId).First(&f).Error
	return
}

func (d *dao) GetApplyRecordByUserId(userId, otherUserId string) (f *FriendShipApply, err error) {
	err = d.db.Query().Where("user_id = ? AND other_user_id = ?", userId, otherUserId).First(&f).Error
	return
}

func (d *dao) ListApplyRecord(userId string, applying bool) (fsaList []*FriendShipApply, err error) {
	m, _ := time.ParseDuration("-72h")
	sqlTxt := "%s = ? AND agree = 0 AND created_at >= ?"
	var who string
	if applying {
		who = "user_id"
	} else {
		who = "other_user_id"
	}
	err = d.db.Query().Where(fmt.Sprintf(sqlTxt, who), userId, time.Now().Add(m)).Find(&fsaList).Error
	return
}

func (d *dao) ListOffsetApplyRecord(userId, recordId string, applying bool) (fsaList []*FriendShipApply, err error) {
	m, _ := time.ParseDuration("-72h")
	sqlTxt := "%s = ? AND agree = 0 AND created_at >= ? AND id > ?"
	var who string
	if applying {
		who = "user_id"
	} else {
		who = "other_user_id"
	}
	err = d.db.Query().Where(fmt.Sprintf(sqlTxt, who), userId, time.Now().Add(m), recordId).Find(&fsaList).Error
	return
}

func (d *dao) UpdateApplyRecord(record *FriendShipApply) error {
	return d.db.GetDB().Where("id = ?", record.ID).Updates(record).Error
}

func (d *dao) SaveApplyRecord(userId, otherUserId, remark string) error {
	user, err := d.GetUser(userId)
	if err != nil {
		return err
	}

	return d.db.GetDB().Create(&FriendShipApply{
		UserId:      userId,
		Username:    user.Username,
		OtherUserId: otherUserId,
		Remark:      remark,
	}).Error
}

// Online status
func (d *dao) IsOnline(userId string) (bool, error) {
	return d.rdb.HExists("ONLINE_MAP", userId)
}


func (d *dao) SetOnline(userId string) error {
	return d.rdb.HSet("ONLINE_MAP", userId, 1)
}

func (d *dao) SetOffline(userId string) error {
	return d.rdb.HDel("ONLINE_MAP", userId)
}
