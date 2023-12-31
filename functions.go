package slicetools

// AllSlice returns true if all elements of the slice return true for the given function.
// Takes a slice and a function that takes the type of the slice elements and returns a boolean.
// Returns a bool.
func AllSlice[T any](s []T, f func(T) bool) bool {
	out := true

	for _, elem := range s {
		out = out && f(elem)
	}

	return out
}

// AnySlice returns true if one or more elements of the slice return true for the given function.
// Takes a slice and a function that takes the type of the slice elements and returns a boolean.
// Returns a bool.
func AnySlice[T any](s []T, f func(T) bool) bool {
	var out bool

	for _, elem := range s {
		out = out || f(elem)
	}

	return out
}

// FilterSlice returns the elements of an input slice for which a given filter function returns true.
// Takes a slice and a function that takes the type of the slice elements and returns a boolean.
// Returns the result as a new slice.
func FilterSlice[T any](s []T, f func(T) bool) []T {
	var out []T

	for _, elem := range s {
		if f(elem) {
			out = append(out, elem)
		}
	}

	return out
}

// MapSlice returns a new slice containing the results of a given mapping function being applied to each element of the input slice.
// Takes a slice and a function that operates on the type of the slice elements.
// Returns a slice containing elements of the type of the mapping function's output type.
func MapSlice[T, U any](s []T, f func(T) U) []U {
	var out []U

	for _, elem := range s {
		out = append(out, f(elem))
	}

	return out
}

// MapReduceSlice returns a single value, representing the result of accumulating (through a reduce function) the results of a mapping function for each element.
// Takes a slice, a mapping function that operates on the type of the slice elements, and a reduction function that operates on the output type of the mapping function.
// Returns a value of the output type of the reduction function.
func MapReduceSlice[T, U, V any](s []T, m func(T) U, r func(U, V) V) V {
	var acc V

	for _, elem := range s {
		acc = r(m(elem), acc)
	}

	return acc
}

// ReduceSlice returns a single value, representing the result of accumulating (through a reduce function) all elements of the input slice.
// Takes a slice and a function that operates on the type of the slice elements.
// Returns a value of the output type of the reduction function.
func ReduceSlice[T, U any](s []T, f func(T, U) U) U {
	var acc U

	for _, elem := range s {
		acc = f(elem, acc)
	}

	return acc
}

// UniqSlice returns a slice containing the unique elements of the input slice.
// Takes a slice of a comparable type.
// Returns a slice of the same type.
func UniqSlice[T comparable](s []T) []T {
	var out []T
	var set map[T]struct{}

	for _, elem := range s {
		if _, present := set[elem]; !present { // If NOT present in the set
			out = append(out, elem)
			set[elem] = struct{}{}
		}
	}

	return out
}
