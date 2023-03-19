package private

import (
	"backend/internal/config"
	"backend/internal/infrastructure/db"
	"backend/internal/logger"
	"go.uber.org/fx"
)

type Connection struct {
	*db.Connection
}

func NewProvider(cfg *config.Config, log *logger.Logger, lc fx.Lifecycle) (*Connection, error) {
	l := log.WithField("module", "db.private")
	connection, err := db.NewConnection(
		cfg.DB.Private.Host,
		cfg.DB.Private.User,
		cfg.DB.Private.Password,
		cfg.DB.Private.Dbname,
		cfg.DB.Private.Sslmode,
		cfg.DB.Private.Port,
	)
	if err != nil {
		return nil, err
	}
	if err := connection.Db.Ping(); err != nil {
		l.Warn("cannot ping database. does it work?")
	}
	return &Connection{Connection: connection}, err
}
