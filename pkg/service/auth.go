package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"web"
	"web/pkg/repo"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "slfnoinrf90h390fnviofkl"
	signingKey = "ahrtylsjdljq2li4hkw"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	ShopId int `json:"shop_id"`
}
type AuthService struct {
	repo repo.Authorization
}

func NewAuthService(repo repo.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateShop(shop web.Shop) (int, error) {
	shop.Password = passHash(shop.Password)
	return s.repo.CreateShop(shop)
}

func (s *AuthService) GenerateToken(shopname, password string) (string, error) {
	shop, err := s.repo.GetShop(shopname, passHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		shop.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesstok string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstok, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.ShopId, nil
}
func passHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
