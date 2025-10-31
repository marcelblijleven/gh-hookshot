package repository

import "github.com/marcelblijleven/gh-hookshot/internal/api"

type dataFetchMsg struct {
	Repo api.Repository
	Err  error
}

type RepositoryDataMsg struct {
	Valid bool
	Err   error
}
