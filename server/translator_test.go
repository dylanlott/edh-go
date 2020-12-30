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
			from: &InputGame{},
			want: &Game{},
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
