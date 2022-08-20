package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddUser(ctx context.Context, user repoModels.User) (uint64, error) {
	query, args, err := squirrel.Insert("users").
		Columns("username").
		Values(user.Username).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Repository.AddUser: to sql: %w", err)
	}
	// row := r.pool.QueryRow(ctx, query, args...)
	row := r.db.QueryRowContext(ctx, query, args...)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Repository.AddUser: to sql: %w", err)
	}
	return id, nil
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
	// rows, err := r.pool.Query(ctx, query, args...)
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return repoModels.User{}, fmt.Errorf("Repository.GetUserById: to sql: %w", err)
	}
	var user repoModels.User
	ok := rows.Next()
	if !ok {
		return repoModels.User{}, errors.New("no user")
	}
	err = rows.StructScan(&user)
	// err = pgxscan.ScanOne(&user, rows)
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
	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("Repository.DeleteUser: to sql: %w", err)
	}
	return nil
}

func (r *repository) ListUsers(ctx context.Context, params repoModels.ListInput) ([]repoModels.User, error) {
	preparedQuery := squirrel.Select("id", "username").
		From("users").
		OrderBy("username " + params.Order).
		Offset(uint64(params.Offset))
	if params.Limit > 0 {
		preparedQuery = preparedQuery.Limit(uint64(params.Limit))
	}
	query, args, err := preparedQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return []repoModels.User{}, fmt.Errorf("Repository.ListUsers: to sql: %w", err)
	}
	var users []repoModels.User
	if err := r.db.SelectContext(ctx, &users, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.ListUsers: to sql: %w", err)
	}
	return users, nil
}

func (r *repository) UpdateUser(ctx context.Context, newUser repoModels.User) error {
	query, args, err := squirrel.Update("users").
		Set("username", newUser.Username).
		Where(
			squirrel.Eq{
				"id": newUser.Id,
			},
		).PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository:UpdateUser: to sql: %w", err)
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository:UpdateUser: to sql: %w", err)
	}
	return nil
}
