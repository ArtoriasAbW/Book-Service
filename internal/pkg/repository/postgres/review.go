package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddReview(ctx context.Context, review repoModels.Review) (uint64, error) {
	query, args, err := squirrel.Insert("reviews").
		Columns("rating", "review_text", "time", "user_id", "book_id").
		Values(review.Rating, review.ReviewText, review.Time, review.UserId, review.BookId).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Repository.AddReview: to sql: %w", err)
	}
	row := r.pool.QueryRow(ctx, query, args...)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Repository.AddReview: to sql: %w", err)
	}
	return id, nil
}

func (r *repository) GetReviewById(ctx context.Context, id uint) (repoModels.Review, error) {
	query, args, err := squirrel.Select("id", "rating", "review_text", "time", "user_id", "book_id").
		From("reviews").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return repoModels.Review{}, fmt.Errorf("Repository.GetReviewById: to sql: %w", err)
	}
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return repoModels.Review{}, fmt.Errorf("Repository.GetReviewById: to sql: %w", err)
	}
	var data struct {
		Id         uint      `db:"id"`
		Rating     uint      `db:"rating"`
		ReviewText string    `db:"review_text"`
		Time       time.Time `db:"time"`
		BookId     uint      `db:"book_id"`
		UserId     uint      `db:"user_id"`
	}
	err = pgxscan.ScanOne(&data, rows)
	if err != nil {
		return repoModels.Review{}, fmt.Errorf("Repository.GetReviewById: to sql: %w", err)
	}
	return repoModels.Review{
		Id:         data.Id,
		Rating:     data.Rating,
		ReviewText: data.ReviewText,
		Time:       data.Time.Format(time.RFC3339),
		BookId:     data.BookId,
		UserId:     data.UserId,
	}, nil
}

func (r *repository) ListReviews(ctx context.Context) ([]repoModels.Review, error) {
	query, args, err := squirrel.Select("*").
		From("reviews").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return []repoModels.Review{}, fmt.Errorf("Repo")
	}
	var reviews []repoModels.Review
	if err := pgxscan.Select(ctx, r.pool, &reviews, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.GetUsersByNiknameOrEmail: select: %w", err)
	}
	return reviews, nil
}
