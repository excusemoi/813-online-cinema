package models

//grpc создаст кастомные типы - пихуй

type User struct {
	Id       string `pg:"id"`
	Name     string `pg:"name"`
	Surname  string `pg:"surname"`
	Login    string `pg:"login"`
	Password string `pg:"password"`
	Status   string `pg:"status"`
}

type MovieUser struct {
	UserId  string `pg:"user_id"`
	MovieId string `pg:"movie_id"`
}

type Movie struct {
	Id         string `pg:"id"`
	TorrentUrl string `pg:"torrent_url"`
}

type MovieStats struct {
	Id      string `pg:"id"`
	Views   int64  `pg:"views"`
	Reviews int64  `pg:"reviews"`
}
