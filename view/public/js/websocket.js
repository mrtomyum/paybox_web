
        var wsUri = "ws://"+window.location.host+"/ws";

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
      doSend('{"Device":"host","Payload":{"type":"request","command":"onhand"}}');
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
                    alert("การทำรายการ "+p['data']);
                    if(p['result']==true){websocket.close();window.location = "index.html";}
                  }else if(p['command']=="cancel"){
                    alert("ยกเลิกรายการ "+JSON.stringify(p['data']));
                    if(p['result']==true){
                        window.location = "index.html";
                    }
                  }
                  console.log("item "+evt.data);
            }else if(pathname[1]=="model.html"){
                  console.log(evt.data);
                  $("#datatext").append(evt.data+"<br>");
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


