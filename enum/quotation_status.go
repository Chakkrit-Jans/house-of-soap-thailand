package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// ========== Cost Group ==========
// CostGroup cost type
type QuotationStatus string

// const available value for enum
const (
	Draft     QuotationStatus = "Draft"
	Completed QuotationStatus = "Completed"
	Approved  QuotationStatus = "Approved"
	Canceled  QuotationStatus = "Canceled"
	Rejected  QuotationStatus = "Rejected"
)

// Value validate enum when set to database
func (t QuotationStatus) Value() (driver.Value, error) {
	switch t {
	case Draft, Completed, Approved, Canceled, Rejected: //valid case
		return string(t), nil
	}
	return nil, errors.New("Invalid quotation status value") //else is invalid
}

// Scan validate enum on read from data base
func (t *QuotationStatus) Scan(value interface{}) error {
	var ct QuotationStatus
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.([]uint8) // if we declare db type as ENUM gorm will scan value as []uint8
	if !ok {
		return errors.New("Invalid data for quotation status")
	}
	ct = QuotationStatus(string(st)) //convert type from string to CostGroup

	switch ct {
	case Draft, Completed, Approved, Canceled, Rejected: //valid case
		*t = ct
		return nil
	}
	return fmt.Errorf("Invalid quotation status value :%s", st) //else is invalid
}
