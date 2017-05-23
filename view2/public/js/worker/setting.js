/**
 * Created by nathaphol on 18/5/2560.
 */
var step = 0;
var sec = 0;
var myVar = "";

//ใช้ addEventListener เพื่อรับ message จาก Main --> wk.postMessage('Hello');
self.addEventListener('message', function (e) {
    if (e.data == 0) {
        step = 0;
    } else {
        if (step >= 5) {
            step = 1;
        } else {
            step += e.data;
        }
        postMessage(step);
    }

}, false);
