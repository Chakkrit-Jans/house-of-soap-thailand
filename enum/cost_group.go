package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// ========== Cost Group ==========
// CostGroup cost type
type CostGroup string

// const available value for enum
const (
	ProductionCost   CostGroup = "Production Cost"
	PackagingCost    CostGroup = "Packaging Cost"
	LabelCost        CostGroup = "Label Cost"
	ExtraFreightCost CostGroup = "Extra freight Cost"
	AdditionCost     CostGroup = "Addition Cost"
)

// Value validate enum when set to database
func (t CostGroup) Value() (driver.Value, error) {
	switch t {
	case ProductionCost, PackagingCost, LabelCost, ExtraFreightCost, AdditionCost: //valid case
		return string(t), nil
	}
	return nil, errors.New("Invalid cost group value") //else is invalid
}

// Scan validate enum on read from data base
func (t *CostGroup) Scan(value interface{}) error {
	var ct CostGroup
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.([]uint8) // if we declare db type as ENUM gorm will scan value as []uint8
	if !ok {
		return errors.New("Invalid data for cost group")
	}
	ct = CostGroup(string(st)) //convert type from string to CostGroup

	switch ct {
	case ProductionCost, PackagingCost, LabelCost, ExtraFreightCost, AdditionCost: //valid case
		*t = ct
		return nil
	}
	return fmt.Errorf("Invalid cost group value :%s", st) //else is invalid
}
