package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/geiqin/xconfig/model"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"time"
)

type OauthHelper struct {
	PrivateKey        string
	RedisAddr         string
	RedisDB           int
	clients           []*models.Client
	accessTokenExp    time.Duration
	refreshTokenExp   time.Duration
	isGenerateRefresh bool
}

var mgr *manage.Manager

func NewOauthHelper(cfg *model.TokenInfo) *OauthHelper {
	if cfg == nil {
		cfg = &model.TokenInfo{}
	}
	if cfg.AccessTokenExp == 0 {
		cfg.AccessTokenExp = 120
	}
	if cfg.RefreshTokenExp == 0 {
		cfg.RefreshTokenExp = 60 * 24 * 7
	}

	return &OauthHelper{
		RedisAddr:         cfg.RedisAddr,
		RedisDB:           cfg.RedisDB,
		PrivateKey:        cfg.PrivateKey,
		accessTokenExp:    time.Duration(cfg.AccessTokenExp) * time.Hour,
		refreshTokenExp:   time.Duration(cfg.RefreshTokenExp) * time.Hour,
		isGenerateRefresh: cfg.IsGenerateRefresh,
	}

}

func (b *OauthHelper) LoadClient() {
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

func (b *OauthHelper) AddClient(client *models.Client) {
	b.clients = append(b.clients, client)
}

func (b *OauthHelper) AddClients(clients []*models.Client) {
	for _, v := range clients {
		b.clients = append(b.clients, &models.Client{
			ID:     v.ID,
			Secret: v.Secret,
			Domain: v.Domain,
			UserID: v.UserID,
		})
	}
}

func (b *OauthHelper) GetConfig() *manage.Config {
	return &manage.Config{
		AccessTokenExp:    b.accessTokenExp,
		RefreshTokenExp:   b.refreshTokenExp,
		IsGenerateRefresh: b.isGenerateRefresh,
	}
}

func (b *OauthHelper) GetManager() *manage.Manager {
	if mgr != nil {
		return mgr
	}
	mgr = manage.NewDefaultManager()
	var err error
	mgr.SetPasswordTokenCfg(manage.DefaultPasswordTokenCfg)
	mgr.MustTokenStorage(store.NewFileTokenStore("data.db"))
	/*
		mgr.MustTokenStorage(oredis.NewRedisStore(&redis.Options{
			Addr: b.RedisAddr,
			DB:   b.RedisDB,
		}, "plaza"), err) //{access_token}

	*/

	if err != nil {
		log.Println("MustTokenStorage err:", err.Error())
	}
	/*
		mgr.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
			Addr: b.RedisAddr,
			DB:   b.RedisDB,
		}))

	*/

	//log.Println("getconfig:", helper.JsonEncode(b.GetConfig()))
	/*
		mgr.SetClientTokenCfg(b.GetConfig())
		mgr.SetPasswordTokenCfg(b.GetConfig())

		mgr.SetRefreshTokenCfg(&manage.RefreshingConfig{
			AccessTokenExp:    b.accessTokenExp,
			RefreshTokenExp:   b.refreshTokenExp,
			IsGenerateRefresh: b.isGenerateRefresh,
		})

	*/

	mgr.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(b.PrivateKey), jwt.SigningMethodHS512))
	return mgr
}

func (b *OauthHelper) GetSrv() *server.Server {
	mgr := b.GetManager()
	srv := server.NewDefaultServer(mgr)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	return srv
}
