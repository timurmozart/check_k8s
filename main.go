package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	resource := flag.String("r", "all", "resource type")
	namespace := flag.String("n", "default", "resource namspace")
	flag.Parse()
	resource_name := flag.Args()[0]
	res, ns := *resource, *namespace
	std := Stdout_to_dictionary(KubectlGet(res, ns))
	service_std := Stdout_to_dictionary(KubectlGet("service", ns))
	ip := "192.168.10.80"
	if check_exist(std, resource_name){
		var status [3]bool

		status[0] = check_service_url(service_std, ip, resource_name)
		status[1] = check_ready(std, resource_name)
		status[2] = check_status(std, resource_name)
		fmt.Println("status", status)
		if status[0] == true && status[1] == true && status[2] == true {
			os.Exit(0)
		} else if status[0] == false && status[1] == true && status[2] == true {
			os.Exit(1)
		} else if status[0] == false && (status[1] == false || status[2] == false) {
			os.Exit(2)
		}
	} else {
		os.Exit(3)
	}
}
