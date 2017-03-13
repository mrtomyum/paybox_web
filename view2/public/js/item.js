function menu_detail(lang,menuId){
	console.log(lang+", "+menuId);
    $.ajax({
            url: "http://"+window.location.host+"/menu/"+(parseInt(menuId)),
          //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                    //console.log(JSON.stringify(result));
                    var items = result[parseInt(lang)].items;;
                   // console.log("new "+JSON.stringify(items));
                    var item = "";
                    for(var i = 0; i < items.length; i++){
                    	var price = items[i].prices;
                    	item += '<div class="block-3"';
                    	item += 'onclick="show_modal(\''+items[i].Id+'\',\''+items[i].name+'\',\''+items[i].menu_seq+'\',\'/img/'+items[i].image+'\',\''+items[i].unit+'\',';
                    	item += '\''+price[0].name+'/'+price[0].price+'\'';
                    	item += ',\''+price[1].name+'/'+price[1].price+'\'';
                    	item += ',\''+price[2].name+'/'+price[2].price+'\')"><a href="#">';
                        item += '<img src="/img/life.png" class="active_menu" style="width:200px; height:100px;">';
                    	item += '<img src="/img/'+items[i].image+'" onError="this.src = \'/img/noimg.jpg\'" class="block-img">';
                    	item += '<span class="item-name">'+items[i].name;
                    	item += '</span><span class="item-price">'+price[0].price+' ฿</span>';
                    	item += '</a></div>';
                    }
                    					//console.log(item);
                    document.getElementById("list_item").innerHTML = item;
                },
                error: function (err){
                    console.log(JSON.stringify(err));
                }
          });

}

function show_modal(id,name,line,img,unit,s,m,l){

	console.log(id+","+name+","+line+","+img+","+unit+","+s+","+m+","+l);


	var Mitem = id+' : '+name;
	var Mimg = '<img src="'+img+'" onError="this.src = \'/img/noimg.jpg\'" width="100%"/>';

	var s = s.split("/");
	var sName = s[0];
	var sPrice = s[1];

	var m = m.split("/");
	var mName = m[0];
	var mPrice = m[1];

	var l = l.split("/");
	var lName = l[0];
	var lPrice = l[1];
	var size = "";
    if(localStorage.OrgCode == 0){
		if( sPrice != "0"){
		size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="'+sName+'"';
		size += 'onclick="active_size(\''+sName+'\',\''+sPrice+'\',0)">';
		size += 'Small</button>';
		}

		if( mPrice != "0"){
   		size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="'+mName+'"';
        size += 'onclick="active_size(\''+mName+'\',\''+mPrice+'\',1)">';
        size += 'Medium</button>';
        }

        if( lPrice != "0"){
        size += '<button class="ui-btn ui-shadow" style="font-size: 24px; color:#fff; background:#66a3ff; border-radius:10px;" id="'+lName+'"';
        size += ' onclick="active_size(\''+lName+'\',\''+lPrice+'\',2)">';
        size += 'large</button>';
        }
        setTimeout(function(){
            if(sPrice!=0){
               active_size(sName,sPrice,0);
            }else if(mPrice!=0){
               active_size(mName,mPrice,1);
            }else if(lPrice!=0){
               active_size(lName,lPrice,2);
            }
        },300);

    }else{

    }

    var totalPrice = 1*sPrice;

    document.getElementById("MitemNo").value = id;
    document.getElementById("MitemName").value = name;
    document.getElementById("line").value = line;
    document.getElementById("Munit").value = unit;
    document.getElementById("Msize").value = sName;
    document.getElementById("mo_qty").value = 1;
	document.getElementById("Mitem_title").innerHTML = Mitem.substring(0, 25)+"...";
	document.getElementById("Mimg").innerHTML = Mimg;
	document.getElementById("menusize").innerHTML = size;
/*	document.getElementById("mo-pri").value = totalPrice+' ฿';
*/
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
	document.getElementById("pop-pri").innerHTML = size_pri*addQty+' ฿';
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
		document.getElementById("pop-pri").innerHTML = size_pri*1+' ฿';
        document.getElementById("mo-pri").value = size_pri*1;
	//	document.getElementById("mo-pri").value = size_pri*1+' ฿';
	}else{
		document.getElementById("mo_qty").value = addQty;
		document.getElementById("pop-pri").innerHTML = size_pri*addQty+' ฿';
        document.getElementById("mo-pri").value = size_pri*addQty;
		//document.getElementById("mo-pri").value = size_pri*addQty+' ฿';
	}
}

function active_size(name,price,id){
	    /*$("h1").removeClass("acsize");
		$("#"+id).addClass("acsize");*/
		console.log(name+","+price+","+id);
		var x = document.getElementsByClassName("ui-shadow");
        for (i = 0; i < x.length; i++) {
             if(id==i){
                x[i].style.background = "#0052cc";
             }else{
                x[i].style.background = "#66a3ff";
             }
                    //   console.log("ui-shadow length "+ x.length)
        }

		var qty = document.getElementById("mo_qty").value;
		document.getElementById("Msize").value = name;

		var totalPrice = qty*price;

		console.log("ราคา " + totalPrice+", nameSize = "+name);
		document.getElementById("pop-pri").innerHTML = totalPrice+" ฿";
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
        document.getElementById("textpri3").innerHTML = price;


        price = price.split(" ");

        console.log("send order : "+itemCode,itemName,line,size,qty,unit,price[0]);
        doSend('{"Device":"host","type":"request","command":"onhand"}');

        var orderType = "coffee";
        var listOrder = '{"line":'+parseInt(line)+',"item_id":'+parseInt(itemCode)+',"item_name":"'+itemName+'","qty":'+parseInt(qty)+',"price_name":"'+size+'","price":'+parseInt(price)+',"unit":"'+unit+'"}';
        var output = '{"total":'+price+',"type":"'+orderType+'","sale_subs":['+listOrder+']}';
        //console.log(parseInt(pri2)+","+parseInt(pri1));
        console.log(output);
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
	    console.log("เปิดการใช้งาน ชำระเงิน");
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
    $.mobile.changePage("#pageone");
    var cancel = '{"device":"host",';
                 cancel += '"type":"request",';
                 cancel += '"command":"cancel"';
                   // cancel += '"result" : true';
                 cancel += '}';
                 doSend(cancel);
}