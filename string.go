package main

import "fmt"

func main() {
	creds := make([]string, 0)
	authKeys := fmt.Sprintf("%s:%s",
		"root",
		"qwe")

	creds = append(creds, "--user", authKeys)

	args := make([]string, 0)
	args = append(args, creds...)
	args = append(args, "--endpoints", "etcd:2379", "endpoint", "health")

	fmt.Println(args)
}
