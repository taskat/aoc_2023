package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

type card rune

var cardValus = map[card]int{
	'2': 2, '3': 3, '4': 4, '5': 5,
	'6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

func Map[T, U any](arr []T, f func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

type hand struct {
	cards    []card
	handType handType
	bid      int
}

func parseHand(input string) hand {
	parts := strings.Split(input, " ")
	cards := make([]card, len(parts[0]))
	for i, c := range parts[0] {
		cards[i] = card(c)
	}
	bid := atoi(parts[1])
	h := hand{cards: cards, bid: bid}
	h.handType = handTypeHashes[h.calculateHandTypeHash()]
	return h
}

func (h *hand) calculateHandTypeHash() string {
	cards := make(map[card]int)
	for _, c := range h.cards {
		cards[c]++
	}
	values := make([]int, 0, len(cards))
	for _, v := range cards {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	return strings.Join(Map(values, func(i int) string { return strconv.Itoa(i) }), "")
}

func (h *hand) String() string {
	return fmt.Sprintf("Cards: %s, handType: %s, bid: %d", string(h.cards), h.handType.String(), h.bid)
}

func (h *hand) isStrongerThan(other hand) bool {
	if h.handType != other.handType {
		return h.handType > other.handType
	}
	for i := range h.cards {
		if h.cards[i] != other.cards[i] {
			return cardValus[h.cards[i]] > cardValus[other.cards[i]]
		}
	}
	return true
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var handTypeHashes = map[string]handType{
	"11111": highCard,
	"2111":  onePair,
	"221":   twoPairs,
	"311":   threeOfAKind,
	"32":    fullHouse,
	"41":    fourOfAKind,
	"5":     fiveOfAKind,
}

func (h *handType) String() string {
	switch *h {
	case highCard:
		return "High Card"
	case onePair:
		return "One Pair"
	case twoPairs:
		return "Two Pairs"
	case threeOfAKind:
		return "Three of a Kind"
	case fullHouse:
		return "Full House"
	case fourOfAKind:
		return "Four of a Kind"
	case fiveOfAKind:
		return "Five of a Kind"
	default:
		panic("Unknown hand type")
	}
}

func parseHands(input string) []hand {
	lines := strings.Split(input, "\n")
	return Map(lines, parseHand)
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	hands := parseHands(input)
	sort.Slice(hands, func(i, j int) bool { return !hands[i].isStrongerThan(hands[j]) })
	winnings := 0
	for i, h := range hands {
		winnings += h.bid * (i + 1)
	}
	return fmt.Sprintf("%d", winnings)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}
