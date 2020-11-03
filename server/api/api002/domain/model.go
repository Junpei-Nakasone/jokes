package domain

// Jokes is for return value
type Jokes struct {
	Jokes []Joke `json:"jokes"`
	// ID   string `json:"id"`
	// Joke string `json:"joke"`
}

// Joke contains each joke
type Joke struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}
