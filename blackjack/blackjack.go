package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	tot := ParseCard(card1) + ParseCard(card2)
	switch {
	case tot == 22:
		return "P"
	case tot == 21 && ParseCard(dealerCard) == 11:
		return "S"
	case tot == 21 && ParseCard(dealerCard) == 10:
		return "S"
	case tot == 21 && ParseCard(dealerCard) != 10:
		return "W"
	case tot >= 17 && tot <= 20:
		return "S"
	case tot >= 12 && tot <= 16 && ParseCard(dealerCard) >= 7:
		return "H"
	case tot >= 12 && tot <= 16 && ParseCard(dealerCard) <= 7:
		return "S"
	default:
		return "H"
	}
}
