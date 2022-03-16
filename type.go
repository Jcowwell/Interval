package interval

import "golang.org/x/exp/constraints"

type Numeric interface {
	constraints.Float | constraints.Integer
}
