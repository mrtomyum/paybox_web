package ctrl

//import (
//	"github.com/mrtomyum/paybox_terminal/model"
//)

//var (
//	o model.Order
//)

//func PostNewOrderSub(ctx *gin.Context) {
//	strItemId := ctx.PostForm("itemId")
//	itemId, _ := strconv.ParseUint(strItemId, 10, 64)
//	strSize := ctx.PostForm("size")
//	size, _ := strconv.ParseInt(strSize, 10, 8)
//	price := ctx.PostForm("price")
//	qty := ctx.PostForm("qty")
//
//	newItem := new(model.OrderSub)
//	newItem.ItemId = itemId
//	newItem.Price = size
//	newItem.Price = price
//	newItem.Qty = qty
//
//	o.Items = append(o.Items, newItem)
//	var total float64 = 0
//	for _, i := range o.Items {
//		sumItem := i.Price * i.Qty
//		total += sumItem
//	}
//	o.Total = total
//}
//
//func DeleteOrder(ctx *gin.Context) {
//	o = nil
//}
//
//func DeleteOrderItem(ctx *gin.Context) {
//	l := ctx.Param("line")
//	line, _ := strconv.ParseUint(l, 10, 64)
//	i := line - 1 // slice index start from 0
//	o.Items = append(o.Items[:i], o.Items[i + 1:]...)
//}
