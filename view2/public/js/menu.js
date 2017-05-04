var num_call = 0;
// document.addEventListener('contextmenu', event => event.preventDefault());
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

  $(".img-header").on("taphold",function(){
    $( "#mySetting" ).popup( "open" );
    document.getElementById("pwd").value = "";
  });
call_websocket();
/*setInterval(function(){
     $.ajax({
             url: "http://"+window.location.host+"/menu/1",
             contentType: "application/json; charset=utf-8",
             dataType: "json",
             type: "GET",
             cache: false,
                 success: function(result){
                   num_call++;
                   console.log("API Sample ready!! " + num_call);

                 },
                 error: function (err){
                    console.log("API Sample Not ready!! ERROR "+ JSON.stringify(err));
                 }
           });
 }, 1000);*/


    setTimeout(function () {
        if (localStorage.language) {
            var id = localStorage.language;
        } else {
            var id = 0;
        }
        switch (parseInt(id)) {
            case 0 :
                onsaythai(id);
                console.log("ไทย");
                break;
            case 1 :
                onsayeng(id);
                console.log("อังกฤษ");
                break;
            case 2 :
                onsaychina(id);
                console.log("จีน");
                break;
        }
        // var id = 2;


        var worker = new Worker('js/worker/time.js');
        worker.onmessage = function (event) {
            document.getElementById('timer').innerHTML = event.data;
            document.getElementById('timer2').innerText = event.data;
        };


        switch (parseInt(id)) {
            case 1:
                document.getElementById("version").innerHTML = "เวอร์ชั่น 0.9";
                document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.9 ";

                //    document.getElementById("bt_back").innerHTML = "ย้อนกลับ";

                document.getElementById("N_time").innerHTML = "เวลา ";
                document.getElementById("Name_time2").innerHTML = "เวลา ";

                break;
            case 2:
                document.getElementById("version").innerHTML = "version 0.9";
                document.getElementById("version2").innerHTML = "version 0.9 ";

                // document.getElementById("bt_back").innerHTML = "back";

                document.getElementById("N_time").innerHTML = "time ";
                document.getElementById("Name_time2").innerHTML = "time ";
                break;
            case 3:
                document.getElementById("version").innerHTML = "版本 0.9";
                document.getElementById("version2").innerHTML = "版本 0.9 ";

                //	document.getElementById("bt_back").innerHTML = "背部";

                document.getElementById("N_time").innerHTML = "時間 ";
                document.getElementById("Name_time2").innerHTML = "時間 ";

                break;
        }

        // console.log(id);
        detailmenu(id);
        console.log("screen width : " + screen.width);
        active_lang(id);

        //** screen server ***//
        //new Worker object
        var wk=new Worker("js/worker/worker.js");
        //ใช้ addEventListener เพื่อรับ message จาก Woker --> self.postMessage('worker got : '+data);
                wk.onmessage = function(oEvent){
                    //document.getElementById('display').textContent = "เวลา " + oEvent.data + " วินาที";
                    //จะ print 'worker got : Hello'
                    if(oEvent.data==300){
                        console.log("screen");
                        window.location = "#page_screen";
                        wk.postMessage(0);
                    }
                };
        //start Worker และส่ง message ให้ Worker ด้วย postMessage


        $('html').click(function(){
            wk.postMessage(0);
        });
        ////////////////////////

    }, 1000);
});

function detailmenu(id){

     $.ajax({
            url: "http://"+window.location.host+"/menu",
          //  data: '{"barcode":"'+barcode+'","docno":"'+DocNo+'","type":"1"}',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            type: "GET",
            cache: false,
                success: function(result){
                  //  console.log(JSON.stringify(result));
                    var listmenu = result[id].menus;
                  //  console.log(JSON.stringify(listmenu));
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

    console.log("active " +localStorage.language);
    menu_detail(localStorage.language,menuId);
    setTimeout(function () {
        voice_say(localStorage.language, mName);
    }, 500);
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
     console.log(content);
    console.log("vioce say " + lang);
     responsiveVoice.setDefaultVoice(language);
     responsiveVoice.speak(content);
}

function onsayeng(id){
    responsiveVoice.setDefaultVoice("UK English Female");
    responsiveVoice.speak("English language");
 
    active_lang(id);
    //console.log("active _ id "+id);

    document.getElementById("version").innerHTML = "version 0.9";
    document.getElementById("version2").innerHTML = "version 0.9 ";

  //  document.getElementById("bt_back").innerHTML = "back";

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

	//document.getElementById("bt_back").innerHTML = "ย้อนกลับ";

    document.getElementById("N_time").innerHTML = "เวลา ";
    document.getElementById("Name_time2").innerHTML = "เวลา ";

    setTimeout(function(){
        localStorage.language = 0;
        detailmenu(id);
    }, 1000);

}

function onsaychina(id){
    responsiveVoice.setDefaultVoice("Chinese Female");
    responsiveVoice.speak("中國");

    active_lang(id);

    document.getElementById("version").innerHTML = "版本 0.9";
    document.getElementById("version2").innerHTML = "版本 0.9 ";

  //  document.getElementById("bt_back").innerHTML = "背部";

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
                    console.log("lang red"+id);
                }else{
                    x[i].style.borderColor = "#fff";
               }
            }

}

function enPwd(number){
    //console.log(number);
    var pwd = document.getElementById("pwd").value;
    pwd += number;

    document.getElementById("pwd").value = pwd;
}

function delete_pwd(){
    var text = document.getElementById("pwd").value;

    var newStr = text.substring(0, text.length-1);
    //console.log(newStr);
    document.getElementById("pwd").value = newStr;
}

function close_set(){
    document.getElementById("pwd").value = "";
    $("#mySetting").popup("close");
}
