# Team Maker API

Team Maker API is a backend service written in Go that helps you manage players and generate fair teams automatically.  
Typical use cases include:

- Splitting a group of players into balanced teams based on skill level
- Managing players and their attributes (name, rank, preferred position, etc.)
- Creating and managing match sessions

This project is also intended as a **portfolio showcase** of how I structure Go services using **Clean Architecture (Gogen)**.

---

## Tech Stack

- **Language:** Go
- **Architecture:** Clean Architecture (Gogen)
- **Config:** JSON-based config
- **Containerization:** Docker

---

## Features (Planned / Implemented)

- Player management (CRUD)
- Team generation logic:
  - Randomized team assignment
  - Can be extended to support rank / skill / position balancing
- Match/session management (grouping players into a specific game session)
- Structured project layout using Gogen:
  - Separation of domain, application, and infrastructure layers

> Note: The feature set is still evolving. This repository is primarily to demonstrate backend structure, not a finished product.

---

## Project Structure

High-level structure:

```bash
.
├── application/                 # Application layer (use cases, services)
├── domain_teammakerappservice/  # Domain layer (entities, interfaces)
├── lib/core/client/             # Shared client libraries / helpers
├── shared/                      # Shared utilities, responses, errors, etc.
├── config/                      # Environment-related configs (if any)
├── config.sample.json           # Sample config file
├── main.go                      # Application entrypoint
├── Dockerfile                   # Container build definition
├── go.mod                       # Go module definition
└── go.sum                       # Dependencies lock file
