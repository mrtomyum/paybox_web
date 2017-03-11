function addZero(i) {
    if (i < 10) {
        i = "0" + i;
    }
    return i;
}

function getMyDateFormat() {
    var d = new Date();
  //  var x = document.getElementById("demo");
    var h = addZero(d.getHours());
    var m = addZero(d.getMinutes());
    var s = addZero(d.getSeconds());
    return h + ":" + m;
}
var time = "";

setInterval(function() {
  postMessage(getMyDateFormat());
}, 100);

