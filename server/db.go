package server

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DbStart() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}
	return dbpool, nil
}

func createDbTransaction(pool *pgxpool.Pool) (pgx.Tx, error) {
	tx, err := pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("start db transaction: %w", err)
	}
	return tx, nil
}

func CreateNewPlantRow(pool *pgxpool.Pool, name, species, location string) error {
	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("generate plant id: %w", err)
	}

	tx, err := createDbTransaction(pool)
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(context.Background())
		}
	}()

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO plants (id, name, species, location) VALUES ($1, $2, $3, $4)",
		id,
		name,
		species,
		location,
	)
	if err != nil {
		return fmt.Errorf("insert plant row: %w", err)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("commit plant row: %w", err)
	}

	committed = true
	return nil
}

func GetPlantRow(pool *pgxpool.Pool, UUID string) error {
	tx, err := createDbTransaction(pool)
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(context.Background())
		}
	}()

	_, err = tx.Exec(
		context.Background(),
		"SELECT * WHERE id = $1", UUID,
	)
	if err != nil {
		return fmt.Errorf("insert plant row: %w", err)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("commit plant row: %w", err)
	}

	committed = true
	return nil
}
