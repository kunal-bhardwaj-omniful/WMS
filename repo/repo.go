package repo

import (
	"github.com/omniful/go_commons/db/sql/postgres"
	"sync"
)

type Repository interface {
}

type repository struct {
	db *postgres.DbCluster
}

var repo *repository
var repoOnce sync.Once

// NewRepository is the constructor function that ensures the Repository is initialized only once.
func NewRepository(db *postgres.DbCluster) Repository {
	repoOnce.Do(func() {
		// Initialize the Repository with a given DbCluster.
		repo = &repository{
			db: db,
		}
	})
	return repo
}
