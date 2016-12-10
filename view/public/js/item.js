var listOrder = [];
$("document").ready(function(){
    var menuId = localStorage.menuId;
	var id = localStorage.language;
	var status = localStorage.action;

    call_websocket();
    //websocket.onopen();
   /* onHend = setInterval(function(){ doSend(`{"job":"onHand"}`);},1000);*/
   // onHend = setInterval(function(){  websocket.onmessage(); },1000);

      var worker = new Worker('/js/time.js');
                 worker.onmessage = function (event) {
                 document.getElementById('timer').innerText =event.data ;
                 document.getElementById('timer2').innerText =event.data;
                 };

	switch(parseInt(id)){
	    case 1: if(status==1){ status = "รับประทานที่ร้าน"}else{ status = "ซื้อกลับบ้าน"}
	            document.getElementById("status").innerHTML = "สถานะ : "+status;

	            document.getElementById("txtTotalPri").innerHTML = "ราคารวม";
	            document.getElementById("txtmacPri").innerHTML = "จำนวนเงิน";
	            document.getElementById("txtUnit").innerHTML = "บาท";
	            document.getElementById("txtUnit2").innerHTML = "บาท";

	            document.getElementById("bt_payment").innerHTML = "ชำระเงิน";
	            document.getElementById("bt_print").innerHTML = "พิมพ์ใบเสร็จ";
	            document.getElementById("bt_cancel").innerHTML = "ยกเลิก";

	            document.getElementById("default_order").innerHTML = "** กรุณาเลือกรายการ **";
	            document.getElementById("title_order").innerHTML = "รายการ";

	            document.getElementById("version").innerHTML = "เวอร์ชั่น 0.1";
	            document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.1 ";

                 document.getElementById("Name_time").innerHTML = "เวลา ";
                 document.getElementById("Name_time2").innerHTML = "เวลา ";

	            break;
	    case 2: if(status==1){ status = "take this"}else{ status = "this out"}
                document.getElementById("status").innerHTML = "status : "+status;

                document.getElementById("txtTotalPri").innerHTML = "Total";
                document.getElementById("txtmacPri").innerHTML = "Payment";
                document.getElementById("txtUnit").innerHTML = "baht";
                document.getElementById("txtUnit2").innerHTML = "baht";

                document.getElementById("bt_payment").innerHTML = "Payment";
                document.getElementById("bt_print").innerHTML = "Print Lipts";
                document.getElementById("bt_cancel").innerHTML = "cancel";

                document.getElementById("default_order").innerHTML = "** Please select an item **";
                document.getElementById("title_order").innerHTML = "Order";

                document.getElementById("version").innerHTML = "version 0.1";
                document.getElementById("version2").innerHTML = "version 0.1 ";

                document.getElementById("Name_time").innerHTML = "time ";
                document.getElementById("Name_time2").innerHTML = "time ";
	            break;
	    case 3: if(status==1){ status = "拿著它"}else{ status = "取出"}
               	document.getElementById("status").innerHTML = "狀態 : "+status;

               	document.getElementById("txtTotalPri").innerHTML = "總價";
               	document.getElementById("txtmacPri").innerHTML = "付款";
               	document.getElementById("txtUnit").innerHTML = "銖";
               	document.getElementById("txtUnit2").innerHTML = "銖";

               	document.getElementById("bt_payment").innerHTML = "付款";
               	document.getElementById("bt_print").innerHTML = "打印收據";
               	document.getElementById("bt_cancel").innerHTML = "取消";

               	document.getElementById("default_order").innerHTML = "** 請選擇一個項目 **";
               	document.getElementById("title_order").innerHTML = "名單";

               	document.getElementById("version").innerHTML = "版本 0.1";
               	document.getElementById("version2").innerHTML = "版本 0.1 ";

               	document.getElementById("Name_time").innerHTML = "時間 ";
                document.getElementById("Name_time2").innerHTML = "時間 ";
	            break;
	}

      responsiveVoice.OnVoiceReady = function() {
          console.log("speech time?");
          if(localStorage.nName!="null"){
          responsiveVoice.setDefaultVoice(localStorage.lName);
          responsiveVoice.speak(localStorage.nName);
          }
          localStorage.lName = null;
          localStorage.nName = null;
        };
    main_menu(id);
    console.log("menuid "+menuId);
    item(id,(parseInt(menuId)-1));
    disabled_payment();

});

