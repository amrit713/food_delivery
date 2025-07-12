# 🍔 Go Fiber Food Delivery API

A full-featured **Food Delivery REST API** built in **Go** using [Fiber](https://github.com/gofiber/fiber), GORM, PostgreSQL, and JWT-based authentication. Designed with a clean architecture and scalable folder structure, this backend powers operations for users, restaurants, menu items, orders, and payments.

---

## 🚀 Features

- 🔐 **JWT Authentication & Cookie-based session**
- 👥 **Role-based User Access** (Admin, Owner, Customer)
- 🍽️ **Restaurant CRUD** (create/edit your own)
- 📋 **Menu & Item Management**
- 🛒 **Order System with Multi-Item Support**
- 💳 **Payment Support Placeholder**
- 🧾 **Review & Rating System**
- 📦 **Clean folder structure** (`cmd`, `internal`, `pkg`)
- 📘 **Swagger API Docs Integration**

---

## 🧱 Tech Stack

| Layer        | Tool                         |
|--------------|------------------------------|
| Language     | Go 1.22+                     |
| Web Framework| [Fiber](https://gofiber.io) |
| ORM          | [GORM](https://gorm.io/)     |
| Database     | PostgreSQL with UUID         |
| Auth         | JWT + Cookies                |
| Docs         | Swagger via gofiber/contrib  |
| Validation   | `go-playground/validator`    |

---

## 📁 Project Structure

