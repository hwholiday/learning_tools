package aggregate

import (
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/repository"
)

type Factory struct {
	*AuthFactory
}

func NewFactory(repo *repository.Repository) *Factory {
	return &Factory{
		&AuthFactory{
			merchantRepo:  repo.Merchant,
			authCodeRepo:  repo.AuthCode,
			authTokenRepo: repo.AuthToken,
		},
	}
}
