package main

import (
	"813-online-cinema/pkg/services/db"
	"813-online-cinema/pkg/services/db/repository/postgres"
	"fmt"
)

func main() {
	rep := postgres.Repository{}
	if err := rep.InitRepository(&db.Cfg); err != nil {
		fmt.Println(err)
	}
}
