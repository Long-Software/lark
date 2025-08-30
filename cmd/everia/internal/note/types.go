package note

import "time"

type Info struct {
	Title        string    `json:"title"`
	LastEditTime string    `json:"lastEditTime"`
	DateTime     time.Time `json:"-"`
}
