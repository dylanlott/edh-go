---
title: "App Update 001"
date: 2020-10-14
slug: "app-update-001"
description: "A quick update on the progress of the app."
keywords: ["app update", "admin", "pre-launch"]
draft: false
tags: ["admin", "app-update", "devlog"]
math: false
toc: false
---

* **Status**: alpha
* **Version**: `0.0.1-alpha`
* **Description**: Boardstates are self updating.

I had a lot of work in the last month, so I've neglected the app. However, I spent a lot of time on it this week and have addeda good chunk of new features. 

I've added shuffling on the server side when a game is created, wired the opponent board states up to a subscription service so that they receive incoming board updates from every other player, and added the ability to fetch and move cards around.

There are only a few things left before we launch to beta, namely: 

- More testing and bug fixing for multi-player modes
- Turn orders need to be represented correctly in turn counter
- Token creation and handling needs to be added
- Labels and counters for players and cards needs to be added
- Add support for fetching cards from libraries, graveyards, and exile.

Once these features are done, we'll be heading into beta.