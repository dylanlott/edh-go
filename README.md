# EDH-Go
> A board state tracker for Magic: The Gathering

[![Build Status](https://travis-ci.org/dylanlott/edh-go.svg?branch=master)](https://travis-ci.org/dylanlott/edh-go)

This library is meant to be a board state tracker (_not_ a game engine) for Magic: The Gathering.

The idea is to facilitate online play without a restrictive game engine, making the game more freeform and casual-competitive without requiring the ton of work to programmatically model a full Magic: The Gathering game engine.

## SQLite database
You'll need to download the SQLite database of MTG cards from [here.](https://mtgjson.com/downloads/all-files/)

Then you'll need to add it to `/persistence` and name it `mtgallcards.sqlite`

Your database should be working after that. I use DBeaver for local database introspection.

# Testing

`go test -v ./...`

I usually use GoConvey to run tests during local development.
