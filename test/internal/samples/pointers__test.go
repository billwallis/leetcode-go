package samples_test

import (
	"github.com/billwallis/leetcode-go/internal/samples"
	"testing"
)

func TestIncrementByValue(t *testing.T) {
	thing := samples.Thing{Value: 1}
	value := samples.IncrementByValue(thing)
	if thing.Value != 1 {
		t.Errorf("thing.Value == %d, want %d", thing.Value, 1)
	}
	if value != 2 {
		t.Errorf("newThing.Value == %d, want %d", value, 2)
	}
}

func TestIncrementByReference(t *testing.T) {
	thing := samples.Thing{Value: 1}
	value := samples.IncrementByReference(&thing)
	if thing.Value != 2 {
		t.Errorf("thing.Value == %d, want %d", thing.Value, 2)
	}
	if value != 2 {
		t.Errorf("newThing.Value == %d, want %d", value, 2)
	}
}

func ExampleDoTheThing() {
	samples.DoTheThing()
	// Output:
	// thing.Value: 1
	// thing.Value: 1, byVal: 2
	// thing.Value: 2, byRef: 2
}
