package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:        ":9090",
		Handler:     http.DefaultServeMux,
		ReadTimeout: 10 * time.Second,
	}
	http.HandleFunc("/", usage)
	http.HandleFunc("/aircon/on", powerOnAircon)
	http.HandleFunc("/aircon/off", powerOffAircon)
	http.HandleFunc("/ceiling/full", fullCeiling)
	http.HandleFunc("/ceiling/on", powerOnCeiling)
	http.HandleFunc("/ceiling/off", powerOffCeiling)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func usage(w http.ResponseWriter, _ *http.Request) {
	usage := "usage:\n" +
		"/aircon/on : power on aircon\n" +
		"/aircon/off : power off aircon\n"
	_, err := w.Write([]byte(usage))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func powerOnAircon(w http.ResponseWriter, _ *http.Request) {
	out, err := exec.Command("/home/pi/bin/irkit/aircon_on.sh").Output()
	fmt.Println(string(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func powerOffAircon(w http.ResponseWriter, _ *http.Request) {
	out, err := exec.Command("/home/pi/bin/irkit/aircon_off.sh").Output()
	fmt.Println(string(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func fullCeiling(w http.ResponseWriter, _ *http.Request) {
	out, err := exec.Command("/home/pi/bin/irkit/ceiling_full.sh").Output()
	fmt.Println(string(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func powerOnCeiling(w http.ResponseWriter, _ *http.Request) {
	out, err := exec.Command("/home/pi/bin/irkit/ceiling_on.sh").Output()
	fmt.Println(string(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func powerOffCeiling(w http.ResponseWriter, _ *http.Request) {
	out, err := exec.Command("/home/pi/bin/irkit/ceiling_off.sh").Output()
	fmt.Println(string(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
