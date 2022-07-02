package repository

import (
	"813-online-cinema/pkg/services/db/config"
	pb "813-online-cinema/pkg/services/db/proto"
	"context"
)

type Repository interface {
	IsUserAuthorized(ctx context.Context, info *pb.UserAuthInfo) (*pb.UserAuthResponse, error)
	GetUserMovieList(ctx context.Context, id *pb.ID) (*pb.MovieList, error)
	GetMovieStats(ctx context.Context, id *pb.ID) (*pb.MovieStats, error)
	InitRepository(cfg *config.Config) error
}
