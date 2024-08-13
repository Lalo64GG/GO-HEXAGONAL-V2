package main

import (
    "go-hexagonal-v2/src/bootstrap"
    "log"
    "github.com/joho/godotenv"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    r := bootstrap.Initialize()
    r.Run(":8080")
}
