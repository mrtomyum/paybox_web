$(document).ready(function(){


    var id = localStorage.language;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");

      responsiveVoice.OnVoiceReady = function() {
              console.log("speech time?");
              if(localStorage.nName!="null"){
              responsiveVoice.setDefaultVoice(localStorage.lName);
              responsiveVoice.speak(localStorage.nName);
              }
              localStorage.lName = null;
              localStorage.nName = null;
            };

            var worker = new Worker('/js/time.js');
                 worker.onmessage = function (event) {
                 document.getElementById('timer').innerText =event.data ;
                 document.getElementById('timer2').innerText =event.data;
                 };


	switch(parseInt(id)){
	    case 1:
	            document.getElementById("version").innerHTML = "เวอร์ชั่น 0.1";
	            document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.1 ";

	            document.getElementById("bt_back").innerHTML = "ย้อนกลับ";

                 document.getElementById("Name_time").innerHTML = "เวลา ";
                 document.getElementById("Name_time2").innerHTML = "เวลา ";

	            break;
	    case 2:
                document.getElementById("version").innerHTML = "version 0.1";
                document.getElementById("version2").innerHTML = "version 0.1 ";

                document.getElementById("bt_back").innerHTML = "back";

                 document.getElementById("Name_time").innerHTML = "time ";
                 document.getElementById("Name_time2").innerHTML = "time ";
	            break;
	    case 3:
               	document.getElementById("version").innerHTML = "版本 0.1";
               	document.getElementById("version2").innerHTML = "版本 0.1 ";

               	document.getElementById("bt_back").innerHTML = "背部";

               	document.getElementById("Name_time").innerHTML = "時間 ";
                document.getElementById("Name_time2").innerHTML = "時間 ";

	            break;
	}

    console.log(id);
     if(localStorage.ColorCode){
                 var nav = document.getElementsByClassName("navbar");
                    for(var i = 0; i < nav.length; i++){
                         nav[i].style.backgroundColor = localStorage.ColorCode;
                    }
         }
     if(localStorage.OrgCode == 0){
         detailmenu(id);
     }else if(localStorage.OrgCode == 1){
       //  window.location = "menu.html";
        var p = document.getElementsByTagName("p");
         p[0].style.display = "none";
         tiket(id);
     }else{
        detailmenu(id);
     }

     if(localStorage.ColorCode){
       localStorage.ColorCode = localStorage.ColorCode;
     }else{
       localStorage.ColorCode = "#0f0f0a";
     }
});

function tiket(id){

    var result = JSON.parse(tiket_menu);
  //  console.log(JSON.stringify(result));
    var listmenu = result[id-1].menu;
    console.log(JSON.stringify(listmenu));
            var menu = "";
            for (var i = 0; i < listmenu.length; i++) {
              menu += '<a href="javascript:active_menu('+listmenu[i].id+',\''+listmenu[i].name+'\',\''+result[id-1].lang_name+'\');">';
              menu += '<div class="block-2">';
              menu += '<img src="/img/'+listmenu[i].image+'" onError="this.src = \'/img/noimg.jpg\'" class="block-img">';
              menu += '<h5 style="margin-top: 0;"><div style="width: 100%; float: left; text-align: center;"><b>'+listmenu[i].name+'</b></div></h5>';
              menu += '</div></a>';
         }

     document.getElementById("menu_data").innerHTML = menu;
     var block2 = document.getElementsByClassName("block-2");
     var blockimg = document.getElementsByClassName("block-img");
     for(var i = 0; i<block2.length; i++){
         block2[i].style.width = "27%";
         block2[0].style.marginLeft = "23%";
         blockimg[i].style.marginBottom = "0";
     }
    // console.log(menu);
}

