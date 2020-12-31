package server

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInputGameTranslator(t *testing.T) {
	var cases = []struct {
		name string
		from *InputGame
		want interface{}
	}{
		{
			// just to skeletonize and set things up
			name: "test translating empty input game",
			from: &InputGame{
				ID: "test123",
			},
			want: Game{
				ID:     "test123",
				Handle: nil,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			to := Game{}
			poly := &polyglot{}
			err := poly.Translate(&to, tt.from, InputGameTranslator)
			if err != nil {
				t.Errorf("failed to translate: %+v\n", err)
			}

			t.Logf("got here: %+v", to)
			if diff := cmp.Diff(tt.want, to); diff != "" {
				t.Errorf("failed to translate: %+v\n", diff)
			}
		})
	}
}