function disabled_payment(){
    if(JSON.stringify(listOrder)=="[]"){
        console.log(JSON.stringify(listOrder));
        console.log("listOrder is isset");
         payment1 = function() {
             alert("ท่านยังไม่ได้เลือกรายการที่ต้องการ");
         }
    }else{
        console.log("listOrder is empty");
        payment1 = function() {
           payment();
        }
    }
}

function main_menu(id){
	//var mydata = jQuery.parseJSON(data);
	console.log("menuid "+id);
     $.ajax({
            url: "http://"+window.location.host+"/menu",
          //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                        var menu = "";
                        //console.log(JSON.stringify(result));
                       // console.log(mydata[0].langId);
                        var listmenu = result[id-1].menus;
                        var sId = localStorage.menuId-1;
                        //console.log(listmenu);
                        for (var i = 0; i < listmenu.length; i++) {
                          if(localStorage.menuId==i){
                            menu += `<li onclick="active(`+i+`)"><a href="#" id="`+i+`">`+listmenu[i].name+`</a></li>`;
                          }else{
                            menu += `<li onclick="active(`+i+`)"><a href="#" id="`+i+`">`+listmenu[i].name+`</a></li>`;
                          }
                        }

                        document.getElementById("li_menu").innerHTML = menu;
                        $("a").removeClass("active");
                        $("#"+parseInt(sId)).addClass("active");
                        var x = document.getElementsByTagName("LI");
                        	for(var i = 0;i < x.length; i++){
                        		x[i].style.background = "#272727";
                        		if(i!=parseInt(sId)){
                        			//console.log(i);
                        			x[i].style.borderBottom = "1px dashed #fff";
                        		}else{
                        			x[i].style.borderBottom = "0px dashed #fff";
                        		}
                        	}
                        //	console.log(sId);
                        $("#"+parseInt(sId)).css("background-color", "#272727");
                        $("#"+parseInt(sId)).css("color", "#fff");
                        document.getElementById("itemlist").style.background = "#272727";
                        document.getElementById("centitem").style.background = "#272727";
               },
                error: function(err){
                    console.log(JSON.stringify(err));
                }
            });

}

