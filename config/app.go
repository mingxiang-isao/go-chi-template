package config

import (
	"database/sql"
	"go_chi_template/config/provider"
	"path"
	"runtime"

	"github.com/go-chi/jwtauth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type App struct {
	env     *provider.EnvProvider
	db      *sql.DB
	redis   *redis.Client
	rootDir string
	logger  *zap.Logger
	jwtAuth *jwtauth.JWTAuth
}

func (app *App) EnvVars() *provider.EnvProvider {
	return app.env
}

func (app *App) DB() *sql.DB {
	return app.db
}

func (app *App) Redis() *redis.Client {
	return app.redis
}

func (app *App) Logger() *zap.Logger {
	return app.logger
}

func (app *App) JWTAuth() *jwtauth.JWTAuth {
	return app.jwtAuth
}

func (app *App) setRootDir() {
	_, b, _, _ := runtime.Caller(0)
	app.rootDir = path.Join(path.Dir(b), "..")
}

func (app *App) UseTestDB() {
	app.db = provider.NewTestDbProvider(app.env)
}

func NewApp() *App {
	app := App{}

	app.setRootDir()
	app.env = provider.NewEnvProvider(app.rootDir)
	app.db = provider.NewDbProvider(app.env)
	app.redis = provider.NewRedisProvider(app.env)
	app.logger = provider.NewLoggerProvider(app.env)
	app.jwtAuth = provider.NewJWTAuth(app.env)

	provider.NewValidationProvider()

	return &app
}
