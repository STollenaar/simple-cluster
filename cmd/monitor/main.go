package main

import (
	"time"

	internal "github.com/STollenaar/simple-cluster/internal"
)

func main() {

	// Our main loop
	for {
		pods := internal.ListPods("default")
		internal.GetErrorLogs(pods)

		time.Sleep(time.Minute * 5)
	}
}
