package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "/tmp/val.txt", "")
	flag.Parse()

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(data), "***")
	var key string
	for _, val := range parts {
		val = strings.Trim(val, "\n")
		val = strings.TrimSpace(val)
		if strings.HasPrefix(key, "/reigstry") {
			key = val
			continue
		}

		_, err := exec.Command("etcdctl", `--endpoints="https://172.16.81.71:2379"`, "put", key, "`"+val+"`").Output()
		if err != nil {
			log.Fatal("etcdctl: %v", err)
		}
		log.Printf("etcdctl put %v OK\n", key)
	}
}
