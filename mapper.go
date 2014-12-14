package utf8mapper

import (
	"errors"
	"unicode/utf8"
)

// Maps a string to an index between lower and upper (inclusive), such that the
// chance of collisions is minimal.
//
// Distribution of characters:
//  - 0000  - 00FF:   50% (ASCII + some Latin1)
//  - 0100  - 01FF:   10% (Latin extended)
//  - 0200  - 1FFF:   10% (Remaining basic languages)
//  - 2000  - FFFF:   10% (Remainder of Plane 0)
//  - 10000 - 1FFFF:   5% (Plane 1)
//  - 20000 - 2FFFF:   5% (Plane 2)
//  - 30000 - 10FFFF: 10% (Planes 3 - 16)
//
// When mapping from 0 to math.MaxInt32, this gives us:
//  - 0000  - 00FF:      256 code points, 4194304 slots per code point
//  - 0100  - 01FF:      256 code points,  838861 slots per code point
//  - 0200  - 1FFF:     7680 code points,   27962 slots per code point
//  - 2000  - FFFF:    57344 code points,    3745 slots per code point
//  - 10000 - 1FFFF:   65536 code points,    1638 slots per code point
//  - 20000 - 2FFFF:   65536 code points,    1638 slots per code point
//  - 30000 - 10FFFF: 917504 code points,     234 slots per code point
func MapString(str string, lower, upper int32) (int32, error) {
	var result int32 = 0

	outputLength := upper - lower

	r, _ := utf8.DecodeRune([]byte(str))
	if r == utf8.RuneError {
		return 0, errors.New("Bad unicode!")
	}
	if r <= '\u00FF' {
		// position = r / 256
		// result = outputLength / 2 * position
		//
		// Order of operations below is swapped to preserve numerical precision
		// and uses a bitwise division for speed..
		//
		// result = (outputLength / 2) * (r / 256)
		// result = outputLength * (1 / 2) * r * (1 / 256)
		// result = outputLength * r / 512
		result = int32((int64(outputLength) * int64(r)) >> 9)
	} else {
		allocation, allocationStart, start, end := rangeParams(r)
		inputLength := end - start
		position := float64(r-start) / float64(inputLength)
		// outputStart = outputLength * allocationStart
		// outputRange = outputLength * allocation
		// result = outputStart + outputRange * position
		//
		// Thus:
		// result = outputLength * allocationStart + outputLength * allocation * position
		//
		// Thus:
		result = int32(float64(outputLength) * (allocationStart + (allocation * position)))
	}
	return result, nil
}

func rangeParams(r rune) (allocation, allocationStart float64, start, end int32) {
	allocation = 0.1
	if r > '\uFFFF' && r <= '\U0002FFFF' {
		allocation = 0.05
	}

	switch {
	case r <= '\u01FF':
		start = '\u0100'
		end = '\u01FF'
		allocationStart = 0.5
	case r <= '\u1FFF':
		start = '\u01FF'
		end = '\u1FFF'
		allocationStart = 0.6
	case r <= '\uFFFF':
		start = '\u1FFF'
		end = '\uFFFF'
		allocationStart = 0.7
	case r <= '\U0001FFFF':
		start = '\uFFFF'
		end = '\U0001FFFF'
		allocationStart = 0.8
	case r <= '\U0002FFFF':
		start = '\U0001FFFF'
		end = '\U0002FFFF'
		allocationStart = 0.85
	default:
		start = '\U0002FFFF'
		end = utf8.MaxRune
		allocationStart = 0.9
	}
	return
}
