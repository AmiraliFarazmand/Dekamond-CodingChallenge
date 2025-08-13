ChatGPT said:
markdown
Copy
Edit

# Dekamond – Go Backend (Gin + GORM + Postgres + Redis)

A backend service written in **Go** using **Gin**, **GORM**, **PostgreSQL**, and **Redis**.

---

## 📦 Prerequisites

- **Go** 1.22+
- **PostgreSQL** 15+
- **Redis** 6.2+
- `make` (optional)

---

## 🚀 Run Locally (Without Docker)

### Have Redis and ProstgreSQL on your local

### Create a .env file and fill these environment variables(like the example below)

<div class="code-file">.env</div>

```python
DSN="host=localhost user=pgUser password=pgPW dbname=dekamond port=5432 sslmode=disable TimeZone=Asia/Shanghai"
SECRET_KEY="some_secret_key_kalfjddfsfsdfsflkfmkmfklsdmfklsdmfkslmdfklsmfdklsmdfh"
REDIS_ADDR="localhost:6379"
```

### run it with $go run cmd/main/main.go

## 🚀 Run With Docker

### execute these commands:

```
docker compose build
docker compose up
```
