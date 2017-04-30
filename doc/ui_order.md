# Order entry waiting for Payment.

## Request Command "order"

* เมื่อ User เลือกสินค้าที่ต้องการ  UI จะแสดงหน้าจอ Order เพื่อปรับแก้รายละเอียด Order Item ที่ต้องการ เช่นปรับราคา หรือจำนวน
* ทั้งนี้ request ทีส่งมา Host จะสั่งเปิดรับธนบัตร และเหรียญไว้ เพื่อรอรับชำระ
* หากมีการหยอดเหรียญ หรือรับธนบัตร Host จะส่ง Event command "payment" เพื่อให้ UI เริ่มทำงานหน้า Payment ต่อไป และ Lock การแก้ไขรายละเอียดสินค้าทันที.

===


```
{
    "device": "host",
    "command": "order",
    "type": "request"
}
```

===

```
{

    "command":"order",
    "type": "response",
    "result": true
}

```