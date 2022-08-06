package api

import (
	"context"
	"errors"
	"fmt"

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
	fmt.Println(in.GetLimit(), in.GetOffset(), in.GetOrder())
	reviews, err := h.service.ListReviews(ctx, models.ReviewListInput{
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
	return nil, errors.New("unimplemented")
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
