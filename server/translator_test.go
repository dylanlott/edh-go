package server

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInputGameTranslator(t *testing.T) {
	var cases = []struct {
		name string
		from *InputGame
		want *Game
	}{
		{
			name: "test translating basic input game",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			to := &Game{}
			poly := &polyglot{}
			err := poly.Translate(to, tt.from, InputGameTranslator)
			if err != nil {
				if diff := cmp.Diff(tt.want, err); diff != "" {
					t.Errorf("wanted: %+v - got: %+v", tt.want, err)
				}
			}

			if diff := cmp.Diff(tt.want, to); diff != "" {
				t.Errorf("wanted: %+v - got %+v", tt.want, to)
			}
		})
	}
}
