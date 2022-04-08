# WebServer

Web Server Backend for Online Programming Judge IEEE Project 2021-2022.

Tech Stack : Go, Gin, Postgres.

# Setup:
    go mod download
    create database (install postgres) and specifiy name in internal/database/database.go
    cd migrations
    go run . init
    go run . up

# Run:
    go run app/main.go

# Curl:
    1) Login/Register: curl -d "Username=adithya&Password=secret123" -X POST localhost:8080/auth/<login OR register>
    2) Create: curl -d "Title=Kadane&Description=Find the maximum subarray sum&TimeLimit=1&MemoryLimit=256" -X POST localhost:8080/problem/create
    3) View: curl -X GET localhost:8080/problem/view/1
