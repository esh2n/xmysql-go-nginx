package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    string `gorm:"primary_key"`
	Name  string `json:"name"`
	Token string
}

func NewUser(name string) *User {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ID := ulid.MustNew(ulid.Timestamp(t), entropy)
	return &User{
		ID:   ID.String(),
		Name: name,
	}
}
