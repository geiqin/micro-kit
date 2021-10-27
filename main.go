package main

import (
	"github.com/geiqin/micro-kit/auth"
	"time"
)

type Say struct{}



func main() {
	oh := auth.NewOauthHelper(&auth.OauthConfig{
		Key:               "123456789",
		RedisAddr:         "172.23.166.100:6379",
		RedisDB:           15,
		AccessTokenExp:    time.Minute * 3,
		RefreshTokenExp:   time.Hour * 24 * 7,
		IsGenerateRefresh: true,
	})
	oh.GetConfig()
}
