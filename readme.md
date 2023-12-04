# slicetools

While working with go I missed a few of the functions I grew to know and love in Elixir's [Enum module](https://hexdocs.pm/elixir/1.15/Enum.html).

This package offers a few of those functions (with generic arguments). Currently, it contains:

```
AllSlice[T any](s []T, f func(T) bool) bool

AnySlice[T any](s []T, f func(T) bool) bool

FilterSlice[T any](s []T, f func(T) bool) []T

MapSlice[T, U any](s []T, f func(T) U) []U

MapReduceSlice[T, U, V any](s []T, m func(T) U, r func(U, V) V) V

ReduceSlice[T, U any](s []T, f func(T, U) U) U

UniqSlice[T comparable](s []T) []T
```