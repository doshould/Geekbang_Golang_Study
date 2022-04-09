package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hello/api/helloworld/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Id     int64
	Title  string
	Geners string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	CreateMovie(ctx context.Context, movie *Greeter) error
	GetMovie(ctx context.Context, id int64) (*Greeter, error)
	UpdateMovie(ctx context.Context, id int64, movie *Greeter) error
	ListMovie(ctx context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
//func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
//
//}

func (uc *GreeterUsecase) List(ctx context.Context) (ps []*Greeter, err error) {
	ps, err = uc.repo.ListMovie(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) Create(ctx context.Context, movie *Greeter) error {
	return uc.repo.CreateMovie(ctx, movie)
}

func (uc *GreeterUsecase) Get(ctx context.Context, id int64) (p *Greeter, err error) {
	p, err = uc.repo.GetMovie(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) Update(ctx context.Context, id int64, movie *Greeter) error {
	return uc.repo.UpdateMovie(ctx, id, movie)
}
