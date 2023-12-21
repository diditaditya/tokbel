# Tokbel

CSGA Final Project

## Getting Started

Clone this repository
```sh
$ git clone https://github.com/diditaditya/tokbel.git
```

Navigate to the folder and install the dependencies
```sh
$ cd tokbel
$ go mod tidy
```

Add a `.env` file containing the following field
```
DB_HOST=
DB_PORT=
DB_NAME=
DB_USER=
DB_PASSWORD=
JWT_SECRET=
```

To run in development mode
```sh
$ go run main.go
```

The documentation should be running at `/swagger/index.html`. Please refer to the [swag repo](https://github.com/swaggo/swag) for further information.