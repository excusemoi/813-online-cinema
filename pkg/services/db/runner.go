package db

import (
	"813-online-cinema/pkg/services/db/config"
	pb "813-online-cinema/pkg/services/db/proto"
	"813-online-cinema/pkg/services/db/repository/postgres"
	"fmt"
)

var Cfg = config.Config{ //пока без cfg
	DbHost:     "localhost",
	DBPort:     5432,
	DbUser:     "postgres",
	DbPassword: "12345",
	DbName:     "OnlineCinema",
}

type Runner struct {
	Repository *postgres.Repository
	pb.UnimplementedDbServer
}

func NewRunner(repository *postgres.Repository) *Runner {
	return &Runner{Repository: repository}
}

func (runner *Runner) InitRunner(cfg *config.Config) error {
	return nil
}

func (runner *Runner) Start() error {
	cfg := &Cfg
	runner.Repository = &postgres.Repository{}
	if err := runner.Repository.InitRepository(cfg); err != nil {
		return err
	}
	fmt.Println(runner)
	return nil
}
