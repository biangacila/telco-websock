package repositories

import (
	"github.com/biangacila/telco-websock/domain/entities"
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"sync"
)

// DashboardRepository defines the interface for storing dashboard information
type DashboardRepository interface {
	Store(info *entities.DashboardInfo)
	Get(userCode valueobjects.UserCode) (*entities.DashboardInfo, bool)
}

// InMemoryDashboardRepository stores dashboard info temporarily
type InMemoryDashboardRepository struct {
	store map[string]*entities.DashboardInfo
	lock  sync.RWMutex
}

// NewInMemoryDashboardRepository creates a new instance of the repository
func NewInMemoryDashboardRepository() *InMemoryDashboardRepository {
	return &InMemoryDashboardRepository{
		store: make(map[string]*entities.DashboardInfo),
	}
}

// Store stores or updates the dashboard info for a user
func (repo *InMemoryDashboardRepository) Store(info *entities.DashboardInfo) {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	repo.store[info.UserCode.Code] = info
}

// Get retrieves the dashboard info for a user
func (repo *InMemoryDashboardRepository) Get(userCode valueobjects.UserCode) (*entities.DashboardInfo, bool) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	info, exists := repo.store[userCode.Code]
	return info, exists
}
