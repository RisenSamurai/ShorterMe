# ShorterMe

> **High-Performance URL Shortener & Link Management Engine**

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Gin Framework](https://img.shields.io/badge/Gin-Web_Framework-008ECF?style=flat&logo=go)](https://gin-gonic.com/)
[![Redis](https://img.shields.io/badge/Redis-In--Memory_Store-DC382D?style=flat&logo=redis)](https://redis.io/)
[![SvelteKit](https://img.shields.io/badge/SvelteKit-Tailwind_v4-FF3E00?style=flat&logo=svelte)](https://svelte.dev/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker)](https://www.docker.com/)

ShorterMe is an ultra-fast URL shortening web service built with a Go (Gin) backend, an in-memory Redis key-value cache, and a sleek SvelteKit frontend styled with TailwindCSS v4.

---

## ⚡ Tech Stack

- **Backend:** Go, Gin HTTP framework, Go-Redis (`github.com/go-redis/redis/v8`), CORS middleware
- **Frontend:** SvelteKit, TypeScript, TailwindCSS v4, Vitest, Testing Library
- **In-Memory Store:** Redis (10-minute TTL per short link, instant lookups)
- **Containerization:** Docker & Docker Compose

---

## 🚀 Key Features

- **Sub-Millisecond Resolution:** Direct in-memory lookup via Redis delivers instant 302 redirects.
- **Collision-Resistant Hash Generation:** Base-36 cryptographic random code generator.
- **Modern SvelteKit Frontend:** Instant URL generation with one-click copy to clipboard.
- **Comprehensive Test Suite:** Client and server tests using Vitest and Testing Library.
- **Zero-Config Deployment:** Ready-to-run multi-container setup via Docker Compose.

---

## 📡 API Endpoints

### Shorten URL
```http
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com/very/long/target/address"
}
```
**Response:**
```json
{
  "shortUrl": "http://localhost:8080/redirect/2f9a1b"
}
```

### Redirect
```http
GET /redirect/:id
```
*Responds with HTTP 302 Found, redirecting immediately to the destination URL.*

---

## 🚦 Quick Start with Docker

```bash
# Clone the repository
git clone https://github.com/RisenSamurai/ShorterMe.git
cd ShorterMe

# Spin up Redis, Go Backend, and SvelteKit Frontend
docker compose up -d --build
```
Open [http://localhost:5173](http://localhost:5173) in your browser.

---

## 📄 License
MIT License.
