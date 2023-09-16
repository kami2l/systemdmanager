package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"systemdmanager/internal/model"
	"systemdmanager/internal/util"
)

// GetStatusAll list all services and their statuses.
//
// systemctl list-units --type=service --all
func GetStatusAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	cmd := exec.Command("systemctl", "list-units", "--type=service", "--all")
	output, err := cmd.CombinedOutput()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)

		return
	}

	var services []model.Service
	for _, serviceStr := range strings.Split(string(output), "\n") {
		if !strings.Contains(serviceStr, ".service") {
			continue
		}

		serviceStr = util.TrimNonAlphaRubbish(serviceStr)
		parts := strings.Fields(serviceStr)

		service := model.Service{
			Unit:   parts[0],
			Load:   parts[1],
			Active: parts[2],
			Sub:    parts[3],
		}

		for i := 4; i < len(parts); i++ {
			service.Description += parts[i] + " "
		}

		service.Description = strings.TrimSpace(service.Description)

		services = append(services, service)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(services)
}

// GetStatus return the status of a specific service.
//
// systemctl list-units --type=service --all -> filter out
func GetStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Service name is required in the path")

		return
	}
	service := pathParts[2]

	cmd := exec.Command("systemctl", "list-units", "--type=service", "--all")
	output, err := cmd.CombinedOutput()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)

		return
	}

	for _, serviceStr := range strings.Split(string(output), "\n") {
		if !strings.Contains(serviceStr, ".service") {
			continue
		}

		serviceStr = util.TrimNonAlphaRubbish(serviceStr)
		parts := strings.Fields(serviceStr)

		if strings.Contains(serviceStr, service) {
			service := model.Service{
				Unit:   parts[0],
				Load:   parts[1],
				Active: parts[2],
				Sub:    parts[3],
			}

			for i := 4; i < len(parts); i++ {
				service.Description += parts[i] + " "
			}

			service.Description = strings.TrimSpace(service.Description)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(service)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = fmt.Fprint(w, fmt.Sprintf("Service %s not found", service))
}
