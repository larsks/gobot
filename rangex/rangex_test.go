package rangex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRange(t *testing.T) {
	items := []struct {
		expr     string
		expected []int
		valid    bool
	}{
		{"1-10", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{"1", []int{1}, true},
		{"1-2-3", []int{}, false},
	}

	for _, item := range items {
		r, err := ParseRange(item.expr)
		if item.valid {
			assert.NoError(t, err, item.expr)
		} else {
			assert.Error(t, err, item.expr)
			continue
		}

		have := r.Expand()
		assert.Equal(t, len(item.expected), len(have), item.expr)
		assert.EqualValues(t, item.expected, have, item.expr)
	}
}

func TestParseRangeExpression(t *testing.T) {
	items := []struct {
		expr     string
		expected []int
		valid    bool
	}{
		{"1-3,4,8-10", []int{1, 2, 3, 4, 8, 9, 10}, true},
		{"1,2,3", []int{1, 2, 3}, true},
		{"1-3,", []int{}, false},
		{"", []int{}, false},
	}

	for _, item := range items {
		x, err := ParseRangeExpression(item.expr)
		if item.valid {
			assert.NoError(t, err, item.expr)
		} else {
			assert.Error(t, err, item.expr)
			continue
		}

		have := []int{}
		for i := range x.Iterate() {
			have = append(have, i)
		}

		assert.Equal(t, len(item.expected), len(have), item.expr)
		assert.EqualValues(t, item.expected, have, item.expr)
	}
}

func TestAddRange(t *testing.T) {
	x := NewRangeExpression()
	x.AddRange(NewRange(1, 3))
	x.AddRange(NewRange(7, 10))

	have := x.Expand()
	assert.EqualValues(t, []int{1, 2, 3, 7, 8, 9, 10}, have)
}
