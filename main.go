package main

import "echo-server/cmd/api"

func main() {
	app := api.NewServer()
	app.Run()
}
