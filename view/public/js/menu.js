$(document).ready(function(){


    var id = localStorage.language;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");

    console.log(id);
    detailmenu(id);
    
});

function detailmenu(id){

     $.ajax({
            url: "http://localhost:8888/menu/",
          //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                  //  console.log(JSON.stringify(result));
                    var listmenu = result[id-1].menus;
                  //  console.log(JSON.stringify(listmenu));
                    var menu = "";
                    for (var i = 0; i < listmenu.length; i++) {
                          menu += `<a href="javascript:active_menu(`+listmenu[i].Id+`,'`+listmenu[i].name+`','`+result[id-1].lang_name+`');">
                                        <div class="block-2">
                                          <img src="/img/`+listmenu[i].image+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
                                            <h5 style="margin-top: 0;"><div style="width: 100%; float: left; text-align: center;"><b>`+listmenu[i].name+`</b></div></h5>
                                        </div>
                                    </a>`;

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

function menu_say(mName,lName){
    responsiveVoice.setDefaultVoice(lName);
    responsiveVoice.speak(mName);
}

function active_menu(menuId,mName,lName){
    console.log("menu_id"+ menuId);
    localStorage.menuId = menuId;
    menu_say(mName,lName);
      window.location = "item.html";

}

function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female");
    responsiveVoice.speak("English language");
 
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    localStorage.language = 1;
    detailmenu(id);
}

function onsaythai(id){
    responsiveVoice.setDefaultVoice("Thai Female");
    responsiveVoice.speak("ภาษาไทย");

    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    localStorage.language = 2;
    detailmenu(id);
}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female");
    responsiveVoice.speak("中國");

    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    localStorage.language = 3;
    detailmenu(id);
}