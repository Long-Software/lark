package token

import "time"

type Maker interface {
	// CreateToken creates a new token for a specific data and duration
	CreateToken(d interface{}, duration time.Duration) (string, *Payload, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
