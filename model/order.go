package model

import (
	//"time"
	"reflect"
	"fmt"
	"errors"
)

type Order struct {
	//Id    uint64
	//HostId uint64
	//Time  *time.Time
	Total     float64
	Payment   float64
	Change    float64
	OrderType string `json:"order_type" db:"order_type"`
	Items     []*OrderSub
}

type OrderSub struct {
	Line     uint64 `json:"line"`
	OrderId  uint64
	ItemId   uint64  `json:"item_id"`
	ItemName string  `json:"item_name"`
	SizeId   int     `json:"size_id"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	fmt.Printf("[func SetField] reflect.ValueOf(obj).Elem() name= %v ,value= %v \n", name, value)
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %o in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %o field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func (o *Order) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(o, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
