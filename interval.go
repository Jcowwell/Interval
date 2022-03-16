package interval

import (
	"math"
)

type IntervalType int // Interval Enum Type

const (
	EmptyInterval       IntervalType = iota
	DegenerateInterval               // [a,a] = {a}
	OpenInterval                     // (a,b) = {x | a < x < b}
	ClosedInterval                   // [a,b] = {x | a <= x <= b}
	OpenClosedInterval               // (a,b] = {x | a < x <= b}
	ClosedOpenInterval               // [a,b) = {x | a <= x < b}
	GreaterThanInterval              // (a,+∞) = {x | x > a}
	AtLeastInterval                  // [a,+∞) = {x | x >= a}
	LessThanInterval                 // (-∞,b) = {x | x < b}
	AtMostInterval                   // (-∞,b] = {x | x <= b}
	UnboundedInterval                // (-∞,+∞) = {x}
)

type Interval[N Numeric] struct {
	LowerBound Point[N] // start point of interval
	UpperBound Point[N] // end point of interval
	Values     []N      // Values of an interval. For unbounded intervals Values will be nil
	Type       IntervalType
}

/*
	Private construction function to create an interval.

	Parameters:
		lo Point[N] LowerBound Value
		hi Point[N]	Higherbound Value
	Return:
		Interval[N] Interval Struct
*/
func createInterval[N Numeric](start, end Point[N]) Interval[N] {
	if start.Value > end.Value {
		panic("The LowerBound endpoint cannot be higher than the UpperBound endpoint")
	}
	interval := Interval[N]{LowerBound: start, UpperBound: end}
	interval.setIntervalType()
	interval.setValues()
	return interval
}

/* Private void method to set an interval's type. */
func (self *Interval[N]) setIntervalType() {
	var loType, hiType PointType = self.LowerBound.Type, self.UpperBound.Type

	/* OpenInterval */
	if loType == OpenPoint && hiType == OpenPoint {
		/* Empty Check: (a,a) = {} */
		if self.LowerBound.Value == self.UpperBound.Value {
			self.Type = DegenerateInterval
			return
		}
		self.Type = OpenInterval

		/* ClosedInterval */
	} else if loType == ClosedPoint && hiType == ClosedPoint {
		/* Degenerate Check */
		if self.LowerBound.Value == self.UpperBound.Value {
			self.Type = DegenerateInterval
			return
		}
		self.Type = ClosedInterval

		/* OpenClosedInterval Interval */
	} else if loType == OpenPoint && hiType == ClosedPoint {
		/* Empty Check: (a,a] = {} */
		if self.LowerBound.Value == self.UpperBound.Value {
			self.Type = DegenerateInterval
			return
		}
		self.Type = OpenClosedInterval

		/* ClosedOpenInterval Interval */
	} else if loType == ClosedPoint && hiType == OpenPoint {
		/* Empty Check: [a,a) = {} */
		if self.LowerBound.Value == self.UpperBound.Value {
			self.Type = DegenerateInterval
			return
		}
		self.Type = ClosedOpenInterval

		/* GreaterThanInterval Interval */
	} else if loType == OpenPoint && hiType == UnboundedPoint {
		self.Type = GreaterThanInterval

		/* AtLeastInterval Interval */
	} else if loType == ClosedPoint && hiType == UnboundedPoint {
		self.Type = AtLeastInterval

		/* LessThanInterval Interval */
	} else if loType == UnboundedPoint && hiType == OpenPoint {
		self.Type = LessThanInterval

		/* AtMostInterval Interval */
	} else if loType == UnboundedPoint && hiType == ClosedPoint {
		self.Type = AtMostInterval

		/* UnboundedInterval Interval */
	} else if loType == UnboundedPoint && hiType == UnboundedPoint {
		self.Type = UnboundedInterval
	}
}

