package connecter

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connecter struct {
	DB *gorm.DB
}

func NewConnecter() *Connecter {
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName,
		os.Getenv("DB_LOC"))
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return &Connecter{
		DB: gormDB,
	}
}

func (c *Connecter) Get() *gorm.DB {
	return c.DB
}

func (c *Connecter) TransactinHandler(txFunc func() (interface{}, error)) (data interface{}, err error) {
	tx := c.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	data, err = txFunc()
	return
}
