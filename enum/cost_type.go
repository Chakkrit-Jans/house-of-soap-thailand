package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// ========== Cost Type ==========
// CostType cost type
type CostType string

// const available value for enum
const (
	None        CostType = ""
	CostPerItem CostType = "Cost per Item"
	CostPerQty  CostType = "Cost per Qty"
)

// Value validate enum when set to database
func (t CostType) Value() (driver.Value, error) {
	switch t {
	case None, CostPerItem, CostPerQty: //valid case
		return string(t), nil
	}
	return nil, errors.New("Invalid cost type value") //else is invalid
}

// Scan validate enum on read from data base
func (t *CostType) Scan(value interface{}) error {
	var ct CostType
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.([]uint8) // if we declare db type as ENUM gorm will scan value as []uint8
	if !ok {
		return errors.New("Invalid data for cost type")
	}
	ct = CostType(string(st)) //convert type from string to CostType

	switch ct {
	case CostPerItem, CostPerQty: //valid case
		*t = ct
		return nil
	}
	return fmt.Errorf("Invalid cost type value :%s", st) //else is invalid
}
