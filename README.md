# EDH-Go
> A Magic: The Gathering board state tracker built with GraphQL and Go.
 [![Build Status](https://travis-ci.org/dylanlott/edh-go.svg?branch=master)](https://travis-ci.org/dylanlott/edh-go)

## Running

Run server:

```
$ docker-compose up -d --build
```

Run Vue app:

```
$ cd frontend
$ npm run start
```

### Postgres 
We currently use SQLite3 for the card database, but this will probably change in the future as we grow. 
For now, these are resources we're using internally to work on the Postgres upgrade.

#### Documentation & Resources:
https://stackoverflow.com/questions/37694987/connecting-to-postgresql-in-a-docker-container-from-outside
https://stackoverflow.com/questions/6842393/import-sql-dump-into-postgresql-database
