package private

import (
	"backend/internal/config"
	"backend/internal/infrastructure/db"
	"backend/internal/logger"
	"context"
	pgxdec "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct {
	Dsn  string
	Pool *pgxpool.Pool
}

func NewProvider(cfg *config.Config, log *logger.Logger) (*Connection, error) {
	l := log.WithField("module", "db.private")

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
		l.Fatalf("something went wrong: %e", err)
		return nil, err
	}

	return &Connection{
		Pool: dbPool,
		Dsn:  dsn,
	}, err
}
