function menu_detail(lang,menuId){
    var api = 0;
    $.ajax({
            url: "http://"+window.location.host+"/menu/"+(parseInt(menuId)),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                    var items = result[parseInt(lang)].items;
                    var item = "";
                    for(var i = 0; i < items.length; i++){
                        var pri = [];
                        var prices = items[i].prices;
                        for (var p = 0; p < prices.length; p++) {
                            if (lang == 0) {
                                pri.push({'id': prices[p].id, 'name': prices[p].name, 'price': prices[p].price});
                            } else if (lang == 1) {
                                pri.push({'id': prices[p].id, 'name': prices[p].name_en, 'price': prices[p].price});
                            } else if (lang == 2) {
                                pri.push({'id': prices[p].id, 'name': prices[p].name_cn, 'price': prices[p].price});
                            }
                        }
                        item += '<div class="block-3" ';
                        item += "onclick='show_modal(" + i + ",\"" + items[i].Id + "\",\"" + items[i].name + "\",\"" + items[i].menu_seq + "\",\"/img/" + items[i].image + "\",\"" + items[i].unit + "\"";
                        item += "," + JSON.stringify(pri) + ")'>";
                        item += '<a href="#"><img src="/img/' + items[i].image + '" onError="this.src = \'/img/noimg.jpg\'" class="block-img">';
                    	item += '<span class="item-name">'+items[i].name;
                        item += '</span><span class="item-price">' + prices[0].price + ' ฿</span>';
                    	item += '</a></div>';
                    }
                    document.getElementById("list_item").innerHTML = item;

                    setTimeout(function(){
                        $.mobile.changePage("#page_item");
                    },100);
                    api = 1;
                },
                error: function (err){
                    console.log(JSON.stringify(err));
                }
          });
          setTimeout(function(){
              if(api==0){
                 alertify.error("กรุณารอการแก้ไขข้อบกพร่องสักครู่");
                 window.location.reload();
              }
          }, 300)
}

function show_modal(c_id, id, name, line, img, unit, price) {
    voice_say(localStorage.language, name);
     var img_ac = document.getElementsByClassName("active_menu");
     for(var i = 0; i < img_ac.length; i++){
         img_ac[c_id].style.display = "block";
     }

	var Mitem = id+' : '+name;
	var Mimg = '<img src="'+img+'" onError="this.src = \'/img/noimg.jpg\'" width="100%"/>';
	var size = "";
    for (s = 0; s < price.length; s++) {
        size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="' + price[s].name + '"';
        size += 'onclick="active_size(\'' + price[s].name + '\',\'' + price[s].price + '\',' + s + ')">';
        size += price[s].name + '</button>';
    }
    for (s = 0; s < price.length; s++) {
        if (price[s].price > 0) {
            var totalPrice = 1 * price[s].price;
            s = price.length;
        }
    }
    setTimeout(function () {
        for (s = 0; s < price.length; s++) {
            if (price[s].price > 0) {
                active_size(price[s].name, price[s].price, 0);
                document.getElementById("Msize").value = price[s].name;
                s = price.length;
            }
        }
    }, 300);
    document.getElementById("MitemNo").value = id;
    document.getElementById("MitemName").value = name;
    document.getElementById("line").value = line;
    document.getElementById("Munit").value = unit;
    document.getElementById("mo_qty").value = 1;
    if (Mitem.length < 25) {
        document.getElementById("Mitem_title").innerHTML = Mitem;
    } else {
        document.getElementById("Mitem_title").innerHTML = Mitem.substring(0, 25) + "...";
    }

    document.getElementById("Mimg").innerHTML = Mimg;
    document.getElementById("menusize").innerHTML = size;
    document.getElementById("mo-pri").value = totalPrice + ' ฿';

    $("#select_item").popup('open');
}


function addQty(){
	var qty = document.getElementById("mo_qty").value;
	var pri = document.getElementById("mo-pri").value;

	var price = "";
	price = pri.split(" ");
	price = price[0];

	var size_pri = parseInt(price)/parseInt(qty);

	var addQty = 0;
	if(qty>=1){
		addQty = parseInt(qty)+1;
	}
	document.getElementById("mo_qty").value = addQty;
	document.getElementById("pop-pri").innerHTML = addCommas(size_pri*addQty)+' ฿';
    document.getElementById("mo-pri").value = size_pri*addQty;
}

