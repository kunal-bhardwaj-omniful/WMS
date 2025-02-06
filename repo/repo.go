package repo

import (
	"context"
	"github.com/omniful/go_commons/db/sql/postgres"
	"sync"
	"wms/domain"
)

type Repository interface {
	GetAllHubs(ctx context.Context) []domain.Hub
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

func (r *repository) GetAllHubs(ctx context.Context) []domain.Hub {

	var hubs []domain.Hub
	//err := r.db.Find(&hubs).Error
	r.db.GetMasterDB(ctx).Find(&hubs)
	return hubs
}
