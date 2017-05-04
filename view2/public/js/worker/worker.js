var data = 0;
var sec = 0;
//ใช้ addEventListener เพื่อรับ message จาก Main --> wk.postMessage('Hello'); 
self.addEventListener('message',function(e){
    data=e.data;
    sec = e.data;
},false); 

setInterval(function() {
  sec += data+1;
  postMessage(sec);
 //console.log("worker "+ sec);
}, 1000);