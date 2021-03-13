#!/bin/bash

# this script starts a database for local dev.
docker run -d \
    -e POSTGRES_PASSWORD=edhgodev \
    -e POSTGRES_USER=edhgo \
    -e POSTGRES_DB=edhgo \
    -p 5432:5432 \
    postgres