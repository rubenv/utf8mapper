package utf8mapper

import "unicode/utf8"

// Maps a string to an index between upper and lower (exclusive), such that the
// chance of collisions is minimal.
//
// Distribution of characters:
//  - 0000  - 00FF:   40% (ASCII + some Latin1)
//  - 0100  - 01FF:   10% (Latin extended)
//  - 0200  - 1FFF:   10% (Remaining basic languages)
//  - 2000  - FFFF:   10% (Remainder of Plane 0)
//  - 10000 - 1FFFF:  10% (Plane 1)
//  - 20000 - 2FFFF:  10% (Plane 2)
//  - 30000 - 10FFFF: 10% (Planes 3 - 16)
//
// When mapping from 0 to math.MaxInt32, this gives us:
//  - 0000  - 00FF:      256 code points, 3355443 slots per code point
//  - 0100  - 01FF:      256 code points,  838861 slots per code point
//  - 0200  - 1FFF:     7680 code points,   27962 slots per code point
//  - 2000  - FFFF:    57344 code points,    3745 slots per code point
//  - 10000 - 1FFFF:   65536 code points,    3277 slots per code point
//  - 20000 - 2FFFF:   65536 code points,    3277 slots per code point
//  - 30000 - 10FFFF: 917504 code points,     234 slots per code point
func MapString(str string, lower, upper int32) (int32, error) {
	var result int32 = 0
	var start float64 = 0
	var end float64 = 0
	var allocation float64 = 0

	r, _ := utf8.DecodeRune([]byte(str))
	if r <= '\u00FF' {
		allocation = 0.4
		start = 0
		end = '\u00FF'
	} else {
		allocation = 0.1
		if r <= '\u01FF' {
			start = 0.4
			end = '\u01FF'
		} else if r <= '\u1FFF' {
			start = 0.5
			end = '\u1FFF'
		} else if r <= '\uFFFF' {
			start = 0.6
			end = '\uFFFF'
		} else if r <= '\U0001FFFF' {
			start = 0.7
			end = '\U0001FFFF'
		} else if r <= '\U0002FFFF' {
			start = 0.8
			end = '\U0002FFFF'
		} else {
			start = 0.9
			end = utf8.MaxRune
		}
	}

	position := float64(r) / end
	outputLength := float64(upper - lower)
	allocationStart := outputLength * start
	assignedPosition := outputLength * allocation * position
	result = int32(allocationStart + assignedPosition)
	return result, nil
}
