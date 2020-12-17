package main

import "consulapi/cmd"

func main() {

	cmd.Execute()
	// 目標(copy) - 對象(paste)
	//consul.CloneKV("http://localhost:8500", "http://192.168.31.229:8500")
}
