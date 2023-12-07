package camelgame

type Composition int

const (
	HighCard Composition = iota
	OnePair
	TwoPair
	ThreeofaKind
	FullHouse
	FourofaKind
	FiveofaKind
)

var compNames = [...]string{"HighCard", "OnePair", "TwoPair", "ThreeofaKind", "FullHouse", "FourofaKind", "FiveofaKind"}

func (c Composition) String() string {
	return compNames[c]
}

var (
	rankToVal = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
	rankwJoker = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}
)

type Hand struct {
	Cards string
	Bid   int
}

func CardComp(h Hand) Composition {
	m := map[string]int{}
	for _, v := range h.Cards {
		m[string(v)]++
	}
	switch len(m) {
	case 1:
		return FiveofaKind
	case 2:
		for _, v := range m {
			if v == 4 {
				return FourofaKind
			}
		}
		return FullHouse
	case 3:
		for _, v := range m {
			if v == 3 {
				return ThreeofaKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}
func CardCompJoker(h Hand) Composition {
	m := map[string]int{}
	for _, v := range h.Cards {
		m[string(v)]++
	}
	switch len(m) {
	case 1:
		return FiveofaKind
	case 2:
		if _, jkr := m["J"]; jkr {
			return FiveofaKind
		}
		for _, v := range m {
			if v == 4 {
				return FourofaKind
			}
		}
		return FullHouse
	case 3:
		numJkr, jkr := m["J"]
		for _, v := range m {
			if v == 3 && jkr {
				return FourofaKind
			} else if v == 3 && !jkr {
				return ThreeofaKind
			}
		}
		if jkr {
			if numJkr == 1 {
				return FullHouse
			}
			return FourofaKind
		}
		return TwoPair
	case 4:
		if _, jkr := m["J"]; jkr {
			return ThreeofaKind
		}
		return OnePair
	default:
		if _, jkr := m["J"]; jkr {
			return OnePair
		}
		return HighCard
	}
}

func CardValue(s string) int {
	return rankToVal[s]
}

func JkrCardValue(s string) int {
	return rankwJoker[s]
}

func highCard(h, j Hand, cardValue func(s string) int) bool {
	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i] == j.Cards[i] {
			continue
		}
		return cardValue(string(h.Cards[i])) < cardValue(string(j.Cards[i]))
	}
	return false
}

func Less(hands []Hand, cardValue func(s string) int, cardComp func(h Hand) Composition) func(i, j int) bool {
	return func(i, j int) bool {
		c1, c2 := cardComp(hands[i]), cardComp(hands[j])
		if c1 != c2 {
			return c1 < c2
		}
		return highCard(hands[i], hands[j], cardValue)
	}
}
