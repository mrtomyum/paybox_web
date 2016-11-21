$("document").ready(function(){
	var menuId = localStorage.menuId; 

	var id = localStorage.language;
	/*console.log("lang " + id);
	console.log("menuId " + menuId);*/

    $("img").removeClass("active_img");
    $("#l"+id).addClass("active_img");
    $("#r"+id).addClass("active_img");
    main_menu(id);
    item(id,menuId);
	active(menuId-1);


});

function main_menu(id){
	var mydata = jQuery.parseJSON(data);
    var menu = "";
    //console.log(JSON.stringify(mydata));
   // console.log(mydata[0].langId);
    var listmenu = mydata[id-1].menu;
    //console.log(listmenu);
    for (var i = 0; i < listmenu.length; i++) {
      menu += `<li onclick="active(`+i+`)"><a href="#" id="`+i+`">`+listmenu[i].name+`</a></li>`;
      
    }

    document.getElementById("li_menu").innerHTML = menu;
}

function item(lang,menuId){
	//console.log(lang+", "+menuId);
	switch( parseInt(menuId)){
		case 0 : var menu_one = jQuery.parseJSON(menu_hot);
				
				var listitem = menu_one[lang-1].itemList;
				//console.log(JSON.stringify(listitem));
				var item = "";
					for(var i = 0; i < listitem.length; i++){
						var size = listitem[i].size;
						item += `
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" 
											onclick="showmodal('`+listitem[i].id+`','`+listitem[i].name+`','/`+listitem[i].img+`',
											'`+size[0].name+'/'+size[0].price+`'
											,'`+size[1].name+'/'+size[1].price+`'
											,'`+size[2].name+'/'+size[2].price+`')">
											<img src="/`+listitem[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
									    	<h5 style="margin-top: 0;"><div style="width: 80%; float: left;"><b>`+listitem[i].name+`
									    	</b></div><div style="width: 20%; float: left;"><b>`+size[0].price+` ฿</b></div></h5>
										</div></a>

								`;	
					}
					//console.log(item);
				document.getElementById("list_item").innerHTML = item;
				break;

		case 1 : var menu_two = jQuery.parseJSON(menu_ice);
				
				var listitem = menu_two[lang-1].itemList;
				//console.log(JSON.stringify(listitem));
				var item = "";
					for(var i = 0; i < listitem.length; i++){
						var size = listitem[i].size;
						item += `
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" 
											onclick="showmodal('`+listitem[i].id+`','`+listitem[i].name+`','/`+listitem[i].img+`',
											'`+size[0].name+'/'+size[0].price+`'
											,'`+size[1].name+'/'+size[1].price+`'
											,'`+size[2].name+'/'+size[2].price+`')">
											<img src="/`+listitem[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
									    	<h5 style="margin-top: 0;"><div style="width: 80%; float: left;"><b>`+listitem[i].name+`
									    	</b></div><div style="width: 20%; float: left;"><b>`+size[0].price+` ฿</b></div></h5>
										</div></a>

								`;	
					}
					//console.log(item);
				document.getElementById("list_item").innerHTML = item;
				break;

		case 2 : var menu_three = jQuery.parseJSON(menu_frappe);
				
				var listitem = menu_three[lang-1].itemList;
				//console.log(JSON.stringify(listitem));
				var item = "";
					for(var i = 0; i < listitem.length; i++){
						var size = listitem[i].size;
						item += `
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" 
											onclick="showmodal('`+listitem[i].id+`','`+listitem[i].name+`','/`+listitem[i].img+`',
											'`+size[0].name+'/'+size[0].price+`'
											,'`+size[1].name+'/'+size[1].price+`'
											,'`+size[2].name+'/'+size[2].price+`')">
											<img src="/`+listitem[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
									    	<h5 style="margin-top: 0;"><div style="width: 80%; float: left;"><b>`+listitem[i].name+`
									    	</b></div><div style="width: 20%; float: left;"><b>`+size[0].price+` ฿</b></div></h5>
										</div></a>

								`;	
					}
					//console.log(item);
				document.getElementById("list_item").innerHTML = item;
				break;

		case 3 : var menu_four = jQuery.parseJSON(menu_kids);
				
				var listitem = menu_four[lang-1].itemList;
				//console.log(JSON.stringify(listitem));
				var item = "";
					for(var i = 0; i < listitem.length; i++){
						var size = listitem[i].size;
						item += `
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" 
											onclick="showmodal('`+listitem[i].id+`','`+listitem[i].name+`','/`+listitem[i].img+`',
											'`+size[0].name+'/'+size[0].price+`'
											,'`+size[1].name+'/'+size[1].price+`'
											,'`+size[2].name+'/'+size[2].price+`')">
											<img src="/`+listitem[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
									    	<h5 style="margin-top: 0;"><div style="width: 80%; float: left;"><b>`+listitem[i].name+`
									    	</b></div><div style="width: 20%; float: left;"><b>`+size[0].price+` ฿</b></div></h5>
										</div></a>

								`;	
					}
					//console.log(item);
				document.getElementById("list_item").innerHTML = item;
				break;

		case 4 : var menu_five = jQuery.parseJSON(menu_snake);
				
				var listitem = menu_five[lang-1].itemList;
				//console.log(JSON.stringify(listitem));
				var item = "";
					for(var i = 0; i < listitem.length; i++){
						var size = listitem[i].size;
						item += `
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" 
											onclick="showmodal('`+listitem[i].id+`','`+listitem[i].name+`','/`+listitem[i].img+`',
											'`+size[0].name+'/'+size[0].price+`'
											,'`+size[1].name+'/'+size[1].price+`'
											,'`+size[2].name+'/'+size[2].price+`')">
											<img src="/`+listitem[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
									    	<h5 style="margin-top: 0;"><div style="width: 80%; float: left;"><b>`+listitem[i].name+`
									    	</b></div><div style="width: 20%; float: left;"><b>`+size[0].price+` ฿</b></div></h5>
										</div></a>

								`;	
					}
					//console.log(item);
				document.getElementById("list_item").innerHTML = item;
				break;

		case 5 : var menu_two = jQuery.parseJSON(menu_ice);
				
				var listitem = menu_two[lang-1].itemList;
				//console.log(JSON.stringify(listitem));
				var item = "";
					for(var i = 0; i < listitem.length; i++){
						var size = listitem[i].size;
						item += `
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" 
											onclick="showmodal('`+listitem[i].id+`','`+listitem[i].name+`','/`+listitem[i].img+`',
											'`+size[0].name+'/'+size[0].price+`'
											,'`+size[1].name+'/'+size[1].price+`'
											,'`+size[2].name+'/'+size[2].price+`')">
											<img src="/`+listitem[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
									    	<h5 style="margin-top: 0;"><div style="width: 80%; float: left;"><b>`+listitem[i].name+`
									    	</b></div><div style="width: 20%; float: left;"><b>`+size[0].price+` ฿</b></div></h5>
										</div></a>

								`;	
					}
					//console.log(item);
				document.getElementById("list_item").innerHTML = item;
				break;

		default : console.log("error menu id "+menuId);
				break;
	}

}


function active_size(id,price) {
 		$("h1").removeClass("acsize");
		$("#"+id).addClass("acsize");
		var qty = document.getElementById("mo_qty").value;
		var totalPrice = qty*price;

		console.log("ราคา " + totalPrice);
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
	if(localStorage.getID!=null){
		console.log(localStorage.getID);
		document.getElementById(localStorage.getID).style.background = "#f3f3f4";
		document.getElementById(localStorage.getID).style.color = "#000";
		localStorage.getID = id;
	}else{
		localStorage.getID = id;
	}	
	console.log("menuId " + localStorage.getID);
	$("a").removeClass("active");
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
	x[id].style.background = "#272727";	
	document.getElementById(id).style.background = "#272727";
	document.getElementById(id).style.color = "#fff";
	document.getElementById("itemlist").style.background = "#272727";
	document.getElementById("centitem").style.background = "#272727";
	item(localStorage.language,localStorage.getID);
}

function showmodal(id,name,img,s,m,l){
	$("h1").removeClass("acsize");
	$("#s").addClass("acsize");

	//console.log(id+", "+name+", "+img+", "+s+", "+m+", "+l);
	

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
    document.getElementById("mo_qty").value = 1;
	document.getElementById("Mitem_title").innerHTML = Mitem;
	document.getElementById("Mimg").innerHTML = Mimg;
	document.getElementById("menusize").innerHTML = size;
	document.getElementById("mo-pri").value = totalPrice+` ฿`;

	$('#myModal').show();
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