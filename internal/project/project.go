package project

import (
	"fmt"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/types"
)

// Generic function instead of method
func GetUnit[T Unit](p *Project, unitType types.UnitType) (T, error) {
	var zero T

	i := slices.IndexFunc(p.Units, func(u Unit) bool {
		return u.UnitType() == unitType
	})

	if i == -1 {
		return zero, fmt.Errorf("unit type %v not found", unitType)
	}

	unit, ok := p.Units[i].(T)
	if !ok {
		return zero, fmt.Errorf("unit is not of type %T", zero)
	}

	return unit, nil
}