function detailmenu(id){

     $.ajax({
            url: "http://"+window.location.host+"/menu",
          //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                    console.log(JSON.stringify(result));
                    var listmenu = result[id-1].menus;
                  //  console.log(JSON.stringify(listmenu));
                    var menu = "";
                    for (var i = 0; i < listmenu.length; i++) {
                          menu += '<a href="javascript:active_menu('+listmenu[i].Id+',\''+listmenu[i].name+'\',\''+result[id-1].lang_name+'\');">';
                          menu += '<div class="block-2">';
                          menu += '<img src="/img/'+listmenu[i].image+'" onError="this.src = \'/img/noimg.jpg\'" class="block-img">';
                          menu += '<h5 style="margin-top: 0;"><div style="width: 100%; float: left; text-align: center;"><b>'+listmenu[i].name+'</b></div></h5>';
                          menu += '</div></a>';

                        }

                    document.getElementById("menu_data").innerHTML = menu;
                },
                error: function(err){
                    console.log(JSON.stringify(err));
                }
            });
  //var mydata = jQuery.parseJSON(data);
   //
    //console.log(JSON.stringify(mydata));
   // console.log(mydata[0].langId);
    //
    //console.log(listmenu);
   /* */
}


function active_menu(menuId,mName,lName){
    console.log("menu_id"+ menuId);
    localStorage.menuId = menuId;
    localStorage.nName = mName;
    localStorage.lName = lName;
    console.log(localStorage.language);
      window.location = "item.html";


}

function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female");
    responsiveVoice.speak("English language");
 
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");

    document.getElementById("version").innerHTML = "version 0.1";
    document.getElementById("version2").innerHTML = "version 0.1 ";

    document.getElementById("bt_back").innerHTML = "back";

    document.getElementById("Name_time").innerHTML = "time ";
    document.getElementById("Name_time2").innerHTML = "time ";

    localStorage.language = 2;
        if(localStorage.OrgCode == 0){
             detailmenu(id);
         }else if(localStorage.OrgCode == 1){
           //  window.location = "menu.html";
            var p = document.getElementsByTagName("p")
             p[0].style.display = "none";
             tiket(id);
         }
}

function onsaythai(id){
    responsiveVoice.setDefaultVoice("Thai Female");
    responsiveVoice.speak("ภาษาไทย");

    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");

    document.getElementById("version").innerHTML = "เวอร์ชั่น 0.1";
	document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.1 ";

	document.getElementById("bt_back").innerHTML = "ย้อนกลับ";

    document.getElementById("Name_time").innerHTML = "เวลา ";
    document.getElementById("Name_time2").innerHTML = "เวลา ";

    setTimeout(function(){
        localStorage.language = 1;
            if(localStorage.OrgCode == 0){
                 detailmenu(id);
             }else if(localStorage.OrgCode == 1){
               //  window.location = "menu.html";
                var p = document.getElementsByTagName("p")
                 p[0].style.display = "none";
                 tiket(id);
             }
    }, 1000);

}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female");
    responsiveVoice.speak("中國");

    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");

    document.getElementById("version").innerHTML = "版本 0.1";
    document.getElementById("version2").innerHTML = "版本 0.1 ";

    document.getElementById("bt_back").innerHTML = "背部";

	document.getElementById("Name_time").innerHTML = "時間 ";
    document.getElementById("Name_time2").innerHTML = "時間 ";

    localStorage.language = 3;
        if(localStorage.OrgCode == 0){
             detailmenu(id);
         }else if(localStorage.OrgCode == 1){
           //  window.location = "menu.html";
            var p = document.getElementsByTagName("p")
             p[0].style.display = "none";
             tiket(id);
         }
}

function input_num(number){
    var text = document.getElementById("pwd_setting");
   // console.log(number);
    text.value += number;
}

function delete_text(){
    var text = document.getElementById("pwd_setting").value;

    var newStr = text.substring(0, text.length-1);
    //console.log(newStr);
    document.getElementById("pwd_setting").value = newStr;
}

function check_setting(){
    window.location = "setting.html"
}