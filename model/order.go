package model

import (
	//"time"
	"reflect"
	"fmt"
	"errors"
	"time"
)

type Order struct {
	Id        int64
	Created   *time.Time
	HostId    string
	Total     float64
	Payment   float64
	Change    float64
	OrderType string `json:"order_type" db:"order_type"`
	IsPosted  bool `json:"is_posted" db:"is_posted"`
	Items     []*OrderSub
}

type OrderSub struct {
	Line     uint64 `json:"line"`
	OrderId  uint64
	ItemId   uint64  `json:"item_id"`
	ItemName string  `json:"item_name"`
	PriceId  int     `json:"price_id"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
	Unit     string
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

func (o *Order) Post() error {
	// Ping Server api.paybox.work:8080/ping
	// if post Error o.IsPosted = false
	// IsNetOnline => Post Order ขึ้น Cloud
	o.IsPosted = true
	return nil
}

func (o *Order) Save() error {
	sql1 := `INSERT INTO order(
		created,
		host_id,
		total,
		payment,
		change,
		order_type
		is_posted
	VALUES (?,?,?,?,?,?)`

	created := time.Now()
	rs, err := db.Exec(sql1,
		created,
		o.HostId,
		o.Total,
		o.Payment,
		o.Change,
		o.OrderType,
		o.IsPosted,
	)
	if err != nil {
		return err
	}
	o.Id, _ = rs.LastInsertId()

	os := OrderSub{}
	sql2 := `INSERT INTO order_sub(
		order_id,
		item_id,
		qty,
		price_id,
		price
	VALUES(?,?,?,?,?)`
	// Todo: Loop til end OrderSub
	rs, err = db.Exec(sql2,
		o.Id,
		os.Line,
		os.ItemId,
		os.ItemName,
		os.PriceId,
		os.Price,
		os.Qty,
		os.Unit,
	)
	if err != nil {
		return err
	}

	fmt.Println("h.OrderSave() run")
	return nil
}
