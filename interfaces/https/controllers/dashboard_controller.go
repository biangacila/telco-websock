package controllers

import (
	"encoding/json"
	"github.com/biangacila/telco-websock/application/services"
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"net/http"
)

type DashboardController struct {
	dashboardService *services.DashboardService
}

// NewDashboardHandler creates a new handler for managing dashboard interactions
func NewDashboardController(service *services.DashboardService) *DashboardController {
	return &DashboardController{dashboardService: service}
}

// PostDashboardInfo handles posting new dashboard info for a user
func (h *DashboardController) PostDashboardInfo(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserCode string                 `json:"userCode"`
		Data     map[string]interface{} `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userCode := valueobjects.UserCode{Code: request.UserCode}
	h.dashboardService.UpdateDashboardInfo(userCode, request.Data)
	w.WriteHeader(http.StatusOK)
}

// GetDashboardUpdate checks for updated dashboard info for a user
func (h *DashboardController) GetDashboardUpdate(w http.ResponseWriter, r *http.Request) {
	userCode := r.URL.Query().Get("userCode")
	if userCode == "" {
		http.Error(w, "Missing userCode", http.StatusBadRequest)
		return
	}

	info, updated := h.dashboardService.CheckForUpdates(valueobjects.UserCode{Code: userCode})
	if !updated {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(info)
}
