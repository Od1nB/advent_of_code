package camelgame

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	h1 = Hand{Cards: "32T3K", Bid: 2456}
	h2 = Hand{Cards: "KK677", Bid: 25}
	h3 = Hand{Cards: "KTJJT", Bid: 25543}
	h4 = Hand{Cards: "T55J5", Bid: 3242}
	h5 = Hand{Cards: "QQQJA", Bid: 2124}
	h6 = Hand{Cards: "KTJ37", Bid: 585}
	h7 = Hand{Cards: "K2J3K", Bid: 585}
	h8 = Hand{Cards: "KJJ37", Bid: 585}
)

func TestHighCard(t *testing.T) {
	assert.Equal(t, highCard(h2, h3, CardValue), false)
	assert.Equal(t, highCard(h3, h2, CardValue), true)
	assert.Equal(t, highCard(h1, h2, CardValue), true)

}
func TestCardComp(t *testing.T) {
	assert.Equal(t, CardComp(h1), OnePair)
	assert.Equal(t, CardComp(h2), TwoPair)
	assert.Equal(t, CardComp(h3), TwoPair)
	assert.Equal(t, CardComp(h4), ThreeofaKind)
	assert.Equal(t, CardComp(h5), ThreeofaKind)
}

func TestHighCardJkr(t *testing.T) {
	assert.Equal(t, highCard(h8, h7, JkrCardValue), true)
}

func TestCardCompJkr(t *testing.T) {
	assert.Equal(t, CardCompJoker(h6), OnePair)
	assert.Equal(t, CardCompJoker(h3), FourofaKind)
	assert.Equal(t, CardCompJoker(h5), FourofaKind)
	assert.Equal(t, CardCompJoker(h2), TwoPair)
}
