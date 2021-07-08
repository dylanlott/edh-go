package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBoardState(t *testing.T) {
	userID := string("0xACAB")

	type args struct {
		ctx   context.Context
		input InputBoardState
	}
	tests := []struct {
		name    string
		args    args
		want    *BoardState
		wantErr bool
	}{
		{
			name: "should update boardstate and notify listeners",
			args: args{
				ctx: context.Background(),
				input: InputBoardState{
					GameID: seedGameID,
					User: &InputUser{
						Username: "shakezula",
						ID:       &userID,
					},
					Life: 38, // test life edits
					Commander: []*InputCard{
						{Name: "Gavi, Nest Warden"},
					},
				},
			},
			wantErr: false,
			want: &BoardState{
				GameID: seedGameID,
				User: &User{
					ID:       userID,
					Username: "shakezula",
				},
				Life: 38,
				Commander: []*Card{
					{Name: "Gavi, Nest Warden"},
				},
			},
		},
		{
			name: "should return an error if game does not exist",
			args: args{
				ctx: context.Background(),
				input: InputBoardState{
					GameID: "doesnotexist",
					User: &InputUser{
						Username: "shakezula",
						ID:       &userID,
					},
					Life: 38, // test life edits
					Commander: []*InputCard{
						{Name: "Gavi, Nest Warden"},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := testAPI(t)
			created, err := s.CreateGame(tt.args.ctx, *seedInputGame)
			if err != nil {
				t.Errorf("failed to create dummy game: %s", err)
			}

			// register the BoardState channel for the userID
			bch, err := s.BoardstateUpdated(
				tt.args.ctx,
				created.ID,
				*tt.args.input.User.ID,
			)
			assert.Equal(t, err, nil)
			got, err := s.UpdateBoardState(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.UpdateBoardState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("graphQLServer.UpdateBoardState() = %v, want %v", got, tt.want)
			}

			// assert on boardstate notifications if no error is expected
			if tt.wantErr == false {
				boardstate := <-bch
				if !reflect.DeepEqual(boardstate, tt.want) {
					diff := cmp.Diff(boardstate, tt.want)
					t.Logf("DIFF: %+v", diff)
					t.Errorf("UpdateBoardState failed to notify listeners")
				} else {
					t.Logf("successfully received boardstate from update: %+v", boardstate)
				}
			}
		})
	}
}

func Test_Boardstates(t *testing.T) {
	s := testAPI(t)
	s.CreateGame(context.Background(), *seedInputGame)
	p1 := string("0xTHEMICRULAH")
	p1commander := string("3269194123946123469")
	_, err := s.BoardstateUpdated(context.Background(), seedGameID, p1)
	if err != nil {
		t.Errorf("failed to register boardstate update channel for p1")
	}
	_, err = s.UpdateBoardState(context.Background(), InputBoardState{
		User: &InputUser{
			Username: "shakezula",
			ID:       &p1,
		},
		GameID: seedGameID,
		Life:   40,
		Commander: []*InputCard{
			{
				ID:   &p1commander,
				Name: "Kykar, Wind's Fury",
			},
		},
	})
	if err != nil {
		t.Errorf("failed to update boardstate: %s", err)
	}
	p2 := string("meatwad")
	p2commander := string("437823498234982349823498")
	_, err = s.BoardstateUpdated(context.Background(), seedGameID, p2)
	if err != nil {
		t.Errorf("failed to register boardstate update channel for p2")
	}
	_, err = s.UpdateBoardState(context.Background(), InputBoardState{
		User: &InputUser{
			Username: "meatwad",
			ID:       &p2,
		},
		GameID: seedGameID,
		Life:   38,
		Commander: []*InputCard{
			{
				ID:   &p2commander,
				Name: "Sidisi, Undead Vizier",
			},
		},
	})
	if err != nil {
		t.Errorf("failed to update boardstate for p2: %s", err)
	}
	type args struct {
		ctx    context.Context
		gameID string
		userID *string
	}
	tests := []struct {
		name    string
		args    args
		want    []*BoardState
		wantErr bool
	}{
		{
			name: "should return a boardstate with a gameID and userID",
			args: args{
				ctx:    context.Background(),
				gameID: seedGameID,
				userID: &p1,
			},
			want: []*BoardState{
				{
					GameID: seedGameID,
					User: &User{
						Username: "shakezula",
						ID:       p1,
					},
					Life: 40,
					Commander: []*Card{
						{
							ID:   p1commander,
							Name: "Kykar, Wind's Fury",
						},
					},
				},
			},
			wantErr: false,
		},
		// {
		// 	name: "should return all boardstate that match the gameID",
		// 	args: args{
		// 		ctx:    context.Background(),
		// 		gameID: seedGameID,
		// 		userID: nil,
		// 	},
		// 	want: []*BoardState{
		// 		{
		// 			GameID: seedGameID,
		// 			User: &User{
		// 				Username: "shakezula",
		// 				ID:       p1,
		// 			},
		// 			Life: 40,
		// 			Commander: []*Card{
		// 				{
		// 					ID:   p1commander,
		// 					Name: "Kykar, Wind's Fury",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			GameID: seedGameID,
		// 			User: &User{
		// 				Username: "meatwad",
		// 				ID:       p2,
		// 			},
		// 			Life: 38,
		// 			Commander: []*Card{
		// 				{
		// 					ID:   p2commander,
		// 					Name: "Sidisi, Undead Vizier",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.userID != nil {
				_, err := s.BoardstateUpdated(tt.args.ctx, tt.args.gameID, *tt.args.userID)
				assert.NoError(t, err, "failed to setup channel listener")
			}
			got, err := s.Boardstates(tt.args.ctx, tt.args.gameID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.Boardstates() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("diff: %s", cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(BoardState{})))
				t.Errorf("graphQLServer.Boardstates() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
