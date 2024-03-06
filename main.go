package main

import (
	"Backend-Golang/src/config"
	"Backend-Golang/src/helper"
	"Backend-Golang/src/routes"
	"fmt"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	// instal linter
	// go get github.com/golangci/golangci-lint/cmd/golangci-lint
	// run linter
	// go test -run=^$ -v ./...
	/*
	config > db.go
	 models > product.go
	 controllers > product.go
	 routes > main.go
	 helper > migration.go
	 ga perlu relation
	*/
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	defer config.DB.Close()
	routes.Router()
	fmt.Print("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}