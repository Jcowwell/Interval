package interval

type PointType int // Point Enum Type

const (
	OpenPoint      PointType = iota // exclusive point
	ClosedPoint                     // inclusive point
	UnboundedPoint                  // infinity
)

/* Point Type to represent a number on a number line. */
type Point[N Numeric] struct {
	Value N
	Type  PointType
}

/* Private function that calculates which start point should be used for interval intersection */
func startPointIntersect[N Numeric](a, b Point[N]) Point[N] {
	if a.Value == b.Value {
		if a.Type < b.Type {
			return a
		} else {
			return b
		}
	} else if a.Value > b.Value {
		return a
	} else {
		return b
	}
}

/* Private function that calculates which end point should be used for interval intersection */
func endPointIntersect[N Numeric](c, d Point[N]) Point[N] {
	if c.Value == d.Value {
		if c.Type < d.Type {
			return c
		} else {
			return d
		}
	} else if c.Value < d.Value {
		return c
	} else {
		return d
	}
}
