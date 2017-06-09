function menu_detail(lang, menuId) {
    console.log(lang + ", " + menuId);
    var api = 0;
    $.ajax({
        url: "http://" + window.location.host + "/menu/" + (parseInt(menuId)),
        //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        type: "GET",
        cache: false,
        success: function (result) {
            //console.log(JSON.stringify(result));
            var items = result[parseInt(lang)].items;
            // console.log("new "+JSON.stringify(items));
            var item = "";
            for (var i = 0; i < items.length; i++) {
                var pri = [];
                var prices = items[i].prices;
                //console.log(JSON.stringify(prices));
                for (var p = 0; p < prices.length; p++) {
                    if (lang == 0) {
                        pri.push({'id': prices[p].id, 'name': prices[p].name, 'price': prices[p].price});
                    } else if (lang == 1) {
                        pri.push({'id': prices[p].id, 'name': prices[p].name_en, 'price': prices[p].price});
                    } else if (lang == 2) {
                        pri.push({'id': prices[p].id, 'name': prices[p].name_cn, 'price': prices[p].price});
                    }

                }

                //size_item(pri);
                //console.log(JSON.stringify(pri));
                item += '<div class="block-3" ';
                // item += "onclick='size_item("+JSON.stringify(items)+")'>";
                item += "onclick='show_modal(" + i + ",\"" + items[i].Id + "\",\"" + items[i].name + "\",\"" + items[i].menu_seq + "\",\"/img/" + items[i].image + "\",\"" + items[i].unit + "\"";
                item += "," + JSON.stringify(pri) + ")'>";

                // item += '<img src="/img/life.png" class="active_menu" style="width:200px; height:100px;">';
                item += '<a href="#"><img src="/img/' + items[i].image + '" onError="this.src = \'/img/noimg.jpg\'" class="block-img">';
                item += '<span class="item-name">' + items[i].name;
                item += '</span><span class="item-price">' + prices[0].price + ' ฿</span>';
                item += '</a></div>';
            }
            //console.log(item);
            document.getElementById("list_item").innerHTML = item;
            // console.log(item);

            setTimeout(function () {
                $.mobile.changePage("#page_item");
                // voice_say(localStorage.language,mName);
            }, 100);
            api = 1;
        },
        error: function (err) {
            console.log(JSON.stringify(err));
        }
    });

    setTimeout(function () {
        if (api == 0) {
            alertify.error("กรุณารอการแก้ไขข้อบกพร่องสักครู่");
            window.location.reload();
        }
    }, 300)
    // console.log("http://"+window.location.host+"/menu/"+(parseInt(menuId)));

}

function show_modal(c_id, id, name, line, img, unit, price) {
// function show_modal(data) {
//     console.log(JSON.stringify(data));
    console.log(c_id + "," + id + "," + name + "," + line + "," + img + "," + unit);
    console.log(JSON.stringify(price));
    setTimeout(function () {
        voice_say(localStorage.language, name);
    }, 300);


    var Mitem = id + ' : ' + name;
    var Mimg = '<img src="' + img + '" onError="this.src = \'/img/noimg.jpg\'" width="100%"/>';
    /*
     var s = s.split("/");
     var sName = s[0];
     var sPrice = s[1];

     var m = m.split("/");
     var mName = m[0];
     var mPrice = m[1];

     var l = l.split("/");
     var lName = l[0];
     var lPrice = l[1];*/
    var size = "";
    //console.log("Org "+localStorage.OrgCode);
    // if(localStorage.OrgCode == 0){
    /*if( sPrice != "0"){
     size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="'+sName+'"';
     size += 'onclick="active_size(\''+sName+'\',\''+sPrice+'\',0)">';
     size += sName+'</button>';
     }

     if( mPrice != "0"){
     size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="'+mName+'"';
     size += 'onclick="active_size(\''+mName+'\',\''+mPrice+'\',1)">';
     size += mName+'</button>';
     }

     if( lPrice != "0"){
     size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="'+lName+'"';
     size += ' onclick="active_size(\''+lName+'\',\''+lPrice+'\',2)">';
     size += lName+'</button>';
     }*/
    for (s = 0; s < price.length; s++) {
        //console.log(s);
        size += '<div class="choice">';
        size += '<img src="img/150.png" class="pop-img" onclick="addQty(' + s + ',' + price[s].price + ')">';
        size += '<div class="name-choice" onclick="addQty(' + s + ',' + price[s].price + ')">' + price[s].name + '</div>';
        size += '<div class="unit">0</div>';
        size += '<div class="removeunit" onclick="removeQty(' + s + ',' + price[s].price + ')">-</div>';
        size += '</div>';
    }

    setTimeout(function () {
        addQty(0, price[0].price);
    }, 300);

    document.getElementById("MitemNo").value = id;
    document.getElementById("MitemName").value = name;
    document.getElementById("line").value = line;
    document.getElementById("Munit").value = unit;
    //document.getElementById("mo_qty").value = 1;

    document.getElementById("Mitem_title").innerHTML = Mitem;


    //document.getElementById("Mimg").innerHTML = Mimg;
    document.getElementById("order-list").innerHTML = size;

    /*document.getElementById("pri3").value = totalPrice;
    document.getElementById("textpri3").innerHTML = addCommas(totalPrice);
     */
    //$("#select_item").popup('open');
    $.mobile.changePage("#page_selectitem");
}


