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

func TestSimplify(t *testing.T) {
	items := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "overlapping ranges",
			input:    "1-4,3-6",
			expected: "1-6",
		},
		{
			name:     "adjacent ranges",
			input:    "1-4,5-10",
			expected: "1-10",
		},
		{
			name:     "multiple overlaps",
			input:    "1-5,3-7,6-10",
			expected: "1-10",
		},
		{
			name:     "no overlaps",
			input:    "1-3,5-7,9-11",
			expected: "1-3,5-7,9-11",
		},
		{
			name:     "unsorted input",
			input:    "10-15,1-5,6-9",
			expected: "1-15",
		},
		{
			name:     "single range",
			input:    "1-5",
			expected: "1-5",
		},
		{
			name:     "completely contained",
			input:    "1-10,3-5",
			expected: "1-10",
		},
		{
			name:     "mixed overlapping and non-overlapping",
			input:    "1-4,3-5,10-15,12-18,25-30",
			expected: "1-5,10-18,25-30",
		},
		{
			name:     "single points",
			input:    "1,2,3,5,7-10",
			expected: "1-3,5,7-10",
		},
		{
			name:     "adjacent single points",
			input:    "1,3,2,4,5",
			expected: "1-5",
		},
		{
			name:     "all ranges merge into one",
			input:    "1-3,4-6,7-9",
			expected: "1-9",
		},
		{
			name:     "duplicate ranges",
			input:    "1-5,1-5,1-5",
			expected: "1-5",
		},
		{
			name:     "ranges with same start",
			input:    "1-3,1-10,1-5",
			expected: "1-10",
		},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			x, err := ParseRangeExpression(item.input)
			assert.NoError(t, err)

			simplified := x.Simplify()

			// Compare by expanding both and checking equality
			expected, err := ParseRangeExpression(item.expected)
			assert.NoError(t, err)
			assert.Equal(t, expected.Expand(), simplified.Expand())
		})
	}
}

func TestSimplifyDoesNotMutate(t *testing.T) {
	original, err := ParseRangeExpression("1-4,3-6,10-15")
	assert.NoError(t, err)
	originalExpanded := original.Expand()

	simplified := original.Simplify()

	// Verify original unchanged
	assert.Equal(t, originalExpanded, original.Expand())
	// Verify simplified is different
	assert.NotEqual(t, originalExpanded, simplified.Expand())
}

func TestSimplifyEmpty(t *testing.T) {
	x := NewRangeExpression()
	simplified := x.Simplify()
	assert.Equal(t, 0, len(simplified.Expand()))
}
