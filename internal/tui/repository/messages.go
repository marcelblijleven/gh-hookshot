package repository

import "github.com/marcelblijleven/gh-hookshot/internal/data"

type dataFetchMsg struct {
	Repo data.Repository
	Err  error
}

type RepositoryDataMsg struct {
	Valid bool
	Err   error
}
