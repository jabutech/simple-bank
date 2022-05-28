package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all function to executr db queries amd transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function withim a database transacton
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Start database transaction
	tx, err := store.db.BeginTx(ctx, nil)
	// If error, return error
	if err != nil {
		return err
	}

	// If no, create new transaction, with argument tx
	q := New(tx)
	// If no error, run query, with argument transaction
	err = fn(q)
	// If error
	if err != nil {
		// Do rollback transaction, and check if rollback error
		if rollbackError := tx.Rollback(); rollbackError != nil {
			// Return error rollback
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rollbackError)
		}
		// If no, return only error transaction
		return err
	}

	// Return commit, if all operation success
	return tx.Commit()
}

// TransferTxParam contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_acount"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

// TransferTX perfoms a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts balance within a single database transaction.
func (store *Store) TransferTX(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update accounts' balance
		// To overcome deadlock, do check
		// if fromAccountID smaller than ToAccount ID
		if arg.FromAccountID < arg.ToAccountID {
			// Update first FromAccountID. And next, update ToAccountID
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			// If no, update first ToAccountID. And next, update FromAccountID
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return nil
	})

	return result, err
}

// Function for add money
func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Accounts, account2 Accounts, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}
