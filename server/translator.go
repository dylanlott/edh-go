package server

import "github.com/zeebo/errs"

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
type Translate func(to, from interface{}, t Translator) error

// **Option 2**: Just make a bunch of interface transformers and pass it and input
// type and let it handle all the custom logic.
// I think I'm going to choose Option 1 for now, because
// it feels like a more fleixble and useful abstraction, but I'm going to leave
// this comment as reference to see if that ends up being the case.
type HandleInputGame func(input interface{}) (*Game, error)
type HandleInputBoardState func(input interface{}) (*BoardState, error)

// # IMPLEMENTATION
// ## Option 1
func InputGameTranslator(input *InputGame) (*Game, error) {
	return nil, errs.New("InputGameTranslator not impl")
}

func TranslateInputGame(to *Game, from *InputGame, t InputGameTranslator) error {
	return errs.New("TranslateInputGame not impl")
}
