
        var wsUri = "ws://"+window.location.host+"/web";
        //var wsUri = "ws://192.168.0.215:8888/ws";

        websocket = new WebSocket(wsUri);
    console.log(window.location.host);
    function call_websocket(){
          websocket.onopen = function(evt) { onOpen(evt) };
          websocket.onclose = function(evt) { onClose(evt) };
          websocket.onmessage = function(evt) { onMessage(evt) };
          websocket.onerror = function(evt) { onError(evt) };
    }

    function onOpen(evt)
    {
      console.log("CONNECTED HOST Websocket");

    }

    function onClose(evt)
    {
      console.log("DISCONNECTED HOST Websocket");
      call_websocket();
    }


    function onMessage(evt)
    {
      console.log('RESPONSE: ' + evt.data);

        var pathname = location.pathname.split("/");
              console.log(pathname[1]);
              var t = JSON.parse(evt.data);
           //   var p = t['payload'];;
              console.log(t);
            //if(pathname[1]=="item.html"){
                  if(t['command']=="onhand"){
                  console.log("ยอดเงินที่ได้รับ "+t['data']);
                    var total = 0;
                    var pri1 = document.getElementById("mo-pri").value;
                    var balance = document.getElementById("pri3").value;
                    var payment =  document.getElementById("pri2").value;
                    console.log("ราคารวม "+pri1);
                    console.log("ยอดค้างชำระ "+balance);

                    total = parseInt(JSON.stringify(t['data']))-parseInt(pri1);
                    console.log(parseInt(JSON.stringify(t['data'])));
                    console.log(parseInt(balance));
                    console.log("ผลลัพย์ total "+total);

                    if(total >= 0){
                        console.log("true");
                        //total = total.toString();
                      /*  if(total.includes("-")){
                            var pay = total.split("-");
                            total = pay[1];
                        }*/
                        document.getElementById("texttotal").innerHTML = "เงินทอน";
                        console.log("document.getElementById('pri2').value = "+t['data']);
                        document.getElementById("pri2").value = t['data'];
                        document.getElementById("textpri2").innerHTML = t['data'];

                        document.getElementById("pri3").value = total;
                        document.getElementById("textpri3").innerHTML = total;

                      /*  if(total!=0){
                            document.getElementById("coinbill").innerHTML = "และเงินทอน "+total+" บาท";
                        }*/
                    }else{
                        console.log("false");
                        total = total.toString();
                       // if(total.includes("-")){
                        var pay = total.split("-");
                             total = pay[1];
                        //}
                        document.getElementById("texttotal").innerHTML = "ค้างชำระ";
                        console.log("document.getElementById('pri2').value = "+t['data']);//
                        document.getElementById("pri2").value = t['data'];
                        document.getElementById("textpri2").innerHTML = t['data'];
                        document.getElementById("textpri3").innerHTML = total;
                        document.getElementById("pri3").value = total;
                    }

                  }else if(t['command']=="order"){
                    Alert7.alert("การทำรายการ "+t['data']);
                    if(t['result']==true){websocket.close(); /*setTimeout(function(){window.location = "index.html";},2000);*/}
                  }else if(t['command']=="cancel"){
                    document.getElementById("list_item").innerHTML = "";
                    alertify.error("ยกเลิกรายการ "+JSON.stringify(t['data']));

                    if(t['result']==true){
                        $.mobile.changePage("#pageone");
                    }else if(t['result']==false){
                        alertify.success(JSON.stringify(t['data']));
                    }
                  }else if(t['command']=='accepted_bill'){
                    var bank = "";
                    var b20 = t['data'].b20;
                    var b50 = t['data'].b50;
                    var b100 = t['data'].b100;
                    var b500 = t['data'].b500;
                    var b1000 = t['data'].b1000;

                    if(b20==true){
                        bank += '<img src="public/img/b20_true.png" class="bank">';
                    }else{
                        bank += '<img src="img/b20_false.png" class="bank">';
                    }

                    if(b50==true){
                        bank += '<img src="img/b50_true.png" class="bank">';
                    }else{
                        bank += '<img src="img/b50_false.png" class="bank">';
                    }

                    if(b100==true){
                        bank += '<img src="img/b100_true.png" class="bank">';
                    }else{
                        bank += '<img src="img/b100_false.png" class="bank">';
                    }

                    if(b500==true){
                        bank += '<img src="img/b500_true.png" class="bank">';
                    }else{
                        bank += '<img src="img/b500_false.png" class="bank">';
                    }

                    if(b1000==true){
                        bank += '<img src="img/b1000_true.png" class="bank">';
                    }else{
                        bank += '<img src="img/b1000_false.png" class="bank">';
                    }
                    document.getElementById("bank_use").innerHTML = bank;

                  }
              //    console.log("item "+evt.data);
            /*}else if(pathname[1]=="model.html"){
                  console.log(evt.data);
                  $("#datatext").append(evt.data+"<br>");
            }*/

            if(t['command']=="warning"){
                alert(t['data']);
            }

            if(t['command']=="change"){
                   $("#pop_payment").popup('close');
                   document.getElementById("W_payment").innerHTML = "กรุณารอรับใบเสร็จ";
                   document.getElementById("W_change").innerHTML = "และเงินทอน "+t['data']+" บาท";
                   setTimeout(function(){
                       $("#pop_bill").popup('open');
                   },300);
            }

            if(t['command']=="payment"){
                   $("#pop_payment").popup('close');
                   document.getElementById("W_payment").innerHTML = "กรุณารอรับใบเสร็จ";
                   setTimeout(function(){
                       $("#pop_bill").popup('open');
                   },300);
            }

            if(t['command']=="print"){
                if(t['data']=="success"){
                   setTimeout(function(){
                    document.getElementById("W_payment").innerHTML = "";
                    document.getElementById("W_change").innerHTML = "";
                    document.getElementById("list_item").innerHTML = "";
                    $.mobile.changePage("#pageone");
                   },1500);
                   console.log("print success");
                }else{
                    alertify.error(t['data']);
                    $.mobile.changePage("#pageone");
                }
            }

    }

    function onError(evt)
    {
      console.log('Error : ' + evt.data);
    }

    function doSend(message)
    {
      console.log("SENT: " + message);
      websocket.send(message);

    }


