// Code generated by sqlc. DO NOT EDIT.

package wordle

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countAccountsByDiscordIdStmt, err = db.PrepareContext(ctx, countAccountsByDiscordId); err != nil {
		return nil, fmt.Errorf("error preparing query CountAccountsByDiscordId: %w", err)
	}
	if q.countNicknameByDiscordIdStmt, err = db.PrepareContext(ctx, countNicknameByDiscordId); err != nil {
		return nil, fmt.Errorf("error preparing query CountNicknameByDiscordId: %w", err)
	}
	if q.countScoresByDiscordIdStmt, err = db.PrepareContext(ctx, countScoresByDiscordId); err != nil {
		return nil, fmt.Errorf("error preparing query CountScoresByDiscordId: %w", err)
	}
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createNicknameStmt, err = db.PrepareContext(ctx, createNickname); err != nil {
		return nil, fmt.Errorf("error preparing query CreateNickname: %w", err)
	}
	if q.createScoreStmt, err = db.PrepareContext(ctx, createScore); err != nil {
		return nil, fmt.Errorf("error preparing query CreateScore: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteNicknameStmt, err = db.PrepareContext(ctx, deleteNickname); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteNickname: %w", err)
	}
	if q.deleteScoresForUserStmt, err = db.PrepareContext(ctx, deleteScoresForUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteScoresForUser: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getNicknameStmt, err = db.PrepareContext(ctx, getNickname); err != nil {
		return nil, fmt.Errorf("error preparing query GetNickname: %w", err)
	}
	if q.getScoreHistoryByAccountStmt, err = db.PrepareContext(ctx, getScoreHistoryByAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetScoreHistoryByAccount: %w", err)
	}
	if q.listAccountsStmt, err = db.PrepareContext(ctx, listAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccounts: %w", err)
	}
	if q.listNicknamesStmt, err = db.PrepareContext(ctx, listNicknames); err != nil {
		return nil, fmt.Errorf("error preparing query ListNicknames: %w", err)
	}
	if q.listScoresStmt, err = db.PrepareContext(ctx, listScores); err != nil {
		return nil, fmt.Errorf("error preparing query ListScores: %w", err)
	}
	if q.updateNicknameStmt, err = db.PrepareContext(ctx, updateNickname); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNickname: %w", err)
	}
	if q.updateScoreStmt, err = db.PrepareContext(ctx, updateScore); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateScore: %w", err)
	}
	if q.updateTimeZoneStmt, err = db.PrepareContext(ctx, updateTimeZone); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTimeZone: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countAccountsByDiscordIdStmt != nil {
		if cerr := q.countAccountsByDiscordIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countAccountsByDiscordIdStmt: %w", cerr)
		}
	}
	if q.countNicknameByDiscordIdStmt != nil {
		if cerr := q.countNicknameByDiscordIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countNicknameByDiscordIdStmt: %w", cerr)
		}
	}
	if q.countScoresByDiscordIdStmt != nil {
		if cerr := q.countScoresByDiscordIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countScoresByDiscordIdStmt: %w", cerr)
		}
	}
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createNicknameStmt != nil {
		if cerr := q.createNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createNicknameStmt: %w", cerr)
		}
	}
	if q.createScoreStmt != nil {
		if cerr := q.createScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScoreStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteNicknameStmt != nil {
		if cerr := q.deleteNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteNicknameStmt: %w", cerr)
		}
	}
	if q.deleteScoresForUserStmt != nil {
		if cerr := q.deleteScoresForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScoresForUserStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getNicknameStmt != nil {
		if cerr := q.getNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNicknameStmt: %w", cerr)
		}
	}
	if q.getScoreHistoryByAccountStmt != nil {
		if cerr := q.getScoreHistoryByAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScoreHistoryByAccountStmt: %w", cerr)
		}
	}
	if q.listAccountsStmt != nil {
		if cerr := q.listAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccountsStmt: %w", cerr)
		}
	}
	if q.listNicknamesStmt != nil {
		if cerr := q.listNicknamesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNicknamesStmt: %w", cerr)
		}
	}
	if q.listScoresStmt != nil {
		if cerr := q.listScoresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listScoresStmt: %w", cerr)
		}
	}
	if q.updateNicknameStmt != nil {
		if cerr := q.updateNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNicknameStmt: %w", cerr)
		}
	}
	if q.updateScoreStmt != nil {
		if cerr := q.updateScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateScoreStmt: %w", cerr)
		}
	}
	if q.updateTimeZoneStmt != nil {
		if cerr := q.updateTimeZoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTimeZoneStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	countAccountsByDiscordIdStmt *sql.Stmt
	countNicknameByDiscordIdStmt *sql.Stmt
	countScoresByDiscordIdStmt   *sql.Stmt
	createAccountStmt            *sql.Stmt
	createNicknameStmt           *sql.Stmt
	createScoreStmt              *sql.Stmt
	deleteAccountStmt            *sql.Stmt
	deleteNicknameStmt           *sql.Stmt
	deleteScoresForUserStmt      *sql.Stmt
	getAccountStmt               *sql.Stmt
	getNicknameStmt              *sql.Stmt
	getScoreHistoryByAccountStmt *sql.Stmt
	listAccountsStmt             *sql.Stmt
	listNicknamesStmt            *sql.Stmt
	listScoresStmt               *sql.Stmt
	updateNicknameStmt           *sql.Stmt
	updateScoreStmt              *sql.Stmt
	updateTimeZoneStmt           *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		countAccountsByDiscordIdStmt: q.countAccountsByDiscordIdStmt,
		countNicknameByDiscordIdStmt: q.countNicknameByDiscordIdStmt,
		countScoresByDiscordIdStmt:   q.countScoresByDiscordIdStmt,
		createAccountStmt:            q.createAccountStmt,
		createNicknameStmt:           q.createNicknameStmt,
		createScoreStmt:              q.createScoreStmt,
		deleteAccountStmt:            q.deleteAccountStmt,
		deleteNicknameStmt:           q.deleteNicknameStmt,
		deleteScoresForUserStmt:      q.deleteScoresForUserStmt,
		getAccountStmt:               q.getAccountStmt,
		getNicknameStmt:              q.getNicknameStmt,
		getScoreHistoryByAccountStmt: q.getScoreHistoryByAccountStmt,
		listAccountsStmt:             q.listAccountsStmt,
		listNicknamesStmt:            q.listNicknamesStmt,
		listScoresStmt:               q.listScoresStmt,
		updateNicknameStmt:           q.updateNicknameStmt,
		updateScoreStmt:              q.updateScoreStmt,
		updateTimeZoneStmt:           q.updateTimeZoneStmt,
	}
}