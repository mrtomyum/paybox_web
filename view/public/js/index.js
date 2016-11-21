$(document).ready(function(){
    var detail =`<a href="javascript:takethiseng()"><img src="img/tackthis.png" class="img_menu"></a>
                 <a href="javascript:takeouteng()"><img src="img/tackout.png" class="img_menu"></a>`;
     document.getElementById('img_bt').innerHTML = detail;
     localStorage.action = 0;
     localStorage.getID = 0;
     localStorage.language = 2;
});
function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female")
    responsiveVoice.speak("English language");
    var detail =`<a href="javascript:takethiseng()"><img src="img/tackthis.png" class="img_menu"></a>
                 <a href="javascript:takeouteng()"><img src="img/tackout.png" class="img_menu"></a>`;
    document.getElementById('img_bt').innerHTML = detail;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    localStorage.language = 1;
}

function onsaythai(id){
    responsiveVoice.setDefaultVoice("Thai Female")
    responsiveVoice.speak("ภาษาไทย");
    var detail = `<a href="javascript:takethisthai()"><img src="img/tackthis_th.png" class="img_menu"></a>
                  <a href="javascript:takeoutthai()"><img src="img/tackout_th.png" class="img_menu"></a>`;
    document.getElementById('img_bt').innerHTML = detail;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    localStorage.language = 2;
}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female")
    responsiveVoice.speak("中國");
    var detail = `<a href="javascript:takethischina()"><img src="img/tackthis_ch.png" class="img_menu"></a>
                  <a href="javascript:takeoutchina()"><img src="img/tackout_ch.png" class="img_menu"></a>`;
    document.getElementById('img_bt').innerHTML = detail;
    $("img").removeClass("active_img");
    $("#"+id).addClass("active_img");
    localStorage.language = 3;
}
/*////////////////// take this  //////////////////////////////*/
function takethiseng(){
    console.log("active uk");
    responsiveVoice.setDefaultVoice("UK English Female")
    responsiveVoice.speak("take this");
    setTimeout(function (){
        window.location = "/html/menu.html";
        localStorage.action = 1;
    },1500);
}

function takethisthai(){
    console.log("active th");
    responsiveVoice.setDefaultVoice("Thai Female")
    responsiveVoice.speak("รับประทานที่ร้าน");
    setTimeout(function (){
        window.location = "/html/menu.html";
        localStorage.action = 1;
    },1500);
}

function takethischina(){
    console.log("active ch");
    responsiveVoice.setDefaultVoice("Chinese Female")
    responsiveVoice.speak("拿著它");
    setTimeout(function (){
        window.location = "/html/menu.html";
        localStorage.action = 1;
    },1500);
}
/*////////////////// take this  //////////////////////////////*/
/*////////////////// take out  //////////////////////////////*/
function takeouteng(){
    console.log("active uk");
    responsiveVoice.setDefaultVoice("UK English Female")
    responsiveVoice.speak("take out");
    setTimeout(function (){
        window.location = "/html/menu.html";
        localStorage.action = 2;
    },1500);
}

function takeoutthai(){
    console.log("active th");
    responsiveVoice.setDefaultVoice("Thai Female")
    responsiveVoice.speak("ซื้อกลับบ้านค่ะ");
    setTimeout(function (){
        window.location = "/html/menu.html";
        localStorage.action = 2;
    },1500);
}

function takeoutchina(){
    console.log("active ch");
    responsiveVoice.setDefaultVoice("Chinese Female")
    responsiveVoice.speak("取出");
    setTimeout(function (){
        window.location = "/html/menu.html";
        localStorage.action = 2;
    },1500);
}
/*////////////////// take out  //////////////////////////////*/
