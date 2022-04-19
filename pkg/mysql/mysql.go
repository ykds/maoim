package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db *gorm.DB
}

func New(conf *Config) *Mysql {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Mysql{db: db}
}

func (m *Mysql) GetDB() *gorm.DB {
	return m.db
}

func (m *Mysql) Query() *gorm.DB {
	return m.db.Scopes(UnDeletedScope)
}
