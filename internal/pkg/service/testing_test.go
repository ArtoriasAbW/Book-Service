package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/mocks"
)

type BaseFixture struct {
	Ctx context.Context
}

type serviceFixture struct {
	BaseFixture
	service *Implementation
	repo    *mock_repository.MockInterface
}

func serviceSetUp(t *testing.T) serviceFixture {
	base := BaseFixture{Ctx: context.Background()}
	f := serviceFixture{BaseFixture: base}
	f.repo = mock_repository.NewMockInterface(gomock.NewController(t))
	deps := Deps{Repository: f.repo}
	f.service = New(deps)
	return f
}