function removeQty(){
	var qty = document.getElementById("mo_qty").value;
	var pri = document.getElementById("mo-pri").value;
	var price = "";
	price = pri.split(" ");
	price = price[0];
	var size_pri = parseInt(price)/parseInt(qty);
	var addQty = 0;
	if(qty>=1){
		addQty = qty-1;
	}
	if(addQty<1){
		document.getElementById("mo_qty").value = 1;
		document.getElementById("pop-pri").innerHTML = addCommas(size_pri*1)+' ฿';
        document.getElementById("mo-pri").value = size_pri*1;
	}else{
		document.getElementById("mo_qty").value = addQty;
		document.getElementById("pop-pri").innerHTML = addCommas(size_pri*addQty)+' ฿';
        document.getElementById("mo-pri").value = size_pri*addQty;
	}
}

function active_size(name,price,id){
		var x = document.getElementsByClassName("ui-shadow");
        for (i = 0; i < x.length; i++) {
             if(id==i){
                x[i].style.background = "#0052cc";
             }else{
                x[i].style.background = "#66a3ff";
             }
        }
		var qty = document.getElementById("mo_qty").value;
		document.getElementById("Msize").value = name;
		var totalPrice = qty*price;
		document.getElementById("pop-pri").innerHTML = addCommas(totalPrice)+" ฿";
		document.getElementById("mo-pri").value = totalPrice;
}

function send_order(){
	$("#select_item").popup("close");
	    var itemCode = document.getElementById("MitemNo").value;
        var itemName = document.getElementById("MitemName").value;
        var line = document.getElementById("line").value;
        var qty = document.getElementById("mo_qty").value;
        var price = document.getElementById("mo-pri").value;
        var unit = document.getElementById("Munit").value;
        var size = document.getElementById("Msize").value;
        document.getElementById("pri3").value = price;
        document.getElementById("textpri3").innerHTML = addCommas(price);
        price = price.split(" ");
        doSend('{"Device":"host","type":"request","command":"onhand"}');
        var orderType = "coffee";
        var listOrder = '{"line":'+parseInt(line)+',"item_id":'+parseInt(itemCode)+',"item_name":"'+itemName+'","qty":'+parseInt(qty)+',"price_name":"'+size+'","price":'+parseInt(price)+',"unit":"'+unit+'"}';
        var output = '{"total":'+price+',"type":"'+orderType+'","sale_subs":['+listOrder+']}';
        $.ajax({
                url: "http://"+window.location.host+"/sale",
                data: output,
                contentType: "application/json; charset=utf-8",
                dataType: "json",
                type: "POST",
                cache: false,
                success: function(res){
                   console.log("payment sale "+JSON.stringify(res.result));
                },
                error: function(err){
                     console.log(JSON.stringify(err));
                }
        });
	setTimeout(function(){
		$("#pop_payment").popup('open');
	},300);
}

function bill(){
	$("#pop_payment").popup('close');
	setTimeout(function(){
		$("#pop_bill").popup('open');
	},100);
}

function cancel(){
    $("#pop_payment").popup('close');
    var cancel = '{"device":"host",';
                 cancel += '"type":"request",';
                 cancel += '"command":"cancel"';
                   // cancel += '"result" : true';
                 cancel += '}';
                 doSend(cancel);
}

function pop_back(){
    $("#select_item").popup("close");
    var img_ac = document.getElementsByClassName("active_menu");
    for(var i = 0; i < img_ac.length; i++){
        img_ac[i].style.display = "none";
    }
}

function addCommas(nStr)
{
    nStr += '';
    x = nStr.split('.');
    x1 = x[0];
    x2 = x.length > 1 ? '.' + x[1] : '';
    var rgx = /(\d+)(\d{3})/;
    while (rgx.test(x1)) {
        x1 = x1.replace(rgx, '$1' + ',' + '$2');
    }
    return x1 + x2;
}
