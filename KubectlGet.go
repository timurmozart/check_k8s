package main

import "os/exec"

func KubectlGet(resource string, namespace string) string {
	app := "/usr/local/nagios/libexec/kubectl"
	arg0 := resource
	arg1 := namespace

	cmd := exec.Command(app, "get", arg0, "-n", arg1)
	stdout, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(stdout)
}
