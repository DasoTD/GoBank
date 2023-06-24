package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Bank struct {
	*Queries
	db *sql.DB
	
}

func NewBank(db *sql.DB) *Bank {
	return &Bank{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (bank *Bank) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := bank.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}


// func (bank *Bank) TransferTx(ctx context.Context, arg TransferTxParams)