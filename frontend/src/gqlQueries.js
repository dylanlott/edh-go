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

export const getBoardstate = gql`
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

export const boardstateSubscription = gql`
  subscription ($boardstate: InputBoardState!) {
    boardUpdate(boardstate: $boardstate) {
      GameID
      User {
        username
      }
    }
  }
`

export const selfStateQuery = gql`
  query($gameID: String!, $userID: String) {
    boardstates(gameID: $gameID, userID: $userID) {
      Life
      User {
        username
      }
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