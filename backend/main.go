package main

import (
	jwtauth "be-tasking/helper"
	"be-tasking/server"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	jwtauth.InitJWTKey([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func main() {
	// run rest server
	srv := server.NewRestServer()
	srv.Run()
}
