<div align="center">
  <img src="https://gitlab.com/xkrishguptaa/go-todo-api/-/raw/main/assets/logo.png" height="100px" width="100px" />
  <br />
  <h1>Go Todo API</h1>
  <p>Supports retrieval, creation, modification and deletion of Todos</p>
  <p>
    <a href="https://go.postman.co/collections/30796221-cdf39375-8b13-4130-95c1-d0db3687e53e"><img src="https://img.shields.io/badge/View%20Postman%20Collection-2965F1?style=for-the-badge" alt="View Postman Collection" /></a>
    <a href="https://hub.docker.com/r/xkrishguptaa/go-todo-api"><img src="https://img.shields.io/badge/Vie%20On%20Dockerhub-2965F1?style=for-the-badge" alt="View On Dockerhub" /></a>
  </p>
</div>

## 📖 Introduction

This is a simple REST API for a Todo application. It supports retrieval, creation, modification and deletion of Todos.

It is written in Go and uses the [Gin](https://github.com/gin-gonic/gin) framework. I was trying to learn Go and this is my first project in Go.

The following is the tutorial I followed to build this API:

[Build a Rest API with GoLang](https://www.youtube.com/embed/d_L64KT3SFM)

## 📦 Built With

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Docker](https://www.docker.com/)

## 🚀 Getting Started

### 🐳 Using Docker

The easiest way to get started is to use Docker. You can pull the image from Dockerhub and run it using the following commands:

```bash
docker pull xkrishguptaa/go-todo-api
docker run -p 3000:3000 xkrishguptaa/go-todo-api
```

### 🛠️ Building From Source

You can also build the API from source. You need to have Go installed on your system. You can then clone the repository and run the following commands:

```bash
go mod download
go build -o main .
./main
```

## 📚 API Documentation

The API documentation is available on [Postman](https://documenter.getpostman.com/view/30796221/2s9YXfcPMz).

## 📝 License

This project is licensed under the GNU GPL v3 License - see the [LICENSE](LICENSE.md) file for details.

## 🙏 Acknowledgments

- [Gopher Icon](https://github.com/egonelbre/gophers/blob/master/vector/superhero/standing.svg)
- [Build a Rest API with GoLang](https://www.youtube.com/watch?v=d_L64KT3SFM)
