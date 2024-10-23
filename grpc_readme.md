
grpcurl -plaintext -d '{"userCode": "UC1022", "data": "{\"gRPC-Push\":\"Updated data\"}"}' localhost:50051 dashboard.DashboardService/UpdateDashboard
