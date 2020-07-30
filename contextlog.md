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
