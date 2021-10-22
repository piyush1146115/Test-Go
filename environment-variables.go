package main

import (
	"fmt"
	"gomodules.xyz/envsubst"
	"log"
)

func main() {
	invokerDataDir := "/var/run/etcd/${HOSTNAME}"
	dataDir, err := envsubst.EvalMap(invokerDataDir, map[string]string{"HOSTNAME": "etcd-0"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataDir)
}
