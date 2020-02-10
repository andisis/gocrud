# GO CRUD

Simple CRUD RESTful API written in Go and MySQL for RDBMS

## Prerequisites

**Install Go v 1.11+**

Please check the [Official Golang Documentation](https://golang.org/doc/install) for installation.

**Install Mockery**

```bash
go get github.com/vektra/mockery/.../
```

**Install MySQL on Docker**

```bash
docker pull mysql
docker run --name <countainer_name> --restart=always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<mysql_password> -d mysql
```

## Installation

**Clone this repository**

```bash
git clone git@github.com:andisis/gocrud.git
# Switch to the repository folder
cd gocrud
```

**Copy the `.env.example` file and make the required configuration changes in the `.env` file**

```bash
cp .env.example .env
```

**Run Application**

```bash
make run
```

## Install on Docker

**Create docker image**

```bash
make image
```

**Create docker container**

```bash
docker run --name <container_name> --rm -d --env-file <path/to/your/.env> -p 8080:8000 gocrud:latest
```

## Unit Testing

```bash
make test
```

If you want to see coverage in an HTML presentation (after the test) just run:

```bash
make coverage
```

## Folders

* `cmd/gocrud` - Contains `main.go`.
* `src/api` - Contains packages which are specific to your project.
* `src/database/config` - Contains database configuration.
* `src/database/query` - Contains database query.
* `src/model` - Contains any object structure.
* `src/router` - Contains router for the project.
* `bin` - Contains binary files from the application.

## Reference

* [Folder Explanation](https://github.com/golang-standards/project-layout)
* [Golang Clean Architecture](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)
* [Go Modules](https://blog.golang.org/using-go-modules)
* [Gorilla Mux](https://www.gorillatoolkit.org/pkg/mux)
* [Gorilla CORS](https://www.gorillatoolkit.org/pkg/handlers#CORS)
* [Logrus](https://github.com/sirupsen/logrus)
* [Mockery](https://github.com/vektra/mockery)
* [SQLMock](https://github.com/DATA-DOG/go-sqlmock)
* [Testify](https://github.com/stretchr/testify)

## Contributing

When contributing to this repository, please note we have a code standards, please follow it in all your interactions with the project.

#### Steps to contribute

1. Clone this repository.
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Submit pull request.

**Note** :

* Please make sure to update tests as appropriate.

* It's recommended to run `make test` command before submit a pull request.

## Deployment

Coming soon