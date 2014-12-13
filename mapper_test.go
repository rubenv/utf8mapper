package utf8mapper_test

import (
	"math"
	"testing"
)
import "github.com/rubenv/utf8mapper"

func TestMapper(t *testing.T) {
	var lower int32 = 0
	var upper int32 = 1000
	result, err := utf8mapper.MapString("test", lower, upper)
	if err != nil {
		t.Error(err)
	}
	if result == 0 || result == 1000 {
		t.Errorf("Result (%d) should be between lower (%d) and upper (%d)", result, lower, upper)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utf8mapper.MapString("hello", 0, math.MaxInt32)
	}
}
