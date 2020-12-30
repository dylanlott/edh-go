import gql from 'graphql-tag';

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
        Username
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
    User {
      Username
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

export const selfStateQuery = gql`
  query($gameID: String!, $userID: String) {
    boardstates(gameID: $gameID, userID: $userID) {
      User {
        Username
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

// gameUpdateQuery powers the TurnTracker and Opponents components.
export const gameUpdateQuery = gql`
subscription($game: InputGame!) {
  gameUpdated(game: $game) {
    ID
    PlayerIDs {
      Username
      ID
    }
  }
} 
`

export const gameQuery = gql`
query ($gameID: String) {
  games(gameID: $gameID) {
    ID
    PlayerIDs {
      Username
      ID
    }
    Turn {
      Player
      Phase
      Number
    }
  }
}
`