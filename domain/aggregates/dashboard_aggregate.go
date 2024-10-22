package aggregates

import (
	"github.com/biangacila/telco-websock/domain/entities"
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"time"
)

type DashboardAggregate struct {
	entity *entities.DashboardInfo
}

func NewDashboardAggregate() *DashboardAggregate {
	return &DashboardAggregate{}
}

// NewDashboardInfo Factory function to create new DashboardInfo
func (a *DashboardAggregate) NewDashboardInfo(userCode valueobjects.UserCode, data map[string]interface{}) *entities.DashboardInfo {
	return &entities.DashboardInfo{
		UserCode:      userCode,
		Data:          data,
		LastUpdatedAt: time.Now(),
	}
}
