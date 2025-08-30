package models

type PasswordEntry struct {
	Site     string `json:"site"`
	Username string `json:"username"`
	Password string `json:"password"`
	Notes    string `json:"notes"`
}