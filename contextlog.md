# EDH-Go Context Log

Established: 23 July 2020

## Table Of Contents

1. What is this?
2. The Context Log
3. Module Documentation
4. To Dos and General Notes

## What is this

This is the context log for this project. The idea is to completely dump my thought process behind the development of this application as a side project.
When I'm starting work for a session on the app, I'll jot down a goal for the day that I want to accomplish. During development I'll keep notes on what I'm working on through the day. At the end of the day, I'll write down a summary of what was accomplished, whether I met my goal for the day, and any relevant links I came across.

## Design Goals

Part of the point of this document is to provide context for the app - EDH-Go - and the vision I have for it. Ideally, if I was to completely stop this project, someone familiar with the tech stack should be able to jump in and start developing on this app just from this design and context log.

_EDH-Go should be:_

- Fast on any device
- Quick to setup
- Free to play
- Easy to use across all device sizes
- Format-agnostic

### What is EDH-Go

EDH-Go is going to be a boardstate emulator. It is not meant to enforce rules, merely aid in representing and tracking them.
That being said, there are some rules we can and should enforce - such as deck size, deck legality, turn orders, etc...

## Logs

### 28 July 2020

SelfState component needs to be passed the props from the Apollo query for selfstate but it's being weird about the mutate and update variables.

### 29 July 2020

Card data is now coming back and being loaded into the Board view component.
Need to get the card data pulling and loading correctly into the Card components.
Once we have the cards showing correctly, we can focus on getting board updates to work.

### 30 July 2020

Working on getting Card data to be shown correctly. Card art is going to be a consideration now. Need to figure out the best way to download the card art on the client side without pushing that heavy-lifting to the server.

- NB: Need to make sure that I'm pulling back ScryfallID from the database for card art images. EDHREC uses the same scryfall ID format so I think that will work for my use case.

Cards are now being populated with data from the server and I fixed the draggable issues.
Commanders are still able to be added to the 99, so that needs to be fixed.

- TO DO:
- [x] Fix card dragging on board view
- [ ] Wire up Scryfall client for card art eventually
- [ ] Add the join-game flow from the perspective of the 2nd, 3rd, and 4th players.
- [ ] Handle attaching equipment and auras to cards.
- [ ] Incorporate vuex into the app for better state management

### 31 July 2020

Figured out that the issue is that we are querying boardstates from the Directory but only updating boardstates from the mutation in the channels, so we need to update boardstates in redis and query them from redis.

This means I'll have to edit the game creation logic to store the initial boardstate in Redis and have the Game object reference that pointer instead of storing the player boardstate there.

- TO DO:
- [x] persist board states to redis
- [x] query board states from redis
- [x] edit game creation logic to store boardstate in redis so that refreshing the page doens't result in losing board state

- NB: We should probably log mutations from the server side instead of having the client send those mutations over the wire for the activity log

### 2 August 2020

There's an issue with board refreshes where the update mutation causes the card data to be lost and only card Name's to be persisted. I think it's an issue with how the board state mutation is passing info back and forth.

Board states are persisting and being fetched from Redis now. There's an interface that each GraphQLServer struct must fulfill called IPersistence that has two purposefully broad functions: `Get` and `Set`.

I wrote it this way so that board states can be persisted in any interface with no changes.

### 26 August 2020 

Figured out what the bugs are with the refresh of state - I think the issue has several factors. 
* The InputBoardState conversion function I wrote on the backend isn't properly converting and maintaining card properties from the input to the output. 
* The client implementation references both `boardstate` and `boardstates` at different points in the cycle. This is a code smell that we need to clean up the handling of the `self` component so that the separation between the `self` boardstate and the other player boardstates is clearer. I think this is the cause of the `ReferenceError` issues I keep having whenever an update occurs on the board. 
  
Fixing these two issues should give us the proper boardstate maintenance and persistence that we're looking for.


### 15 September 2020
Working on adding realtime boardstate updates to the app. 

*References*
https://gist.github.com/gorbypark/91917cf19d1245f52e025b42508344b1

*Task List*
* Account for Turn Ordering and tracking in BoardState subscriptions.
* Write a GraphQL resolver for returning only opponent boardstates.

*Notes*
One of the real benefits of GraphQL is that the client can change their appetite themselves. They can take in the whole massive data object or they can build more complex interactions with specific pieces of data from different areas choosing to return only what they need.

The downside of this approach is that the long-tail amplification can be pretty bad if any part of the chain is slow.

InputCreateGame has BoardStates tied to it. In the future, we should remove this coupling and have the client create a BoardState and a Game independently of each other and then go to the Board route to start collecting data. This will keep BoardStates represented more cleanly by only a Player's Username and ID fields rather than toting around the BoardStates on the Game object.