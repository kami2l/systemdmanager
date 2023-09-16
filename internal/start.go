package internal

import (
	"fmt"
	"net/http"
	"os/exec"
)

// StartService starts a specific service
//
// sudo systemctl start <service>
func StartService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	service := r.URL.Query().Get("service")

	cmd := exec.Command("sudo", "systemctl", "start", service)
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, fmt.Sprintf("Service %s started", service))
}