function addQty(id, price) {
    //console.log(id+","+price);
    var unit = document.getElementsByClassName("unit");
    var runit = document.getElementsByClassName("removeunit");
    unit[id].innerHTML = parseInt(unit[id].innerText) + 1;
    for (var i = 0; i < unit.length; i++) {
        if (i == id) {
            unit[i].style.display = "block";
            runit[i].style.display = "block";
        } else {
            unit[i].style.display = "none";
            runit[i].style.display = "none";
            unit[i].innerHTML = 0;
        }

    }
    if (parseInt(unit[id].innerText) == 1) {
        runit[id].style.display = "none";
    }

    var totalPrice = parseInt(unit[id].innerText) * parseInt(price);
    document.getElementById("textpri3").innerHTML = addCommas(totalPrice);
    document.getElementById("pri3").value = totalPrice;

}

function removeQty(id, price) {
    var unit = document.getElementsByClassName("unit");
    var runit = document.getElementsByClassName("removeunit");
    unit[id].innerHTML = parseInt(unit[id].innerText) - 1;
    for (var i = 0; i < unit.length; i++) {
        if (i == id) {
            unit[i].style.display = "block";
            runit[i].style.display = "block";
        } else {
            unit[i].style.display = "none";
            runit[i].style.display = "none";
            unit[i].innerHTML = 0;
        }

    }
    if (parseInt(unit[id].innerText) == 1) {
        runit[id].style.display = "none";
    }
    var totalPrice = parseInt(unit[id].innerText) * parseInt(price);
    document.getElementById("textpri3").innerHTML = addCommas(totalPrice);
    document.getElementById("pri3").value = totalPrice;
}

function active_size() {
    console.log(id + "," + price);
    /* document.getElementById("textpri3").innerHTML = addCommas(totalPrice);
     document.getElementById("pri3").value = totalPrice;*/
}

function send_order() {
    $("#select_item").popup("close");
    var itemCode = document.getElementById("MitemNo").value;
    var itemName = document.getElementById("MitemName").value;
    var line = document.getElementById("line").value;
    var qty = document.getElementById("mo_qty").value;
    var price = document.getElementById("pri3").value;
    var unit = document.getElementById("Munit").value;
    var size = document.getElementById("Msize").value;

    document.getElementById("pri3").value = price;
    document.getElementById("textpri3").innerHTML = addCommas(price);


    price = price.split(" ");

    console.log("send order : " + itemCode, itemName, line, size, qty, unit, price[0]);
    doSend('{"Device":"host","type":"request","command":"onhand"}');

    var orderType = "coffee";
    var listOrder = '{"line":' + parseInt(line) + ',"item_id":' + parseInt(itemCode) + ',"item_name":"' + itemName + '","qty":' + parseInt(qty) + ',"price_name":"' + size + '","price":' + parseInt(price) + ',"unit":"' + unit + '"}';
    var output = '{"total":' + price + ',"type":"' + orderType + '","sale_subs":[' + listOrder + ']}';
    //console.log(parseInt(pri2)+","+parseInt(pri1));
    console.log(output);
    $.ajax({
        url: "http://" + window.location.host + "/sale",
        data: output,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        type: "POST",
        cache: false,
        success: function (res) {
            console.log("payment sale " + JSON.stringify(res.result));
        },
        error: function (err) {
            console.log(JSON.stringify(err));
        }
    });

    setTimeout(function () {
        console.log("เปิดการใช้งาน ชำระเงิน");
        $("#pop_payment").popup('open');
    }, 300);
}

function bill() {
    $("#pop_payment").popup('close');
    setTimeout(function () {
        $("#pop_bill").popup('open');
    }, 100);
}

function cancel() {
    //$.mobile.changePage("#pageone");
    $("#pop_payment").popup('close');
    var cancel = '{"device":"host",';
    cancel += '"type":"request",';
    cancel += '"command":"cancel"';
    // cancel += '"result" : true';
    cancel += '}';
    doSend(cancel);
}

function pop_back() {
    $("#select_item").popup("close");
    var img_ac = document.getElementsByClassName("active_menu");
    for (var i = 0; i < img_ac.length; i++) {
        img_ac[i].style.display = "none";
    }
}

function addCommas(nStr) {
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
