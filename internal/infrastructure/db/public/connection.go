package public

import (
	"backend/internal/infrastructure/config"
	"context"
	"github.com/jackc/pgx/v5/pgtype"

	"backend/internal/infrastructure/db"
	"backend/internal/logger"
	pgxdec "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

type Connection struct {
	Dsn     string
	Pool    *pgxpool.Pool
	TypeMap *pgtype.Map
}

func NewProvider(cfg *config.Config, log *logger.Logger, lc fx.Lifecycle) (*Connection, error) {
	l := log.WithField("module", "db.private")

	dsn := db.GetDsn(
		cfg.DB.Public.Host,
		cfg.DB.Public.User,
		cfg.DB.Public.Password,
		cfg.DB.Public.Dbname,
		cfg.DB.Public.Schema,
		cfg.DB.Public.Sslmode,
		cfg.DB.Public.Port,
	)

	dbCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		Dsn: dsn,
	}

	dbCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdec.Register(conn.TypeMap())

		return nil
	}

	connection.Pool, err = pgxpool.NewWithConfig(context.Background(), dbCfg)
	if err != nil {
		l.Fatalf("something went wrong: %e", err)
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			connection.Pool.Close()

			return nil
		},
	})

	return connection, err
}
