package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/locpham24/go-coffee/app/model/redis"
	"github.com/locpham24/go-coffee/config"
	"github.com/twinj/uuid"
	"time"
)

type UserClaims struct {
	UserId      int    `json:"userId"`
	Role        string `json:"role"`
	AccessUuid  string `json:"accessUuid"`
	RefreshUuid string `json:"refreshUuid"`
	jwt.StandardClaims
}

func CreateToken(userId int) (*redis.TokenDetails, error) {
	var err error
	redisTokenDetail := &redis.TokenDetails{}

	// access token will expire after 15 minute
	redisTokenDetail.AtExpires = time.Now().Add(time.Duration(config.Get().JwtToken.AccessTokenMaxAge) * time.Second).Unix()
	redisTokenDetail.AccessUuid = uuid.NewV4().String()

	// refresh token will expire after 7 days
	redisTokenDetail.RtExpires = time.Now().Add(time.Duration(config.Get().JwtToken.RefreshTokenMaxAge) * time.Second).Unix()
	redisTokenDetail.RefreshUuid = uuid.NewV4().String()

	accessKey := []byte(config.Get().JwtToken.AccessTokenSecretKey)
	refreshKey := []byte(config.Get().JwtToken.RefreshTokenSecretKey)

	// Create the Claims
	accessTokenClaims := UserClaims{
		UserId:     userId,
		AccessUuid: redisTokenDetail.AccessUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: redisTokenDetail.AtExpires,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	redisTokenDetail.AccessToken, err = token.SignedString(accessKey)

	refreshTokenClaims := UserClaims{
		UserId:      userId,
		RefreshUuid: redisTokenDetail.RefreshUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: redisTokenDetail.RtExpires,
			Issuer:    "test",
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	redisTokenDetail.RefreshToken, err = token.SignedString(refreshKey)

	if err != nil {
		return nil, err
	}
	return redisTokenDetail, nil
}
