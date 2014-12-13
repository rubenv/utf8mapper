package utf8mapper

import "unicode/utf8"

// Maps a string to an index between upper and lower (exclusive), such that the
// chance of collisions is minimal.
func MapString(str string, upper, lower int32) (int32, error) {
	r, _ := utf8.DecodeRune([]byte(str))
	return r, nil
}
