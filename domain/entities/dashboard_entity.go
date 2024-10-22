package entities

import (
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"time"
)

type DashboardInfo struct {
	UserCode      valueobjects.UserCode
	Data          map[string]interface{} // This will store the dashboard info in a generic way
	LastUpdatedAt time.Time
}
