package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/abhibansal530/advent-of-code/utils"
)

type HandType int

const (
	handSize      = 5
	numDigitCards = (9 - 2 + 1)

	// Ordered by strength (strong to weak).
	_ HandType = iota
	// AAAAA
	FiveOfKind
	// AAAAB
	FourOfKind
	// AAABB
	FullHouse
	// AAABC
	ThreeOfKind
	// AABBC
	TwoPair
	// AABCD
	OnePair
	// ABCDE
	HighCard
)

// For non-digit cards, this tells the relative rank (w.r.t. 2).
var cardsRelativeRankMap = map[byte]int{
	'T': numDigitCards,
	'J': numDigitCards + 1,
	'Q': numDigitCards + 2,
	'K': numDigitCards + 3,
	'A': numDigitCards + 4,
}

var cardsRelativeRankWithJokerMap = map[byte]int{
	'J': -1,
	'T': numDigitCards,
	'Q': numDigitCards + 1,
	'K': numDigitCards + 2,
	'A': numDigitCards + 3,
}

func getCardRank(x byte, nonDigitCardsRank map[byte]int) int {
	if unicode.IsDigit(rune(x)) {
		return int(x - '2')
	}
	return nonDigitCardsRank[x]
}

func findHandType(cards string, useJAsWild bool) HandType {
	cardsMap := make(map[byte]int)
	for i := range cards {
		cardsMap[cards[i]]++
	}

	if cardsMap['J'] > 0 && useJAsWild {
		return findHandTypeWithJAsWild(cards, cardsMap)
	}

	switch len(cardsMap) {
	case 1:
		return FiveOfKind
	case 2:
		if utils.MapHasValue(cardsMap, 4) {
			return FourOfKind
		}
		return FullHouse
	case 3:
		if utils.MapHasValue(cardsMap, 3) {
			return ThreeOfKind
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func findHandTypeWithJAsWild(cards string, cardsMap map[byte]int) HandType {
	remainingDistinct := len(cardsMap) - 1
	delete(cardsMap, 'J')

	switch remainingDistinct {
	case 0, 1:
		return FiveOfKind
	case 2:
		if utils.MapHasValue(cardsMap, 1) {
			return FourOfKind
		}
		return FullHouse
	case 3:
		return ThreeOfKind
	}
	return OnePair
}

type CamelHand struct {
	cards    string
	handType HandType
	bid      int64
}

func NewHand(cards string, bid int64, useJAsWild bool) CamelHand {
	if len(cards) != handSize {
		panic(fmt.Sprintf("Invalid hand of size: %d", len(cards)))
	}
	return CamelHand{cards: cards, bid: bid, handType: findHandType(cards, useJAsWild)}
}

// If left < right.
func isLess(left, right CamelHand) bool {
	if left.handType != right.handType {
		return left.handType > right.handType
	}

	for i := range left.cards {
		if left.cards[i] != right.cards[i] {
			return getCardRank(left.cards[i], cardsRelativeRankMap) < getCardRank(right.cards[i], cardsRelativeRankMap)
		}
	}
	return false
}

// If left < right considering 'J' as wildcard.
func isLessWithJoker(left, right CamelHand) bool {
	if left.handType != right.handType {
		return left.handType > right.handType
	}

	for i := range left.cards {
		if left.cards[i] != right.cards[i] {
			return getCardRank(left.cards[i], cardsRelativeRankWithJokerMap) < getCardRank(right.cards[i], cardsRelativeRankWithJokerMap)
		}
	}
	return false
}

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)

	var hands, handsWithJAsWild []CamelHand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards, bidStr, _ := strings.Cut(scanner.Text(), " ")
		hands = append(hands, NewHand(cards, utils.ToInt64(bidStr), false))
		handsWithJAsWild = append(handsWithJAsWild, NewHand(cards, utils.ToInt64(bidStr), true))
	}

	sort.Slice(hands, func(i, j int) bool {
		return isLess(hands[i], hands[j])
	})

	var answer int64
	for i, h := range hands {
		answer += int64(i+1) * (h.bid)
	}

	fmt.Printf("Answer for first part is: %d\n", answer)

	sort.Slice(handsWithJAsWild, func(i, j int) bool {
		return isLessWithJoker(handsWithJAsWild[i], handsWithJAsWild[j])
	})

	var answer2 int64
	for i, h := range handsWithJAsWild {
		answer2 += int64(i+1) * (h.bid)
	}

	fmt.Printf("Answer for second part is: %d\n", answer2)
}
