package server

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestCreateGame(t *testing.T) {
	var cases = []struct {
		name  string
		input *InputCreateGame
		want  *Game
		err   error
	}{
		{
			name: "happy path creation",
			input: &InputCreateGame{
				Players: []*InputBoardState{
					{
						User: &InputUser{
							ID:       randomUserID(),
							Username: "shakezula",
						},
						Life:     40,
						Decklist: decklist(),
						Commander: []*InputCard{
							{
								Name: "Gavi, Nest Warden",
							},
						},
					},
				},
			},
			want: nil,
			err:  nil,
		},
		{
			name: "should allow for two commanders",
			input: &InputCreateGame{
				Players: []*InputBoardState{
					{
						User: &InputUser{
							ID:       randomUserID(),
							Username: "shakezula",
						},
						Life:     40,
						Decklist: decklist(),
						Commander: []*InputCard{
							{
								Name: "Gavi, Nest Warden",
							},
							{
								Name: "Jarad, Golgari Lich Lord",
							},
						},
					},
				},
			},
			want: nil, // TODO: make this test actually compare output results.
			err:  nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := getNewServer(t)
			result, err := s.CreateGame(context.Background(), *tt.input)
			if result.ID == "" {
				t.Errorf("games must have an ID")
			}
			if len(result.PlayerIDs) != len(tt.input.Players) {
				t.Errorf("failed to add correct amount of players to the game")
			}

			diff := cmp.Diff(tt.want, result)
			if diff != "" {
				t.Errorf("failed to create game: %s", diff)
			}
			if tt.err != nil {
				if diff := cmp.Diff(tt.err, err); diff != "" {
					t.Errorf("wanted error: %+v - got error: %+v", tt.err, err)
				}
			}
		})
	}
}

func TestJoinGame(t *testing.T) {
	var cases = []struct {
		name  string
		input InputJoinGame
		want  interface{}
		err   error
	}{
		{
			name: "join game happy path",
			input: InputJoinGame{
				ID:       "", // has to be set later on, so leave it blank
				Decklist: decklist(),
				BoardState: &InputBoardState{
					Commander: []*InputCard{
						{
							Name: "Gavi, Nest Warden",
						},
					},
				},
				User: &InputUser{
					ID:       randomUserID(),
					Username: "joingameid",
				},
			},
			err: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := getNewServer(t)
			input := InputCreateGame{
				Players: []*InputBoardState{
					{
						User: &InputUser{
							ID:       randomUserID(),
							Username: "shakezula",
						},
						Life:     40,
						Decklist: decklist(),
						Commander: []*InputCard{
							{
								Name: "Gavi, Nest Warden",
							},
						},
					},
				},
			}

			host, err := s.CreateGame(context.Background(), input)
			if err != nil {
				t.Errorf("failed to get host game: %+v\n", err)
			}

			t.Logf("joining host game: %+v\n", host)

			// hack to emulated client knowledge by setting input Game ID to host ID
			tt.input.ID = host.ID

			joined, err := s.JoinGame(context.Background(), &tt.input)
			if err != nil {
				fmt.Printf("JOIN GAME ERROR: %+v\n", err)
				if tt.err != nil {
					if diff := cmp.Diff(tt.err, err); diff != "" {
						t.Errorf("wanted: %+v - got: %+v\n", tt.want, err)
					}
				}
			}

			if tt.want != nil {
				// set result.CreatedAt to now for testing comparison, since
				// we only use it reference.
				joined.CreatedAt = time.Time{}
				if diff := cmp.Diff(tt.want, joined); diff != "" {
					t.Errorf("wanted: %+v - got: %+v\n", tt.want, diff)
				}
			}

			t.Logf("joined successfully: %+v\n", joined)
		})
	}
}

func TestGameUpdated(t *testing.T) {
	s := getNewServer(t)

	testGame, err := s.CreateGame(context.Background(), InputCreateGame{
		Players: []*InputBoardState{
			{
				User: &InputUser{
					ID:       randomUserID(),
					Username: "shakezula",
				},
				Life:     40,
				Decklist: decklist(),
				Commander: []*InputCard{
					{
						Name: "Gavi, Nest Warden",
					},
				},
			},
		},
	})
	if err != nil {
		t.Errorf("failed to create test game for TestGameUpdated: %s", err)
	}

	now := time.Now()

	var cases = []struct {
		name  string
		input InputGame
		want  interface{}
		err   error
	}{
		{
			name: "test happy path game updated",
			input: InputGame{
				ID:        testGame.ID,
				CreatedAt: &now,
				PlayerIDs: []*InputUser{
					{
						ID:       randomUserID(),
						Username: "joingameid",
					},
				},
				Turn: &InputTurn{
					Number: 1,
					Phase:  "start",
					Player: "shakezula",
				},
			},
		},
	}

	for _, tt := range cases {
		_, err := s.GameUpdated(context.Background(), tt.input)
		if err != nil {
			t.Errorf("failed to update game: %+v", err)
		}

		// get game from directory and then check it's values
		g, ok := s.Directory[tt.input.ID]
		if !ok {
			t.Errorf("failed to set game ID correctly into Directory")
		}

		t.Logf("game after update: %+v\n", g)
	}
}

