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
export DEFAULT_PORT=
export REDIS_URL=
export DATABASE_URL=
export LOG_LEVEL=
export JWT_SECRET=
```

Then run the server with our Makefile.
The server will attempt to run all migrations and then start up. 
If it can't run migrations, it will rollback the database and noisily fail. 

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
How to connect to a postgres instance inside of docker
https://stackoverflow.com/questions/37694987/connecting-to-postgresql-in-a-docker-container-from-outside

How to import an SQL dump into Postgres
https://stackoverflow.com/questions/6842393/import-sql-dump-into-postgresql-database

Make sure when you rows.Scan() you don't point it at a nil value
https://stackoverflow.com/questions/44670212/scan-sql-null-values-in-golang/46753197

# Deployment 

`heroku static:deploy` will deploy the vue app to heroku. 
You can also run `npm run deploy` and this will run the same thing.

## Server
Heroku will build and deploy any push to GitHub.