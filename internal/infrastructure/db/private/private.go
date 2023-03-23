package private

import (
	"context"

	"backend/internal/config"
	"backend/internal/infrastructure/db"
	"backend/internal/logger"

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
	logF := log.WithField("module", "db.private")

	dsn := db.GetDsn(
		cfg.DB.Private.Host,
		cfg.DB.Private.User,
		cfg.DB.Private.Password,
		cfg.DB.Private.Dbname,
		cfg.DB.Private.Sslmode,
		cfg.DB.Private.Port,
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
		logF.Fatalf("something went wrong: %e", err)

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
