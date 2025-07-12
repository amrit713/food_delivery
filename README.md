# 🍔 Go Fiber Food Delivery API

A full-featured **Food Delivery REST API** built in **Go** using [Fiber](https://github.com/gofiber/fiber), GORM, PostgreSQL, and JWT-based authentication. Designed with a clean architecture and scalable folder structure, this backend powers operations for users, restaurants, menu items, orders, and payments.

---

## 🚀 Features

* 🔐 **JWT Authentication & Cookie-based session**
* 👥 **Role-based User Access** (Admin, Owner, Customer)
* 🍽️ **Restaurant CRUD** (create/edit your own)
* 📋 **Menu & Item Management**
* 🛒 **Order System with Multi-Item Support**
* 💳 **Payment Support Placeholder**
* 📟 **Review & Rating System**
* 📦 **Clean folder structure** (`cmd`, `internal`, `pkg`)
* 📘 **Swagger API Docs Integration**

---

## 🧱 Tech Stack

| Layer         | Tool                        |
| ------------- | --------------------------- |
| Language      | Go 1.22+                    |
| Web Framework | [Fiber](https://gofiber.io) |
| ORM           | [GORM](https://gorm.io/)    |
| Database      | PostgreSQL with UUID        |
| Auth          | JWT + Cookies               |
| Docs          | Swagger via gofiber/contrib |
| Validation    | `go-playground/validator`   |

---

## 📁 Project Structure

```
food-delivery/
│
├── cmd/                    # App entrypoint (main.go)
├── internal/               # Main application packages
│   ├── controllers/        # Fiber route handlers
│   ├── services/           # Business logic
│   ├── repositories/       # DB access via GORM
│   ├── models/             # GORM models
│   ├── routes/             # API versioned routes
│   └── middlewares/        # JWT auth, protected routes
│
├── pkg/                    # Shared utils and packages
│   └── utils/              # Password hash, JWT, etc
│
├── docs/                   # Swagger docs
├── go.mod / go.sum         # Dependencies
└── README.md               # You're here
```

---

## 🛠️ Setup Instructions

### 1. Clone and install

```bash
git clone https://github.com/amrit713/food_delivery.git
cd food-delivery
go mod tidy
```

### 2. Setup PostgreSQL

Make sure you have PostgreSQL running locally and create a database:

```sql
CREATE DATABASE food_delivery;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
```

Update your DB connection string in `config/db.go` or `.env`.

### 3. Run the App

```bash
go run cmd/main.go
```

The server will run at: `http://localhost:4000`

---



## 🔐 Authentication Flow

* **Register/Login** creates a JWT token and stores it in a secure cookie
* Protected routes require `middleware.Protected(userRepo)` to validate the token and load user into `ctx.Locals("user")`

---

## ✅ TODO / Coming Soon

* [ ] Stripe Payment Integration
* [ ] Order Status Updates
* [ ] Admin Dashboard Controls
* [ ] Upload Image to Cloudinary or S3
* [ ] Docker Support
* [ ] React Native Frontend

---

## 🤝 Contributing

Feel free to open issues or PRs for enhancements and bugfixes.

---

## 📄 License

MIT License. © 2025 [Amrit Ghimire](https://github.com/amrit713)

---
