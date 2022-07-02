package postgres

import (
	"813-online-cinema/pkg/services/db/config"
	pb "813-online-cinema/pkg/services/db/proto"
	"context"
	"errors"
	"github.com/jackc/pgx"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{conn: conn}
}

func (r *Repository) InitRepository(cfg *config.Config) error {
	var err error
	r.conn, err = pgx.Connect(pgx.ConnConfig{
		Host:     cfg.DbHost,
		Port:     cfg.DBPort,
		Database: cfg.DbName,
		User:     cfg.DbUser,
		Password: cfg.DbPassword,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) IsUserAuthorized(ctx context.Context, info *pb.UserAuthInfo) (*pb.UserAuthResponse, error) {
	if r.conn == nil {
		return nil, errors.New("connection with db was not established") //TODO:cfg
	}
	return nil, nil
}
func (r *Repository) GetUserMovieList(ctx context.Context, id *pb.ID) (*pb.MovieList, error) {
	return nil, nil
}
func (r *Repository) GetMovieStats(ctx context.Context, id *pb.ID) (*pb.MovieStats, error) {
	return nil, nil
}
