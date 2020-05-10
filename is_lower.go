package StringCase

var (
	lower = map[rune]interface{}{
		'a' + 0:  nil,
		'a' + 1:  nil,
		'a' + 2:  nil,
		'a' + 3:  nil,
		'a' + 4:  nil,
		'a' + 5:  nil,
		'a' + 6:  nil,
		'a' + 7:  nil,
		'a' + 8:  nil,
		'a' + 9:  nil,
		'a' + 10: nil,
		'a' + 11: nil,
		'a' + 12: nil,
		'a' + 13: nil,
		'a' + 14: nil,
		'a' + 15: nil,
		'a' + 16: nil,
		'a' + 17: nil,
		'a' + 18: nil,
		'a' + 19: nil,
		'a' + 20: nil,
		'a' + 21: nil,
		'a' + 22: nil,
		'a' + 23: nil,
		'a' + 24: nil,
		'a' + 25: nil,
	}
)

func IsLower(r rune) bool {
	_, is := lower[r]
	return is
}
