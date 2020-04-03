package game

func (p *PlayerState) AddToBattlefield(cards CardList)        {}
func (p *PlayerState) RemoveFromBattlefield(cards CardList)   {}
func (p *PlayerState) AddCounters(updates map[string]Counter) {}
func (p *PlayerState) Reveal(cards CardList)                  {}
func (p *PlayerState) Draw(library CardList)                  {}
func (p *PlayerState) Shuffle(cards CardList)                 {}
func (p *PlayerState) Discard(card CardList)                  {}
func (p *PlayerState) Fetch(card Card)                        {}
func (p *PlayerState) AddToLibrary(card Card, pos int)        {}
