package services

import (
	"encoding/json"
	"github.com/biangacila/telco-websock/domain/aggregates"
	"github.com/biangacila/telco-websock/domain/entities"
	"github.com/biangacila/telco-websock/domain/repositories"
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"github.com/biangacila/telco-websock/infrastructure/websockets"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// DashboardService manages dashboard interactions
type DashboardService struct {
	repo      repositories.DashboardRepository
	wsManager *websockets.WebSocketManager
}

// NewDashboardService creates a new service
func NewDashboardService(repo repositories.DashboardRepository, wsManager *websockets.WebSocketManager) *DashboardService {
	return &DashboardService{repo: repo, wsManager: wsManager}
}

// UpdateDashboardInfo updates the dashboard info for a given user
func (s *DashboardService) UpdateDashboardInfo(userCode valueobjects.UserCode, data map[string]interface{}) {
	agg := aggregates.NewDashboardAggregate()
	dashboardInfo := agg.NewDashboardInfo(userCode, data)
	s.repo.Store(dashboardInfo)

	// Notify the WebSocket client
	conn, exists := s.wsManager.GetConnection(userCode.Code)
	if exists {
		msg, err := json.Marshal(dashboardInfo)
		if err == nil {
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Failed to send message to WebSocket client: %v", err)
				s.wsManager.RemoveConnection(userCode.Code)
			}
		}
	}
}

// CheckForUpdates checks if the user has updated dashboard info
func (s *DashboardService) CheckForUpdates(userCode valueobjects.UserCode) (*entities.DashboardInfo, bool) {
	info, exists := s.repo.Get(userCode)
	if !exists || time.Since(info.LastUpdatedAt) < time.Minute {
		return nil, false // No update available or recently updated
	}
	return info, true
}
