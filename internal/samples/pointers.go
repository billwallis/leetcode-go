package samples

import "fmt"

// Thing just holds an integer value.
type Thing struct {
	Value int
}

// IncrementByValue increments the value of _a copy of_ the Thing by 1
// and returns the new Thing.
// This corresponds to the `ByVal` keyword in Visual Basic.
func IncrementByValue(thing Thing) int {
	thing.Value++
	return thing.Value
}

// IncrementByReference increments the value of the Thing by 1 and
// returns the same Thing.
// This corresponds to the `ByRef` keyword in Visual Basic.
func IncrementByReference(thing *Thing) int {
	thing.Value++
	return thing.Value
}

func DoTheThing() {
	thing := Thing{Value: 1}
	fmt.Printf("thing.Value: %d\n", thing.Value)

	byVal := IncrementByValue(thing)
	fmt.Printf("thing.Value: %d, byVal: %d\n", thing.Value, byVal)

	byRef := IncrementByReference(&thing)
	fmt.Printf("thing.Value: %d, byRef: %d\n", thing.Value, byRef)
}
