package user

import (
	"fmt"
	"gorm.io/gorm"
	"maoim/pkg/merror"
	"maoim/pkg/mysql"
	"maoim/pkg/redis"
)

var _ Dao = (*dao)(nil)

type Dao interface {
	// User
	SaveUser(u *User) error
	GetUser(userId string) (*User, error)
	BatchGetUser(userIds []string) ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	DeleteUser(userId string) error
	UpdateUser(u *User) error

	// Friend
	AddFriend(userId, otherUserId string) error
	RemoveFriend(userId, otherUserId string) error
	GetFriendList(userId string) ([]*FriendShip, error)
	IsFriend(userId, otherUserId string) (bool, error)

	// Apply Record
	GetApplyRecord(recordId string) (*FriendShipApply, error)
	GetApplyRecordByUserId(userId, otherUserId string) (*FriendShipApply, error)
	SaveApplyRecord(userId, otherUserId *User, remark string) error
	ListApplyRecord(userId string, applying bool) ([]*FriendShipApply, error)
	ListOffsetApplyRecord(userId, recordId string, applying bool) ([]*FriendShipApply, error)
	UpdateApplyRecord(record *FriendShipApply) error
	BatchUpdateApplyRecord(records []*FriendShipApply) error

	// Online status
	SetOnline(userId string) error
	SetOffline(userId string) error
	IsOnline(userId string) (bool, error)
}

type dao struct {
	rdb *redis.Redis
	db  *mysql.Mysql
}

func (d *dao) UpdateUser(u *User) error {
	err := d.db.GetDB().Updates(u).Error
	if err != nil {
		err = merror.Wrap(err, "用户更新失败")
	}
	return err
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
	err := d.db.GetDB().Create(u).Error
	if err != nil {
		err = merror.Wrap(err, "新增用户失败")
	}
	return err
}

// GetUser 获取用户
func (d *dao) GetUser(userId string) (u *User, err error) {
	err = d.db.GetDB().Where("id = ?", userId).First(&u).Error
	if err != nil {
		err = merror.Wrap(err, "获取用户失败")
	}
	return
}

// BatchGetUser 批量获取用户
func (d *dao) BatchGetUser(userIds []string) (u []*User, err error) {
	err = d.db.GetDB().Where("id in ?", userIds).Find(&u).Error
	if err != nil {
		err = merror.Wrap(err, "批量获取用户失败")
	}
	return
}

func (d *dao) GetUserByUsername(username string) (u *User, err error) {
	err = d.db.GetDB().Where("username = ?", username).First(&u).Error
	if err != nil {
		err = merror.Wrap(err, "根据用户名查询用户失败")
	}
	return
}

// DeleteUser 删除用户
func (d *dao) DeleteUser(userId string) error {
	err := d.db.GetDB().Delete("id = ?", userId).Error
	if err != nil {
		err = merror.Wrap(err, "删除用户失败")
	}
	return err
}

// Friend Module
// AddFriend 添加好友
func (d *dao) AddFriend(userId, otherUserId string) error {
	err := d.db.GetDB().Create(&FriendShip{UserId: userId, FUserId: otherUserId}).Error
	if err != nil {
		err = merror.Wrap(err, "添加好友失败")
	}
	return err
}

// RemoveFriend 删除好友
func (d *dao) RemoveFriend(userId, otherUserId string) error {
	err := d.db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete("user_id = ? AND f_user_id = ?", userId, otherUserId).Error; err != nil {
			return err
		}
		if err := tx.Delete("user_id = ? AND f_user_id = ?", otherUserId, userId).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		err = merror.Wrap(err, "删除好友失败")
	}
	return err
}

// GetFriendList 获取好友列表
func (d *dao) GetFriendList(userId string) (fs []*FriendShip, err error) {
	err = d.db.GetDB().Where("user_id = ?", userId).Find(&fs).Error
	if err != nil {
		err = merror.Wrap(err, "获取好友列表失败")
	}
	return
}

// IsFriend 查询是否好友关系
func (d *dao) IsFriend(userId, friendId string) (bool, error) {
	var count int64
	err := d.db.GetDB().Model(&FriendShip{}).Where("user_id = ? AND f_user_id = ?", userId, friendId).Count(&count).Error
	if err != nil {
		err = merror.Wrap(err, "查询好友关系失败")
	}
	return count > 0, err
}

// Apply Record Module
// GetApplyRecord 获取好友申请记录
func (d *dao) GetApplyRecord(recordId string) (f *FriendShipApply, err error) {
	err = d.db.GetDB().Where("id = ? AND agree = 0", recordId).First(&f).Error
	if err != nil {
		err = merror.Wrap(err, "获取好友申请记录失败")
	}
	return
}

// GetApplyRecordByUserId 根据 userid 获取好友申请记录
func (d *dao) GetApplyRecordByUserId(userId, otherUserId string) (f *FriendShipApply, err error) {
	err = d.db.GetDB().Where("user_id = ? AND other_user_id = ?", userId, otherUserId).First(&f).Error
	if err != nil {
		err = merror.Wrap(err, "获取好友申请记录失败")
	}
	return
}

// ListApplyRecord 获取好友申请列表
func (d *dao) ListApplyRecord(userId string, applying bool) (fsaList []*FriendShipApply, err error) {
	sqlTxt := "%s = ? AND agree = 0"
	var who string
	if applying {
		who = "user_id"
	} else {
		who = "other_user_id"
	}
	err = d.db.GetDB().Where(fmt.Sprintf(sqlTxt, who), userId).Find(&fsaList).Error
	if err != nil {
		err = merror.Wrap(err, "获取好友申请记录列表失败")
	}
	return
}

//
func (d *dao) ListOffsetApplyRecord(userId, recordId string, applying bool) (fsaList []*FriendShipApply, err error) {
	sqlTxt := "%s = ? AND agree = 0 AND id > ?"
	var who string
	if applying {
		who = "user_id"
	} else {
		who = "other_user_id"
	}
	err = d.db.GetDB().Where(fmt.Sprintf(sqlTxt, who), userId, recordId).Find(&fsaList).Error
	if err != nil {
		err = merror.Wrap(err, "增量获取好友申请记录列表失败")
	}
	return
}

func (d *dao) UpdateApplyRecord(record *FriendShipApply) error {
	return d.db.GetDB().Where("id = ?", record.ID).Updates(record).Error
}

func (d *dao) SaveApplyRecord(me, other *User, remark string) error {
	err := d.db.GetDB().Create(&FriendShipApply{
		UserId:        me.ID,
		Username:      me.Username,
		OtherUserId:   other.ID,
		OtherUsername: other.Nickname,
		Remark:        remark,
		Status:        WAIT_VERY,
	}).Error
	if err != nil {
		err = merror.Wrap(err, "保存好友申请记录失败")
	}
	return err
}

func (d *dao) BatchUpdateApplyRecord(records []*FriendShipApply) error {
	err := d.db.GetDB().Updates(records).Error
	if err != nil {
		err = merror.Wrap(err, "批量更新申请记录失败")
	}
	return err
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
