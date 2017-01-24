$(document).ready(function(){
    var detail = '<a href="javascript:takethisthai()"><img src="/img/tackthis_th.png" class="img_menu"></a>';
        detail += '<a href="javascript:takeoutthai()"><img src="/img/tackout_th.png" class="img_menu"></a>';
     document.getElementById('img_bt').innerHTML = detail;
     localStorage.action = 0;
     localStorage.getID = 0;
     localStorage.language = 1;

     document.getElementById("version").innerHTML = "เวอร์ชั่น 0.1";
     document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.1 ";

     var worker = new Worker('/js/time.js');
     worker.onmessage = function (event) {
     document.getElementById('timer').innerText =event.data ;
     document.getElementById('timer2').innerText =event.data;
     };

     document.getElementById("Name_time").innerHTML = "เวลา ";
     document.getElementById("Name_time2").innerHTML = "เวลา ";
     console.log(localStorage.OrgCode);
     if(localStorage.ColorCode){
             var nav = document.getElementsByClassName("navbar");
                for(var i = 0; i < nav.length; i++){
                     nav[i].style.backgroundColor = localStorage.ColorCode;
                }
     }

     if(localStorage.ColorCode){
        localStorage.ColorCode = localStorage.ColorCode;
     }else{
        localStorage.ColorCode = "#0f0f0a";
     }
     if(localStorage.OrgCode){
         if(localStorage.OrgCode == 0){
           // window.location = "index.html";
         }else if(localStorage.OrgCode == 1){
            window.location = "menu.html";
         }
     }else{
        window.location = "setting.html";
     }

});

function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female")
    responsiveVoice.speak("English language");
    var detail ='<a href="javascript:takethiseng()"><img src="/img/tackthis.png" class="img_menu"></a>';
        detail += '<a href="javascript:takeouteng()"><img src="/img/tackout.png" class="img_menu"></a>';
    document.getElementById('img_bt').innerHTML = detail;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    document.getElementById("version").innerHTML = "version 0.1";
    document.getElementById("version2").innerHTML = "version 0.1 ";

     document.getElementById("Name_time").innerHTML = "time ";
     document.getElementById("Name_time2").innerHTML = "time ";
    localStorage.language = 2;
}

function onsaythai(id){
    responsiveVoice.setDefaultVoice("Thai Female")
    responsiveVoice.speak("ภาษาไทย");
    var detail = '<a href="javascript:takethisthai()"><img src="/img/tackthis_th.png" class="img_menu"></a>';
        detail += '<a href="javascript:takeoutthai()"><img src="/img/tackout_th.png" class="img_menu"></a>';
    document.getElementById('img_bt').innerHTML = detail;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
         document.getElementById("version").innerHTML = "เวอร์ชั่น 0.1";
         document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.1 ";
     document.getElementById("Name_time").innerHTML = "เวลา ";
     document.getElementById("Name_time2").innerHTML = "เวลา ";

    localStorage.language = 1;
}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female")
    responsiveVoice.speak("中國");
    var detail = '<a href="javascript:takethischina()"><img src="/img/tackthis_ch.png" class="img_menu"></a>';
        detail += '<a href="javascript:takeoutchina()"><img src="/img/tackout_ch.png" class="img_menu"></a>';
    document.getElementById('img_bt').innerHTML = detail;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    document.getElementById("version").innerHTML = "版本 0.1";
    document.getElementById("version2").innerHTML = "版本 0.1 ";

     document.getElementById("Name_time").innerHTML = "時間 ";
     document.getElementById("Name_time2").innerHTML = "時間 ";

    localStorage.language = 3;
}
/*////////////////// take this  //////////////////////////////*/
function takethiseng(){
    console.log("active uk");
        localStorage.lName = "UK English Female";
        localStorage.nName = "take this";
        window.location = "menu.html";
        localStorage.action = 1;
}

function takethisthai(){
    console.log("active th");
    localStorage.lName = "Thai Female";
    localStorage.nName = "รับประทานที่ร้าน";
        window.location = "menu.html";
        localStorage.action = 1;

}

function takethischina(){
    console.log("active ch");
    localStorage.lName = "Chinese Female";
    localStorage.nName = "拿著它";
        window.location = "menu.html";
        localStorage.action = 1;
}
/*////////////////// take this  //////////////////////////////*/
/*////////////////// take out  //////////////////////////////*/
function takeouteng(){
    console.log("active uk");
    localStorage.lName = "UK English Female";
    localStorage.nName = "take out";
        window.location = "menu.html";
        localStorage.action = 2;
}

function takeoutthai(){
    console.log("active th");
    localStorage.lName = "Thai Female";
    localStorage.nName = "ซื้อกลับบ้านค่ะ";
        window.location = "menu.html";
        localStorage.action = 2;
}

function takeoutchina(){
    console.log("active ch");
    localStorage.lName = "Chinese Female";
    localStorage.nName = "取出";
        window.location = "menu.html";
        localStorage.action = 2;
}
/*////////////////// take out  //////////////////////////////*/

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