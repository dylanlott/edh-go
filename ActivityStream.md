# Activity Stream

This document is a brain dump for how the ActivityStream should be setup. 
It will be a sensitive feature that could break easily, so I want to document it as well as possible. 

## 1. BoardState change is recognized on client side.

- Vue \$watch method detects any changes in the board.
- _**TODO**_: We need to watch for changes in the Game, including life and turns.

## 2. InputBoardState is generated from this and sent up to server

- When server receives this event, it should kick off a Goroutine that handles the event processing.
- Use something like <https://github.com/r3labs/diff>

## 3. This triggers other players board state updates to be triggered and updated.

- Other players get their updated board states sent to them

## 4. ActivityStream item is added and pushed to players as well with new boardstate.
