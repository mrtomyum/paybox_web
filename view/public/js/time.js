function getMyDateFormat() {
   var date = new Date();

   var year = date.getFullYear();
   var month = date.getMonth() + 1;
   var day = date.getDate();
   var hours = date.getHours();
   var minutes = date.getMinutes();
   var seconds = date.getSeconds();

  return hours+":"+minutes;
}

setInterval(function() {
postMessage(getMyDateFormat());
}, 1000);