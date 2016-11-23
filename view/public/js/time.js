function getMyDateFormat(date) {
    var d = date ? new Date(date) : new Date;
    var tm = [d.getHours(), d.getMinutes()].join(":");
    return tm;
}

setInterval(function() {
postMessage(getMyDateFormat());
}, 1000);