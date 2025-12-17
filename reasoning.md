# Reasoning & Design Decisions

## Overview
This project implements a RESTful API in Go for managing users with name and date of birth.
Age is calculated dynamically and is not stored in the database.

The application follows a layered architecture to ensure clean separation of concerns.

## Architecture
- **Handlers**: Handle HTTP request/response logic
- **Services**: Contain business logic such as age calculation
- **Repositories**: Abstract database access using SQLC
- **Database**: PostgreSQL

Request flow:
HTTP → Routes → Handler → Service → Repository → Database

## Database Design
The database stores only:
- id
- name
- dob

Age is not stored because it is derived data and changes over time.

## SQLC Usage
SQLC is used to generate type-safe database access code from raw SQL.
This avoids runtime SQL errors and keeps database logic explicit and readable.

The pgx driver was chosen for better PostgreSQL support and performance.

## Age Calculation
Age is calculated dynamically using Go’s time package by comparing the current date with the user’s date of birth.
This ensures correctness without redundancy.

## Validation & Error Handling
Request validation is done using go-playground/validator.
Appropriate HTTP status codes are returned for different error cases.

## Logging & Middleware
Uber Zap is used for structured logging.
Middleware is added to:
- Inject a request ID
- Log request duration

## Conclusion
The project focuses on simplicity, correctness, and clean architecture while following best backend development practices.
