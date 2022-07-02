package postgres

import (
	"813-online-cinema/pkg/services/db/config"
	pb "813-online-cinema/pkg/services/db/proto"
	"context"
	"encoding/base64"
	"errors"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/jackc/pgx"
	"log"
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

/*func (r *Repository) IsUserAuthorized(ctx context.Context, info *pb.UserAuthInfo) (*pb.UserAuthResponse, error) {
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
}*/

func (r *Repository) selectFromDatabase(context context.Context, request *pb.SelectRequest) (*pb.SelectReply, error) {
	result_reply := pb.ResultReply{Result: true, Info: "Entry successfully found."}
	var any_messages []*any.Any
	var rows *pgx.Rows
	var err error

	switch {
	case request.SelectFrom == pb.Table_PROFILE:
		switch {
		case request.SelectBy == pb.Field_ID:
			rows, err = r.conn.Query(context.Background(), "select * from profile where id = $1", request.Integer)
		case request.SelectBy == pb.Field_NAME:
			rows, err = r.conn.Query(context.Background(), "select * from profile where login = $1", request.Str)
		case request.SelectBy == pb.Field_SURNAME:
			rows, err = r.conn.Query(context.Background(), "select * from profile where login = $1", request.Str)
		case request.SelectBy == pb.Field_LOGIN:
			rows, err = r.conn.Query(context.Background(), "select * from profile where login = $1", request.Str)
		case request.SelectBy == pb.Field_HASHED_PASSWORD:
			rows, err = r.conn.Query(context.Background(), "select * from profile where password = $1", base64.StdEncoding.EncodeToString(request.Bytes))
		}
		if rows != nil {
			for rows.Next() {
				var profile pb.Profile
				var any_message any.Any

				err = rows.Scan(&profile.Id, &profile.Name, &profile.Surname, &profile.Login, &profile.HashedPassword)
				if err != nil {
					log.Fatal(err)
				}
				any_message.MarshalFrom(&profile)
				any_messages = append(any_messages, &any_message)
			}
		}
	case request.SelectFrom == pb.Table_PROFILE_STATUS:
		switch {
		case request.SelectBy == pb.Field_ID:
			rows, err = r.conn.Query(context.Background(), "select * from profile_status where id = $1", request.Integer)
		case request.SelectBy == pb.Field_NAME:
			rows, err = r.conn.Query(context.Background(), "select * from profile_status where name = $1", request.Str)
		}
		if rows != nil {
			for rows.Next() {
				var profile_status pb.ProfileStatus
				var any_message any.Any

				err = rows.Scan(&profile_status.Id, &profile_status.Name)
				if err != nil {
					log.Fatal(err)
				}
				any_message.MarshalFrom(&profile_status)
				any_messages = append(any_messages, &any_message)
			}
		}
	case request.SelectFrom == pb.Table_MOVIE:
		switch {
		case request.SelectBy == pb.Field_ID:
			rows, err = r.conn.Query(context.Background(), "select * from movie where id = $1", request.Integer)
		case request.SelectBy == pb.Field_MAGNET_LINK:
			rows, err = r.conn.Query(context.Background(), "select * from movie where magnet_link = $1", request.Str)
		case request.SelectBy == pb.Field_VIEWS_COUNT:
			rows, err = r.conn.Query(context.Background(), "select * from movie where views_count = $1", request.Integer)
		case request.SelectBy == pb.Field_LIKES_COUNT:
			rows, err = r.conn.Query(context.Background(), "select * from movie where likes_count = $1", request.Integer)
		case request.SelectBy == pb.Field_DISLIKES_COUNT:
			rows, err = r.conn.Query(context.Background(), "select * from movie where dislikes_count = $1", request.Integer)
		}
		if rows != nil {
			for rows.Next() {
				var movie pb.Movie
				var any_message any.Any

				err = rows.Scan(&movie.Id, &movie.MagnetLink, &movie.ViewsCount, &movie.LikesCount, &movie.DislikesCount)
				if err != nil {
					log.Fatal(err)
				}
				any_message.MarshalFrom(&movie)
				any_messages = append(any_messages, &any_message)
			}
		}
	case request.SelectFrom == pb.Table_PROFILE_MOVIE:
		switch {
		case request.SelectBy == pb.Field_ID:
			rows, err = r.conn.Query(context.Background(), "select * from profile_movie where id = $1", request.Integer)
		case request.SelectBy == pb.Field_PROFILE_ID:
			rows, err = r.conn.Query(context.Background(), "select * from profile_movie where profile_id = $1", request.Integer)
		case request.SelectBy == pb.Field_MOVIE_ID:
			rows, err = r.conn.Query(context.Background(), "select * from profile_movie where movie_id = $1", request.Integer)
		}
		if rows != nil {
			for rows.Next() {
				var profile_movie pb.Profile_Movie
				var any_message any.Any

				err = rows.Scan(&profile_movie.Id)
				if err != nil {
					log.Fatal(err)
				}
				any_message.MarshalFrom(&profile_movie)
				any_messages = append(any_messages, &any_message)
			}
		}
	}

	if rows == nil {
		result_reply.Result = false
		result_reply.Info = "ERROR: Entry not found."
	}
	return &pb.SelectReply{
		Result: &result_reply,
		Entry:  any_messages,
	}, err
}
func (r *Repository) deleteFromDatabase(context context.Context, request *pb.DeleteRequest) (*pb.ResultReply, error) {
	result_reply := pb.ResultReply{Result: true, Info: "Entry successfully deleted."}
	var row *pgx.Row
	var err error

	switch {
	case request.DeleteFrom == pb.Table_PROFILE:
		switch {
		case request.DeleteBy == pb.Field_ID:
			row = r.conn.QueryRow(context.Background(), "delete * from profile where id = $1", request.Integer)
		case request.DeleteBy == pb.Field_NAME:
			row = r.conn.QueryRow(context.Background(), "delete * from profile where login = $1", request.Str)
		case request.DeleteBy == pb.Field_SURNAME:
			row = r.conn.QueryRow(context.Background(), "delete * from profile where login = $1", request.Str)
		case request.DeleteBy == pb.Field_LOGIN:
			row = r.conn.QueryRow(context.Background(), "delete * from profile where login = $1", request.Str)
		case request.DeleteBy == pb.Field_HASHED_PASSWORD:
			row = r.conn.QueryRow(context.Background(), "delete * from profile where password = $1", base64.StdEncoding.EncodeToString(request.Bytes))
		}
	case request.DeleteFrom == pb.Table_PROFILE_STATUS:
		switch {
		case request.DeleteBy == pb.Field_ID:
			row = r.conn.QueryRow(context.Background(), "select * from profile_status where id = $1", request.Integer)
		case request.DeleteBy == pb.Field_NAME:
			row = r.conn.QueryRow(context.Background(), "select * from profile_status where name = $1", request.Str)
		}
	case request.DeleteFrom == pb.Table_MOVIE:
		switch {
		case request.DeleteBy == pb.Field_ID:
			row = r.conn.QueryRow(context.Background(), "select * from movie where id = $1", request.Integer)
		case request.DeleteBy == pb.Field_MAGNET_LINK:
			row = r.conn.QueryRow(context.Background(), "select * from movie where magnet_link = $1", request.Str)
		case request.DeleteBy == pb.Field_VIEWS_COUNT:
			row = r.conn.QueryRow(context.Background(), "select * from movie where views_count = $1", request.Integer)
		case request.DeleteBy == pb.Field_LIKES_COUNT:
			row = r.conn.QueryRow(context.Background(), "select * from movie where likes_count = $1", request.Integer)
		case request.DeleteBy == pb.Field_DISLIKES_COUNT:
			row = r.conn.QueryRow(context.Background(), "select * from movie where dislikes_count = $1", request.Integer)
		}
	case request.DeleteFrom == pb.Table_PROFILE_MOVIE:
		switch {
		case request.DeleteBy == pb.Field_ID:
			row = r.conn.QueryRow(context.Background(), "select * from profile_movie where id = $1", request.Integer)
		case request.DeleteBy == pb.Field_PROFILE_ID:
			row = r.conn.QueryRow(context.Background(), "select * from profile_movie where profile_id = $1", request.Integer)
		case request.DeleteBy == pb.Field_MOVIE_ID:
			row = r.conn.QueryRow(context.Background(), "select * from profile_movie where movie_id = $1", request.Integer)
		}
	}

	if row == nil {
		result_reply.Result = false
		result_reply.Info = "ERROR: Entry not deleted."
	}
	return &result_reply, err
}
func (r *Repository) updateInDatabase(context context.Context, request *pb.UpdateRequest) (*pb.ResultReply, error) {
	result_reply := pb.ResultReply{Result: true, Info: "Entry successfully deleted."}
	var row *pgx.Row
	var err error

	switch {
	case request.UpdateIn == pb.Table_PROFILE:
		var profile pb.Profile
		err = request.Entry.UnmarshalTo(&profile)
		if err != nil {
			return nil, err
		}
		row = r.conn.QueryRow(context.Background(), "update profile set name=$2, surname=$3, login=$4, hashed_password=$5 where id = $1", profile.Id, profile.Name, profile.Surname, profile.Login, profile.HashedPassword)
	case request.UpdateIn == pb.Table_PROFILE_STATUS:
		var profile_status pb.ProfileStatus
		err = request.Entry.UnmarshalTo(&profile_status)
		if err != nil {
			return nil, err
		}
		row = r.conn.QueryRow(context.Background(), "update profile_status set name=$2 where id = $1", profile_status.Id, profile_status.Name)
	case request.UpdateIn == pb.Table_MOVIE:
		var movie pb.Movie
		err = request.Entry.UnmarshalTo(&movie)
		if err != nil {
			return nil, err
		}
		row = r.conn.QueryRow(context.Background(), "update movie set views_count=$2, likes_count=$3, dislikes_count=$4 where id = $1", movie.Id, movie.ViewsCount, movie.LikesCount, movie.DislikesCount)
	}

	if row == nil {
		result_reply.Result = false
		result_reply.Info = "ERROR: Entry not updated."
	}
	return &result_reply, err
}
func (r *Repository) insertToDatabase(context context.Context, request *pb.InsertRequest) (*pb.ResultReply, error) {
	result_reply := pb.ResultReply{Result: true, Info: "Entry successfully deleted."}
	var row *pgx.Row
	var err error

	switch {
	case request.InsertTo == pb.Table_PROFILE:
		var profile pb.Profile
		err = request.Entry.UnmarshalTo(&profile)
		if err != nil {
			return nil, err
		}
		row = r.conn.QueryRow(context.Background(), "insert into profile values ($1, $2, $3, $4, $5)", profile.Id, profile.Name, profile.Surname, profile.Login, profile.HashedPassword)
	case request.InsertTo == pb.Table_PROFILE_STATUS:
		var profile_status pb.ProfileStatus
		err = request.Entry.UnmarshalTo(&profile_status)
		if err != nil {
			return nil, err
		}
		row = r.conn.QueryRow(context.Background(), "insert into profile_status values ($1, $2)", profile_status.Id, profile_status.Name)
	case request.InsertTo == pb.Table_MOVIE:
		var movie pb.Movie
		err = request.Entry.UnmarshalTo(&movie)
		if err != nil {
			return nil, err
		}
		row = r.conn.QueryRow(context.Background(), "insert into movie values ($1, $2, $3, $4))", movie.Id, movie.ViewsCount, movie.LikesCount, movie.DislikesCount)
	}

	if row == nil {
		result_reply.Result = false
		result_reply.Info = "ERROR: Entry not inserted."
	}
	return &result_reply, err
}
