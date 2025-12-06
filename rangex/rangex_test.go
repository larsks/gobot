package rangex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRange(t *testing.T) {
	items := []struct {
		name     string
		expr     string
		expected []int
		valid    bool
	}{
		{"simple range", "1-10", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{"single item range", "1", []int{1}, true},
		{"invalid range", "1-2-3", []int{}, false},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			r, err := ParseRange(item.expr)
			if item.valid {
				assert.NoError(t, err, item.expr)
			} else {
				assert.Error(t, err, item.expr)
				return
			}

			have := r.Expand()
			assert.Equal(t, len(item.expected), len(have), item.expr)
			assert.EqualValues(t, item.expected, have, item.expr)
		})
	}
}

func TestParseRangeExpression(t *testing.T) {
	items := []struct {
		name     string
		expr     string
		expected []int
		valid    bool
	}{
		{"mix of simple and single digit ranges", "1-3,4,8-10", []int{1, 2, 3, 4, 8, 9, 10}, true},
		{"only single digits", "1,2,3", []int{1, 2, 3}, true},
		{"invalid trailing comma", "1-3,", []int{}, false},
		{"empty expression", "", []int{}, false},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			x, err := ParseRangeExpression(item.expr)
			if item.valid {
				assert.NoError(t, err, item.expr)
			} else {
				assert.Error(t, err, item.expr)
				return
			}

			have := []int{}
			for i := range x.Iterate() {
				have = append(have, i)
			}

			assert.Equal(t, len(item.expected), len(have), item.expr)
			assert.EqualValues(t, item.expected, have, item.expr)
		})
	}
}

func TestAddRange(t *testing.T) {
	x := NewRangeExpression()
	x.AddRange(NewRange(1, 3))
	x.AddRange(NewRange(7, 10))

	have := x.Expand()
	assert.EqualValues(t, []int{1, 2, 3, 7, 8, 9, 10}, have)
}

func TestRangeContains(t *testing.T) {
	r := NewRange(1, 1000)

	assert.True(t, r.Contains(100))
	assert.False(t, r.Contains(2000))
}

func TestRangeExpressionContains(t *testing.T) {
	x, err := ParseRangeExpression("1-20,50-100")
	assert.NoError(t, err)
	assert.True(t, x.Contains(10))
	assert.True(t, x.Contains(60))
	assert.False(t, x.Contains(30))
	assert.False(t, x.Contains(101))
}
