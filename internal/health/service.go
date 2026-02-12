package health

type HealthResponse struct {
	Status  string `json:"status"`
	Env     string `json:"env"`
	Version string `json:"version"`
}

func GetHealth() HealthResponse {
	return HealthResponse{
		Status:  "ok",
		Env:     "dev",
		Version: "1.0.0",
	}
}
