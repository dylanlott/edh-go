# EDH-Go
> A Magic: The Gathering board state tracker built with GraphQL and Go.
 [![Build Status](https://travis-ci.org/dylanlott/edh-go.svg?branch=master)](https://travis-ci.org/dylanlott/edh-go)

## Running

### Go Server
1. Prerequisites:
- Make
- Go v1.15

2. Set Environment variables 
Here's a template for setting the appropriate env vars in `~/.bashrc`
```
exprot DEFAULT_PORT=
exprot REDIS_URL=
exprot DATABASE_URL=
exprot LOG_LEVEL=
```

Then run the server with our Makefile.
The server will attempt to run all migrations and then start up. 
If it can't run migrations, it will rollback the database and fail. 

```
$ make run 
```

You can run the server as if it's in prod with this same config, so you 
can switch between local and prod as long as you've configured your environment
variables correctly.
### Front End 
Run Vue app:

```
$ cd frontend
$ npm run start
```

## Testing 
To run the full test suite, you will need a PostgresDB and Redis instance locally running. 

To run only the unit tests that don't need external sources, you can run 

```
$ make short 
```

Once you have a Redis and Postgres instance running locally, you can run your 
tests. You may need to change the config values in games_test.go to start your tests, or configure your environment to fit with the provided 

## Stack
Postgres for the application data 
SQLite is used for querying card data.
Redis is used for fast and efficient storing of BoardStates and Games.
Using an SQL based storage solution for the BoardStates would have been clunky.


## Documentation & Resources:
https://stackoverflow.com/questions/37694987/connecting-to-postgresql-in-a-docker-container-from-outside

https://stackoverflow.com/questions/6842393/import-sql-dump-into-postgresql-database
