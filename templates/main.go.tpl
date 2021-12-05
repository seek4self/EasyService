package main

import (
    "fmt"

    "{{ .Name }}/initialize"
)

// @title {{ .Name }} API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /{{ .Name }}
func main() {
    initialize.RunServer()
}
