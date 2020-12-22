package server

// Translator is a function for taking a value and returning any other type.
type Translator func(value interface{}) (interface{}, error)

// Merge will apply the Translator to the From value and marshal it to the To value.
type Merge func(to, from interface{}, t Translator) error
