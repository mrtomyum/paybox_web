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
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
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
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
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
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
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
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
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
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
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
										<a href="#"><div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
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


function active_size(id) {
 		$("h1").removeClass("acsize");
		$("#"+id).addClass("acsize");
		//alert('$("#'+id+'").addClass("acsize")');
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

function showmodal(){
	$("h1").removeClass("acsize");
	$("#s").addClass("acsize");
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