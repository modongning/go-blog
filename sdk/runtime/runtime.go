package runtime

import (
	"gorm.io/gorm"
	"net/http"
)

type Runtime interface {
	SetDB(*gorm.DB)
	GetDB() *gorm.DB
	SetEngine(http.Handler)
	GetEngine() http.Handler
}
