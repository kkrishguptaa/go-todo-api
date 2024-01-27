> **NOTICE**: This `learning` project has been completed ✅ So I'm saying bye to it and archiving it! You can find my future projects on [GitHub](https://github.com/xkrishguptaa)

<div align="center">
  <img src="https://github.com/xkrishguptaa/go-todo-api/raw/main/assets/logo.png" height="100px" width="100px" />
  <br />
  <h1>Go Todo API</h1>
  <p>Supports retrieval, creation, modification and deletion of Todos</p>
  <p>
    <a href="https://go.postman.co/collections/30796221-e1bec2da-b843-4502-8057-c2a3cb46327c"><img src="https://img.shields.io/badge/View%20Postman%20Collection-2965F1?style=for-the-badge" alt="View Postman Collection" /></a>
    <a href="https://hub.docker.com/r/xkrishguptaa/go-todo-api"><img src="https://img.shields.io/badge/View%20On%20Dockerhub-2965F1?style=for-the-badge" alt="View On Dockerhub" /></a>
  </p>
</div>

## 📖 Introduction

This is a simple REST API for a Todo application. It supports retrieval, creation, modification and deletion of Todos.

It is written in Go and uses the [Gin](https://github.com/gin-gonic/gin) framework. I was trying to learn Go and this is my first project in Go.

## 📦 Built With

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/)

## 🌐 Check out Deployed

It is deployed at [go-todo-api-sa9e.onrender.com](https://go-todo-api-sa9e.onrender.com)

## 🚀 Getting Started

You need to create a `.env` file in the root directory of the project and add the following environment variables:

```bash
DATABASE_URL=<MongoDB Connection String>
DATABASE_NAME=<MongoDB Database Name>
```

Or you can set the environment variables directly.

Make sure that the MongoDB database is running and the connection string is correct. The database needs to have a collection named `Todos`.

### 🐳 Using Docker

The easiest way to get started is to use Docker. You can pull the image from Dockerhub and run it using the following commands:

```bash
docker run -p 8080:8080 --env-file .env xkrishguptaa/go-todo-api
```

### 🛠️ Building From Source

You can also build the API from source. You need to have Go installed on your system. You can then clone the repository and run the following commands:

```bash
go mod download
go mod verify
chmod +x ./scripts/build.sh
APP_NAME=go-todo-api ./scripts/build.sh
./bin/go-todo-api
```

## 📚 API Documentation

The API documentation is available on [Postman](https://documenter.getpostman.com/view/30796221/2s9YXfcPMz).

## 📝 License

This project is licensed under the GNU GPL v3 License - see the [LICENSE](LICENSE.md) file for details.

## 🙏 Acknowledgments

- [Gopher Icon](https://github.com/egonelbre/gophers/blob/master/vector/superhero/standing.svg)
- [Build a Rest API with GoLang](https://www.youtube.com/watch?v=d_L64KT3SFM)
