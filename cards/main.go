package main

func main() {
	cards := deck{newCard(), "Ace of diamonds"}
	cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