/*
	Private void method to set an interval's set of Values.

	NOTE:
	The initial implementaiton of this method populated endpoints with the interval's start and stop *inclusive* endpoints
	for a user to use (i.e iteration). This was problematic when met with edge cases such as (1,2) where,
	according to the previous logic, the endpoint slice would be []int{2,1}.
	This would defeat the original purpose of the method.
	Instead it will be a method to directy create a slice that holds the inclusive Values of valid bounded intervals.
*/
func (self *Interval[N]) setValues() {
	switch self.Type {
	case DegenerateInterval:
		self.Values = append(self.Values, self.LowerBound.Value)
	case OpenInterval:
		for n := self.LowerBound.Value + 1; n < self.UpperBound.Value; n++ {
			self.Values = append(self.Values, n)
		}
		break
	case ClosedInterval:
		for n := self.LowerBound.Value; n <= self.UpperBound.Value; n++ {
			self.Values = append(self.Values, n)
		}
		break
	case OpenClosedInterval:
		for n := self.LowerBound.Value + 1; n <= self.UpperBound.Value; n++ {
			self.Values = append(self.Values, n)
		}
		break
	case ClosedOpenInterval:
		for n := self.LowerBound.Value; n < self.UpperBound.Value; n++ {
			self.Values = append(self.Values, n)
		}
		break
	/* Unbounded (and Unbounded Hybrid) types cannot set a Value type since they are not finite */
	case EmptyInterval, GreaterThanInterval, AtLeastInterval, LessThanInterval, AtMostInterval, UnboundedInterval:
		break
	}
}

/* SECTION: Interval Generation Functions */

/*
	Public Construction Function to generate an interval.

	Parameters:
		LowerBound Point[N] 	Start endpoint of interval.
		HigherBound Point[N]	End endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateInterval[N Numeric](LowerBound, UpperBound Point[N]) Interval[N] {
	return createInterval(LowerBound, UpperBound)
}

/*

	Public Function to generate an Empty Interval.

	Return:
		Interval[N] Interval Struct

	FIXME: Add start and end endpoint parameters to function.
	Empty Intervals can be created when there is an intersection between
	two intervals that do not intersect (b_start > a_end || a_start > b_end)
	or if the interval intersection is  at the same point but one point is Open
	and the other Closed ex: {(2,3] ∩ (-1, 2] => (2,2]}
*/
func GenerateEmptyInterval[N Numeric]() Interval[N] {
	return Interval[N]{}
}

/*
	Public Function to generate an Open Interval.

	Parameters:
		start N 	Start endpoint of interval.
		end N		End endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateOpenInterval[N Numeric](start, end N) Interval[N] {
	LowerBound := Point[N]{Value: start, Type: OpenPoint}
	higherBound := Point[N]{Value: end, Type: OpenPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a Closed Interval.

	Parameters:
		start N 	Start endpoint of interval.
		end N		End endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateClosedInterval[N Numeric](start, end N) Interval[N] {
	LowerBound := Point[N]{Value: start, Type: ClosedPoint}
	higherBound := Point[N]{Value: end, Type: ClosedPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a OpenClosed Interval.

	Parameters:
		start N 	Start endpoint of interval.
		end N		End endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateOpenClosedInterval[N Numeric](start, end N) Interval[N] {
	LowerBound := Point[N]{Value: start, Type: OpenPoint}
	higherBound := Point[N]{Value: end, Type: ClosedPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a ClosedOpen Interval.

	Parameters:
		start N 	Start endpoint of interval.
		end N		End endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateClosedOpenInterval[N Numeric](start, end N) Interval[N] {
	LowerBound := Point[N]{Value: start, Type: ClosedPoint}
	higherBound := Point[N]{Value: end, Type: OpenPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a GreaterThan Interval.

	Parameters:
		start N 	Start endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateGreaterThanInterval[N Numeric](start N) Interval[N] {
	LowerBound := Point[N]{Value: start, Type: OpenPoint}
	higherBound := Point[N]{Value: N(math.Inf(0)), Type: UnboundedPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a AtLeast Interval.

	Parameters:
		start N 	Start endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateAtLeastInterval[N Numeric](start N) Interval[N] {
	LowerBound := Point[N]{Value: start, Type: ClosedPoint}
	higherBound := Point[N]{Value: N(math.Inf(0)), Type: UnboundedPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a LessThan Interval.

	Parameters:
		start N 	Start endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateLessThanInterval[N Numeric](end N) Interval[N] {
	LowerBound := Point[N]{Value: N(math.Inf(-1)), Type: UnboundedPoint}
	higherBound := Point[N]{Value: end, Type: OpenPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a AtMost Interval.

	Parameters:
		start N 	Start endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateAtMostInterval[N Numeric](end N) Interval[N] {
	LowerBound := Point[N]{Value: N(math.Inf(-1)), Type: UnboundedPoint}
	higherBound := Point[N]{Value: end, Type: ClosedPoint}
	return createInterval(LowerBound, higherBound)
}

/*
	Public Function to generate a Unbounded Interval.

	Parameters:
		start N 	Start endpoint of interval.
	Return:
		Interval[N] Interval Struct
*/
func GenerateUnboundedInterval[N Numeric]() Interval[N] {
	LowerBound := Point[N]{Value: N(math.Inf(-1)), Type: UnboundedPoint}
	higherBound := Point[N]{Value: N(math.Inf(0)), Type: UnboundedPoint}
	return createInterval(LowerBound, higherBound)
}

