package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"hello/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewgreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (mv *greeterRepo) ListMovie(ctx context.Context) ([]*biz.Greeter, error) {
	pa, err := mv.data.db.Movie.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	li := make([]*biz.Greeter, 0)
	for _, p := range pa {
		li = append(li, &biz.Greeter{
			Id:     p.ID,
			Title:  p.Title,
			Geners: p.Geners,
		})
	}
	return li, nil
}

func (mv *greeterRepo) GetMovie(ctx context.Context, id int64) (*biz.Greeter, error) {
	pa, err := mv.data.db.Movie.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Greeter{
		Id:     pa.ID,
		Title:  pa.Title,
		Geners: pa.Geners,
	}, nil
}

func (mv *greeterRepo) CreateMovie(ctx context.Context, movie *biz.Greeter) error {
	_, err := mv.data.db.Movie.
		Create().
		SetTitle(movie.Title).
		SetGeners(movie.Geners).
		Save(ctx)
	return err
}

func (mv *greeterRepo) UpdateMovie(ctx context.Context, id int64, movie *biz.Greeter) error {
	pa, err := mv.data.db.Movie.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = pa.Update().
		SetTitle(movie.Title).
		SetGeners(movie.Geners).
		Save(ctx)
	return err
}
