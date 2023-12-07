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
	assert.Equal(t, highCard(h2, h3), false)
	assert.Equal(t, highCard(h3, h2), true)
	assert.Equal(t, highCard(h1, h2), true)

}
func TestCardComp(t *testing.T) {
	assert.Equal(t, h1.CardComp(), OnePair)
	assert.Equal(t, h2.CardComp(), TwoPair)
	assert.Equal(t, h3.CardComp(), TwoPair)
	assert.Equal(t, h4.CardComp(), ThreeofaKind)
	assert.Equal(t, h5.CardComp(), ThreeofaKind)
}

func TestHighCardJkr(t *testing.T) {
	assert.Equal(t, jkrHighCard(h8, h7), true)
}

func TestCardCompJkr(t *testing.T) {
	assert.Equal(t, h6.CardCompJoker(), OnePair)
	assert.Equal(t, h3.CardCompJoker(), FourofaKind)
	assert.Equal(t, h5.CardCompJoker(), FourofaKind)
	assert.Equal(t, h2.CardCompJoker(), TwoPair)
}
