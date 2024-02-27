## Description
A real neat application based on the [Alex Edwards](https://www.alexedwards.net/) book — ["Let's Go"](https://lets-go.alexedwards.net/). 
It was my starting point in Golang as I had never worked with this language before. At least so closely. I finished these book and app in February 2024.


## How to run
For first start the MySQL database in Docker container:
```sh
docker-compose up db -d
```

Then run the application:
```sh
# on default 4000 port
go run ./cmd/web 

# on specific port
go run ./cmd/web -addr=":9000"
```

For tests use:
```sh
docker-compose up test-db -d
go test ./cmd/web -v
```

## Environment
- MacOS Sonoma 14.2.1
- Go 1.21.6
- Docker 24.0.6
