# utf8mapper

[![Build Status](https://travis-ci.org/rubenv/utf8mapper.svg?branch=master)](https://travis-ci.org/rubenv/utf8mapper) [![GoDoc](https://godoc.org/github.com/rubenv/utf8mapper?status.png)](https://godoc.org/github.com/rubenv/utf8mapper)

Maps strings onto an integer output range. Can be used to calculate an ordering
parameter for a list of items (based on e.g. their name), while attempting to
minimize the chances of collisions when new items are inserted.

## Installation
```
go get github.com/rubenv/utf8mapper
```

Import into your application with:

```go
import "github.com/rubenv/utf8mapper"
```

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

## License 

    (The MIT License)

    Copyright (C) 2014 by Ruben Vermeersch <ruben@rocketeer.be>

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.
