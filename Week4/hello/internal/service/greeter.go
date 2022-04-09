package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "hello/api/helloworld/v1"
	"hello/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	pb.UnimplementedGreeterServer
	movie *biz.GreeterUsecase
	log   *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(movie *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{
		movie: movie,
		log:   log.NewHelper(logger),
	}
}

func (s *GreeterService) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieReply, error) {
	s.log.Infof("input %v", req)
	err := s.movie.Create(ctx, &biz.Greeter{
		Title:  req.Title,
		Geners: req.Geners,
	})
	return &pb.CreateMovieReply{}, err
}

func (s *GreeterService) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieReply, error) {
	s.log.Infof("input %v", req)
	err := s.movie.Update(ctx, req.Id, &biz.Greeter{
		Title:  req.Title,
		Geners: req.Geners,
	})
	return &pb.UpdateMovieReply{}, err
}

func (s *GreeterService) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieReply, error) {
	pa, err := s.movie.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetMovieReply{Movie: &pb.Movie{Id: pa.Id, Title: pa.Title, Geners: pa.Geners}}, nil
}

func (s *GreeterService) ListMovie(ctx context.Context, _ *pb.ListMovieRequest) (*pb.ListMovieReply, error) {
	pa, err := s.movie.List(ctx)
	reply := &pb.ListMovieReply{}
	for _, p := range pa {
		reply.Result = append(reply.Result, &pb.Movie{
			Id:     p.Id,
			Title:  p.Title,
			Geners: p.Geners,
		})
	}
	return reply, err
}
