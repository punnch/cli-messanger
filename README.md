# cli-messanger

A real-time CLI messenger built with Go, PostgreSQL, and Docker.  
Multiple users can register, join rooms, and chat with each other over raw TCP.

## Stack

| Layer | Technology |
|---|---|
| Language | Go 1.26 |
| Transport | TCP (`net` package) |
| Database | PostgreSQL 16 |
| DB driver | `pgx/v5` |
| Migrations | `golang-migrate` |
| Logging | `zap` |
| Containerization | Docker + Docker Compose |
| Architecture | Monolith |

## Project Structure

```
cli-messanger/
├── cmd/
│   ├── server/
│   │   ├── main.go          ← DI entrypoint: repo → service → transport
│   │   └── Dockerfile       ← multi-stage build
│   └── client/
│       └── main.go          ← stdin reader, TCP dial
├── internal/
│   ├── core/
│   │   ├── config/          ← app config from env
│   │   ├── domain/          ← User, Room, Member, Message
│   │   ├── errors/          ← shared error types
│   │   ├── logger/          ← zap wrapper
│   │   ├── repository/postgres/pool/ ← Pool interface + pgx impl
│   │   └── transport/tcp/
│   │       ├── handler/     ← connection loop, command dispatch
│   │       ├── middleware/  ← Logger, Recover
│   │       └── server/      ← net.Listener, graceful shutdown
│   └── features/
│       ├── hub/             ← broadcast goroutine, room → []conn map
│       ├── auth/            ← register, login, bcrypt
│       ├── rooms/           ← join, list rooms
│       └── messages/        ← send, history
├── migrations/
│   ├── 000001_users.{up,down}.sql
│   ├── 000002_rooms.{up,down}.sql
│   └── 000003_messages.{up,down}.sql
├── docker-compose.yaml
├── Makefile
├── go.mod
└── .env.example
```

## Requirements

- [Go 1.26+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [golang-migrate](https://github.com/golang-migrate/migrate)

## Setup

**1. Clone the repo**

```bash
git clone https://github.com/punnch/cli-messanger.git
cd cli-messanger
```

**2. Configure environment**

```bash
cp .env.example .env
```

Fill in `.env`:

```env
TCP_ADDR=:9000

POSTGRES_USER=messenger
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=messenger
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG

TIME_ZONE=UTC
```

**3. Start PostgreSQL**

```bash
make env-up
```

**4. Run migrations**

```bash
make migrate-up
```

**5. Start the server**

```bash
make messanger-run
```

**6. Connect a client** (in a separate terminal)

```bash
make client-run
```

## Docker Deployment

Run everything in Docker:

```bash
make messanger-deploy
make migrate-up
```

Stop:

```bash
make messanger-undeploy
```

## Commands

| Command | Description |
|---|---|
| `/register <username> <password>` | Create a new account |
| `/login <username> <password>` | Authenticate |
| `/rooms` | List available rooms |
| `/join <room>` | Join or create a room |
| `/msg <text>` | Send a message to the current room |
| `/history` | Show last 50 messages in the current room |
| `/help` | Show all commands |

## Example Session

```
$ make client-run

Welcome! Use /register <username> <password> or /login <username> <password>

/register dima password123
OK registered and logged in as dima

/join golang
OK joined in room golang

/rooms
1. golang

/msg hello everyone!
[dima] hello everyone!

/history
[dima] hello everyone!
```

## Makefile Targets

| Target | Description |
|---|---|
| `make env-up` | Start PostgreSQL container |
| `make env-down` | Stop PostgreSQL container |
| `make env-cleanup` | Remove PostgreSQL container and data |
| `make env-port-forward` | Forward PostgreSQL port to localhost |
| `make migrate-up` | Apply all migrations |
| `make migrate-down` | Roll back all migrations |
| `make messanger-run` | Run server locally |
| `make messanger-deploy` | Build and start server in Docker |
| `make messanger-undeploy` | Stop server container |
| `make client-run` | Run CLI client |
| `make logs-cleanup` | Delete log files |
| `make ps` | Show running containers |

## Architecture

The server accepts TCP connections and spawns a goroutine per client.  
Session state (authenticated user, current room) lives on the goroutine stack.  
A central `Hub` manages room membership and broadcasts messages to all connections in a room via a buffered channel.

```
Client (TCP) → TCPServer → Handler (goroutine per conn)
                                   ↓
                          UsersService → PostgreSQL
                          RoomsService → PostgreSQL
                          MessagesService → PostgreSQL
                          Hub → broadcast to room conns
```
