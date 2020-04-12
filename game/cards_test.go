package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeckList(t *testing.T) {
	list, errors := NewDecklist(testdata)
	assert.Equal(t, len(errors), 0)
	assert.Equal(t, len(list), 3)

	for _, card := range list {
		assert.NotEqual(t, card.Name, "")
	}
}

const testdata = `
Karlov of the Ghost Council
Teysa, Envoy of Ghosts
Alela, Artful Provocateur
`

// TODO: Make sure this works
const exampledeck = `
Warlord's Fury
Goblin War Drums
Hordeling Outburst
Battle Squadron
Cleaver Riot
Goblin Diplomats
Goblin Wardriver
Relentless Assault
Makeshift Munitions
Vance's Blasting Cannons // Spitfire Bastion
By Force
Insult // Injury
Hazoret's Monument
Hellrider
Gratuitous Violence
Sulfuric Vortex
Goblin Dark-Dwellers
Goblin Sledder
Reckless One
Hellraiser Goblin
"Krenko, Mob Boss"
Ogre Battledriver
Lightning Volley
Hammer of Purphoros
Five-Alarm Fire
Arms Dealer
Fervor
Battle Hymn
Burn at the Stake
Puppet Strings
Quest for the Goblin Lord
Horde of Boggarts
"Ib Halfheart, Goblin Tactician"
Magewright's Stone
Dogpile
Mountain
"Ben-Ben, Akki Hermit"
Goblin War Strike
Goblin Lookout
Jandor's Saddlebags
Mob Justice
Volley Veteran
Trumpet Blast
Goblin Instigator
Goblin Cratermaker
Goblin Banneret
Raid Bombardment
Desperate Ritual
Burn Bright
Cavalcade of Calamity
"Krenko, Tin Street Kingpin"
Goblin War Party
Goblin Ringleader
"Squee, the Immortal"
Warstorm Surge
Heraldic Banner
Sol Ring
Goblin Charbelcher
Smelt
Skirk Prospector
Outnumber
Impact Tremors
Goblin Warchief
Goblin Oriflamme
Goblin Motivator
Goblin Fireslinger
Goblin Assault
Crash Through
Cathartic Reunion
`
