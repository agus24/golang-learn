//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Orders struct {
	ID           int64 `sql:"primary_key"`
	Date         time.Time
	OrderNumber  string
	GrandTotal   int32
	CustomerName string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
