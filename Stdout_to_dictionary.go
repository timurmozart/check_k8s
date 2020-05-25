package main

import (
	"strings"
)

func Stdout_to_dictionary(stdin string) map[string][]string {
	stdout1 := strings.Split(string(stdin), "\n")
	stdout_dictionary := make(map[string][]string)
	for n, m := range strings.Fields(stdout1[0]) {
		stdout_dictionary[m] = []string{}
		for i := 1; i < len(stdout1)-1; i ++ {
			stdout_dictionary[m] = append(stdout_dictionary[m], strings.Fields(stdout1[i])[n])
		}
	}
	return stdout_dictionary
}