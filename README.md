# gordo
A CoRE Resource Directory in Go!

# Dockerised DB

- Start up
```
docker-compose up -d
```

- Shut down
```
docker-compose down
```

- If you need the command line:
```
brew install postgresql
```

- Get a session (password is "123")
```
psql -p 15432  -U postgres -h localhost -d gordo
```



