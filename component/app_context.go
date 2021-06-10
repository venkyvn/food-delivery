package component

import (
	"go-food-delivery/component/uploadprovider"
	"go-food-delivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pubSub         pubsub.PubSub
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider uploadprovider.UploadProvider,
	secretKey string,
	pubSub pubsub.PubSub,
) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
		secretKey:      secretKey,
		pubSub:         pubSub,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.pubSub
}
