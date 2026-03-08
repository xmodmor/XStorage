# XStorage

Lightweight, S3-inspired object storage for self-hosted apps. Multi-app support, API key auth, bucket management, and a web dashboard -- all in a single `docker compose up`.

**By [xmodmor](https://xofdev.com)**

---

## Features

- S3-like REST API (buckets, objects, upload, download)
- Multi-app / multi-tenant with isolated API keys
- Pluggable storage backend (filesystem now, S3/Ceph later)
- JWT-authenticated admin dashboard
- PostgreSQL metadata with versioned migrations and seeder
- Dockerized -- runs with one command

---

## Stack

| Layer     | Tech              |
| --------- | ----------------- |
| Backend   | Go (Gin + GORM)   |
| Dashboard | Nuxt 4 + shadcn   |
| Database  | PostgreSQL 15      |
| Deploy    | Docker Compose     |

---

## Quick Start

```bash
git clone https://github.com/xmodmor/XStorage.git
cd XStorage
cp .env.example .env
docker compose up -d
```

| Service   | URL                     |
| --------- | ----------------------- |
| API       | http://localhost:8080    |
| Dashboard | http://localhost:3000    |
| Postgres  | localhost:5432           |

Default admin credentials (created by seeder):

```
email:    admin@xstorage.local
password: admin123
```

---

## Environment Variables

Copy `.env.example` to `.env` and adjust as needed:

```
POSTGRES_USER=xstorage
POSTGRES_PASSWORD=xstorage
POSTGRES_DB=xstorage
JWT_SECRET=change-me-in-production
```

---

## Project Structure

```
XStorage/
├── backend/                        # Go API server
│   ├── cmd/api/main.go             # Entry point, dependency wiring
│   ├── internal/
│   │   ├── config/                 # Environment configuration
│   │   ├── database/               # PostgreSQL connection + migrations
│   │   ├── domain/                 # GORM models (User, App, APIKey, Bucket, Object)
│   │   ├── dto/                    # Request/Response structs
│   │   ├── repository/             # Data access interfaces + implementations
│   │   ├── service/                # Business logic
│   │   ├── handler/                # HTTP handlers
│   │   ├── middleware/             # JWT auth, API key auth, CORS, rate limit
│   │   ├── router/                 # Route registration
│   │   ├── storage/                # Storage interface + filesystem implementation
│   │   └── response/               # Standardized JSON response helpers
│   ├── migrations/                 # Versioned SQL migrations
│   ├── seed/                       # Database seeder
│   └── Dockerfile
├── dashboard/                      # Nuxt 4 admin panel
│   ├── app/pages/                  # Dashboard, Apps, Buckets, Objects, API Keys
│   ├── app/layouts/                # Navigation layout
│   └── Dockerfile
├── docker-compose.yml
├── .env.example
└── .gitignore
```

---

## Architecture

```
Dashboard (Nuxt 4)
       │
   REST API
       │
  ┌────▼────┐
  │ Backend  │──── Middleware (Auth / Rate Limit)
  │   (Go)   │
  └────┬─────┘
       │
  ┌────▼─────────────────────┐
  │                          │
  ▼                          ▼
PostgreSQL             Filesystem (/data)
(metadata)             (file blobs)
```

Backend follows a layered architecture with constructor-based dependency injection:

**Handler** (parse request) → **Service** (business logic) → **Repository** (database) / **Storage** (files)

---

## API

All routes are versioned under `/api/v1`.

### Auth

| Action | Method | Path                |
| ------ | ------ | ------------------- |
| Login  | POST   | `/api/v1/auth/login` |

### Apps (JWT auth)

| Action | Method | Path              |
| ------ | ------ | ----------------- |
| Create | POST   | `/api/v1/apps`     |
| List   | GET    | `/api/v1/apps`     |
| Get    | GET    | `/api/v1/apps/:id` |
| Delete | DELETE | `/api/v1/apps/:id` |

### API Keys (JWT auth)

| Action | Method | Path                        |
| ------ | ------ | --------------------------- |
| Create | POST   | `/api/v1/apps/:id/keys`      |
| List   | GET    | `/api/v1/apps/:id/keys`      |
| Delete | DELETE | `/api/v1/apps/:id/keys/:keyId` |

### Buckets (API key auth)

| Action | Method | Path                       |
| ------ | ------ | -------------------------- |
| Create | POST   | `/api/v1/buckets`           |
| List   | GET    | `/api/v1/buckets`           |
| Delete | DELETE | `/api/v1/buckets/:bucket`   |

### Objects (API key auth)

| Action   | Method | Path                                     |
| -------- | ------ | ---------------------------------------- |
| Upload   | PUT    | `/api/v1/buckets/:bucket/objects/*key`    |
| Download | GET    | `/api/v1/buckets/:bucket/objects/*key`    |
| List     | GET    | `/api/v1/buckets/:bucket/objects`         |
| Delete   | DELETE | `/api/v1/buckets/:bucket/objects/*key`    |

**API key auth:** Send `X-Access-Key` and `X-Secret-Key` headers.

**Response format:**

```json
{ "success": true, "data": { ... } }
{ "success": false, "error": { "code": "NOT_FOUND", "message": "..." } }
```

---

## Local Development (without Docker)

### Backend

Requires Go 1.25+ and a running PostgreSQL instance.

```bash
cd backend
export DATABASE_URL="postgres://xstorage:xstorage@localhost:5432/xstorage?sslmode=disable"
export JWT_SECRET="dev-secret"
export STORAGE_PATH="./data"
go run ./cmd/api
```

### Dashboard

Requires Node 22+.

```bash
cd dashboard
npm install
npm run dev
```

---

## Database

Schema is managed via [golang-migrate](https://github.com/golang-migrate/migrate). Migration files live in `backend/migrations/`. Migrations run automatically on backend startup.

### Tables

- **users** -- admin panel accounts
- **apps** -- registered applications (multi-tenant)
- **api_keys** -- per-app access/secret key pairs
- **buckets** -- storage containers scoped to an app
- **objects** -- file metadata (key, size, mime, checksum, storage path)

---

## Roadmap

- [ ] Signed / pre-signed URLs
- [ ] Multipart upload
- [ ] Object versioning
- [ ] Lifecycle rules (auto-delete after N days)
- [ ] Webhooks (`object.uploaded`, `object.deleted`)
- [ ] S3-compatible storage backend
- [ ] Image processing (resize)
- [ ] CDN support
- [ ] Dashboard UI implementation

---

## License

MIT
