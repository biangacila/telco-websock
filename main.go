package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/biangacila/telco-websock/application/services"
	"github.com/biangacila/telco-websock/domain/repositories"
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"github.com/biangacila/telco-websock/infrastructure/websockets"
	"github.com/biangacila/telco-websock/interfaces/grpc"
	"github.com/biangacila/telco-websock/interfaces/wss"
	"github.com/biangacila/telco-websock/utils"
	"log"
	"net/http"
)

type Config struct {
	Port         int
	Env          string
	Backend      string
	DbHostServer string
	DbName       string
	Version      string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Backend     string `json:"backend"`
}

func main() {
	var cfg Config
	flag.IntVar(&cfg.Port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development | production)")
	flag.StringVar(&cfg.Backend, "backend", "/backend-telcowebsocket/api", "Application environment webservice dispatch main route")
	flag.StringVar(&cfg.Version, "version", "1.0", "API version")
	flag.Parse()
	// Initialize WebSocketManager
	wsManager := websockets.NewWebSocketManager()

	// Initialize repository (in-memory or database-backed, depending on your implementation)
	dashboardRepo := repositories.NewInMemoryDashboardRepository() // Replace with actual implementation

	// Initialize DashboardService
	dashboardService := services.NewDashboardService(dashboardRepo, wsManager)

	// Create WebSocket handler
	wsHandler := wss.NewWebSocketHandler(wsManager)

	// System health controller
	http.HandleFunc(cfg.Backend+"/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.Env,
			Version:     cfg.Version,
			Backend:     cfg.Backend,
		}
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err1 := w.Write(js)
		if err1 != nil {
			return
		}
	})

	// HTTP route to serve WebSocket connections
	http.HandleFunc(cfg.Backend+"/ws", wsHandler.ServeWebSocket)

	// Example additional HTTP routes to submit dashboard information (via HTTP POST)
	http.HandleFunc(cfg.Backend+"/dashboard/update", func(w http.ResponseWriter, r *http.Request) {
		userCode := r.URL.Query().Get("userCode")
		if userCode == "" {
			http.Error(w, "userCode is required", http.StatusBadRequest)
			return
		}

		// You would normally parse and validate the data here

		data, err := utils.FetchPayloadData(r)
		if err != nil {
			http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
			return
		}

		// Update dashboard info using the DashboardService
		dashboardService.UpdateDashboardInfo(valueobjects.NewUserCode(userCode), data)
		_, _ = w.Write([]byte("Dashboard info updated"))
	})

	// Start the HTTP server
	go func() {
		log.Println("Starting telco websocket at port", fmt.Sprintf(":%d", cfg.Port), " ...")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Start the gRPC server
	grpc.StartGRPCServer(*dashboardService)
}
