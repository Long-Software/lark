package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	Data     interface{}
	ExpireAt time.Time
}

func NewPayload(d interface{}, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		Data:     d,
		ExpireAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpiredToken
	}
	return nil
}
