# utf8mapper
--
    import "github.com/rubenv/utf8mapper"

[![Build Status](https://travis-ci.org/rubenv/utf8mapper.svg?branch=master)](https://travis-ci.org/rubenv/utf8mapper) [![GoDoc](https://godoc.org/github.com/rubenv/utf8mapper?status.png)](https://godoc.org/github.com/rubenv/utf8mapper)

Maps strings onto an integer output range. Can be used to calculate an ordering
parameter for a list of items (based on e.g. their name), while attempting to
minimize the chances of collisions when new items are inserted.

## Usage

#### func  MapString

```go
func MapString(str string, lower, upper int32) (int32, error)
```
Maps a string to an index between lower and upper (inclusive), such that the
chance of collisions is minimal.

Distribution of characters:

    - 0000  - 00FF:   50% (ASCII + some Latin1)
    - 0100  - 01FF:   10% (Latin extended)
    - 0200  - 1FFF:   10% (Remaining basic languages)
    - 2000  - FFFF:   10% (Remainder of Plane 0)
    - 10000 - 1FFFF:   5% (Plane 1)
    - 20000 - 2FFFF:   5% (Plane 2)
    - 30000 - 10FFFF: 10% (Planes 3 - 16)

When mapping from 0 to math.MaxInt32, this gives us:

    - 0000  - 00FF:      256 code points, 4194304 slots per code point
    - 0100  - 01FF:      256 code points,  838861 slots per code point
    - 0200  - 1FFF:     7680 code points,   27962 slots per code point
    - 2000  - FFFF:    57344 code points,    3745 slots per code point
    - 10000 - 1FFFF:   65536 code points,    1638 slots per code point
    - 20000 - 2FFFF:   65536 code points,    1638 slots per code point
    - 30000 - 10FFFF: 917504 code points,     234 slots per code point
