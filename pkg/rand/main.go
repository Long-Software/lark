package rand

import (
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/interfaces"
	"github.com/go-faker/faker/v4/pkg/options"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Int return a random integer between start and end inclusive
// If start is larger than end then it will return 0
func Int[T int | int8 | int16 | int32 | int64](start T, end T) T {
	if start >= end {
		return 0
	}
	type data struct {
		Int T
	}
	a := data{}
	boundary := interfaces.RandomIntegerBoundary{Start: int(start), End: int(end)}
	if err := faker.FakeData(&a, options.WithRandomIntegerBoundaries(boundary)); err != nil {
		return 0
	}
	return a.Int
}

// String return a random string with the len specified
// If the len is equal or less than 0 it will create an empty string
func String(n int) string {
	if n <= 0 {
		return ""
	}
	type data struct {
		SString string
	}
	a := data{}
	if err := faker.FakeData(&a, options.WithRandomStringLength(uint(n))); err != nil {
		return ""
	}
	return a.SString
}

// Float return a random float between start and end inclusive
// If start is larger than end then it will return 0
func Float(start float64, end float64) float64 {
	if start >= end {
		return 0
	}
	type data struct {
		Float float64
	}
	a := data{}
	boundary := interfaces.RandomFloatBoundary{Start: start, End: end}
	if err := faker.FakeData(&a, options.WithRandomFloatBoundaries(boundary)); err != nil {
		return 0
	}
	return a.Float
}
