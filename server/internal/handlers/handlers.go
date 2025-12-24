package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"go.racktop.io/monorepo/server/internal/models/response"
	"go.racktop.io/monorepo/server/internal/services"
)

// RegisterRoutes registers all API routes on the given ServeMux
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", handleHealth)
	mux.HandleFunc("GET /api/info", handleInfo)
}

// @Summary Health check
// @Tags system
// @Produce json
// @Success 200 {object} response.Health
// @Router /api/health [get]
func handleHealth(w http.ResponseWriter, r *http.Request) {
	status := "unhealthy"
	if services.HealthCheck() {
		status = "ok"
	}
	writeJSON(w, r, http.StatusOK, response.Health{
		Status: status,
	})
}

// @Summary Service info
// @Tags system
// @Produce json
// @Success 200 {object} response.Info
// @Failure 500 {object} response.Error
// @Router /api/info [get]
func handleInfo(w http.ResponseWriter, r *http.Request) {
	name, version, err := services.GetServiceInfo()
	if err != nil {
		slog.Error("Error getting service info", "error", err)
		writeError(w, r, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, r, http.StatusOK, response.Info{
		Name:    name,
		Version: version,
		Time:    time.Now().UTC().Format(time.RFC3339),
	})
}

func writeJSON(w http.ResponseWriter, r *http.Request, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("Error encoding JSON", "error", err)
		writeError(w, r, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(status)
	w.Write(jsonData)
}

func writeError(w http.ResponseWriter, r *http.Request, status int, err error) {
	accept := strings.ToLower(r.Header.Get("Accept"))
	if strings.Contains(accept, "application/json") {
		w.Header().Set("Content-Type", "application/json")
		b, marshalErr := json.Marshal(response.Error{Error: err.Error()})
		if marshalErr == nil {
			w.WriteHeader(status)
			_, _ = w.Write(b)
			return
		}
		slog.Error("Error encoding JSON", "error", marshalErr)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	_, _ = w.Write([]byte("error: " + err.Error()))
}
