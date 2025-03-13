# 🚀 Golang Clean Architecture Project

Welcome to my Golang repository! This project is built with **Clean Architecture** to ensure scalability, maintainability, and high performance. If you're into **well-structured** and **highly optimized** Golang code, you're in the right place! 💪⚡

## 📌 Tech Stack

- **Golang** 🐹 – High-performance backend language
- **Gin Gonic** 🍸 – Fast and lightweight web framework
- **GORM** 🛢 – Powerful ORM for database management
- **PostgreSQL/MySQL** 🗄 – Relational database support
- **Clean Architecture** 🏗 – Scalable and maintainable project structure
- **Docker** 🐳 – Containerized deployment
- **Swagger** 📜 – API documentation

## 📂 Project Structure

```
📦 project-root
 ┣ 📂 cmd            # Application entry point
 ┣ 📂 config         # Configuration files
 ┣ 📂 internal       # Core business logic (UseCase, Repository, etc.)
 ┃ ┣ 📂 entities     # Data models
 ┃ ┣ 📂 repository   # Database interactions
 ┃ ┣ 📂 usecase      # Business logic
 ┣ 📂 delivery       # HTTP handlers, controllers
 ┣ 📂 migrations     # Database migrations
 ┣ 📂 docs          # API Documentation (Swagger, etc.)
 ┣ 📜 .env.example  # Environment variables template
 ┣ 📜 Dockerfile    # Docker setup
 ┣ 📜 README.md     # You are here! 🧐
```

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/BagusRianKusuma/golang-playground.git
cd golang-playground
```

### 2️⃣ Install Dependencies

```sh
go mod tidy
```

### 3️⃣ Setup Environment Variables

Rename `.env.example` to `.env` and configure your database connection.

### 4️⃣ Run the Application

```sh
go run main.go
```

### 5️⃣ Run with Docker (Optional)

```sh
docker-compose up --build
```

## 📖 API Documentation

After running the project, open **Swagger** docs at:

```
http://localhost:8080/swagger/index.html
```

## 🎯 Features

✅ RESTful API with Clean Architecture  
✅ Secure authentication & authorization  
✅ Database integration with GORM  
✅ Swagger API Documentation  
✅ Scalable and modular structure  
✅ Docker support for easy deployment

## 🏆 Contribution

Feel free to fork this project, submit issues, and make pull requests! Let's build something awesome together. 💡🔥

---

💻 **Author:** [Bagus Rian Kusuma](https://github.com/BagusRianKusuma)  
📧 **Contact:** Bagusriankusuma3@gmail.com  
📌 **Follow me for more updates!** 🚀
