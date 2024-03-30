# 04 Accessing a relational database
Tutorial from https://go.dev/doc/tutorial/database-access

## Instuctions
### 1. Bringing the DB up
Before running the code, you need to bring up the database, for that you'll need to execute form the `db` folder:

```zsh
docker-compose up
```

### 2. Running the Go app
Once the db is up and listening for connections, you can run your Go app. For that switch to `data-access` folder and run:

```zsh
export DBUSER=someuser
export DBPASS=somepass
go run .
```

(replace `someuser` and `somepass` by the db username and password specified in `./db/docker-compose.yml` file).
