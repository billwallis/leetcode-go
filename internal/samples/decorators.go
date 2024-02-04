package samples

import "fmt"

/*
Sample code for decorators, adapted from:

- https://youtu.be/GipAZwKFgoA?si=qag6tl1bak3k6T1x
*/

// SomeExternalFunction calls a function (suppose it comes from an
// external package)
func SomeExternalFunction(fn func(string)) {
	fn("Hello, world!")
}

type Store interface {
	Save(string) error
}

type PostgresStore struct{}

func (pgStore PostgresStore) Save(data string) error {
	fmt.Println("Saving to Postgres:", data)
	return nil
}

func ConcreteFunction(s string) {
	fmt.Println("Concrete function called with:", s)
}

func Decorator(fn func(string), store Store) func(string) {
	return func(s string) {
		// Call the concrete function
		fn(s)

		// Save the data to the store (which is the "decoration")
		if err := store.Save(s); err != nil {
			fmt.Println("Error while saving to store:", err)
		}
	}
}

func Main() {
	pgStore := PostgresStore{}
	decoratedFunction := Decorator(ConcreteFunction, pgStore)
	SomeExternalFunction(decoratedFunction)
}
