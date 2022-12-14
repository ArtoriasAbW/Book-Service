package api

import (
	"context"
	"errors"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ReviewGet(ctx context.Context, in *pb.ReviewGetRequest) (*pb.ReviewGetResponse, error) {
	review, err := h.service.GetReview(ctx, uint(in.GetId()))
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewGetResponse{
		Id:         uint64(review.Id),
		Rating:     uint32(review.Rating),
		ReviewText: review.ReviewText,
		Time:       uint64(review.Time),
		BookId:     uint64(review.BookId),
		UserId:     uint64(review.UserId),
	}, nil
}
func (h *handler) ReviewCreate(ctx context.Context, in *pb.ReviewCreateRequest) (*pb.ReviewCreateResponse, error) {
	id, err := h.service.AddReview(ctx, models.Review{
		Rating:     uint(in.GetRating()),
		ReviewText: in.GetReviewText(),
		BookId:     uint(in.GetBookId()),
		UserId:     uint(in.GetUserId()),
	})
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewCreateResponse{
		Id: id,
	}, nil
}

func (h *handler) ReviewList(ctx context.Context, in *pb.ReviewListRequest) (*pb.ReviewListResponse, error) {
	reviews, err := h.service.ListReviews(ctx, models.ListInput{
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
		Order:  in.GetOrder(),
	})
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.ReviewListResponse_Review, 0, len(reviews))
	for _, review := range reviews {
		result = append(result, &pb.ReviewListResponse_Review{
			Id:         uint64(review.Id),
			Rating:     uint32(review.Rating),
			ReviewText: review.ReviewText,
			Time:       uint64(review.Time),
			BookId:     uint64(review.BookId),
			UserId:     uint64(review.UserId),
		})
	}
	return &pb.ReviewListResponse{
		Reviews: result,
	}, nil
}

func (h *handler) ReviewUpdate(ctx context.Context, in *pb.ReviewUpdateRequest) (*pb.ReviewUpdateResponse, error) {
	if err := h.service.UpdateReview(ctx, models.Review{
		Id:         uint(in.GetId()),
		Rating:     uint(in.GetRating()),
		ReviewText: in.GetReviewText(),
		BookId:     uint(in.GetBookId()),
		UserId:     uint(in.GetUserId()),
	}); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewUpdateResponse{}, nil

}
func (h *handler) ReviewDelete(ctx context.Context, in *pb.ReviewDeleteRequest) (*pb.ReviewDeleteResponse, error) {
	err := h.service.DeleteReview(ctx, uint(in.GetId()))
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewDeleteResponse{}, nil
}
