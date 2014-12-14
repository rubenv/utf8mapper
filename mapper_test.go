package utf8mapper_test

import (
	"math"
	"testing"
)
import "github.com/rubenv/utf8mapper"

func testMapping(t *testing.T, input string, lower, upper, expected int32) {
	result, err := utf8mapper.MapString(input, lower, upper)
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Result %d should be %d (input: %+q)", result, expected, input)
	}
}

func TestMapper(t *testing.T) {
	var lower int32 = 0
	var upper int32 = math.MaxInt32
	result, err := utf8mapper.MapString("test", lower, upper)
	if err != nil {
		t.Error(err)
	}
	if result == 0 || result == upper {
		t.Errorf("Result (%d) should be between lower (%d) and upper (%d)", result, lower, upper)
	}
}

func TestLower(t *testing.T) {
	testMapping(t, "\x00", 0, math.MaxInt32, 0)
}

func TestUpper(t *testing.T) {
	testMapping(t, "\U0010FFFF", 0, math.MaxInt32, math.MaxInt32)
}

func TestLatin(t *testing.T) {
	for i := 0; i <= '\xFF'; i++ {
		testMapping(t, string(i), 0, 512, int32(i))
	}
}

func TestEdges(t *testing.T) {
	edges := []rune{'\u00FF', '\u01FF', '\u1FFF', '\uFFFF', '\U0001FFFF', '\U0002FFFF'}
	var lower int32 = 0
	var upper int32 = math.MaxInt32
	for edge := range edges {
		result, _ := utf8mapper.MapString(string(edge), lower, upper)
		result2, _ := utf8mapper.MapString(string(edge+1), lower, upper)
		if result >= result2 {
			t.Errorf("Overlapping results around edges!")
		}
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utf8mapper.MapString("hello", 0, math.MaxInt32)
	}
}
