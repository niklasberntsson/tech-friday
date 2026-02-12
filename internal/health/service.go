package health

import (
	"math/rand"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Env     string `json:"env"`
	Version string `json:"version"`
	Quote   string `json:"quote"`
}

var quotes = []string{
	"It seems today that all you see is violence in movies and sex on TV.",
	"Shut up, Meg.",
	"Giggity giggity goo!",
	"Victory is mine!",
	"Road House.",
	"Oh, have you not heard? It was my understanding that everyone had heard... that the bird is the word!",
	"Whoa, whoa, whoa, whoa, whoa, whoa... Lois, this is not my Batman glass.",
	"I'm not a drunk, I'm a social drinker. Which means I have to drink when I'm around people.",
	"A boat's a boat, but the mystery box could be anything. It could even be a boat!",
	"Perhaps a game of Hungry Hungry Hippos will settle our differences.",
	"He's a guy who over-explains his jokes! Get it? Because I just told you what he does!",
	"I'm the most intelligent person in this house, and quite possibly the world.",
	"Cool Hwhip.",
	"By the way, Meg, you're adopted. No, I'm just kidding. Who'd want you?",
	"I'm not just a dog, Brian. I'm a person. I'm a person with a dog's body.",
	"Peter, what are you doing? I'm crack.",
	"Brian, there's a message in my Alpha-Bits. It says 'Oooooo'.",
	"Holy crap!",
	"Everything's better with a bag of weed!",
	"Gosh, it's nice to be out of that bear.",
}

func GetHealth() HealthResponse {
	randomQuote := quotes[rand.Intn(len(quotes))]

	return HealthResponse{
		Status:  "ok",
		Env:     "dev",
		Version: "1.0.0",
		Quote:   randomQuote,
	}
}
