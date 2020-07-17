package main

import (
	"consulapi/consul"
)

func main() {
	consul.CloneKV("http://192.168.8.40:8500", "http://localhost:8500")
}
