package main

import (
	"log"
	"net/http"
	"systemdmanager/internal"
)

func main() {
	http.HandleFunc("/status/all", internal.GetStatusAll)
	http.HandleFunc("/status/", internal.GetStatus)
	http.HandleFunc("/start", internal.StartService)
	http.HandleFunc("/stop", internal.StopService)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
