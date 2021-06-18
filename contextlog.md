# EDH-Go Context Log

Established: 23 July 2020

## References 
[Vue-Apollo Subscription Examples](https://hasura.io/learn/graphql/vue/realtime-feed/2-sync-todo/) Super helpful and better than the actual vue apollo docs themselves. 


## Table Of Contents

1. What is this?
2. The Context Log
3. Module Documentation
4. To Dos and General Notes

Note: To dos are ephemeral in terms of where they fall in the log and may be moved around. If I add it to the to do list in one date entry, I might carry it with me
down to another date entry to remind myself to do it. This is essentially a living document to be viewed through the lens of git as a way to track the evolution of
this project log. Consider this like patchnotes or a changelog but tied to dates instead of deployments or versions.

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

_EDH-Go should aim to be_

- Intuitive to use
- Resource efficient
- Light

### What is EDH-Go

EDH-Go is going to be a boardstate emulator. It is not meant to enforce rules, merely aid in representing and tracking them.
That being said, there are some rules we can and should enforce - such as deck size, deck legality, turn orders, etc...

### What is EDH-GO _not_? 

EDH Go is _not_ meant to be

- A rules engine. We're trying to facilitate online play with the spirit of the social contract.
- Strict. We're trying to embrace Rule 0 in a digital way that's still accurate to real table games.
- Official. We're not trying to enforce rules, we're trying to showcase games.

Any rules imposed by our app should be (eventually) able to be changed.

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
https://gist.github.com/gorbypark/91917cf19d1245f52e025b42508344b1 For vue apollo subscriptions
https://en.wikipedia.org/wiki/Magic:_The_Gathering - Card sizes should be proportionally sized to the real life dimensions. Aspect ratio of "approximately 63 × 88 mm in size (2.5 by 3.5 inches)"
https://css-tricks.com/scaled-proportional-blocks-with-css-and-javascript/ We can use an approach like this to maintain size and allow users to scale their size up or down and maintain their aspect ratio.


*Task List*
- [*] Account for Turn Ordering and tracking in Game subscriptions.
- [x] Write a GraphQL resolver for returning only opponent boardstates. Update: Instead, I'm just going to handle this at the component level by requesting them individually to be bandwidth optimized.
- [*] Only reference players by ID and username on Games.
- [*] Decouple BoardStates from Game model

*Notes*
One of the real benefits of GraphQL is that the client can change their appetite themselves. They can take in the whole massive data object or they can build more complex interactions with specific pieces of data from different areas choosing to return only what they need.

The downside of this approach is that the long-tail amplification can be pretty bad if any part of the chain is slow. For queries that all run on the same cluster, this means that individually slow queries are going be harm all your graphql returns. For microservice architectures using GraphQL to tie services together, this is even more important as a single slow microservice could bring everything down much more.

InputCreateGame has BoardStates tied to it. In the future, we should remove this coupling and have the client create a BoardState and a Game independently of each other and then go to the Board route to start collecting data. This will keep BoardStates represented more cleanly by only a Player's Username and ID fields rather than toting around the BoardStates on the Game object.

GraphQL modeling can be a real foot-meet-shotgun problem, however. One change of a value can have sweeing consequences, so make sure you spend appropriate time decoupling data wherever you can before committing it to code. For example, I originally had BoardStates tied to Games as a required field, and this made sense initially, but after working on some improvements to the game, I decided I wanted BoardStates to be a separate entity from Games to decouple them and instead the Game model should only reference User's, and if a BoardState was needed it should be queried as a separate resource. However this meant that I had to refactor the create game logic because of how BoardStates and Games were created, and I had to refactor how Games were queried and updated for subscriptions, too. This is definitely my fault, but it's something you need to be hyper aware of when designing schema for your own personal apps.

At some point I should update the GraphQL schema to remove `Input` from types and instead use `Create` or `Update` based on what type the resource is going to be applied. For example `InputDeck` should be `CreateDeck` and `InputTurn` should be `UpdateTurn` because Turns are always idempotent and so can be treated like a PUT operation, but decks are only created once and thus should be treated like a POST operation.

Need to add a loading animation to the Game creation page. 
Join Game page needs to be created.

*Where I Left Off*
Test that Game updates and subscriptions are working so that BoardStates for Opponents can be written. 
We're getting closer to implementing Opponent Boardstates. 

_Opponent Component_ This should be fed a list of opponent IDs and then use those ID's to generate BoardState components for them that can't be updated or effected by the Self user.

### 16 Sep 2020

Currently wiring up the turn tracker to persist to the backend state.

* cards() and card() graphql endpoints should be refactored into one single endpoint 
* card searches in backend should utilize WHERE IN queries
* working on wiring up the turn tracker to game updates so they're reflected in game data accordingly.

*Opponent Component Design Notes*
Opponent comp. will need to listen for updates to the Game and detect if a player joins, adjusting accordingly. 
When a new player joins, we should launch a notification of some sort. 
This is where Ready checks should be gathered and displayed as well. 

### 18 Sep 2020

I've got a cute little state bug brewing somewhere. 

As far as I can tell, when the Boardstate gets refreshed at some point, there's a race condition between what the state gets set to by selfstatequery and what updateBoardState sets it to. If there's nothing set, the state breaks because it gets into a place where it doesn't know what it's supposed to be. Somewhere along the line we're resetting our own BoardState and so when it goes to grab it again, it's not there in its correct form.

Note: Maybe it's time to start looking into a rough state implementation with vuex or something.

I have officially hit the point where I have run into the exact moment where I needed a state system and almost considered writing my own setup for it but instead, will just use vuex like a good, ethical person. (Please, God, not another state management library.)

*Updated* 
I think the issue was actually in a variable not being set with the correct case. I'm attempting to switch all queries to capital case to make sure.

Okay, so `boardstates()` subscription in Board.vue is sending a `GameID` of `undefined` which is overwriting the Game in the backend and simultaneously causing it to be 404'd. I think it is being mounted before the route is being computed and thus is sending off the Game query before the GameID has been returned from createGame, which tells the Board component what ID to load. 

*Tasks*
- [ ] Setup rough GET patterns for something with VueX and call it a day.
- [ ] Get opponent state fetching boardstates correctly
- [ ] Investigate the Gavi decklist I've been using - I'm noticing it generated a decklist of 112 cards here, need to make sure that's still functioning right.

References for GraphQL VueX Implementation:
* https://markus.oberlehner.net/blog/combining-graphql-and-vuex/

### 19 Sep 2020

Turns out, there were several bugs. Game objects were being overwritten by the UpdateGame function in both the mutation and subscription endpoints. But to make it worse, the front end was also send incomplete or incorrect payloads so PlayerIDs weren't being persisted correctly, causing BoardUpdates to not be fetched correctly.

There's no way this is going to scale without a better way of querying the boardstate. There should probably be a Board and a Game store in VueX, with mutations and actions for both, and they should each have a connection to User models.

*Tasks*
> Completed in November 2020
- [x] Add initial setup for state management
- [x] Factor out the rest of the gql queries

### 12 Dec 2020
- Will eventually need a stronger support for combining two structs of different types for GraphQL. Something like https://play.golang.org/p/UBCq0waIEe should eventually be used.
- Opponent boardstate loading was broken at some point, so I need to fix that.
- Game updates should be separated more cleanly from Board updates
- Game updates should be made more granular.

### 21 Dec 2020 
When I query for a game, there's no associated boardstate. Need to make sure that boardstates are accessible by GameID as well.

Figured out that the `subscribeToMore` method in the `game()` apollo query in Board.vue was clobbering the PlayerIDs state of the subsequent `Game` query requests, meaning that all of those queries had `[]` set as their PlayerIDs, causing a whole bunch of issues. 

I fixed that, but now I need to make it so that the `game()` query in Board.vue properly fetches the game's PlayerIDs array and queries based on that. Right now, we have a chicken and the egg problem where the Board.vue tries to subscribe to the Game object but can't because the PlayerIDs haven't loaded yet. Maybe turning this into a method that we call in `created()` lifecycle hook would fix the problem?

Ultimately, I need to introduce a state management solution to the app via VueX to fix these issues. This will be a pretty large refactor but I suspect would lend itself heavily to the overall quality and reliability of the app.

`add-vuex` branch has some work that does exactly this on it.

### 22 Dec 2020
**TODO**: 
- [x] Need to start grapplevining VueX into the app and get opponent boardstates working.
- [x] Get opponent boardstates pulling up however we have to manage it
- [x] Finish the Join Game user flow

### 24 December 2020 
Christmas Eve! 

The app needs to set the opponent boardstates and susbcribe to updates. I have a rough query working now, it just needs to be refined to pull only the opponent states back.

### 27 Dec 2020
Working on the Join Game flow. Currently writing the join game mutation for the view with the same name.

Once this is done, I can perfect the Opponent boardstate subscription to correctly return Opponent boardstates upon `gameUpdated` events.

**TODO**
- [ ] `handleJoinGame` mutation working
- [ ] `Board.vue` subscription working with gameUpdated events.
- [ ] `JoinGame` function wired up

### 29 Dec 2020
Okay so while working on the Join Game flow, it occurred to me that it would be better served to just write a custom JoinGame service method through the GraphQL API to handle any of our custom logic. Joining a game is a special event and probably should be handled as such.

**TODO**
- [x] Regenerate GraphQL Schema
- [x] Write the JoinGame method in games.go
- [x] Wire up the `handleJoinGame` method jin `JoinGame.vue` method to poin to that endpoint instead.
- [x] Add tests for JoinGame and CreateGame // CreateGame tests are in progress

### 30 Dec 2020 
Working on Join Game functionality still. Found a bug in decklists and library creation where tab characters werent' being processed right. Table testing the CreateGame function exposed the bug, so that was pretty sick. Now the decklist will be a regression test itself for handling tab characters. Neat!

Working on the front end for the Join Game flow now, since the backend has been mostly figured out thanks to the tests. This is a personal lesson in how much faster front end development can go with a proper backend test suite. I already know the exact payloads to send, and GraphQL makes it fast to wire up a very precise request and turn it into a feature.

#### The Translator Problem 
I came up with two implementation ideas, and decided to puruse the first one
because I felt it was the better concept.

I'm quite happy with how this Polyglot interface is coming together.
It's easy to test, it's functional, it will be thread safe, and it's an easy 
way to handle the deeper intricacies that we can face with Board and Game state 
changes that will require subtle and specific handling.

**TODO**
- [x] Wire up `JoinGame` mutation to `handleJoinGame()` function.
- [x] Test the subscriptions on `Board.vue` to see if we're even listening for Game Update events.

### 31 Dec 2020 
New Years Eve

Backend is working much more smoothly now that I've discovered a cute little json encoding
hack to get around type issues between GraphQL and Go.

### 9 Jan 2021 
Carrying over the TO DO list from New Years Eve 2020 entry. Need to work on the Join Game query for an Opponent now.

**TODO**
- [x] Wire up JoinGame mutation to Front end 

### 14 Jan 2021

Added vue-cookies to solve the user ID issue on the front end. This will get around auth for now but we'll need
a more reliable way to sign up users and track them. We can keep it lite for now though.

### 20 Jan 2021
Need to look into best way to handle auth. 

** TODO ** 
- [ ] Add json/encoding hack to the Boardstates logic
- [ ] Get Game subscription working
- [ ] Test boardstate subscriptions for opponents in other views
- [ ] Work on Decktester feature
- [ ] Add vuex state for BoardState
- [ ] Add vuex state for Game
- [ ] Tie in to the `gameUpdated` subscription events so that we can detect game changes on the front end.
- [ ] Persist the Game directory to Redis


13 Feb 2021
===========

*Notes*
- Need to remove Directory from GraphQL server struct and make it use Redis or whatever the KV store is using. 
- Auth needs to be added and we need to figure out a graceful way to attach user data to our requests.
    - I think for auth, we can use something like [this](https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/authentication.md)
- Need to write more tests to shore up backend functionality.
- The app has a lot of cruft code I should get rid of right now. 
    - I should slim things down and remove all the mental overhead of it so that I can focus on a launch priority better.


15 Feb 2021
===========

- [x] Setup docker compose and migrations. 
- [ ] Get sign up and login working 
- [ ] Pass authentication info to server context
- [ ] Make games aware of user contexts
- [ ] Simplify game handling & modeling 
    - [ ] Remove redundant user info in boardstates and games
    - [ ] Make it so games have no concept of users and only see boardstates with users attached


27 Feb 2021
===========

My test suite has made refactoring and changes a lot faster, and for future me: I Should really make sure I start all personal projects
with a test harness. It makes it much easier to developer on longer time scales with tests. It's also much easier to setup pipelines and 
make small changes with a test harness. I can be much more confident about my code and bug fixes when I have passing tests.

*Notes*
- Working on migrations, they should be done soon.
- Having to finish auth tests, but migrations were a blocker to that because we had no way to reliably update the sqlite3 database.
- Once auth tests are done we can wire up the front end to accept different users and load multiple users into a single game.
- Need to figure out the fast-path to MVP and take it. I keep getting side tracked on small quality of life features. 
    - While they are helpful, we're coming up on a year after we started this project and I want to launch a beta.
- I should write down a launch plan

7 Mar 2021
===========

Postgres is working with migrations and we're back to a rapid test driven development cycle. 

Need to add postgres support for JSON and then switch out Redis with Postgres.
https://www.cloudbees.com/blog/unleash-the-power-of-storing-json-in-postgres/

The game log also needs to just push to Postgres stores. 
Game log should be committed on each "resolution". But we'll need to figure out targeting to make resolution work.

First up is getting auth working. 


13 March 2021
==============

* Spruced up my terminal with powerlevel10k.
* Got the Signup tests passing.
* Adding login and auth tests 

5 April 2021
============

* Started major refactor of the state management and front end interaction 
    * Completely blew away all of the component state management and instead force everything to go through vuex actions.
* There's some bugs in the Game update logic
    - I should probably write some tests for this part rather than trying to hand-test it.


12 Apr 2021
===========

* Need to fix a CORS issue that got introduced somewhere in our updates. Not sure how or why but it did. 
* Boardstate subscriptions are up next. 
    * Once we have boardstates updating we can start to build out the components of the app. 
    * We can focus on just battlefield and hand for now.

16 APr 2021
===========

GF is out of town, so I'm trying to get some serious time put into the app for a production beta push. 

I've been doing some research as well: 
https://www.npmjs.com/package/magic-card-parser

General Google Query: https://www.google.com/search?client=firefox-b-1-d&q=parse+mtg+card+text+ 

https://github.com/grilix/mtg-parser

https://github.com/Zannick/demystify

Here's a good Blog post https://hudecekpetr.cz/a-formal-grammar-for-magic-the-gathering/
And an accompanying tool https://soothsilver.github.io/mtg-grammar/


A reddit post with some discussion on the matter 
https://www.reddit.com/r/magicTCG/comments/ho69ph/parser_for_magic_cards/

Which led me here https://github.com/rmmilewi/mtgcompiler 

Debugging notes: 
* Something is happening with Games being saved where a player that was added after the fact disappears after about a second. 
* Need to write tests for Boardstate and Game methods, so that we have some idea of what's happening.
* Need to setup tests for how the game will look in real life - with several people subscribing to each other's boardstate. Write tests for that and how it handles it.
* Need better test driven development all around 
* Need to include the card files for the production server deployments, too. 
    * Write a script that grabs the latest MTGJSON files and loads them down directly to our repo for easier deployments.


*GameUpdated Function*
Okay so we have better tests now. 
I had a realization about using GQLGen. 

You treat the mutation functions like the setters. They act on the persistence layer and actually cause the mutation. 
You treat the subscription/notification/posted functions like emitters. They are only taking a value and pushing it to the rest of the listeners. They don't act on the database or persist anything.

For example, UpdateGame should act on the database, and then it should
_call_ GameUpdated. GameUpdated should then independently read from the database or emit the object that was passed to it.

I think in most cases it makes sense to have it take an updated value, rather than making it access persistence a second time. Cuts down on concurrency bugs and doesn't require another lock on the persistence layer.

## Apr 19, 2021

Flaky tests are caused by migrations m.Up() and m.Down() behaving unexpectedly. 
I think we need to error handle better on the migration run. 
But I still can't figure out why half the time it correctly applies migrations and other times it doesn't.



## 30 APR 2021

Game subscription is being updated correctly, however the Game is either
not being stored correctly or not being retrieved correctly. 

We need to make sure that we only store and retrieve game data and board
state data from redis. 

This has the benefit of giving us hard restart persistence through Redis.
This has the drawback of requiring transaction handling for redis updates.

## 1 May 2021

Okay so basically decided I need to completely remove any concept of the Directory from the app to clear up any confusion around where or when games and boardstates are being stored and retrieved

I'll be working on that refactor today, addressing the points from yesterday's log entry while removing the Directory from graphQL server.

NB: https://en.wikipedia.org/wiki/Hexspeak

* Fixed issues in Game update lifecycle
* Removed the entire concept of the Game Directory from the app 
* Made all Game checks go through redis instead of the Directory 
   * The disconnect between Directory and Redis was already causing issues

Future Ideas: 
* Might be worth treating the Game Log events as a PubSub through Redis
* Would handle performance pressure easier that way 
* But would couple us more tightly to Redis
* Might not even be necessary for a while if we just interact directly with Postgres
* Another option could be a simple JSON serverless AWS function

### Daily Notes
Ran a deploy today after I got all the tests passing. This is a pretty sturdy 
dev environment for a hobby project, and I have a pretty high confidence in doing deploys right now. 

Next up: Need to update BoardStates UI and add back the gameplay interface.


6 May 2021
==========

Yesterday, I got realtime updates working for the Self boardstate! Now boardstates are able to be updated 
on the server, and broadcast back down to all clients.

The key to our architecture for this is that we:
1. Always push to the server to update state.
2. Always pull from the server for our local state.

Additionally, I hooked up the Card elements and all the board elements like Battlefield, Hand, and Library
up to the new state components. 

In the server, I fixed two race conditions with goroutines that were causing flakey state updates in 
both Game and BoardState persistence. The key takewaway from this was that 

**Next Steps**
- Need to fix the bug in the Turn Tracker where Turn Number does not increase.
- Clean up Card layouts and looks
- Add an error toast pop up
- Make Join functions work correctly
    - This is probably an issue with Cookies and Authentication

8 May 2021
==========

- Deployed the vue app as a Heroku static build 
- Got automatic deploys working 

9 May 2021
==========

Auth needs to be added now to the front end, since we have working boardstates, and I think that the 
issues I'm seeing with Join functionality are related to the authentication issues I'm having.

* Added the signup logic to the app
* Had an issue with websocket urls and firefox throwing a fit about them.
    * https://github.com/apollographql/subscriptions-transport-ws/issues/332 
    * https://stackoverflow.com/questions/11768221/firefox-websocket-security-issue/12042843#12042843 
* Updated the build process to account for environment variables. 

14 Jun 2021
===========

* Need to simplify and condense the boardstate update and fetching logic. 
* It needs to be bi directional, with the boardstate components only getting display information from pulls, and the back end only receiving mutations from pushes.
* I stashed the work I had but I suspect it was mostly nonsense. There's a lot of changes in that stash though. Need to look through them and figure it out.

18 Jun 2021
===========

* Working on real time updates for Game 
    * I need to get one realtime component working, and game object seems to be the best target for this.
* I fixed a test in the JoinGame logic and added better assertions to that logic. 
* I'm in the process of overhauling the UI to correctly work with the realtime logic 
* Turns out, real time data is a hard problem.
* I think the server is in a decent place right now, but the UI needs a lot of work. 
