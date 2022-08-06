package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddUser(ctx context.Context, user models.User) error {
	query, args, err := squirrel.Insert("users").
		Columns("username").
		Values(user.Username).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.AddUser: to sql: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository.AddUser: to sql: %w", err)
	}
	return nil
}

func (r *repository) GetUserById(ctx context.Context, id uint) (models.User, error) {
	query, args, err := squirrel.Select("users").
		Columns("id", "username").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return models.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	var user models.User
	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		return models.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	return user, nil
}

func (r *repository) DeleteUser() error {
	return errors.New("delete user: not implemented")
}

func (r *repository) ListUsers() error {
	return errors.New("list users: not implemented")
}
