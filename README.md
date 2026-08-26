# Movies Ticketing

Movies Ticketing is a RESTful API built with Go and Gin for online movie ticket booking. It features JWT based authentication, clean 3 layer architecture (Handler, Service, Repository), and race-safe ticket purchases backed by PostgreSQL transactions with row-level locking.

## Tech Stack

- Go
- Gin
- PostgreSQL
- Docker

## Setup

To set up the project locally, follow these steps:

1. Clone the repository: `git clone https://github.com/ahmadilham22/movies-ticket.git`
2. Navigate to the project directory: `cd movies-ticket`
3. Install dependencies: `go mod download`
4. Copy the example environment file: `cp .env.example .env`
5. Set up the database: `docker-compose up -d`

## API Endpoints

- `POST /users`: Register a new user
- `POST /login`: Login and obtain a JWT token
- `POST /tickets`: Book a ticket
- `POST /tickets/create`: Create an event
- `GET /users`: Get a list of users
- `GET /tickets`: Get a list of tickets

## Authentication

Login -> Get a token -> Send Authorization header with Bearer token for protected endpoints -> Use token for protected endpoints -> Send protected requests -> Receive protected responses
