package rangex

import (
	"fmt"
	"iter"
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

func NewRange(left, right int) Range {
	return Range{
		left:  left,
		right: right,
	}
}

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

func (r Range) Iterate() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r.left; i <= r.right; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func (r Range) Expand() []int {
	val := []int{}

	for i := range r.Iterate() {
		val = append(val, i)
	}

	return val
}

func (r Range) Contains(val int) bool {
	return val >= r.left && val <= r.right
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

		r.ranges = append(r.ranges, parsed)
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
