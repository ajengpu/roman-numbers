package model

type RomanNumber struct {
	Value int
	Rank  int
	IsOne bool
}

var romans = map[string]RomanNumber{
	"I": RomanNumber{
		Value: 1,
		Rank:  1,
		IsOne: true,
	},
	"V": RomanNumber{
		Value: 5,
		Rank:  2,
		IsOne: false,
	},
	"X": RomanNumber{
		Value: 10,
		Rank:  3,
		IsOne: true,
	},
	"L": RomanNumber{
		Value: 50,
		Rank:  4,
		IsOne: false,
	},
	"C": RomanNumber{
		Value: 100,
		Rank:  5,
		IsOne: true,
	},
	"D": RomanNumber{
		Value: 500,
		Rank:  6,
		IsOne: false,
	},
	"M": RomanNumber{
		Value: 1000,
		Rank:  7,
		IsOne: true,
	},
}

func GetRomanNumber(symbol string) (RomanNumber, error) {
	if romans[symbol] != (RomanNumber{}) {
		return romans[symbol], nil
	}
	return RomanNumber{}, ErrInvalidRomanNumber
}

func GetDecimalFromRoman(roman string) (int, error) {
	res := 0
	rankCounter := 0

	for i := 0; i < len(roman); i++ {
		curNum, err := GetRomanNumber(string(roman[i]))
		if err != nil {
			return 0, ErrInvalidRomanNumber
		}

		if i == len(roman)-1 {
			res += curNum.Value
			break
		}

		nextNum, err := GetRomanNumber(string(roman[i+1]))
		if err != nil {
			return 0, ErrInvalidRomanNumber
		}

		if curNum.Rank == nextNum.Rank {
			if !curNum.IsOne {
				return 0, ErrInvalidRomanNumber
			}
			rankCounter++
			if rankCounter > 2 {
				return 0, ErrInvalidRomanNumber
			}
		} else {
			rankCounter = 0
		}

		if curNum.Rank < nextNum.Rank {
			if (curNum.Rank+1 != nextNum.Rank && curNum.Rank+2 != nextNum.Rank) || !curNum.IsOne {
				return 0, ErrInvalidRomanNumber
			}

			if i > 0 {
				pNum, err := GetRomanNumber(string(roman[i-1]))
				if err != nil {
					return 0, ErrInvalidRomanNumber
				}
				if (!nextNum.IsOne && pNum.Rank <= nextNum.Rank) || (nextNum.IsOne && pNum.Rank < nextNum.Rank) {
					return 0, ErrInvalidRomanNumber
				}
			}

			if i != len(roman)-2 {
				fNum, err := GetRomanNumber(string(roman[i+2]))
				if err != nil {
					return 0, ErrInvalidRomanNumber
				}
				if (nextNum.IsOne && fNum.Rank > nextNum.Rank-3) || (!nextNum.IsOne && fNum.Rank > nextNum.Rank-2) {
					return 0, ErrInvalidRomanNumber
				}
			}
			res = res - curNum.Value + nextNum.Value
			i++
		} else {
			if i == len(roman)-1 {
				res += curNum.Value + nextNum.Value
				break
			}
			res += curNum.Value
		}
	}
	return res, nil
}
