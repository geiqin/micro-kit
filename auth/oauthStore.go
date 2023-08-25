package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/geiqin/xconfig/model"
	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"time"
)

type OauthStore struct {
	PrivateKey        string
	RedisAddr         string
	RedisDB           int
	clients           []*models.Client
	accessTokenExp    time.Duration
	refreshTokenExp   time.Duration
	isGenerateRefresh bool
}

var globalMgr *manage.Manager

func NewOauthStore(cfg *model.TokenInfo) *OauthStore {
	if cfg == nil {
		cfg = &model.TokenInfo{}
	}
	if cfg.AccessTokenExp == 0 {
		cfg.AccessTokenExp = 120
	}
	if cfg.RefreshTokenExp == 0 {
		cfg.RefreshTokenExp = 60 * 24 * 7
	}

	return &OauthStore{
		RedisAddr:         cfg.RedisAddr,
		RedisDB:           cfg.RedisDB,
		PrivateKey:        cfg.PrivateKey,
		accessTokenExp:    time.Duration(cfg.AccessTokenExp) * time.Hour,
		refreshTokenExp:   time.Duration(cfg.RefreshTokenExp) * time.Hour,
		isGenerateRefresh: cfg.IsGenerateRefresh,
	}

}

func (b *OauthStore) LoadClient() {
	manager := b.GetManager()
	clientStore := store.NewClientStore()
	for _, v := range b.clients {
		clientStore.Set(v.ID, &models.Client{
			ID:     v.ID,
			Secret: v.Secret,
			Domain: v.Domain,
			UserID: v.UserID,
		})
	}
	manager.MapClientStorage(clientStore)
}

func (b *OauthStore) AddClient(client *models.Client) {
	b.clients = append(b.clients, client)
}

func (b *OauthStore) AddClients(clients []*models.Client) {
	for _, v := range clients {
		b.clients = append(b.clients, &models.Client{
			ID:     v.ID,
			Secret: v.Secret,
			Domain: v.Domain,
			UserID: v.UserID,
		})
	}
}

func (b *OauthStore) GetConfig() *manage.Config {
	return &manage.Config{
		AccessTokenExp:    b.accessTokenExp,
		RefreshTokenExp:   b.refreshTokenExp,
		IsGenerateRefresh: b.isGenerateRefresh,
	}
}

func (b *OauthStore) GetManager() *manage.Manager {
	if globalMgr != nil {
		return globalMgr
	}

	globalMgr = manage.NewDefaultManager()

	globalMgr.SetPasswordTokenCfg(manage.DefaultPasswordTokenCfg)

	// use redis token store
	globalMgr.MapTokenStorage(nil)

	globalMgr.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: b.RedisAddr,
		DB:   b.RedisDB,
	}))

	globalMgr.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(b.PrivateKey), jwt.SigningMethodHS512))

	return globalMgr
}

func (b *OauthStore) GetSrv() *server.Server {
	mgr := b.GetManager()
	srv := server.NewDefaultServer(mgr)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	return srv
}
