package rangex

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"
)

type (
	Range struct {
		left  int
		right int
	}

	RangeExpression struct {
		ranges []Range
	}
)

// NewRange returns a new Range value with the given left/right bounds.
func NewRange(left, right int) Range {
	return Range{
		left:  left,
		right: right,
	}
}

// ParseRange parses a range expression (<left>-<right>) and returns
// a new Range value.
func ParseRange(expr string) (Range, error) {
	parts := strings.Split(expr, "-")

	switch len(parts) {
	case 1:
		parts = append(parts, parts[0])
	case 2:
	default:
		return Range{}, fmt.Errorf("invalid range expression")
	}

	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return Range{}, fmt.Errorf("invalid lhs (%s)", parts[0])
	}

	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return Range{}, fmt.Errorf("invalid rhs (%s)", parts[1])
	}

	return NewRange(left, right), nil
}

func (r Range) String() string {
	if r.left == r.right {
		return fmt.Sprintf("<Range %d>", r.left)
	}
	return fmt.Sprintf("<Range %d-%d>", r.left, r.right)
}

// Iterate is an iterator over a range:
//
//	for i := r.Iterate() { ... }
func (r Range) Iterate() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r.left; i <= r.right; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Expand returns a list of all the values contained in the range.
func (r Range) Expand() []int {
	val := []int{}

	for i := range r.Iterate() {
		val = append(val, i)
	}

	return val
}

// Contains returns true if the given value is contained in the range,
// false otherwise.
func (r Range) Contains(val int) bool {
	return val >= r.left && val <= r.right
}

func (r Range) Left() int {
	return r.left
}

func (r Range) Right() int {
	return r.right
}

// Size returns the size (numer of integers) contained in the Range.
func (r Range) Size() int {
	return r.right - r.left + 1
}

func NewRangeExpression() *RangeExpression {
	return &RangeExpression{}
}

func ParseRangeExpression(expr string) (*RangeExpression, error) {
	r := RangeExpression{}

	parts := strings.SplitSeq(expr, ",")
	for part := range parts {
		parsed, err := ParseRange(part)
		if err != nil {
			return nil, fmt.Errorf("invalid range expression: %w", err)
		}

		r.AddRange(parsed)
	}

	return &r, nil
}

func (x *RangeExpression) Iterate() iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, r := range x.ranges {
			for i := range r.Iterate() {
				if !yield(i) {
					return
				}
			}
		}
	}
}

func (x *RangeExpression) AddRange(r Range) {
	x.ranges = append(x.ranges, r)
}

func (x *RangeExpression) Expand() []int {
	val := []int{}

	for i := range x.Iterate() {
		val = append(val, i)
	}

	return val
}

func (x *RangeExpression) Contains(val int) bool {
	for _, r := range x.ranges {
		if r.Contains(val) {
			return true
		}
	}

	return false
}

// Simplify returns a new RangeExpression with overlapping or adjacent ranges
// coalesced.
func (x *RangeExpression) Simplify() *RangeExpression {
	result := NewRangeExpression()

	if len(x.ranges) == 0 {
		return result
	}

	if len(x.ranges) == 1 {
		result.AddRange(x.ranges[0])
		return result
	}

	// Copy and sort ranges by left bound
	sorted := make([]Range, len(x.ranges))
	copy(sorted, x.ranges)
	slices.SortFunc(sorted, func(a, b Range) int {
		return a.Left() - b.Left()
	})

	// Merge overlapping and adjacent ranges
	merged := []Range{}
	current := sorted[0]

	for i := 1; i < len(sorted); i++ {
		next := sorted[i]

		// Check if ranges overlap or are adjacent
		if current.Right() >= next.Left()-1 {
			// Merge by extending current range
			if next.Right() > current.Right() {
				current = NewRange(current.Left(), next.Right())
			}
		} else {
			// Ranges don't overlap or touch, save current and move to next
			merged = append(merged, current)
			current = next
		}
	}

	// Add the final range
	merged = append(merged, current)

	for _, r := range merged {
		result.AddRange(r)
	}

	return result
}

// Size returns the size (numer of integers) contained in the RangeExpression.
func (x *RangeExpression) Size() int {
	size := 0
	for _, r := range x.ranges {
		size += r.Size()
	}
	return size
}
