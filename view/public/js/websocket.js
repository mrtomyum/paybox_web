
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
      console.log("CONNECTED");

    }

    function onClose(evt)
    {
      console.log("DISCONNECTED");
    }

    function onMessage(evt)
    {
      console.log('RESPONSE: ' + evt.data);

        var pathname = location.pathname.split("/");
              console.log(pathname[1]);
              var t = JSON.parse(evt.data);
           //   var p = t['payload'];;
              console.log(t);
            if(pathname[1]=="item.html"){
                  if(t['command']=="onhand"){
                    var total = 0;
                    var pri1 = document.getElementById("pri1").value;
                    var balance = document.getElementById("pri3").value;
                    var payment =  document.getElementById("pri2").value;
                    total = parseInt(JSON.stringify(t['data'])-parseInt(balance));
                    console.log(parseInt(JSON.stringify(t['data'])));
                    console.log(parseInt(balance));
                    console.log("ยอดเงิน "+parseInt(payment)+" ชำระ "+total);
                    if(parseInt(payment)<= total){
                        console.log("true");
                        total = total.toString();
                        if(total.includes("-")){
                            var pay = total.split("-");
                            total = pay[1];
                        }
                        document.getElementById("texttotal").innerHTML = "เงินถอน";
                        document.getElementById("pri2").value = t['data'];
                        document.getElementById("pri3").value = total;
                    }else{
                        console.log("false");
                        total = total.toString();
                        if(total.includes("-")){
                        var pay = total.split("-");
                             total = pay[1];
                        }
                        document.getElementById("texttotal").innerHTML = "ค้างชำระ";
                        document.getElementById("pri2").value = t['data'];
                        document.getElementById("pri3").value = total;
                    }

                  }else if(t['command']=="order"){
                    Alert7.alert("การทำรายการ "+t['data']);
                    if(t['result']==true){websocket.close(); /*setTimeout(function(){window.location = "index.html";},2000);*/}
                  }else if(t['command']=="cancel"){
                    Alert7.alert("ยกเลิกรายการ "+JSON.stringify(p['data']));
                    if(t['result']==true){
                        window.location = "index.html";
                    }
                  }
                  console.log("item "+evt.data);
            }else if(pathname[1]=="model.html"){
                  console.log(evt.data);
                  $("#datatext").append(evt.data+"<br>");
            }

            if(t['command']=="status"){
                var m = t['data'];
                Alert7.alert(m['message']);
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


