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
		Values(review.Rating, review.ReviewText, review.Time.Format(time.RFC3339), review.UserId, review.BookId).
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
	var review repoModels.Review
	err = pgxscan.ScanOne(&review, rows)
	if err != nil {
		return repoModels.Review{}, fmt.Errorf("Repository.GetReviewById: to sql: %w", err)
	}
	return review, nil
}

func (r *repository) ListReviews(ctx context.Context, params repoModels.ReviewListInput) ([]repoModels.Review, error) {
	preparedQuery := squirrel.Select("id", "rating", "review_text", "time", "user_id", "book_id").
		From("reviews").
		OrderBy("time " + params.Order).
		Offset(uint64(params.Offset))
	if params.Limit > 0 {
		preparedQuery = preparedQuery.Limit(uint64(params.Limit))
	}
	query, args, err := preparedQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return []repoModels.Review{}, fmt.Errorf("Repo")
	}
	var reviews []repoModels.Review
	if err := pgxscan.Select(ctx, r.pool, &reviews, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.GetUsersByNiknameOrEmail: select: %w", err)
	}
	return reviews, nil
}
