$(document).ready(function(){


    var id = localStorage.language;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");

    console.log(id);
    detailmenu(id);
    
});

function detailmenu(id){
  var mydata = jQuery.parseJSON(data);
    var menu = "";
    //console.log(JSON.stringify(mydata));
   // console.log(mydata[0].langId);
    var listmenu = mydata[id-1].menu;
    //console.log(listmenu);
    for (var i = 0; i < listmenu.length; i++) {
      menu += `<a href="javascript:active_menu(`+listmenu[i].id+`,'`+listmenu[i].name+`','`+mydata[id-1].langName+`');">
                    <div class="block-2">            
                      <img src="/`+listmenu[i].img+`" onError="this.src = '/img/noimg.jpg'" class="block-img">
                        <h5 style="margin-top: 0;"><div style="width: 100%; float: left; text-align: center;"><b>`+listmenu[i].name+`</b></div></h5>
                    </div> 
                </a>`;
      
    }

    document.getElementById("menu_data").innerHTML = menu;
}

function menu_say(mName,lName){
    responsiveVoice.setDefaultVoice(lName);
    responsiveVoice.speak(mName);
}

function active_menu(menuId,mName,lName){
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