package slicetools

// All returns true if all elements of the slice return true for the given function.
// Takes a slice and a function that takes the type of the slice elements and returns a boolean.
// Returns a bool.
func All[T any](s []T, f func(T) bool) bool {
	out := true

	for _, elem := range s {
		out = out && f(elem)
	}

	return out
}

// Any returns true if one or more elements of the slice return true for the given function.
// Takes a slice and a function that takes the type of the slice elements and returns a boolean.
// Returns a bool.
func Any[T any](s []T, f func(T) bool) bool {
	var out bool

	for _, elem := range s {
		out = out || f(elem)
		if out {
			return out
		}
	}

	return out
}

// Chunk divides a slice of elements into a slice containing subslices of a given length. In other words, it 'chunks' the input slice every n elements.
// Chunk takes a slice of any type and an int, and returns a slice of slices containing the input type.
func Chunk[T any](s []T, l int) [][]T {
	var subSlices int

	if len(s) != 0 {
		subSlices = (len(s) / l)
		if len(s)%l != 0 {
			subSlices += 1
		}
	}

	out := make([][]T, subSlices)

	for i, elem := range s {
		out[i/l] = append(out[i/l], elem)
	}

	return out
}

// Count returns the number of elements in a slice for which a provided function returns true.
// Takes a slice of type T and a function that takes type T and returns a bool. Count returns an int.
func Count[T any](s []T, f func(T) bool) int {
	var out int

	for _, elem := range s {
		if f(elem) {
			out++
		}
	}

	return out
}

// Dedup will compress consecutive identical elements into one element.
// For example: []int{1, 2, 2, 3, 2} would become []int{1, 2, 3, 2}.
// Dedup takes and returns a slice of a comparable type.
func Dedup[T comparable](s []T) []T {
	out := make([]T, 0)

	for i, elem := range s {
		if i > 0 {
			if elem == s[i-1] {
				continue
			} else {
				out = append(out, elem)
			}
		} else {
			out = append(out, elem)
		}
	}

	return out
}

// DedupFunc behaves like Dedup(), but works on types that don't implement the comparable interface.
// DedupFunc takes a slice of elements and a comparison function that returns a bool to indicate whether two items are equal.
func DedupFunc[T any](s []T, compareFunc func(a T, b T) bool) []T {
	out := make([]T, 0)

	for i, elem := range s {
		if i > 0 {
			if compareFunc(elem, s[i-1]) {
				continue
			} else {
				out = append(out, elem)
			}
		} else {
			out = append(out, elem)
		}
	}

	return out
}

// Filter returns the elements of an input slice for which a given filter function returns true.
// Takes a slice and a function that takes the type of the slice elements and returns a boolean.
// Returns the result as a new slice.
func Filter[T any](s []T, f func(T) bool) []T {
	var out []T

	for _, elem := range s {
		if f(elem) {
			out = append(out, elem)
		}
	}

	return out
}

// Map returns a new slice containing the results of a given mapping function being applied to each element of the input slice.
// Takes a slice and a function that operates on the type of the slice elements.
// Returns a slice containing elements of the type of the mapping function's output type.
func Map[T, U any](s []T, f func(T) U) []U {
	var out []U

	for _, elem := range s {
		out = append(out, f(elem))
	}

	return out
}

// MapReduce returns a single value, representing the result of accumulating (through a reduce function) the results of a mapping function for each element.
// Takes a slice, a mapping function that operates on the type of the slice elements, and a reduction function that operates on the output type of the mapping function.
// Returns a value of the output type of the reduction function. Will return the zero value of the return type if nothing is accumulated.
func MapReduce[T, U, V any](s []T, m func(T) U, r func(U, V) V) V {
	var acc V

	for _, elem := range s {
		acc = r(m(elem), acc)
	}

	return acc
}

// Reduce returns a single value, representing the result of accumulating (through a reduce function) all elements of the input slice.
// Takes a slice and a function that operates on the type of the slice elements.
// Returns a value of the output type of the reduction function.
func Reduce[T, U any](s []T, f func(T, U) U) U {
	var acc U

	for _, elem := range s {
		acc = f(elem, acc)
	}

	return acc
}

// Uniq returns a slice containing the unique elements of the input slice.
// Takes a slice of a comparable type.
// Returns a slice of the same type.
func Uniq[T comparable](s []T) []T {
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
