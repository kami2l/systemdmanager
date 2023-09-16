package internal

import (
	"fmt"
	"net/http"
	"os/exec"
)

// StopService stop a specific service
//
// sudo systemctl stop <service>
func StopService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	service := r.URL.Query().Get("service")

	cmd := exec.Command("sudo", "systemctl", "stop", service)
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, fmt.Sprintf("Service %s stopped", service))
}
