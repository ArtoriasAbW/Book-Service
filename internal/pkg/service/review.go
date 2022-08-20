package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func (s *Implementation) GetReview(ctx context.Context, id uint) (models.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	reviewResponce, err := s.Repository.ReviewGet(ctx, &pb.ReviewGetRequest{
		Id: uint64(id),
	})
	if err != nil {
		return models.Review{}, err
	}
	return models.Review{
		Id:         uint(reviewResponce.GetId()),
		Rating:     uint(reviewResponce.GetRating()),
		ReviewText: reviewResponce.GetReviewText(),
		Time:       uint(reviewResponce.GetTime()),
		BookId:     uint(reviewResponce.GetBookId()),
		UserId:     uint(reviewResponce.GetUserId()),
	}, err
}

func (s *Implementation) AddReview(ctx context.Context, reviewInput models.Review) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.UserGet(ctx, &pb.UserGetRequest{
		Id: uint64(reviewInput.UserId),
	})
	if err != nil {
		return 0, errors.Wrap(ErrValidation, "user with this id doesn't exist")
	}
	_, err = s.Repository.BookGet(ctx, &pb.BookGetRequest{
		Id: uint64(reviewInput.BookId),
	})
	if err != nil {
		return 0, errors.Wrap(ErrValidation, "book with this id doesn't exist")
	}
	if reviewInput.Rating > 10 {
		return 0, errors.Wrap(ErrValidation, "rating must be less than 10")
	}
	reviewResponce, err := s.Repository.ReviewCreate(ctx, &pb.ReviewCreateRequest{
		Rating:     uint32(reviewInput.Rating),
		ReviewText: reviewInput.ReviewText,
		BookId:     uint64(reviewInput.BookId),
		UserId:     uint64(reviewInput.UserId),
	})
	if err != nil {
		return 0, err
	}
	return reviewResponce.GetId(), err
}

func (s *Implementation) ListReviews(ctx context.Context, params models.ListInput) ([]models.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if strings.ToLower(params.Order) == "desc" {
		params.Order = "DESC"
	} else {
		params.Order = "ASC"
	}
	reviewsResponce, err := s.Repository.ReviewList(ctx, &pb.ReviewListRequest{
		Offset: params.Offset,
		Limit:  params.Limit,
		Order:  params.Order,
	})
	if err != nil {
		return nil, err
	}
	reviewsData := reviewsResponce.Reviews
	reviews := make([]models.Review, len(reviewsData))
	for i, review := range reviewsData {
		reviews[i] = models.Review{
			Id:         uint(review.GetId()),
			Rating:     uint(review.GetRating()),
			ReviewText: review.GetReviewText(),
			Time:       uint(review.GetTime()),
			BookId:     uint(review.GetBookId()),
			UserId:     uint(review.GetUserId()),
		}
	}
	return reviews, nil
}

func (s *Implementation) DeleteReview(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.ReviewDelete(ctx, &pb.ReviewDeleteRequest{
		Id: uint64(id),
	})
	return err
}

func (s *Implementation) UpdateReview(ctx context.Context, reviewInput models.Review) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.ReviewGet(ctx, &pb.ReviewGetRequest{
		Id: uint64(reviewInput.Id),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "review with this id doesn't exist")
	}
	_, err = s.Repository.UserGet(ctx, &pb.UserGetRequest{
		Id: uint64(reviewInput.UserId),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "user with this id doesn't exist")
	}
	_, err = s.Repository.BookGet(ctx, &pb.BookGetRequest{
		Id: uint64(reviewInput.BookId),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "book with this id doesn't exist")
	}
	if reviewInput.Rating > 10 {
		return errors.Wrap(ErrValidation, "rating must be less than 10")
	}
	_, err = s.Repository.ReviewUpdate(ctx, &pb.ReviewUpdateRequest{
		Id:         uint64(reviewInput.Id),
		Rating:     uint32(reviewInput.Rating),
		ReviewText: reviewInput.ReviewText,
		UserId:     uint64(reviewInput.UserId),
		BookId:     uint64(reviewInput.BookId),
	})
	return err
}
