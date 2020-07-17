package main

import (
	"consulapi/consul"
)

func main() {

	// 目標 - 對象
	consul.CloneKV("http://localhost:8500", "http://192.168.31.229:8500")
}
