package api

import (
	"context"
	"time"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ReviewGet(ctx context.Context, in *pb.ReviewGetRequest) (*pb.ReviewGetResponse, error) {
	review, err := h.repo.GetReviewByID(ctx, uint(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewGetResponse{
		Id:         uint64(review.ID),
		Rating:     uint32(review.Rating),
		ReviewText: review.ReviewText,
		Time:       uint64(review.Time.Unix()),
		BookId:     uint64(review.BookID),
		UserId:     uint64(review.UserID),
	}, nil
}
func (h *handler) ReviewCreate(ctx context.Context, in *pb.ReviewCreateRequest) (*pb.ReviewCreateResponse, error) {
	id, err := h.repo.AddReview(ctx, models.Review{
		Rating:     uint(in.GetRating()),
		ReviewText: in.GetReviewText(),
		Time:       time.Now(),
		BookID:     uint(in.GetBookId()),
		UserID:     uint(in.GetUserId()),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewCreateResponse{
		Id: id,
	}, nil
}

func (h *handler) ReviewList(ctx context.Context, in *pb.ReviewListRequest) (*pb.ReviewListResponse, error) {
	reviews, err := h.repo.ListReviews(ctx, models.ListInput{
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
		Order:  in.GetOrder(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.ReviewListResponse_Review, 0, len(reviews))
	for _, review := range reviews {
		result = append(result, &pb.ReviewListResponse_Review{
			Id:         uint64(review.ID),
			Rating:     uint32(review.Rating),
			ReviewText: review.ReviewText,
			Time:       uint64(review.Time.Unix()),
			BookId:     uint64(review.BookID),
			UserId:     uint64(review.UserID),
		})
	}
	return &pb.ReviewListResponse{
		Reviews: result,
	}, nil
}

func (h *handler) ReviewUpdate(ctx context.Context, in *pb.ReviewUpdateRequest) (*pb.ReviewUpdateResponse, error) {
	if err := h.repo.UpdateReview(ctx, models.Review{
		ID:         uint(in.GetId()),
		Rating:     uint(in.GetRating()),
		ReviewText: in.GetReviewText(),
		Time:       time.Now(),
		BookID:     uint(in.GetBookId()),
		UserID:     uint(in.GetUserId()),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewUpdateResponse{}, nil

}
func (h *handler) ReviewDelete(ctx context.Context, in *pb.ReviewDeleteRequest) (*pb.ReviewDeleteResponse, error) {
	err := h.repo.DeleteReview(ctx, uint(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReviewDeleteResponse{}, nil
}