func getNewServer(t *testing.T) *graphQLServer {
	cfg := Conf{}
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to find homedir: %s", err)
	}
	t.Logf("test server path: %+v", path)
	cardDB, err := persistence.NewSQLite("../persistence/AllPrintings.sqlite")
	if err != nil {
		t.Errorf("failed to open cardDB for games_test: %s", err)
	}

	appDB, err := persistence.NewSQLite("../persistence/db.sqlite")
	if err != nil {
		t.Errorf("failed to open appDB for games_test: %s", err)
	}

	s, err := NewGraphQLServer(nil, appDB, cardDB, cfg)
	if err != nil {
		t.Errorf("failed to create new test server: %+v", err)
	}

	return s
}

// NB: necessary to get a reference to a string because of this problem
// https://stackoverflow.com/questions/10535743/address-of-a-temporary-in-go

// randomUserID returns a stringified UUID that is used for Player generation
func randomUserID() *string {
	var id = uuid.New().String()
	return &id
}

// returns a csv formatted string using a premade decklist that tests up to the
// Theros set
func decklist() *string {
	var deck = string(`1,Floodwaters
	1,Astral Slide
	1,Possibility Storm
	1,Braid of Fire
	1,Complicate
	1,Miscalculation
	1,Glint-Horn Buccaneer
	1,Alhammarret's Archive
	1,Soothsaying
	1,Drannith Stinger
	1,"Vadrok, Apex of Thunder"
	1,Raugrin Crystal
	1,Flourishing Fox
	1,Neutralize
	1,Reconnaissance Mission
	1,Boon of the Wish-Giver
	1,"Yidaro, Wandering Monster"
	1,Arcane Signet
	1,Shark Typhoon
	1,"Rielle, the Everwise"
	1,Rooting Moloch
	1,Ominous Seas
	1,"Shabraz, the Skyshark"
	1,"Brallin, Skyshark Rider"
	1,Spellpyre Phoenix
	1,Crystalline Resonance
	1,Herald of the Forgotten
	1,Dismantling Wave
	1,Astral Drift
	1,Eternal Dragon
	1,Fluctuator
	1,Surly Badgersaur
	1,Savai Thundermane
	1,Lightning Rift
	1,Splendor Mare
	1,Decree of Justice
	1,Akroma's Vengeance
	1,Smoldering Crater
	1,Skycloud Expanse
	1,Secluded Steppe
	1,Remote Isle
	1,Prairie Stream
	1,Reliquary Tower
	1,Myriad Landscape
	1,Mystic Monastery
	1,Lonely Sandbar
	1,Izzet Boilerworks
	1,Exotic Orchard
	1,Forgotten Cave
	1,Drifting Meadow
	1,Azorius Chancery
	1,Psychosis Crawler
	1,Izzet Signet
	1,Azorius Signet
	1,Boros Signet
	1,Migratory Route
	1,"Niv-Mizzet, the Firemind"
	1,Starstorm
	1,Slice and Dice
	1,"Chandra, Flamecaller"
	1,Windfall
	1,Zenith Flare
	1,Sanctuary Smasher
	1,Unpredictable Cyclone
	1,Raugrin Triome
	1,Imposing Vantasaur
	1,Teferi's Ageless Insight
	1,Thriving Isle
	1,Ash Barrens
	1,"Gavi, Nest Warden"
	1,Forsake the Worldly
	1,Hieroglyphic Illumination
	1,Drake Haven
	1,Countervailing Winds
	1,Curator of Mysteries
	1,Irrigated Farmland
	1,Desert of the Fervent
	1,Desert of the Mindful
	1,Desert of the True
	1,Abandoned Sarcophagus
	1,The Locust God
	1,Sol Ring
	1,Command Tower
	3,Mountain
	4,Island
	3,Plains
	1,Swords to Plowshares
	1,Hallowed Fountain
	1,Shivan Reef
	1,Steam Vents
	1,Cloud of Faeries
	1,Nimble Obstructionist
	1,Sun Titan
	1,Valiant Rescuer
	1,Vizier of Tumbling Sands
	1,"Ephara, God of the Polis"
	1,Talisman of Creativity
	1,Commander's Sphere
	1,Radiant's Judgment
	1,Idyllic Tutor
	1,Cast Out
	1,New Perspectives
	1,Tectonic Reformation
	1,Decree of Silence
	1,Fierce Guardianship`)

	return &deck
}
