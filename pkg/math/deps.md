# Linear Algebra

# Scalar
```go
type Scalar float32

type IScalar interface {
  AddScalarWithScalar(a, b Scalar) Scalar
  SubScalarWithScalar(a, b Scalar) Scalar
  MulScalarWithScalar(a, b Scalar) Scalar
  DivScalarWithScalar(a, b Scalar) Scalar
}

```

# Vector
```go
type Vector []Scalar

type IVector interface {
  Add(a *Vector) (Vector, error)  // deps: MapWithVector, AddScalarWithScalar
  Dot(a *Vector) (Scalar, error)  // deps: MapWithVector, MulScalarWithScalar
  Norm() Scalar // deps: Dot
  // helper function
  Reduce(f func(a, current Scalar) Scalar, initial Scalar) Scalar
  Map(f func(a Scalar) Scalar) Vector
  MapWithVector(a Vector, f func(b, c Scalar) Scalar) (Vector, error)
}


```