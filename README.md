# Go User API

A RESTful API built with Go, Fiber, PostgreSQL, and SQLC to manage users with dynamically calculated age.

## Features
- Create, update, delete users
- Fetch users with dynamically calculated age
- SQLC for type-safe database access
- PostgreSQL database
- Structured logging with Zap

## Run Locally
1. Setup PostgreSQL and create database
2. Update DATABASE_URL in config
3. Run:
   go run cmd/server/main.go
