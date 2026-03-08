# XStorage

Lightweight, S3-inspired object storage for self-hosted apps. API-compatible design, multi-app buckets, and a simple admin dashboard.

**By [xmodmor](https://xofdev.com)**

---

## Features

- S3-like API (buckets, objects, pre-signed URLs)
- API key auth · Web dashboard · PostgreSQL + filesystem
- Docker-ready · Modular and scalable

---

## Stack

| Layer        | Tech           |
| ------------ | -------------- |
| Backend      | Go             |
| Dashboard    | Nuxt 4         |
| Database     | PostgreSQL     |
| Deploy       | Docker Compose |

---

## Quick Start

```bash
git clone https://github.com/xmodmor/XStorage.git && cd XStorage
docker compose up -d
```

- **API:** http://localhost:8080  
- **Dashboard:** http://localhost:3000  

---

## API (overview)

| Action        | Method / Path                              |
| ------------- | ------------------------------------------ |
| Create bucket| `POST /buckets`                             |
| Upload       | `PUT /buckets/{bucket}/objects/{key}`       |
| Download     | `GET /buckets/{bucket}/objects/{key}`       |
| List         | `GET /buckets/{bucket}/objects`             |
| Delete       | `DELETE /buckets/{bucket}/objects/{key}`    |

**Auth:** send `X-Access-Key` and `X-Secret-Key` in headers.

---

## License

MIT
