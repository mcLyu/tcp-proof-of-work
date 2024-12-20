package quotes

import "math/rand/v2"

var quotes = [8]string{
	"Never be an insincere friend, never be manipulative, one day you will be discovered and lose everything",
	"A bad friend secretly plots your downfall",
	"Keep quiet and the enemy will reveal himself",
	"Nothing exists, all is shadow and illusion",
	"Don't tell your friend you are fighting with your wife, it gives him pleasure",
	"Don’t fight with your enemy’s brains, fight with his heart",
	"Riches draw friends as corpses draw vultures",
	"Nature tells the truth as it is; it has no euphemism",
}

func GetRandomQuote() string {
	return quotes[rand.IntN(8)]
}
