
        var wsUri = "ws://"+window.location.host+"/ws";
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
              var p = t['payload'];;
              console.log(p);
            if(pathname[1]=="item.html"){
                  if(p['command']=="onhand"){
                        document.getElementById("pri2").value = p['data'];


                  }else if(p['command']=="billing"){
                    Alert7.alert("การทำรายการ "+p['data']);
                    if(p['result']==true){websocket.close(); setTimeout(function(){window.location = "index.html";},2000);}
                  }else if(p['command']=="cancel"){
                    Alert7.alert("ยกเลิกรายการ "+JSON.stringify(p['data']));
                    if(p['result']==true){
                        window.location = "index.html";
                    }
                  }
                  console.log("item "+evt.data);
            }else if(pathname[1]=="model.html"){
                  console.log(evt.data);
                  $("#datatext").append(evt.data+"<br>");
            }

            if(p['command']=="status"){
                var m = p['data'];
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


