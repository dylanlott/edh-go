import gql from 'graphql-tag';

export const stackQuery = gql`
query($gameID: String!) {
	games(gameID: $gameID) {
    stack {
      Name
      Text
    }
	}
}
`

export const stackMutation = gql`
  mutation($inputGame: InputGame!) {
    updateGame(input: $inputGame) {
      Stack {
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Name
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
    }
  }
`

export const createGameMutation = gql`mutation ($inputGame: InputGame!) {
  createGame(input: $inputGame){
     id
    created_at
    stack {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID  
    }
    turn {
      Number
      Player
      Phase
    }
    players {
      GameID
      Commander {
        Name
        ID
      }
    }
  }
}`

export const boardstates = gql`
query($gameID: String!) {
  boardstates(gameID: $gameID) {
    Life
    Commander {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID  
    }
    Library {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID  
    }
    Graveyard {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID  
    }
    Exiled {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID   
    }
    Field {
      Name 
      ID 
      Tapped
      Flipped
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID  
    }
    Hand {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID 
    }
    Revealed {
      Name 
      ID 
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID 
    }
    Controlled {
      Name 
      ID 
      Tapped
      Flipped
      Colors 
      ColorIdentity 
      ManaCost 
      Power 
      Toughness 
      CMC 
      Text 
      Types 
      Subtypes 
      Supertypes 
      IsTextless 
      TCGID 
      ScryfallID 
    }
  }
}
`

export const updateBoardStateQuery = gql`
  mutation ($boardstate: InputBoardState!) {
    updateBoardState(input: $boardstate) {
      User {
        username
      }
      GameID
      Life
      Commander { 
        Name 
        ID 
        Tapped
        Flipped
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Library { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Graveyard { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Exiled { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Field { 
        Name 
        ID 
        Tapped
        Flipped
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Hand { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Revealed { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Controlled { 
        Name 
        ID 
        Tapped
        Flipped
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      } 
    }
  }
`

export const boardstatesSubscription = gql`
subscription($boardstate: InputBoardState!) {
  boardUpdate(boardstate: $boardstate) {
    GameID
    Life
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
    Field {
      Name
    }
    Hand {
      Name
    }
    Revealed {
      Name
    }
    Controlled {
     Name
    }
  }
}
`

export const selfStateQuery = gql`
  query($gameID: String!, $userID: String) {
    boardstates(gameID: $gameID, userID: $userID) {
      User {
        username
      }
      Life
      Commander { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Library { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Graveyard { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Exiled { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Field { 
        Name 
        ID 
        Tapped
        Flipped
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Hand { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Revealed { 
        Name 
        ID 
        Tapped
        Flipped
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      }
      Controlled { 
        Name 
        ID 
        Colors 
        ColorIdentity 
        ManaCost 
        Power 
        Toughness 
        CMC 
        Text 
        Types 
        Subtypes 
        Supertypes 
        IsTextless 
        TCGID 
        ScryfallID 
      } 
    }
  }
`