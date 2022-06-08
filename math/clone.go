package math 

// Deep clone a one-dimensional array.
func Clone[T any](xs []T) []T {
  ys := make([]T, len(xs))
  for i := 0; i < len(xs); i++ {
    ys[i] = xs[i]
  }

  return ys
}

// Deep clone a two-dimensional array.
func Clone2d[T any](xs [][]T) [][]T {
	ys := make([][]T, len(xs))
	for i := 0; i < len(xs); i++ {
		arr := xs[i]
		ys[i] = arr
	}

	return ys
}
