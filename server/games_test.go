package server

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

func TestCreateGame(t *testing.T) {
	id := string("0xACAB")
	var cases = []struct {
		name  string
		input *InputCreateGame
		want  *Game
		err   error
	}{
		{
			name: "happy path creation",
			input: &InputCreateGame{
				ID: "deadbeef",
				Players: []*InputBoardState{
					{
						User: &InputUser{
							ID:       &id,
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
				Turn: &InputTurn{
					Player: "shakezula",
					Phase:  "pregame",
					Number: 0,
				},
			},
			want: &Game{
				ID: "deadbeef",
				PlayerIDs: []*User{
					{
						ID:       "0xACAB",
						Username: "shakezula",
					},
				},
				Turn: &Turn{
					Player: "shakezula",
					Phase:  "pregame",
					Number: 0,
				},
				Rules: []*Rule{
					{Name: "format", Value: "EDH"},
					{Name: "deck_size", Value: "99"},
				},
			},
			err: nil,
		},
		{
			name: "should allow for two commanders",
			input: &InputCreateGame{
				ID: "deadbeef",
				Players: []*InputBoardState{
					{
						User: &InputUser{
							ID:       &id,
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
				Turn: &InputTurn{
					Player: "shakezula",
					Phase:  "pregame",
					Number: 0,
				},
			},
			want: &Game{
				ID: "deadbeef",
				PlayerIDs: []*User{
					{
						ID:       "0xACAB",
						Username: "shakezula",
					},
				},
				Turn: &Turn{
					Player: "shakezula",
					Phase:  "pregame",
					Number: 0,
				},
				Rules: []*Rule{
					{Name: "format", Value: "EDH"},
					{Name: "deck_size", Value: "99"},
				}, // TODO: make this test actually compare output results.
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := testAPI(t)
			result, err := s.CreateGame(context.Background(), *tt.input)
			if result.ID == "" {
				t.Errorf("games must have an ID")
			}
			if len(result.PlayerIDs) != len(tt.input.Players) {
				t.Errorf("failed to add correct amount of players to the game")
			}

			diff := cmp.Diff(tt.want, result, cmpopts.IgnoreFields(
				Game{},
				"CreatedAt",
			))
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
	gameID := "0xbeefbeef"
	userID := "shakezulathemicrulah"
	userID2 := "abc123"

	var cases = []struct {
		name  string
		input InputJoinGame
		want  interface{}
		err   error
	}{
		{
			name: "join game happy path",
			input: InputJoinGame{
				ID:       "0xbeefbeef", // has to be set later on, so leave it blank
				Decklist: decklist(),
				BoardState: &InputBoardState{
					Commander: []*InputCard{
						{
							Name: "Gavi, Nest Warden",
						},
					},
				},
				User: &InputUser{
					ID:       &userID2,
					Username: "meatwad",
				},
			},
			err: nil,
			want: &Game{
				ID: gameID,
				Rules: []*Rule{
					{Name: "format", Value: "EDH"},
					{Name: "deck_size", Value: "99"},
				},
				Turn: &Turn{
					Phase:  "pregame",
					Number: 0,
					Player: "shakezula",
				},
				PlayerIDs: []*User{
					{
						ID:       userID,
						Username: "shakezula",
					},
					{
						ID:       userID2,
						Username: "meatwad",
					},
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := testAPI(t)

			_, err := s.CreateGame(context.Background(), InputCreateGame{
				ID: gameID,
				Players: []*InputBoardState{
					{
						User: &InputUser{
							ID:       &userID,
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
				Turn: &InputTurn{
					Player: "shakezula",
					Number: 0,
					Phase:  "pregame",
				},
			})
			if err != nil {
				t.Errorf("failed to get host game: %+v\n", err)
			}

			joined, err := s.JoinGame(context.Background(), &tt.input)
			// if err != nil {
			// 	if tt.err != nil {
			// 		if diff := cmp.Diff(tt.err, err); diff != "" {
			// 			t.Errorf("wanted: %+v - got: %+v\n", tt.want, err)
			// 		}
			// 	}
			// }

			if tt.want != nil {
				if diff := cmp.Diff(tt.want, joined, cmpopts.IgnoreFields(Game{}, "CreatedAt")); diff != "" {
					t.Errorf("wanted: %+v - got: %+v\n", tt.want, diff)
				}
			}

			t.Logf("joined successfully: %+v\n", joined)
		})
	}
}

func TestGameUpdated(t *testing.T) {
	userID := string("deadbeef")
	gameID := string("beefdead")
	now := time.Now()

	var cases = []struct {
		name  string
		input InputGame
		err   error
		want  *Game
	}{
		{
			name: "test happy path game updated",
			input: InputGame{
				ID:        gameID,
				CreatedAt: &now,
				PlayerIDs: []*InputUser{
					{
						ID:       &userID,
						Username: "shakezula",
					},
				},
				Turn: &InputTurn{
					Number: 0,
					Phase:  "pregame",
					Player: "shakezula",
				},
				Rules: []*InputRule{
					{Name: "format", Value: "EDH"},
					{Name: "deck_size", Value: "99"},
				},
			},
			err: nil,
			want: &Game{
				ID:        gameID,
				CreatedAt: now,
				PlayerIDs: []*User{
					{
						ID:       userID,
						Username: "shakezula",
					},
				},
				Turn: &Turn{
					Number: 0,
					Phase:  "pregame",
					Player: "shakezula",
				},
				Rules: []*Rule{
					{Name: "format", Value: "EDH"},
					{Name: "deck_size", Value: "99"},
				},
			},
		},
	}

	for _, tt := range cases {
		s := testAPI(t)
		testGame, err := s.CreateGame(context.Background(), InputCreateGame{
			ID: gameID,
			Players: []*InputBoardState{
				{
					User: &InputUser{
						ID:       &userID,
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
			Turn: &InputTurn{
				Number: 0,
				Phase:  "pregame",
				Player: "shakezula",
			},
		})
		if err != nil {
			t.Errorf("failed to create test game for TestGameUpdated: %s", err)
		}
		log.Printf("created testGame: %+v", testGame)

		resp, err := s.GameUpdated(context.Background(), tt.input)
		if err != nil {
			t.Errorf("failed to update game: %+v", err)
		}
		g := <-resp
		diff := cmp.Diff(g, tt.want)
		if diff != "" {
			log.Printf("[DIFF]: %s", diff)
			t.Errorf("GameUpdated() wanted: %+v - got: %+v", tt.want, g)
		}
	}
}

func TestUpdateGame(t *testing.T) {
	userID := string("deadbeef")
	userID2 := string("deadbeef2")

	type args struct {
		ctx context.Context
		new InputGame
	}
	tests := []struct {
		name    string
		args    args
		want    *Game
		wantErr bool
	}{
		{
			name: "should update game and alert gameChannels",
			args: args{
				ctx: context.Background(),
				new: InputGame{
					ID:        seedGameID,
					CreatedAt: &time.Time{},
					PlayerIDs: []*InputUser{
						{
							Username: "shakezula",
							ID:       &userID,
						},
						{
							Username: "meatwad",
							ID:       &userID2,
						},
					},
					Rules: []*InputRule{
						{Name: "format", Value: "EDH"},
						{Name: "deck_size", Value: "99"},
					},
					Turn: &InputTurn{
						Number: 3,
						Phase:  "the after party",
						Player: "meatwad",
					},
				},
			},
			wantErr: false,
			want: &Game{
				ID: seedGameID,
				PlayerIDs: []*User{
					{
						Username: "shakezula",
						ID:       userID,
					},
					{
						Username: "meatwad",
						ID:       userID2,
					},
				},
				Rules: []*Rule{
					{Name: "format", Value: "EDH"},
					{Name: "deck_size", Value: "99"},
				},
				Turn: &Turn{
					Number: 3,
					Phase:  "the after party",
					Player: "meatwad",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := testAPI(t)
			g, err := s.CreateGame(tt.args.ctx, *seedInputGame)

			got, err := s.UpdateGame(tt.args.ctx, tt.args.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.UpdateGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(Game{}, "CreatedAt"))
			if diff != "" {
				log.Printf("diff: %s", diff)
				t.Errorf("UpdateGame wanted: %+v - got %+v", tt.want, got)
			}

			// test game that was emitted
			game := <-s.gameChannels[g.ID]
			log.Printf("GAME: %+v", game)
			diff2 := cmp.Diff(game, tt.want, cmpopts.IgnoreFields(Game{}, "CreatedAt"))
			if diff2 != "" {
				t.Logf("[DIFF] %s", diff)
				t.Errorf("failed to emit game on channels correctly")
			}
		})
	}
}

func testAPI(t *testing.T) *graphQLServer {
	cfg := Conf{
		RedisURL:    "redis://localhost:6379",
		DefaultPort: 8080,
	}
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to find homedir: %s", err)
	}
	t.Logf("test server path: %+v", path)
	cardDB, err := persistence.NewSQLite("../persistence/AllPrintings.sqlite")
	if err != nil {
		t.Errorf("failed to open cardDB for games_test: %s", err)
	}
	kv, err := persistence.NewRedis("redis://localhost:6379", "", persistence.Config{})
	if err != nil {
		t.Errorf("failed to get kv from redis: %s", err)
	}
	appDB, err := persistence.NewSQLite("../persistence/db.sqlite")
	if err != nil {
		t.Errorf("failed to open appDB for games_test: %s", err)
	}
	s, err := NewGraphQLServer(kv, appDB, cardDB, cfg)
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

// Seed values for tests
var (
	seedUserID   string = "xfeedbeefx"
	seedGameID   string = "xdeadbeefx"
	seedUsername string = "shakezula"
)

// seedInputGame is a bare minimum game input that passes validation
var seedInputGame = &InputCreateGame{
	ID: seedGameID,
	Players: []*InputBoardState{
		{
			User: &InputUser{
				ID:       &seedUserID,
				Username: seedUsername,
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
	Turn: &InputTurn{
		Player: seedUsername,
		Phase:  "pregame",
		Number: 0,
	},
}
