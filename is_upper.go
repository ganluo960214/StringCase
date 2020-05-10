package StringCase

var (
	upper = map[rune]interface{}{
		'A' + 0:  nil,
		'A' + 1:  nil,
		'A' + 2:  nil,
		'A' + 3:  nil,
		'A' + 4:  nil,
		'A' + 5:  nil,
		'A' + 6:  nil,
		'A' + 7:  nil,
		'A' + 8:  nil,
		'A' + 9:  nil,
		'A' + 10: nil,
		'A' + 11: nil,
		'A' + 12: nil,
		'A' + 13: nil,
		'A' + 14: nil,
		'A' + 15: nil,
		'A' + 16: nil,
		'A' + 17: nil,
		'A' + 18: nil,
		'A' + 19: nil,
		'A' + 20: nil,
		'A' + 21: nil,
		'A' + 22: nil,
		'A' + 23: nil,
		'A' + 24: nil,
		'A' + 25: nil,
	}
)

func IsUpper(r rune) bool {

	_, is := upper[r]
	return is
}
