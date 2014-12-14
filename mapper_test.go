package utf8mapper_test

import (
	"math"
	"testing"
	"unicode/utf8"
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

func TestOutputRange(t *testing.T) {
	var lower int32 = 0
	var upper int32 = math.MaxInt32
	for i := 0; i < utf8.MaxRune; i++ {
		result, _ := utf8mapper.MapString(string(i), lower, upper)
		if result < lower || result > upper {
			t.Fatalf("Result not in output range: %d (%d)", i, result)
		}
	}
}

func TestIncreasing(t *testing.T) {
	var lower int32 = 0
	var upper int32 = math.MaxInt32
	for i := 0; i < utf8.MaxRune; i++ {
		result, err1 := utf8mapper.MapString(string(i), lower, upper)
		result2, err2 := utf8mapper.MapString(string(i+1), lower, upper)
		if err1 != nil || err2 != nil {
			continue // Skip bad Unicode
		}
		if result >= result2 {
			t.Fatalf("Overlapping results around %d: %d - %d", i, result, result2)
		}
	}
}

func TestRecursing(t *testing.T) {
	var lower int32 = 0
	var upper int32 = math.MaxInt32
	result, err := utf8mapper.MapString("d1", lower, upper)
	if err != nil {
		t.Fatal()
	}

	result2, err := utf8mapper.MapString("d2", lower, upper)
	if err != nil {
		t.Fatal()
	}

	if result >= result2 {
		t.Fatalf("Equal mappings: %d - %d", result, result2)
	}
}

func TestRecursing2(t *testing.T) {
	var lower int32 = 0
	var upper int32 = math.MaxInt32
	result, err := utf8mapper.MapString("doca", lower, upper)
	if err != nil {
		t.Fatal()
	}

	result2, err := utf8mapper.MapString("docb", lower, upper)
	if err != nil {
		t.Fatal()
	}

	if result >= result2 {
		t.Fatalf("Equal mappings: %d - %d", result, result2)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utf8mapper.MapString("Hello", 0, math.MaxInt32)
	}
}
