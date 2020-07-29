const updateBoardStateQuery = gql`
  mutation ($boardstate: InputBoardState!) {
    updateBoardState(input: $boardstate) {
      User {
        username
      }
      GameID
      Commander {
        Name
      }
      Library {
        Name
      }
      Graveyard {
        Name
      }
      Exiled {
        Name
      }
      Revealed {
        Name
      }
    }
  }
`

const getBoardstate = gql`
  query($gameID: String!) {
    boardstates(gameID: $gameID) {
      User {
        id
      }
      Library {
        Name
        ID
      }
      Graveyard {
        Name
        ID
      }
      Exiled {
        Name
        ID
      }
      Field {
        Name
        ID
      }
      Hand {
        Name
        ID
      }
      Revealed {
        Name
        ID
      }
      Controlled {
        Name
        ID
      }
    }
  }
`

const boardstateSubscription = gql`
  subscription ($boardstate: InputBoardState!) {
    boardUpdate(boardstate: $boardstate) {
      GameID
      User {
        username
      }
    }
  }
`

// export all of the queries we defined.
export default {
  boardstateSubscription: boardstateSubscription
}