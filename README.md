# Lab: Docker & Docker Compose

Runs:
- MySQL (**mysql_primary**) + phpMyAdmin
- PostgreSQL (**postgres**) + pgAdmin
- Simple static HTML site (**web**, multi-stage build)
- **Scaling demo:** 3 MySQL replicas (**mysql_replicas**)
- **Persistence:** named volumes for MySQL/Postgres
- **Healthcheck + depends_on:** pgAdmin waits for Postgres

## 1) Prereqs
- Docker Desktop running
- Copy `.env.example` → **`.env`** and fill values (do not commit `.env`)

## 2) How to run
```bash
docker compose --compatibility up -d --build
docker compose ps
3) URLs

phpMyAdmin: http://localhost:8081

Server: mysql_primary

User/Pass: from .env

pgAdmin: http://localhost:8082

Login: PGADMIN_EMAIL / PGADMIN_PASSWORD from .env

Add server inside pgAdmin: Host postgres, Port 5432, User/Pass from .env

Static site: http://localhost:8080

4) What each service is for

mysql_primary: single MySQL with named volume (mysql_data) → used by phpMyAdmin

phpmyadmin: GUI for MySQL (connects to mysql_primary)

mysql_replicas: 3 replicas (no volume) → scaling demo only

postgres: PostgreSQL with named volume (postgres_data) + healthcheck

pgadmin: GUI for Postgres; depends_on Postgres health

web: static HTML page served by Nginx, built via a multi-stage Dockerfile

Note: DB ports aren’t published to the host. UIs connect via the internal network using service names.