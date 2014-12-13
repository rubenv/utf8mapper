# utf8mapper
--
    import "github.com/rubenv/utf8mapper"

Work in progress

## Usage

#### func  MapString

```go
func MapString(str string, lower, upper int32) (int32, error)
```
Maps a string to an index between upper and lower (exclusive), such that the
chance of collisions is minimal.

Distribution of characterss:

    - 0000  - 00FF:   40% (ASCII + some Latin1)
    - 0100  - 01FF:   10% (Latin extended)
    - 0200  - 1FFF:   10% (Remaining basic languages)
    - 2000  - FFFF:   10% (Remainder of Plane 0)
    - 10000 - 1FFFF:  10% (Plane 1)
    - 20000 - 2FFFF:  10% (Plane 2)
    - 30000 - 10FFFF: 10% (Planes 3 - 16)

When mapping from 0 to math.MaxInt32, this gives us:

    - 0000  - 00FF:      256 code points, 3355443 slots per code point
    - 0100  - 01FF:      256 code points,  838861 slots per code point
    - 0200  - 1FFF:     7680 code points,   27962 slots per code point
    - 2000  - FFFF:    57344 code points,    3745 slots per code point
    - 10000 - 1FFFF:   65536 code points,    3277 slots per code point
    - 20000 - 2FFFF:   65536 code points,    3277 slots per code point
    - 30000 - 10FFFF: 917504 code points,     234 slots per code point