function item(lang,menuId){
	console.log(lang+", "+menuId);
    $.ajax({
            url: "http://"+window.location.host+"/menu/"+(parseInt(menuId)+1),
          //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                    //console.log(JSON.stringify(result));
                    var items = result[parseInt(lang)-1].items;;
                    console.log("new "+JSON.stringify(items));
                    var item = "";
                    for(var i = 0; i < items.length; i++){
                    	var size = items[i].sizes;
                    	item += `<a href="#"><div class="block-3"
                    			onclick="showmodal('`+items[i].id+`','`+items[i].name+`','/img/`+items[i].image+`','`+items[i].unit+`',
                    			'`+size[0].name+'/'+size[0].price+`'
                    			,'`+size[1].name+'/'+size[1].price+`'
                    			,'`+size[2].name+'/'+size[2].price+`')">
                    			<img src="/img/`+items[i].image+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
                    			<h5 style="margin-top: 0;">
                    			<div style="width: 80%; float: left;">`+items[i].name+`
                    			</div><div style="width: 20%; float: left; padding:0;">`+size[0].price+` ฿</div></h5>
                    			</div></a>`;
                    }
                    					//console.log(item);
                    document.getElementById("list_item").innerHTML = item;
                },
                error: function (err){
                    console.log(JSON.stringify(err));
                }
          });

}


function active_size(id,price) {
 		$("h1").removeClass("acsize");
		$("#"+id).addClass("acsize");
		var qty = document.getElementById("mo_qty").value;
		document.getElementById("Msize").value = id;

		var totalPrice = qty*price;


		console.log("ราคา " + totalPrice+", nameSize = "+id);
		document.getElementById("mo-pri").value = totalPrice+` ฿`;
		//alert('$("#'+id+'").addClass("acsize")');
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
	document.getElementById("mo-pri").value = size_pri*addQty+` ฿`;
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
		document.getElementById("mo-pri").value = size_pri*1+` ฿`;
	}else{		
		document.getElementById("mo_qty").value = addQty;
		document.getElementById("mo-pri").value = size_pri*addQty+` ฿`;
	}
}


function active(id) {
console.log("active id "+id);
	if(localStorage.getID!=null){
		console.log(localStorage.getID);
		document.getElementById(localStorage.getID).style.background = "#f3f3f4";
		document.getElementById(localStorage.getID).style.color = "#000";
		/*$("#"+localStorage.getID).css("background-color", "#f3f3f4");
		$("#"+localStorage.getID).css("color", "#000");*/
		localStorage.getID = id;
		localStorage.menuId = id+1;
	}else{
		localStorage.getID = id;
		localStorage.menuId = id+1;
	}
	console.log("menuId " + localStorage.getID);
	$("a").removeClass("active");
	$("a").removeAttr("style");
	$("li").removeAttr("style");
	$("#"+id).addClass("active");
	var x = document.getElementsByTagName("LI");

	for(var i = 0;i < x.length; i++){
		x[i].style.background = "#272727";
		if(i!=id){
			//console.log(i);
			x[i].style.borderBottom = "1px dashed #fff";
		}else{
			x[i].style.borderBottom = "0px dashed #fff";
		}
	}
	//x[id].style.background = "#272727";
	/*document.getElementById(id).style.background = "#272727";
	document.getElementById(id).style.color = "#fff";*/
	$("#"+id).css("background-color", "#272727");
    $("#"+id).css("color", "#fff");
	document.getElementById("itemlist").style.background = "#272727";
	document.getElementById("centitem").style.background = "#272727";
	item(localStorage.language,localStorage.getID);
}

function showmodal(id,name,img,unit,s,m,l){
	$("h1").removeClass("acsize");
	$("#s").addClass("acsize");

	//console.log(id+", "+name+", "+img+", "+unit+", "+s+", "+m+", "+l);
	

	var Mitem = id+` : `+name;
	var Mimg = `<img class="crop" src="`+img+`" onError="this.src = '/img/noimg.jpg'" width="100%"/>`;

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

		if( sPrice != "0"){
		size += `<a href="#"><h1 id="`+sName+`" onclick="active_size('`+sName+`','`+sPrice+`')" class="acsize">
					<img src="/img/s.png" class="img-size"><b>Small </b></h1>
				</a>`;
		}

		if( mPrice != "0"){
   		size += `<a href="#"><h1 id="`+mName+`" onclick="active_size('`+mName+`','`+mPrice+`')">
          			<img src="/img/m.png" class="img-size"><b>Medium </b></h1>
          		</a>`;
        }

        if( lPrice != "0"){
        size += `<a href="#"><h1 id="`+lName+`" onclick="active_size('`+lName+`','`+lPrice+`')">
          			<img src="/img/l.png" class="img-size"><b>large </b></h1>
          		</a>`;
        }
    var totalPrice = 1*sPrice;

    document.getElementById("MitemNo").value = id;
    document.getElementById("MitemName").value = name;
    document.getElementById("Munit").value = unit;
    document.getElementById("Msize").value = sName;
    document.getElementById("mo_qty").value = 1;
	document.getElementById("Mitem_title").innerHTML = Mitem;
	document.getElementById("Mimg").innerHTML = Mimg;
	document.getElementById("menusize").innerHTML = size;
	document.getElementById("mo-pri").value = totalPrice+` ฿`;

	 $('#myModal').modal('show');
}

function send_order(){
    var itemCode = document.getElementById("MitemNo").value;
    var itemName = document.getElementById("MitemName").value;
    var qty = document.getElementById("mo_qty").value;
    var price = document.getElementById("mo-pri").value;
    var unit = document.getElementById("Munit").value;
    var size = document.getElementById("Msize").value;

    price = price.split(" ");

    order_list(itemCode,itemName,size,qty,unit,price[0]);
    $('#myModal').modal('toggle');
   // console.log("itemCode = "+itemCode+", itemName = "+itemName+", qty = "+qty+", size ="+size+", unit = "+unit+", price = "+price);
}

function order_list(itemCode,itemName,size,qty,unit,price){
   // console.log(itemCode+","+itemName+","+size+","+qty+","+unit+","+price);
   listOrder.push({"item_code":itemCode,"item_name":itemName,"item_size":size,"qty":qty,"unit":unit,"price":price});

   console.log(JSON.stringify(listOrder));
   var list = "";
   var totalPrice = 0;
   for(var i = 0; i < listOrder.length; i++){
       list += `
                  <label class="orderlist">
                       <div class="ordername" onclick="return false;"> `+listOrder[i].item_name+` `+listOrder[i].item_size+`</div>
                       <div class="orderqty" onclick="return false;">`+listOrder[i].qty+` `+listOrder[i].unit+`</div>
                       <div class="orderprice" onclick="return false;">`+listOrder[i].price+` ฿</div>
                       <div class="ordercancel">
                       <button class="btn btn-danger btn-xs"  onclick="item_cancel(`+i+`)" style="padding-left: 22.5%; padding-right: 22.5%;">-
                       </button>
                     </div>
                  </label>
                `;
       totalPrice += parseInt(listOrder[i].price);
   }
    console.log(formatMoney(totalPrice));
    var ttPrice = formatMoney(totalPrice);
   document.getElementById("pri1").value = ttPrice;
   document.getElementById("order_list").innerHTML = list;
   console.log(list);
   disabled_payment();
}

function item_cancel(index){
   listOrder.splice(index, 1);
   alert("ยกเลิกรายการที่ท่านต้องการเรียบร้อย");

     var list = "";
      var totalPrice = 0;
      for(var i = 0; i < listOrder.length; i++){
          list += `
                     <label class="orderlist" onclick="return false">
                          <div class="ordername"> `+listOrder[i].item_name+` `+listOrder[i].item_size+`</div>
                          <div class="orderqty">`+listOrder[i].qty+` `+listOrder[i].unit+`</div>
                          <div class="orderprice">`+listOrder[i].price+` ฿</div>
                          <div class="ordercancel">
                          <button class="btn btn-danger btn-xs" onclick="item_cancel(`+i+`)" style="padding-left: 22.5%; padding-right: 22.5%;">-
                          </button>
                        </div>
                     </label>
                   `;
          totalPrice += parseInt(listOrder[i].price);
      }
       console.log(formatMoney(totalPrice));
       var ttPrice = formatMoney(totalPrice);
      document.getElementById("pri1").value = ttPrice;
         if(list!=""){
              list = list;
         }else{
              list = `<label class="orderlist"><h4 style='color:red; text-align:center; padding:0; width:95%;'>** กรุณาเลือกรายการ **</h4></label>`;
         }
      document.getElementById("order_list").innerHTML = list;
     // console.log(list);
     disabled_payment();

}

function formatMoney(inum){  // ฟังก์ชันสำหรับแปลงค่าตัวเลขให้อยู่ในรูปแบบ เงิน
    var s_inum=new String(inum);
    var num2=s_inum.split(".");
    var n_inum="";
    if(num2[0]!=undefined){
        var l_inum=num2[0].length;
        for(i=0;i<l_inum;i++){
            if(parseInt(l_inum-i)%3==0){
                if(i==0){
                    n_inum+=s_inum.charAt(i);
                }else{
                    n_inum+=","+s_inum.charAt(i);
                }
            }else{
                n_inum+=s_inum.charAt(i);
            }
        }
    }else{
        n_inum=inum;
    }
    if(num2[1]!=undefined){
        n_inum+="."+num2[1];
    }
    return n_inum;
}

function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female");
    responsiveVoice.speak("English language");
 
    $("img").removeClass("active_img");
    $("#l"+id).addClass("active_img");
    $("#r"+id).addClass("active_img");
    localStorage.language = 1;

    main_menu(localStorage.language);
	active(localStorage.getID);
	item(localStorage.language,localStorage.getID);
    //detailmenu(id);
}

function onsaythai(id){
    responsiveVoice.setDefaultVoice("Thai Female");
    responsiveVoice.speak("ภาษาไทย");

    $("img").removeClass("active_img");
    $("#l"+id).addClass("active_img");
    $("#r"+id).addClass("active_img");
    localStorage.language = 2;

    main_menu(localStorage.language);
	active(localStorage.getID);
	item(localStorage.language,localStorage.getID);
    //detailmenu(id);
}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female");
    responsiveVoice.speak("中國"); 

    $("img").removeClass("active_img");
    $("#l"+id).addClass("active_img");
    $("#r"+id).addClass("active_img");
    localStorage.language = 3;

    main_menu(localStorage.language);
	active(localStorage.getID);
	item(localStorage.language,localStorage.getID);
    //detailmenu(id);
}
var onHend = "";
function payment(){
    var bt_payment = document.getElementsByClassName("Payment");
    var bt_print = document.getElementsByClassName("print");

    var block = document.getElementsByClassName("block-3");
    for(var i = 0; i < block.length; i++){
        block[i].style.pointerEvents = 'none';
    }

    $(".btn-xs").removeAttr("onclick");
    $(".btn-xs").addClass("disabled");


    anchors = document.querySelectorAll('.ul_menu li'),
    elems = {};
    [].forEach.call(anchors, function(anchor) {
        elems[anchor.id] = anchor.getAttribute('onclick');
        anchor.setAttribute('onclick', '');
    });

    bt_payment[0].style.display = "none";
    bt_print[0].style.display = "inline-block";
  /*  call_websocket();
    websocket.onopen();
    onHend = setInterval(function(){ doSend(`{"job":"onHand"}`);},1000);*/


    //console.log("kkk");
}

function print(){
    var pri1 = document.getElementById("pri1").value;
    var pri2 = document.getElementById("pri2").value;

    var changeMoney = parseInt(pri2)-parseInt(pri1);
   //console.log(changeMoney);
    if(localStorage.action==2){
        var take_home = true;
    }else{
        var take_home = false;
    }
    var output = [];
    output.push({"device":window.location.host,"payload":{"type":"request","command":"billing","result": true,"data":{"total":pri1,"payment":pri2,"change":changeMoney,"take_home":take_home,"items":listOrder}}});
    console.log(JSON.stringify(output));
    if(changeMoney<0){
        alert("ยอดเงินไม่พอชำระ");
    }else{
        alert(JSON.stringify(output));
        var str = JSON.stringify(output);
        var i = str.length-1;
        var res = str.substring(1,i);
        doSend(res);
        clearInterval(onHend);
    }

}

function cancel_menu(){
    var cancel = '{"device": "'+window.location.host+'",';
        cancel += '"payload":';
        cancel += '{';
        cancel += '"type" : "request",';
        cancel += '"command" : "cancel"';
       // cancel += '"result" : true';
        cancel += '}';
        cancel += '}';
    doSend(cancel);
}