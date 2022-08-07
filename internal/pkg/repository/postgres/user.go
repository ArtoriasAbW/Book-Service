package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddUser(ctx context.Context, user repoModels.User) error {
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

func (r *repository) GetUserById(ctx context.Context, id uint) (repoModels.User, error) {
	query, args, err := squirrel.Select("id", "username").
		From("users").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return repoModels.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return repoModels.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	var user repoModels.User
	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		return repoModels.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	return user, nil
}

func (r *repository) DeleteUser(ctx context.Context, id uint) error {
	query, args, err := squirrel.Delete("users").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.DeleteUser: to sql: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("Repository.DeleteUser: to sql: %w", err)
	}
	return nil
}

func (r *repository) ListUsers() error {
	return errors.New("list users: not implemented")
}
