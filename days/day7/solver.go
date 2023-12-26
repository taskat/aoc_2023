package day7

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/maps"
	"aoc_2023/utils/stringutils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(lines []string, extraParams ...any) string {
	hands := parseHands(lines, false)
	return fmt.Sprintf("%d", calculateWinnings(hands, cardValues))
}

func (*Solver) SolvePart2(lines []string, extraParams ...any) string {
	hands := parseHands(lines, true)
	return fmt.Sprintf("%d", calculateWinnings(hands, cardValues_withJoker))
}

type card rune

func (c card) isSame(other card) bool {
	return c == other
}

func (c card) isStrongerThan(other card, cardValues map[card]int) bool {
	return cardValues[c] > cardValues[other]
}

var cardValues = map[card]int{
	'2': 2, '3': 3, '4': 4, '5': 5,
	'6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

var cardValues_withJoker = map[card]int{
	'J': 1, '2': 2, '3': 3, '4': 4,
	'5': 5, '6': 6, '7': 7, '8': 8,
	'9': 9, 'T': 10, 'Q': 12, 'K': 13, 'A': 14,
}

type hand struct {
	cards    []card
	handType handType
	bid      int
}

func parseHand(input string, withJokers bool) hand {
	parts := strings.Split(input, " ")
	cards := arrays.Map([]rune(parts[0]), func(c rune) card { return card(c) })
	bid := stringutils.Atoi(parts[1])
	h := hand{cards: cards, bid: bid}
	h.handType = handTypeHashes[h.calculateHandTypeHash(withJokers)]
	return h
}

func (h *hand) calculateHandTypeHash(withJoker bool) string {
	cards := make(map[card]int)
	for _, c := range h.cards {
		cards[c]++
	}
	jokers := 0
	if withJoker {
		jokers = cards['J']
		delete(cards, 'J')
		if len(cards) == 0 {
			return "5"
		}
	}
	values := maps.MapToArray(cards, func(k card, v int) int { return v })
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	if withJoker {
		values[0] += jokers
	}
	return strings.Join(arrays.Map(values, strconv.Itoa), "")
}

func (h *hand) String() string {
	return fmt.Sprintf("Cards: %s, handType: %s, bid: %d", string(h.cards), h.handType.String(), h.bid)
}

func (h *hand) isStrongerThan(other hand, cardValues map[card]int) bool {
	if h.handType != other.handType {
		return h.handType > other.handType
	}
	idx, card, ok := arrays.Find_i(h.cards, func(i int, c card) bool { return !c.isSame(other.cards[i]) })
	if !ok {
		return false
	}
	return card.isStrongerThan(other.cards[idx], cardValues)
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

func parseHands(lines []string, withJokers bool) []hand {
	return arrays.Map(lines, func(line string) hand { return parseHand(line, withJokers) })
}

func calculateWinnings(hands []hand, cardValues map[card]int) int {
	sort.Slice(hands, func(i, j int) bool { return !hands[i].isStrongerThan(hands[j], cardValues) })
	winnings := arrays.Accumulate_i(hands, 0, func(acc, idx int, h hand) int { return acc + h.bid*(idx+1) })
	return winnings
}
