package bootstrap

import (
	"fertigation-Service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/bootstrap/db"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env   *Env
	DB    *pg.DB
	Log   *log.LogGRPCImpl
	Repos *repo.Repositories
}

func App() *Application {
	env := Env{}
	NewEnv(&env)
	logConfig := log.NewConfig()
	log := log.InitLogGRPC(logConfig, zapcore.DebugLevel, env.IsProduction())
	db := db.NewPostgresDB(db.ConfigDB{
		URL:  env.UrlDb,
		Mode: env.NodeEnv,
	})

	repos := repo.NewRepositories(db)

	return &Application{
		Env:   &env,
		DB:    db,
		Log:   log,
		Repos: repos,
	}
}
