package math

import "math"

type Vector []Scalar
type Vec struct{}

var vec Vec

func (v *Vec) Add(a, b Vector) (Vector, error) {
	result, err := a.MapWithVector(b, AddScalarWithScalar)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (v *Vec) Sub(a, b Vector) (Vector, error) {
	result, err := a.MapWithVector(b, SubScalarWithScalar)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (v *Vec) Dot(a, b Vector) (Scalar, error) {
	vec, err := a.MapWithVector(b, MulScalarWithScalar)
	if err != nil {
		return 0, err
	}
	return vec.Reduce(AddScalarWithScalar, 0), nil
}

func (v *Vec) Distance(a, b Vector) (Scalar, error) {
	d, err := vec.Sub(a, b)
	if err != nil {
		return 0, err
	}
	return d.Norm(), nil
}
func (v *Vec) Proj(a, b Vector) (Vector, error) {
	uv, err := vec.Dot(a, b)
	if err != nil {
		return nil, err
	}
	uu, err := vec.Dot(a, a)
	if err != nil {
		return nil, err
	}
	return a.Times(uv / uu), nil
}

func (v *Vec) CosDel(a, b Vector) (Scalar, error) {
	uv, err := vec.Dot(a, b)
	if err != nil {
		return 0, err
	}
	return uv / a.Norm() / b.Norm(), nil
}

func (v *Vec) IsOrthogonal(a, b Vector) (bool, error) {
	uv, err := vec.Dot(a, b)
	if err != nil {
		return false, err
	}
	return uv == 0, nil
}

func NewVector(scals []Scalar) Vector { return Vector(scals) }
func (v *Vector) Times(a Scalar) Vector {
	return v.Map(func(b Scalar) Scalar { return a * b })
}
func (v *Vector) Norm() Scalar {
	result, err := vec.Dot(*v, *v)
	if err != nil {
		return 0
	}
	return Scalar(math.Sqrt(float64(result)))
}
func (v *Vector) Unit() Vector { return v.Times(1 / v.Norm()) }

// sim(x,y) = x.y/(||x||||y||)
// func (v Vector) Sim(a *Vector) (Scalar, error) {
// 	dot_prod, err := v.Dot(a)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return dot_prod / (v.Norm() * a.Norm()), nil
// }

func IsCompatible(a Vector, b Vector) bool {
	return len(a) == len(b)
}

func (v Vector) MapWithVector(a Vector, f func(b, c Scalar) Scalar) (Vector, error) {
	if !IsCompatible(v, a) {
		return nil, InCompatibleLengthError
	}
	result := make([]Scalar, len(v))
	for i := range v {
		result[i] = f(v[i], a[i])
	}
	return result, nil
}
func (v Vector) Map(f func(a Scalar) Scalar) Vector {
	result := make([]Scalar, len(v))
	for i := range v {
		result[i] = f(v[i])
	}
	return result
}
func (v Vector) Reduce(f func(a, current Scalar) Scalar, initial Scalar) Scalar {
	for _, v := range v {
		initial = f(v, initial)
	}
	return initial
}
