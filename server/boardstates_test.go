package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			boardstateChannel, err := s.BoardstateUpdated(tt.args.ctx, created.ID, *tt.args.input.User.ID)
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
				boardstate := <-boardstateChannel
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
