package access_token

import (
	"github.com/ShagaleevAlexey/go-clean/internal/usecases"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiration   int64  `json:"expiration"`
}

type AccessTokenClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func NewAccessTokenClaims(uid uuid.UUID, exp time.Time, issuer string) *AccessTokenClaims {
	return &AccessTokenClaims{
		uid.String(),
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    issuer,
		},
	}
}

type RefreshTokenClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func NewRefreshTokenClaims(uid uuid.UUID, exp time.Time, issuer string) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		uid.String(),
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    issuer,
		},
	}
}

func GenerateTokens(uid uuid.UUID, sig []byte) (*AuthTokens, error) {
	exp := time.Now().Add(time.Hour * 24 * 30)
	aClaims := NewAccessTokenClaims(uid, exp, "test")

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS512, aClaims)
	accessToken, err := aToken.SignedString(sig)
	if err != nil {
		return nil, err
	}

	rClaims := NewRefreshTokenClaims(uid, time.Now().Add(time.Hour*24*30*10), "test")

	rToken := jwt.NewWithClaims(jwt.SigningMethodHS512, rClaims)
	refreshToken, err := rToken.SignedString(sig)
	if err != nil {
		return nil, err
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiration:   exp.Unix(),
	}, nil
}

func ValidAccessToken(token string, sig []byte) (*AccessTokenClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sig, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*AccessTokenClaims); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, usecases.TokenIsInvalidError
	}
}
