package main

import (
	"net/http"
	"strings"
)

func check_exist(pod_map map[string][]string, resource_name string) bool {
	for _, m := range pod_map["NAME"] {
		if strings.Contains(m, resource_name){
			return true
		}
	}
	return false
}
func check_service_url(service_std map[string][]string, ip string, resource_name string) bool {
	url := "http://192.168.10.80:30001/login"
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func check_status(pod_map map[string][]string, resource_name string) bool {
	for n, _ := range pod_map["NAME"] {
		if pod_map["STATUS"][n] == "Running" {
			return true
		}

	}
	return false
}

func check_ready(pod_map map[string][]string, resource_name string) bool {
	for n, m := range pod_map["NAME"] {
		if strings.Contains(m, resource_name) {
			ready := (strings.Split(pod_map["READY"][n], "/"))
			if ready[0] == ready[1] {
				return true
			}
		}
	}

	return false
}


