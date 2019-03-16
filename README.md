# gordo
A CoRE Resource Directory in Go

## Dependencies

- [Go](https://golang.org/dl/)
- [Docker and Compose](https://hub.docker.com/editions/community/docker-ce-desktop-mac)
- (optional) PostgreSQL CLI (`brew install postgresql`)
- A CoAP CLI, e.g.: [CoAP-CLI](https://www.npmjs.com/package/coap-cli)

## Build gordo

- Get a copy of the repo:
```
go get github.com/thomas-fossati/gordo
```

- Move to the source directory:
```
cd $(go env GOPATH)/src/github.com/thomas-fossati/gordo
```

- Install dependencies:
```
go get -u ./...
```

- Build gordo:
```
go build
```

## Run gordo

- Start up the dockerized PostgreSQL backend:
```
docker-compose up -d
```

- Start gordo:
```
./gordo
```

- Run a query or two:
```
coap get 'coap://localhost/rd-lookup/res?ct=41'
[...]

coap get 'coap://localhost/rd-lookup/res?rt=temperature'
[...]
```

- To shut PostgreSQL down:
```
docker-compose down
```

# Fiddling around

- Get a session on the DB (password is "123")
```
psql -p 15432  -U postgres -h localhost -d gordo
```
