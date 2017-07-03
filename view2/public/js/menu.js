var num_call = 0;
$(document).ready(function(){
  if(navigator.onLine)
  {
    document.getElementById("net_sta").innerHTML = '<img src="/img/inter_connect.png">';
    document.getElementById("al_marq").innerHTML = '** ระบบพร้อมทำงานแล้ว';
  }
  else
  {
    document.getElementById("net_sta").innerHTML = '<img src="/img/inter_notconnect.png">';
    document.getElementById("al_marq").innerHTML = '** ไม่มีการเชื่อมต่อ Internet';
     setTimeout(function(){
        alertify.error("ระบบกำลังพยายามเชื่อมต่อ internet");
        setTimeout(function(){
            window.location.reload();
        },1000);
     },10000);
  }

call_websocket();
    setTimeout(function () {
        if (localStorage.language) {
            var id = localStorage.language;
        } else {
            var id = 0;
        }
        switch (parseInt(id)) {
            case 0 :
                onsaythai(id);
                break;
            case 1 :
                onsayeng(id);
                break;
            case 2 :
                onsaychina(id);
                break;
        }
        var worker = new Worker('js/worker/time.js');
        worker.onmessage = function (event) {
            document.getElementById('timer').innerHTML = event.data;
            document.getElementById('timer2').innerText = event.data;
        };
        switch (parseInt(id)) {
            case 1:
                document.getElementById("version").innerHTML = "เวอร์ชั่น 0.9";
                document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.9 ";
                document.getElementById("N_time").innerHTML = "เวลา ";
                document.getElementById("Name_time2").innerHTML = "เวลา ";

                break;
            case 2:
                document.getElementById("version").innerHTML = "version 0.9";
                document.getElementById("version2").innerHTML = "version 0.9 ";
                document.getElementById("N_time").innerHTML = "time ";
                document.getElementById("Name_time2").innerHTML = "time ";
                break;
            case 3:
                document.getElementById("version").innerHTML = "版本 0.9";
                document.getElementById("version2").innerHTML = "版本 0.9 ";
                document.getElementById("N_time").innerHTML = "時間 ";
                document.getElementById("Name_time2").innerHTML = "時間 ";

                break;
        }
        detailmenu(id);
        active_lang(id);
        var wk = new Worker("js/worker/worker.js");
                wk.onmessage = function(oEvent){
                    if(oEvent.data==300){
                        $.mobile.changePage("#page_screen");
                        wk.postMessage(0);
                    }
                };
        $('html').click(function(){
            wk.postMessage(0);
        });
    }, 500);

});

var wksetting = new Worker("js/worker/setting.js");
wksetting.onmessage = function (oEvent) {
    document.getElementById("stepst").innerHTML = oEvent.data;
    showDialog();
    if (oEvent.data >= 5) {
        $("#mySetting").popup("open");
        document.getElementById("pwd").value = "";
        closeDialog();
        clearInterval(timeclick);
        stepclick = 0;
    }
};
var stepclick = 0;
var timeclick = "";
function setting_ck() {
    wksetting.postMessage(1);
    stepclick = 0;
    clearInterval(timeclick);
    timeclick = setInterval(function () {
        stepclick += 1;
        if (stepclick == 2) {
            stepclick = 0;
            clearInterval(timeclick);
            closeDialog();
            wksetting.postMessage(0);
        }
        console.log("stepck " + stepclick);
    }, 1000);
}

var x = document.getElementById("myDialog");

function showDialog() {
    x.show();
}

function closeDialog() {
    x.close({hide: {effect: "fade", duration: 200}});
}

function detailmenu(id){
     $.ajax({
            url: "http://"+window.location.host+"/menu",
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                    var listmenu = result[id].menus;
                    var menu = '<ul class="block-2">';
                    for (var i = 0; i < listmenu.length; i++) {
                          menu += '<li>';
                          menu += '<a href="javascript:active_menu('+listmenu[i].Id+',\''+listmenu[i].name+'\',\''+result[id].lang_name+'\');">';
                        menu += '<img src="/img/' + listmenu[i].image + '" onError="this.src = \'/img/noimg.jpg\'" class="block-img">';
                          menu += '<p>'+listmenu[i].name+'</p>';
                          menu += '</a>';
                          menu += '</li>';
                        }
                    menu += '</ul>';
                    document.getElementById("list_menu").innerHTML = menu;
                },
                error: function(err){
                    console.log(JSON.stringify(err));
                }
            });
}

function active_menu(menuId,mName,lName){
    localStorage.menuId = menuId;
    localStorage.nName = mName;
    localStorage.lName = lName;
    menu_detail(localStorage.language,menuId);
    voice_say(localStorage.language, mName);
}

function voice_say(lang,content){
     switch(parseInt(lang)){
         case 0 : var language = "Thai Female";
                break;
         case 1 : var language = "UK English Female";
                break;
         case 2 : var language = "Chinese Female";
                break;
     }
     responsiveVoice.setDefaultVoice(language);
     responsiveVoice.speak(content);
}

function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female");
    responsiveVoice.speak("English language");
    active_lang(id);
    document.getElementById("version").innerHTML = "version 0.9";
    document.getElementById("version2").innerHTML = "version 0.9 ";
    document.getElementById("N_time").innerHTML = "time ";
    document.getElementById("Name_time2").innerHTML = "time ";
    localStorage.language = 1;
    detailmenu(id);
}

function onsaythai(id){
    responsiveVoice.setDefaultVoice("Thai Female");
    responsiveVoice.speak("ภาษาไทย");
    active_lang(id);
    document.getElementById("version").innerHTML = "เวอร์ชั่น 0.9";
    document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.9 ";
    document.getElementById("N_time").innerHTML = "เวลา ";
    document.getElementById("Name_time2").innerHTML = "เวลา ";
    localStorage.language = 0;
    detailmenu(id);
}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female");
    responsiveVoice.speak("中國");
    active_lang(id);
    document.getElementById("version").innerHTML = "版本 0.9";
    document.getElementById("version2").innerHTML = "版本 0.9 ";
	document.getElementById("N_time").innerHTML = "時間 ";
    document.getElementById("Name_time2").innerHTML = "時間 ";
    localStorage.language = 2;
    detailmenu(id);
}

function active_lang(id){
        var x = document.getElementsByClassName("lang");
        for (i = 0; i < x.length; i++) {
                if(id==i){
                    x[i].style.borderColor = "#f00";
                }else{
                    x[i].style.borderColor = "#fff";
               }
            }
}

function enPwd(number){
    var pwd = document.getElementById("pwd").value;
    pwd += number;
    document.getElementById("pwd").value = pwd;
}

function delete_pwd(){
    var text = document.getElementById("pwd").value;
    var newStr = text.substring(0, text.length-1);
    document.getElementById("pwd").value = newStr;
}

function close_set(){
    document.getElementById("pwd").value = "";
    $("#mySetting").popup("close");
}

function setting_page() {
    $.mobile.changePage("#page_setting");
}
