# slicetools

While working with go I missed a few of the functions I grew to know and love in Elixir's [Enum module](https://hexdocs.pm/elixir/1.15/Enum.html).

This package offers a few of those functions (with generic arguments). Currently, it contains:

```
All[T any](s []T, f func(T) bool) bool

Any[T any](s []T, f func(T) bool) bool

Filter[T any](s []T, f func(T) bool) []T

Map[T, U any](s []T, f func(T) U) []U

MapReduce[T, U, V any](s []T, m func(T) U, r func(U, V) V) V

Reduce[T, U any](s []T, f func(T, U) U) U

Uniq[T comparable](s []T) []T
```