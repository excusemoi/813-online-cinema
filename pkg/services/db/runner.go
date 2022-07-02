package db

import (
	"813-online-cinema/pkg/services/db/config"
	"813-online-cinema/pkg/services/db/repository"
)

type Runner struct {
	config *config.Config
	pg     *repository.PostgresRepository
}

func Start() error {
	return nil
}