/* !SECTION: Interval Generation Functions */

/*
	Public Boolean Method that returns true if a Value is within an interval. False otherwise.

	Parameters:
		Value N 	Value of type Integer
	Return:
		bool
*/
func (self *Interval[N]) Contains(Value N) bool {
	switch self.Type {
	case EmptyInterval:
		return false
	case OpenInterval:
		return self.LowerBound.Value < Value && self.UpperBound.Value > Value
	case ClosedInterval:
		return self.LowerBound.Value <= Value && self.UpperBound.Value >= Value
	case OpenClosedInterval:
		return self.LowerBound.Value < Value && self.UpperBound.Value >= Value
	case ClosedOpenInterval:
		return self.LowerBound.Value <= Value && self.UpperBound.Value > Value
	case GreaterThanInterval:
		return self.LowerBound.Value < Value
	case AtLeastInterval:
		return self.LowerBound.Value <= Value
	case LessThanInterval:
		return self.UpperBound.Value > Value
	case AtMostInterval:
		return self.UpperBound.Value >= Value
	case UnboundedInterval:
		return true
	}
	return false
}

/*
	Public Integer Method that returns the amount of numbers in an interval. Unbounded intervals
	returns math.inf

	Return:
		int
*/
func (self *Interval[N]) Count() int {
	switch self.Type {
	case EmptyInterval:
		return 0
	case DegenerateInterval:
		return 1
	case OpenInterval, ClosedInterval, OpenClosedInterval, ClosedOpenInterval:
		return len(self.Values)
	case GreaterThanInterval, AtLeastInterval, LessThanInterval, AtMostInterval, UnboundedInterval:
		return int(math.Inf(0))
	}
	return 0
}

/*
	Public Interval Function that returns the intersect (∩) between two intervals.

	Parameters:
		a Interval[N]
		b Interval[N]
	Return:
		c Interval[N]	a ∩ b

*/
func Intersect[N Numeric](a, b Interval[N]) Interval[N] {
	/* an intersection involving an Empty Interval always results in an Empty Interval */
	if a.Type == EmptyInterval || b.Type == EmptyInterval {
		return GenerateEmptyInterval[N]()
	}
	/* adopted from: https://stackoverflow.com/a/325964/6427171. Checks if the two intervals intersect. */
	if Max(a.LowerBound.Value, b.LowerBound.Value) > Min(a.UpperBound.Value, b.UpperBound.Value) {
		return GenerateEmptyInterval[N]()
	} else { /* the intervals intersect */
		start, end := startPointIntersect(a.LowerBound, b.LowerBound), endPointIntersect(a.UpperBound, b.UpperBound)
		return GenerateInterval(start, end)
	}
}
