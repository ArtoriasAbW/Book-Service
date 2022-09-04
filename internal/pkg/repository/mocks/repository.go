// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	grpc "google.golang.org/grpc"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AuthorCreate mocks base method.
func (m *MockInterface) AuthorCreate(ctx context.Context, in *api.AuthorCreateRequest, opts ...grpc.CallOption) (*api.AuthorCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorCreate", varargs...)
	ret0, _ := ret[0].(*api.AuthorCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorCreate indicates an expected call of AuthorCreate.
func (mr *MockInterfaceMockRecorder) AuthorCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorCreate", reflect.TypeOf((*MockInterface)(nil).AuthorCreate), varargs...)
}

// AuthorDelete mocks base method.
func (m *MockInterface) AuthorDelete(ctx context.Context, in *api.AuthorDeleteRequest, opts ...grpc.CallOption) (*api.AuthorDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorDelete", varargs...)
	ret0, _ := ret[0].(*api.AuthorDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorDelete indicates an expected call of AuthorDelete.
func (mr *MockInterfaceMockRecorder) AuthorDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorDelete", reflect.TypeOf((*MockInterface)(nil).AuthorDelete), varargs...)
}

// AuthorGet mocks base method.
func (m *MockInterface) AuthorGet(ctx context.Context, in *api.AuthorGetRequest, opts ...grpc.CallOption) (*api.AuthorGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorGet", varargs...)
	ret0, _ := ret[0].(*api.AuthorGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorGet indicates an expected call of AuthorGet.
func (mr *MockInterfaceMockRecorder) AuthorGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorGet", reflect.TypeOf((*MockInterface)(nil).AuthorGet), varargs...)
}

// AuthorList mocks base method.
func (m *MockInterface) AuthorList(ctx context.Context, in *api.AuthorListRequest, opts ...grpc.CallOption) (*api.AuthorListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorList", varargs...)
	ret0, _ := ret[0].(*api.AuthorListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorList indicates an expected call of AuthorList.
func (mr *MockInterfaceMockRecorder) AuthorList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorList", reflect.TypeOf((*MockInterface)(nil).AuthorList), varargs...)
}

// AuthorUpdate mocks base method.
func (m *MockInterface) AuthorUpdate(ctx context.Context, in *api.AuthorUpdateRequest, opts ...grpc.CallOption) (*api.AuthorUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorUpdate", varargs...)
	ret0, _ := ret[0].(*api.AuthorUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorUpdate indicates an expected call of AuthorUpdate.
func (mr *MockInterfaceMockRecorder) AuthorUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorUpdate", reflect.TypeOf((*MockInterface)(nil).AuthorUpdate), varargs...)
}

// BookCreate mocks base method.
func (m *MockInterface) BookCreate(ctx context.Context, in *api.BookCreateRequest, opts ...grpc.CallOption) (*api.BookCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookCreate", varargs...)
	ret0, _ := ret[0].(*api.BookCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookCreate indicates an expected call of BookCreate.
func (mr *MockInterfaceMockRecorder) BookCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookCreate", reflect.TypeOf((*MockInterface)(nil).BookCreate), varargs...)
}

// BookDelete mocks base method.
func (m *MockInterface) BookDelete(ctx context.Context, in *api.BookDeleteRequest, opts ...grpc.CallOption) (*api.BookDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookDelete", varargs...)
	ret0, _ := ret[0].(*api.BookDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookDelete indicates an expected call of BookDelete.
func (mr *MockInterfaceMockRecorder) BookDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookDelete", reflect.TypeOf((*MockInterface)(nil).BookDelete), varargs...)
}

// BookGet mocks base method.
func (m *MockInterface) BookGet(ctx context.Context, in *api.BookGetRequest, opts ...grpc.CallOption) (*api.BookGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookGet", varargs...)
	ret0, _ := ret[0].(*api.BookGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookGet indicates an expected call of BookGet.
func (mr *MockInterfaceMockRecorder) BookGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookGet", reflect.TypeOf((*MockInterface)(nil).BookGet), varargs...)
}

// BookList mocks base method.
func (m *MockInterface) BookList(ctx context.Context, in *api.BookListRequest, opts ...grpc.CallOption) (*api.BookListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookList", varargs...)
	ret0, _ := ret[0].(*api.BookListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookList indicates an expected call of BookList.
func (mr *MockInterfaceMockRecorder) BookList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookList", reflect.TypeOf((*MockInterface)(nil).BookList), varargs...)
}

// BookUpdate mocks base method.
func (m *MockInterface) BookUpdate(ctx context.Context, in *api.BookUpdateRequest, opts ...grpc.CallOption) (*api.BookUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookUpdate", varargs...)
	ret0, _ := ret[0].(*api.BookUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookUpdate indicates an expected call of BookUpdate.
func (mr *MockInterfaceMockRecorder) BookUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookUpdate", reflect.TypeOf((*MockInterface)(nil).BookUpdate), varargs...)
}

// ReviewCreate mocks base method.
func (m *MockInterface) ReviewCreate(ctx context.Context, in *api.ReviewCreateRequest, opts ...grpc.CallOption) (*api.ReviewCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewCreate", varargs...)
	ret0, _ := ret[0].(*api.ReviewCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewCreate indicates an expected call of ReviewCreate.
func (mr *MockInterfaceMockRecorder) ReviewCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewCreate", reflect.TypeOf((*MockInterface)(nil).ReviewCreate), varargs...)
}

// ReviewDelete mocks base method.
func (m *MockInterface) ReviewDelete(ctx context.Context, in *api.ReviewDeleteRequest, opts ...grpc.CallOption) (*api.ReviewDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewDelete", varargs...)
	ret0, _ := ret[0].(*api.ReviewDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewDelete indicates an expected call of ReviewDelete.
func (mr *MockInterfaceMockRecorder) ReviewDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewDelete", reflect.TypeOf((*MockInterface)(nil).ReviewDelete), varargs...)
}

// ReviewGet mocks base method.
func (m *MockInterface) ReviewGet(ctx context.Context, in *api.ReviewGetRequest, opts ...grpc.CallOption) (*api.ReviewGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewGet", varargs...)
	ret0, _ := ret[0].(*api.ReviewGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewGet indicates an expected call of ReviewGet.
func (mr *MockInterfaceMockRecorder) ReviewGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewGet", reflect.TypeOf((*MockInterface)(nil).ReviewGet), varargs...)
}

// ReviewList mocks base method.
func (m *MockInterface) ReviewList(ctx context.Context, in *api.ReviewListRequest, opts ...grpc.CallOption) (*api.ReviewListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewList", varargs...)
	ret0, _ := ret[0].(*api.ReviewListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewList indicates an expected call of ReviewList.
func (mr *MockInterfaceMockRecorder) ReviewList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewList", reflect.TypeOf((*MockInterface)(nil).ReviewList), varargs...)
}

// ReviewUpdate mocks base method.
func (m *MockInterface) ReviewUpdate(ctx context.Context, in *api.ReviewUpdateRequest, opts ...grpc.CallOption) (*api.ReviewUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewUpdate", varargs...)
	ret0, _ := ret[0].(*api.ReviewUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewUpdate indicates an expected call of ReviewUpdate.
func (mr *MockInterfaceMockRecorder) ReviewUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewUpdate", reflect.TypeOf((*MockInterface)(nil).ReviewUpdate), varargs...)
}

// UserCreate mocks base method.
func (m *MockInterface) UserCreate(ctx context.Context, in *api.UserCreateRequest, opts ...grpc.CallOption) (*api.UserCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserCreate", varargs...)
	ret0, _ := ret[0].(*api.UserCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCreate indicates an expected call of UserCreate.
func (mr *MockInterfaceMockRecorder) UserCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*MockInterface)(nil).UserCreate), varargs...)
}

// UserDelete mocks base method.
func (m *MockInterface) UserDelete(ctx context.Context, in *api.UserDeleteRequest, opts ...grpc.CallOption) (*api.UserDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserDelete", varargs...)
	ret0, _ := ret[0].(*api.UserDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDelete indicates an expected call of UserDelete.
func (mr *MockInterfaceMockRecorder) UserDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDelete", reflect.TypeOf((*MockInterface)(nil).UserDelete), varargs...)
}

// UserGet mocks base method.
func (m *MockInterface) UserGet(ctx context.Context, in *api.UserGetRequest, opts ...grpc.CallOption) (*api.UserGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserGet", varargs...)
	ret0, _ := ret[0].(*api.UserGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGet indicates an expected call of UserGet.
func (mr *MockInterfaceMockRecorder) UserGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*MockInterface)(nil).UserGet), varargs...)
}

// UserList mocks base method.
func (m *MockInterface) UserList(ctx context.Context, in *api.UserListRequest, opts ...grpc.CallOption) (*api.UserListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserList", varargs...)
	ret0, _ := ret[0].(*api.UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserList indicates an expected call of UserList.
func (mr *MockInterfaceMockRecorder) UserList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*MockInterface)(nil).UserList), varargs...)
}

// UserUpdate mocks base method.
func (m *MockInterface) UserUpdate(ctx context.Context, in *api.UserUpdateRequest, opts ...grpc.CallOption) (*api.UserUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserUpdate", varargs...)
	ret0, _ := ret[0].(*api.UserUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserUpdate indicates an expected call of UserUpdate.
func (mr *MockInterfaceMockRecorder) UserUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdate", reflect.TypeOf((*MockInterface)(nil).UserUpdate), varargs...)
}

// Mockbook is a mock of book interface.
type Mockbook struct {
	ctrl     *gomock.Controller
	recorder *MockbookMockRecorder
}

// MockbookMockRecorder is the mock recorder for Mockbook.
type MockbookMockRecorder struct {
	mock *Mockbook
}

// NewMockbook creates a new mock instance.
func NewMockbook(ctrl *gomock.Controller) *Mockbook {
	mock := &Mockbook{ctrl: ctrl}
	mock.recorder = &MockbookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockbook) EXPECT() *MockbookMockRecorder {
	return m.recorder
}

// BookCreate mocks base method.
func (m *Mockbook) BookCreate(ctx context.Context, in *api.BookCreateRequest, opts ...grpc.CallOption) (*api.BookCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookCreate", varargs...)
	ret0, _ := ret[0].(*api.BookCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookCreate indicates an expected call of BookCreate.
func (mr *MockbookMockRecorder) BookCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookCreate", reflect.TypeOf((*Mockbook)(nil).BookCreate), varargs...)
}

// BookDelete mocks base method.
func (m *Mockbook) BookDelete(ctx context.Context, in *api.BookDeleteRequest, opts ...grpc.CallOption) (*api.BookDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookDelete", varargs...)
	ret0, _ := ret[0].(*api.BookDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookDelete indicates an expected call of BookDelete.
func (mr *MockbookMockRecorder) BookDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookDelete", reflect.TypeOf((*Mockbook)(nil).BookDelete), varargs...)
}

// BookGet mocks base method.
func (m *Mockbook) BookGet(ctx context.Context, in *api.BookGetRequest, opts ...grpc.CallOption) (*api.BookGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookGet", varargs...)
	ret0, _ := ret[0].(*api.BookGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookGet indicates an expected call of BookGet.
func (mr *MockbookMockRecorder) BookGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookGet", reflect.TypeOf((*Mockbook)(nil).BookGet), varargs...)
}

// BookList mocks base method.
func (m *Mockbook) BookList(ctx context.Context, in *api.BookListRequest, opts ...grpc.CallOption) (*api.BookListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookList", varargs...)
	ret0, _ := ret[0].(*api.BookListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookList indicates an expected call of BookList.
func (mr *MockbookMockRecorder) BookList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookList", reflect.TypeOf((*Mockbook)(nil).BookList), varargs...)
}

// BookUpdate mocks base method.
func (m *Mockbook) BookUpdate(ctx context.Context, in *api.BookUpdateRequest, opts ...grpc.CallOption) (*api.BookUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BookUpdate", varargs...)
	ret0, _ := ret[0].(*api.BookUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookUpdate indicates an expected call of BookUpdate.
func (mr *MockbookMockRecorder) BookUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookUpdate", reflect.TypeOf((*Mockbook)(nil).BookUpdate), varargs...)
}

// Mockuser is a mock of user interface.
type Mockuser struct {
	ctrl     *gomock.Controller
	recorder *MockuserMockRecorder
}

// MockuserMockRecorder is the mock recorder for Mockuser.
type MockuserMockRecorder struct {
	mock *Mockuser
}

// NewMockuser creates a new mock instance.
func NewMockuser(ctrl *gomock.Controller) *Mockuser {
	mock := &Mockuser{ctrl: ctrl}
	mock.recorder = &MockuserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockuser) EXPECT() *MockuserMockRecorder {
	return m.recorder
}

// UserCreate mocks base method.
func (m *Mockuser) UserCreate(ctx context.Context, in *api.UserCreateRequest, opts ...grpc.CallOption) (*api.UserCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserCreate", varargs...)
	ret0, _ := ret[0].(*api.UserCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCreate indicates an expected call of UserCreate.
func (mr *MockuserMockRecorder) UserCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*Mockuser)(nil).UserCreate), varargs...)
}

// UserDelete mocks base method.
func (m *Mockuser) UserDelete(ctx context.Context, in *api.UserDeleteRequest, opts ...grpc.CallOption) (*api.UserDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserDelete", varargs...)
	ret0, _ := ret[0].(*api.UserDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDelete indicates an expected call of UserDelete.
func (mr *MockuserMockRecorder) UserDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDelete", reflect.TypeOf((*Mockuser)(nil).UserDelete), varargs...)
}

// UserGet mocks base method.
func (m *Mockuser) UserGet(ctx context.Context, in *api.UserGetRequest, opts ...grpc.CallOption) (*api.UserGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserGet", varargs...)
	ret0, _ := ret[0].(*api.UserGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGet indicates an expected call of UserGet.
func (mr *MockuserMockRecorder) UserGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*Mockuser)(nil).UserGet), varargs...)
}

// UserList mocks base method.
func (m *Mockuser) UserList(ctx context.Context, in *api.UserListRequest, opts ...grpc.CallOption) (*api.UserListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserList", varargs...)
	ret0, _ := ret[0].(*api.UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserList indicates an expected call of UserList.
func (mr *MockuserMockRecorder) UserList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*Mockuser)(nil).UserList), varargs...)
}

// UserUpdate mocks base method.
func (m *Mockuser) UserUpdate(ctx context.Context, in *api.UserUpdateRequest, opts ...grpc.CallOption) (*api.UserUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserUpdate", varargs...)
	ret0, _ := ret[0].(*api.UserUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserUpdate indicates an expected call of UserUpdate.
func (mr *MockuserMockRecorder) UserUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdate", reflect.TypeOf((*Mockuser)(nil).UserUpdate), varargs...)
}

// Mockauthor is a mock of author interface.
type Mockauthor struct {
	ctrl     *gomock.Controller
	recorder *MockauthorMockRecorder
}

// MockauthorMockRecorder is the mock recorder for Mockauthor.
type MockauthorMockRecorder struct {
	mock *Mockauthor
}

// NewMockauthor creates a new mock instance.
func NewMockauthor(ctrl *gomock.Controller) *Mockauthor {
	mock := &Mockauthor{ctrl: ctrl}
	mock.recorder = &MockauthorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockauthor) EXPECT() *MockauthorMockRecorder {
	return m.recorder
}

// AuthorCreate mocks base method.
func (m *Mockauthor) AuthorCreate(ctx context.Context, in *api.AuthorCreateRequest, opts ...grpc.CallOption) (*api.AuthorCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorCreate", varargs...)
	ret0, _ := ret[0].(*api.AuthorCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorCreate indicates an expected call of AuthorCreate.
func (mr *MockauthorMockRecorder) AuthorCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorCreate", reflect.TypeOf((*Mockauthor)(nil).AuthorCreate), varargs...)
}

// AuthorDelete mocks base method.
func (m *Mockauthor) AuthorDelete(ctx context.Context, in *api.AuthorDeleteRequest, opts ...grpc.CallOption) (*api.AuthorDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorDelete", varargs...)
	ret0, _ := ret[0].(*api.AuthorDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorDelete indicates an expected call of AuthorDelete.
func (mr *MockauthorMockRecorder) AuthorDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorDelete", reflect.TypeOf((*Mockauthor)(nil).AuthorDelete), varargs...)
}

// AuthorGet mocks base method.
func (m *Mockauthor) AuthorGet(ctx context.Context, in *api.AuthorGetRequest, opts ...grpc.CallOption) (*api.AuthorGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorGet", varargs...)
	ret0, _ := ret[0].(*api.AuthorGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorGet indicates an expected call of AuthorGet.
func (mr *MockauthorMockRecorder) AuthorGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorGet", reflect.TypeOf((*Mockauthor)(nil).AuthorGet), varargs...)
}

// AuthorList mocks base method.
func (m *Mockauthor) AuthorList(ctx context.Context, in *api.AuthorListRequest, opts ...grpc.CallOption) (*api.AuthorListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorList", varargs...)
	ret0, _ := ret[0].(*api.AuthorListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorList indicates an expected call of AuthorList.
func (mr *MockauthorMockRecorder) AuthorList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorList", reflect.TypeOf((*Mockauthor)(nil).AuthorList), varargs...)
}

// AuthorUpdate mocks base method.
func (m *Mockauthor) AuthorUpdate(ctx context.Context, in *api.AuthorUpdateRequest, opts ...grpc.CallOption) (*api.AuthorUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorUpdate", varargs...)
	ret0, _ := ret[0].(*api.AuthorUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorUpdate indicates an expected call of AuthorUpdate.
func (mr *MockauthorMockRecorder) AuthorUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorUpdate", reflect.TypeOf((*Mockauthor)(nil).AuthorUpdate), varargs...)
}

// Mockreview is a mock of review interface.
type Mockreview struct {
	ctrl     *gomock.Controller
	recorder *MockreviewMockRecorder
}

// MockreviewMockRecorder is the mock recorder for Mockreview.
type MockreviewMockRecorder struct {
	mock *Mockreview
}

// NewMockreview creates a new mock instance.
func NewMockreview(ctrl *gomock.Controller) *Mockreview {
	mock := &Mockreview{ctrl: ctrl}
	mock.recorder = &MockreviewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockreview) EXPECT() *MockreviewMockRecorder {
	return m.recorder
}

// ReviewCreate mocks base method.
func (m *Mockreview) ReviewCreate(ctx context.Context, in *api.ReviewCreateRequest, opts ...grpc.CallOption) (*api.ReviewCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewCreate", varargs...)
	ret0, _ := ret[0].(*api.ReviewCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewCreate indicates an expected call of ReviewCreate.
func (mr *MockreviewMockRecorder) ReviewCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewCreate", reflect.TypeOf((*Mockreview)(nil).ReviewCreate), varargs...)
}

// ReviewDelete mocks base method.
func (m *Mockreview) ReviewDelete(ctx context.Context, in *api.ReviewDeleteRequest, opts ...grpc.CallOption) (*api.ReviewDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewDelete", varargs...)
	ret0, _ := ret[0].(*api.ReviewDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewDelete indicates an expected call of ReviewDelete.
func (mr *MockreviewMockRecorder) ReviewDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewDelete", reflect.TypeOf((*Mockreview)(nil).ReviewDelete), varargs...)
}

// ReviewGet mocks base method.
func (m *Mockreview) ReviewGet(ctx context.Context, in *api.ReviewGetRequest, opts ...grpc.CallOption) (*api.ReviewGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewGet", varargs...)
	ret0, _ := ret[0].(*api.ReviewGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewGet indicates an expected call of ReviewGet.
func (mr *MockreviewMockRecorder) ReviewGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewGet", reflect.TypeOf((*Mockreview)(nil).ReviewGet), varargs...)
}

// ReviewList mocks base method.
func (m *Mockreview) ReviewList(ctx context.Context, in *api.ReviewListRequest, opts ...grpc.CallOption) (*api.ReviewListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewList", varargs...)
	ret0, _ := ret[0].(*api.ReviewListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewList indicates an expected call of ReviewList.
func (mr *MockreviewMockRecorder) ReviewList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewList", reflect.TypeOf((*Mockreview)(nil).ReviewList), varargs...)
}

// ReviewUpdate mocks base method.
func (m *Mockreview) ReviewUpdate(ctx context.Context, in *api.ReviewUpdateRequest, opts ...grpc.CallOption) (*api.ReviewUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewUpdate", varargs...)
	ret0, _ := ret[0].(*api.ReviewUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewUpdate indicates an expected call of ReviewUpdate.
func (mr *MockreviewMockRecorder) ReviewUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewUpdate", reflect.TypeOf((*Mockreview)(nil).ReviewUpdate), varargs...)
}
