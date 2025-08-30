package math

type Scalar float64

func NewScalar(val float64) Scalar { return Scalar(val) }
func (s Scalar) Get() float64      { return float64(s) }

// func (s Scalar) Add(a Scalar) Scalar { return s + a }
// func (s Scalar) Sub(a Scalar) Scalar { return s - a }
// func (s Scalar) Mul(a Scalar) Scalar { return s * a }
// func (s Scalar) Div(a Scalar) Scalar { return s / a }

func AddScalarWithScalar(a, b Scalar) Scalar { return a + b }
func SubScalarWithScalar(a, b Scalar) Scalar { return a - b }
func MulScalarWithScalar(a, b Scalar) Scalar { return a * b }
func DivScalarWithScalar(a, b Scalar) Scalar { return a / b }
