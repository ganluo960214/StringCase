package StringCase

var (
	upper = map[rune]byte{
		'A' + 0:  0,
		'A' + 1:  0,
		'A' + 2:  0,
		'A' + 3:  0,
		'A' + 4:  0,
		'A' + 5:  0,
		'A' + 6:  0,
		'A' + 7:  0,
		'A' + 8:  0,
		'A' + 9:  0,
		'A' + 10: 0,
		'A' + 11: 0,
		'A' + 12: 0,
		'A' + 13: 0,
		'A' + 14: 0,
		'A' + 15: 0,
		'A' + 16: 0,
		'A' + 17: 0,
		'A' + 18: 0,
		'A' + 19: 0,
		'A' + 20: 0,
		'A' + 21: 0,
		'A' + 22: 0,
		'A' + 23: 0,
		'A' + 24: 0,
		'A' + 25: 0,
	}
)

func IsUpper(r rune) bool {

	_, is := upper[r]
	return is
}
