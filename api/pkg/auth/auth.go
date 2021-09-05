package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/aead/chacha20poly1305"
	u "github.com/esh2n/xmysql-go-nginx/api/pkg/domain/user"
	"github.com/o1egl/paseto"
)

type Payload struct {
	User      *u.User   `json:"user"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	footer       []byte
}

func NewPasetoMaker() (*PasetoMaker, error) {
	symmetricKey := os.Getenv("SYMMETRICKEY")
	footer := os.Getenv("FOOTERKEY")
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return &PasetoMaker{}, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		footer:       []byte(footer),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateTokenString(user *u.User) (string, error) {
	now := time.Now()
	exp := now.Add(time.Hour * 24 * 15)

	token, err := maker.paseto.Encrypt(maker.symmetricKey, &Payload{
		User:      user,
		IssuedAt:  now,
		ExpiredAt: exp,
	}, maker.footer)
	if err != nil {
		return "", fmt.Errorf("%d", err)
	}
	return token, err
}
