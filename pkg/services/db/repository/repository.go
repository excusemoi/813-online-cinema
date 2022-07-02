package repository

import "github.com/jackc/pgx"

type Repository interface {
	//authorizeUser(login, password string) (bool, error)
	//getUserMovieList(id int64) ([]Movie, error)
	//getMovieStats(id int64) (MovieStats, error)
}

type PostgresRepository struct { //should implement interface above methods
	conn *pgx.Conn
}
