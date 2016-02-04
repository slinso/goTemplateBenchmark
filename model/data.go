package model

type User struct {
	FirstName      string
	Email string
	FavoriteColors []string
	RawContent string
	EscapedContent string
}

type Navigation struct {
	Item string
	Link string
}