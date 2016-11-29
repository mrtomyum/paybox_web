
        var wsUri = "ws://localhost:8888/ws";
          websocket = new WebSocket(wsUri);

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
            if(pathname[1]=="item.html"){

                      var t = JSON.parse(evt.data);
                      document.getElementById("pri2").value = t['OnhandAmount'];

                  console.log("item "+evt.data);
            }else if(pathname[1]=="model.html"){
                  console.log(evt.data);
                  $("#datatext").append(evt.data);
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


