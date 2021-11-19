package runtime

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
)

/*
Application 实现Runtime
 */
type Application struct {
	DB  *gorm.DB
	mux sync.RWMutex
	Logger *log.Logger
	Handler http.Handler
}

func (a *Application) SetEngine(engine http.Handler) {
	a.Handler = engine
}

func (a *Application) GetEngine() http.Handler {
	return a.Handler
}

func NewApplication() Runtime {
	return &Application{
		DB: nil,
	}
}

func (a *Application) SetDB(db *gorm.DB) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.DB = db
	return
}

func (a *Application) GetDB() *gorm.DB  {
	return a.DB
}

func (a *Application) SetLogger(logger *log.Logger) {
	a.Logger = logger
}

func (a *Application) GetLogger() *log.Logger {
	return a.Logger
}
