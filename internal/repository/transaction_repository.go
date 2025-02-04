package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-REST/internal/model"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Deposit(userID int, amount float64) error {
	query := `
        INSERT INTO transactions (user_id, amount, type)
        VALUES ($1, $2, 'deposit')
    `
	_, err := r.db.Exec(context.Background(), query, userID, amount)
	return err
}

func (r *TransactionRepository) Transfer(senderID, recipientID int, amount float64) error {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	querySender := `
        INSERT INTO transactions (user_id, amount, type)
        VALUES ($1, -$2, 'transfer_out')
    `
	_, err = tx.Exec(context.Background(), querySender, senderID, amount)
	if err != nil {
		return err
	}

	queryRecipient := `
        INSERT INTO transactions (user_id, amount, type)
        VALUES ($1, $2, 'transfer_in')
    `
	_, err = tx.Exec(context.Background(), queryRecipient, recipientID, amount)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (r *TransactionRepository) GetLastTransactions(userID int) ([]model.Transaction, error) {
	query := `
        SELECT id, user_id, amount, type, created_at
        FROM transactions
        WHERE user_id = $1
        ORDER BY created_at DESC
        LIMIT 10
    `
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Type, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
