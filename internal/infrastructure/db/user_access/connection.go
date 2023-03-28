package user_access

import (
	"backend/internal/config"
	"backend/internal/infrastructure/db"
	"backend/internal/logger"
	"context"
	pgxdec "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

type Connection struct {
	Dsn  string
	Pool *pgxpool.Pool
}

func NewProvider(cfg *config.Config, log *logger.Logger, lc fx.Lifecycle) (*Connection, error) {
	l := log.WithField("module", "db.private")

	dsn := db.GetDsn(
		cfg.DB.UserAccess.Host,
		cfg.DB.UserAccess.User,
		cfg.DB.UserAccess.Password,
		cfg.DB.UserAccess.Dbname,
		cfg.DB.UserAccess.Sslmode,
		cfg.DB.UserAccess.Port,
	)

	dbCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	dbCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdec.Register(conn.TypeMap())

		return nil
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbCfg)
	if err != nil {
		l.Fatalf("something went wrong: %e", err)

		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			dbPool.Close()

			return nil
		},
	})

	return &Connection{
		Pool: dbPool,
		Dsn:  dsn,
	}, err
}
