# Post Sale to Cloud.

This can be run with `silk -silk.url="http://paybox.work/api/v1/vending/sell"`

## POST /sell

Post sale transaction to cloud api.

===

* `Status: 200`
* `Content-Type: "application/json;charset=utf-8"`
* Accept: "application/json"

```
{
	"total": 120,
	"payment": 200,
	"change": 80,
	"type": "take_home",
	"sale_pay": {
		"B50": 2,
		"B100": 1
	},
	"sale_subs": [{
		"line": 1,
		"item_id": 0,
		"item_name": "Capuchino Hot",
		"qty": 1,
		"price_name": "small",
		"price": 60,
		"unit": "แก้ว"
	}, {
		"line": 2,
		"item_id": 1,
		"item_name": "Capuchino Ice",
		"qty": 2,
		"price_name": "large",
		"price": 80,
		"unit": "แก้ว"
	}]
}
```

===
* Status:200
* Content-Type: "application/json"
```
{
    "status":200,
    "message":"success",
    "result":{
        ...
    }
}

```