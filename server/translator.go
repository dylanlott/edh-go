package server

import (
	"github.com/zeebo/errs"
)

// # DESIGN
// Translator
// So the problem that keeps coming up in this app and that I've managed to hack
// around until now is that GraphQL Input Types clash a lot with the actual
// Model types. It basically comes down to a problem of how to elegantly
// combine structs of different types but similar compositions in Go.

// **Option 1**: This is my first idea, which takes a functional,
// interface driven approach that you can pass a Translator func,
// input type, and output type to.
// I think this is probably the best way to sovle it.

// Translator is a function for taking a value and returning any other type.
type Translator func(value interface{}) (interface{}, error)

// Translate will apply the Translator to the From value and marshal it to the To value.
// type Translate func(to, from interface{}, t Translator) error

// Polyglot binds the two together above imlementations together with an interface
type Polyglot interface {
	// Translate should be an immutable, thread safe function that applies the
	// Translator to the from value onto the to value and returns an error, if any.
	Translate(to, from interface{}, t Translator) error
}

// # IMPLEMENTATION
// ## Option 1

// polyglot fulfills Polyglot to translate structs around in a functional way
type polyglot struct{}

// Translate applies the Translator to the received to interface
func (p *polyglot) Translate(to, from interface{}, t Translator) error {
	to, err := t(from)
	if err != nil {
		return errs.New("failed to translate: %+v", err)
	}

	return nil
}

// InputGameTranslator fulfills the Translator interface to be used in
// the game subscription logic.
func InputGameTranslator(value interface{}) (interface{}, error) {
	g := &Game{}
	input := value.(*InputGame)

	g.ID = input.ID

	if input.Handle != nil {
		g.Handle = input.Handle
	}

	if input.CreatedAt != nil {
		g.CreatedAt = *input.CreatedAt
	}

	for _, p := range input.PlayerIDs {
		u := &User{}
		if p.ID != nil {
			u.ID = *p.ID
		}
	}

	if input.Turn != nil {
		g.Turn.Number = input.Turn.Number
		g.Turn.Phase = input.Turn.Phase
		g.Turn.Player = input.Turn.Player
	}

	return g, nil
}

// **Option 2**: Just make a bunch of interface transformers and pass it and input
// type and let it handle all the custom logic.
// I think I'm going to choose Option 1 for now, because
// it feels like a more fleixble and useful abstraction, but I'm going to leave
// this comment as reference to see if that ends up being the case.

// Orator is an example implementation of Option 2. I think it's clunkier than
// the first option, however there's an issue with scaling here, and this interface
// would likely grow much larger over time. Large interfaces aren't Go idiomatic,
// we want thin interfaces that we can compose in lots of different ways.
// This app in particular will have a lot of data manipulation needs that
// Option 1 does a better job of flexibly solving.

// Orator implmements Option 2 as an example.
type Orator interface {
	HandleInputGame(input *InputGame) (*Game, error)
	HandleInputBoardState(input *InputBoardState) (*BoardState, error)
}
