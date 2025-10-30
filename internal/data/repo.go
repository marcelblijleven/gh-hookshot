package data

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

type Repository struct {
	Owner       Owner       `json:"owner"`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Description string      `json:"description"`
	Permissions Permissions `json:"permissions"`
}

type Owner struct {
	Login string `json:"login"`
}

type Permissions struct {
	Admin bool `json:"admin"`
}

func (r Repository) IsAdmin() bool {
	return r.Permissions.Admin
}

func GetRepo(owner, repo string) (Repository, error) {
	var resp Repository

	client, err := api.DefaultRESTClient()
	if err != nil {
		return resp, err
	}

	err = client.Get(fmt.Sprintf("repos/%s/%s", owner, repo), &resp)

	return resp, err
}
