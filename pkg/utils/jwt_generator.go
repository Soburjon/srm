package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/work/SRM/configs"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var config = configs.Config()

// Tokens struct to describe tokens object.
type Tokens struct {
	Access    string
	AccExpire int64
	Refresh   string
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id string, credentials map[string]string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, expire, err := generateNewAccessToken(id, credentials)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:    accessToken,
		Refresh:   refreshToken,
		AccExpire: expire,
	}, nil
}

func generateNewAccessToken(id string, credentials map[string]string) (string, int64, error) {

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["role"] = credentials["role"]

	// in local server access token ttl = 31 days
	if config.Environment == "develop" {
		claims["expires"] = time.Now().Add(time.Hour * 24 * 31).Unix()
	} else {
		// in staging server access token ttl = day
		claims["expires"] = time.Now().Add(time.Minute * 2 * time.Duration(config.JWTSecretKeyExpireMinutes)).Unix()
	}
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(config.JWTSecretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", 0, err
	}

	return t, claims["expires"].(int64), nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	sha256 := sha256.New()

	// Create a new now date and time string with salt.
	refresh := config.JWTRefreshKey + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := sha256.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(config.JWTRefreshKeyExpireHours)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(sha256.Sum(nil)) + "." + expireTime

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
