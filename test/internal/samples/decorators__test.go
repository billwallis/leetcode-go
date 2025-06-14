package samples_test

import "github.com/billwallis/leetcode-go/internal/samples"

func ExamplePostgresStore_Save() {
	pgStore := samples.PostgresStore{}
	_ = pgStore.Save("some test data")
	// Output: Saving to Postgres: some test data
}

func ExampleConcreteFunction() {
	samples.ConcreteFunction("some test data")
	// Output: Concrete function called with: some test data
}

func ExampleDecorator() {
	pgStore := samples.PostgresStore{}
	decoratedFunction := samples.Decorator(func(s string) {}, pgStore)
	decoratedFunction("some test data")
	// Output: Saving to Postgres: some test data
}

func ExampleMain() {
	samples.Main()
	// Output: Concrete function called with: Hello, world!
	// Saving to Postgres: Hello, world!
}
