package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrRepositoryNotFound          = errors.New("repository not found")
	ErrRepositoryAlreadyRegistered = errors.New("repository already registered")
	ErrTransactionNotStarted       = errors.New("transaction not started")
	ErrTransactionAlreadyStarted   = errors.New("transaction already started")
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fn RepositoryFactory) error
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	Unregister(name string) error
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(db *sql.DB) UowInterface {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fn RepositoryFactory) error {
	if _, ok := u.Repositories[name]; ok {
		return ErrRepositoryAlreadyRegistered
	}
	u.Repositories[name] = fn
	return nil
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if _, ok := u.Repositories[name]; !ok {
		return nil, ErrRepositoryNotFound
	}
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow UowInterface) error) error {
	if u.Tx != nil {
		return ErrTransactionAlreadyStarted
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx
	err = fn(u) // execute the function passed as argument in the UOW scope
	if err != nil {
		return u.rollbackAndWrapError(err)
	}
	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	if u.Tx == nil {
		return ErrTransactionNotStarted
	}
	err := u.Tx.Commit()
	if err != nil {
		return u.rollbackAndWrapError(err)
	}
	u.Tx = nil
	return nil
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return ErrTransactionNotStarted
	}
	err := u.Tx.Rollback()
	if err != nil {
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) rollbackAndWrapError(originalErr error) error {
	if errRollback := u.Rollback(); errRollback != nil {
		return fmt.Errorf("rollback error: %v, original error: %w", errRollback, originalErr)
	}
	return originalErr
}

func (u *Uow) Unregister(name string) error {
	if _, ok := u.Repositories[name]; !ok {
		return ErrRepositoryNotFound
	}
	delete(u.Repositories, name)
	return nil
}
