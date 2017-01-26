# Post Sale from web ui.

This can be run with `silk -silk.url="http://localhost:8888/sale"`

## POST /sale

Post sale transaction from web ui.

===

* `Status: 200`
* `Content-Type: "application/json;charset=utf-8"`
* Accept: "application/json"

```
{
  	"total": 120,
  	"type": "take_home",
  	"sale_subs": [{
  		"line": 1,
  		"item_id": 1,
  		"item_name": "Capuchino Hot",
  		"price_id": 1,
  		"price": 60,
  		"qty": 1,
  		"unit": "แก้ว"
  	}, {
  		"line": 2,
  		"item_id": 2,
  		"item_name": "Capuchino Ice",
  		"price_id": 2,
  		"price": 80,
  		"qty ": 2,
  		"unit": "แก้ว"
  	}]
  }
```

===
* Status:201
* Content-Type: "applicaiton/json"
```
{
    "command":"sale",
    "result":"success",
    "data": {
      	"total": 120,
      	"payment": 150,
      	"change": 30,
      	"type": "take_home",
      	"sale_subs": [{
      		"line": 1,
      		"item_id": 1,
      		"item_name": "Capuchino Hot",
      		"price_id": 1,
      		"price": 60,
      		"qty": 1,
      		"unit": "แก้ว"
      	}, {
      		"line": 2,
      		"item_id": 2,
      		"item_name": "Capuchino Ice",
      		"price_id": 2,
      		"price": 60,
      		"qty ": 2,
      		"unit": "แก้ว"
      	}]
    }
}

```