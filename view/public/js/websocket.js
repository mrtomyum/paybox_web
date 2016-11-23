function payment_socket(){
    var wsUri = "ws://echo.websocket.org/";
         var output;

         function init() {
            output = document.getElementById("output");
            testWebSocket();
         }

         function testWebSocket() {
            websocket = new WebSocket(wsUri);

            websocket.onopen = function(evt) {
               onOpen(evt)
            };
         }

         function onOpen(evt) {
            console.log("CONNECTED");
         }

         window.addEventListener("load", init, false);
}