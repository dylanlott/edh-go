#!/bin/bash

# this script starts a database for local dev.
docker run -d \
    --name edhgo-postgres \
    -e POSTGRES_PASSWORD=edhgodev \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v data:/var/lib/postgresql/data \
    postgres