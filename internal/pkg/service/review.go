package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func (s *service) GetReview(ctx context.Context, id uint) (models.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	review, err := s.Repository.GetReviewById(ctx, id)
	if err != nil {
		return models.Review{}, err
	}
	return models.Review{
		Id:         review.Id,
		Rating:     review.Rating,
		ReviewText: review.ReviewText,
		Time:       uint(review.Time.Unix()),
		BookId:     review.BookId,
		UserId:     review.UserId,
	}, err
}

func (s *service) AddReview(ctx context.Context, reviewInput models.Review) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var review repoModels.Review
	_, err := s.Repository.GetUserById(ctx, reviewInput.UserId)
	if err != nil {
		return 0, errors.Wrap(ErrValidation, "user with this id doesn't exist")
	}
	review.UserId = reviewInput.UserId
	_, err = s.Repository.GetBookById(ctx, reviewInput.BookId)
	if err != nil {
		return 0, errors.Wrap(ErrValidation, "book with this id doesn't exist")
	}
	review.BookId = reviewInput.BookId
	if reviewInput.Rating > 10 {
		return 0, errors.Wrap(ErrValidation, "rating must be less than 10")
	}
	review.Rating = reviewInput.Rating
	review.Time = time.Now()
	review.ReviewText = reviewInput.ReviewText
	id, err := s.Repository.AddReview(ctx, review)
	return id, err
}

func (s *service) ListReviews(ctx context.Context, params models.ListInput) ([]models.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var repoParams repoModels.ListInput
	repoParams.Offset = params.Offset
	repoParams.Limit = params.Limit
	if strings.ToLower(params.Order) == "desc" {
		repoParams.Order = "DESC"
	} else {
		repoParams.Order = "ASC"
	}
	reviewsRepo, err := s.Repository.ListReviews(ctx, repoParams)
	if err != nil {
		return nil, err
	}
	reviews := make([]models.Review, len(reviewsRepo))
	for i, review := range reviewsRepo {
		reviews[i] = models.Review{
			Id:         review.Id,
			Rating:     review.Rating,
			ReviewText: review.ReviewText,
			Time:       uint(review.Time.Unix()),
			BookId:     review.BookId,
			UserId:     review.UserId,
		}
	}
	return reviews, nil
}

func (s *service) DeleteReview(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	return s.Repository.DeleteReview(ctx, id)
}

func (s *service) UpdateReview(ctx context.Context, reviewInput models.Review) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var review repoModels.Review
	_, err := s.Repository.GetReviewById(ctx, reviewInput.Id)
	if err != nil {
		return errors.Wrap(ErrValidation, "review with this id doesn't exist")
	}
	review.Id = reviewInput.Id
	_, err = s.Repository.GetUserById(ctx, reviewInput.UserId)
	if err != nil {
		return errors.Wrap(ErrValidation, "user with this id doesn't exist")
	}
	review.UserId = reviewInput.UserId
	_, err = s.Repository.GetBookById(ctx, reviewInput.BookId)
	if err != nil {
		return errors.Wrap(ErrValidation, "book with this id doesn't exist")
	}
	review.BookId = reviewInput.BookId
	if reviewInput.Rating > 10 {
		return errors.Wrap(ErrValidation, "rating must be less than 10")
	}
	review.Rating = reviewInput.Rating
	review.Time = time.Now()
	review.ReviewText = reviewInput.ReviewText
	err = s.Repository.UpdateReview(ctx, review)
	return err
}
