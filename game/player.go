package game

/*
  player.go holds all of the board state updates for an individual player.
*/

func (p *PlayerState) CreateDeck(cards CardList)              {}
func (p *PlayerState) AddToBattlefield(cards CardList)        {}
func (p *PlayerState) RemoveFromBattlefield(cards CardList)   {}
func (p *PlayerState) AddCounters(updates map[string]Counter) {}
func (p *PlayerState) Reveal(cards CardList)                  {}
func (p *PlayerState) Draw(library CardList)                  {}
func (p *PlayerState) Shuffle(cards CardList)                 {}
func (p *PlayerState) Discard(card CardList)                  {}
func (p *PlayerState) DiscardAtRandom(card CardList)          {}
func (p *PlayerState) Fetch(card Card)                        {}
func (p *PlayerState) AddToLibrary(card Card, pos int)        {}
func (p *PlayerState) AddToGraveyard(card Card)               {}
func (p *PlayerState) AddToExile(card Card)                   {}
func (p *PlayerState) Scry(num int)                           {}
