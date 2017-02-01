$(document).ready(function(){
     localStorage.action = 0;
     localStorage.getID = 0;
     localStorage.language = 1;

     document.getElementById("Scde").innerHTML = screen.width+" X "+screen.height;

     document.getElementById("version").innerHTML = "เวอร์ชั่น 0.1";
     document.getElementById("version2").innerHTML = "เวอร์ชั่น 0.1 ";

     var worker = new Worker('/js/time.js');
     worker.onmessage = function (event) {
     document.getElementById('timer').innerText =event.data ;
     document.getElementById('timer2').innerText =event.data;
     };

     document.getElementById("Name_time").innerHTML = "เวลา ";
     document.getElementById("Name_time2").innerHTML = "เวลา ";

     if(localStorage.ColorCode){
        document.getElementById("color_theme").value = localStorage.ColorCode;
            document.getElementById("bgview").style.backgroundColor = localStorage.ColorCode;
        var nav = document.getElementsByClassName("navbar");
           for(var i = 0; i < nav.length; i++){
                nav[i].style.backgroundColor = localStorage.ColorCode;
           }
     }

});


function confirm(){
    var org_code = $('#org_style').find(":selected").val();
    var color_code = document.getElementById("color_theme").value;

    console.log("org code: "+org_code+" , colorcode : "+color_code);
    localStorage.ColorCode = color_code;
    localStorage.OrgCode = org_code;
    alert("บันทึกเรียบร้อย");
    window.location = "index.html";
}

// Pure JS...no plugin > based on but altered/recoded from http://www.cssscript.com/demo/pick-a-color-from-an-image-using-canvas-and-javascript/

// vars > _underscore is the getElement wildcard
var img = _('.thumb img'),
  cv = _('#canv'),
  colorVal = _('.colorVal'),
  bgview = _('.bgview'),
  x = '',
  y = '';

// img click function
img.addEventListener('click', function(e) {
  // for chrome
  if (e.offsetX) {
    x = e.offsetX;
    y = e.offsetY;
  }
  // for firefox
  else if (e.layerX) {
    x = e.layerX;
    y = e.layerY;
  }
  grabCanvas(cv, img, function() {
    // image data
    var $ = cv.getContext('2d')
      .getImageData(x, y, 1, 1).data;
    // show info
    document.getElementById("bgview").style.backgroundColor = rgbToHex($[0], $[1], $[2]);
    document.getElementById("color_theme").value = rgbToHex($[0], $[1], $[2]);
  });
}, false);

// preview color possibilities as mouse moves over image
img.addEventListener('mousemove', function(e) {
  // chrome
  if (e.offsetX) {
    x = e.offsetX;
    y = e.offsetY;
  }
  // firefox
  else if (e.layerX) {
    x = e.layerX;
    y = e.layerY;
  }

  grabCanvas(cv, img, function() {

    // image data
    var $ = cv.getContext('2d')
      .getImageData(x, y, 1, 1).data;
    // preview color
  });
}, false);

// canvas function
function grabCanvas(el, image, callback) {
  el.width = image.width; // element / img width
  el.height = image.height; // element /img height
  // draw image in canvas tag
  el.getContext('2d')
    .drawImage(image, 0, 0, image.width, image.height);
  return callback();
}
// querySelector
function _(el) {
  return document.querySelector(el);
};

/* convert rgba to hex > reference http://stackoverflow.com/questions/5623838/rgb-to-hex-and-hex-to-rgb */
function componentToHex(c) {
  var hex = c.toString(16);
  return hex.length == 1 ? "0" + hex : hex;
}

function rgbToHex(r, g, b) {
  return "#" + componentToHex(r) + componentToHex(g) + componentToHex(b);
}

function findPos(obj) {
  var curleft = 0,
    curtop = 0;
  if (obj.offsetParent) {
    do {
      curleft += obj.offsetLeft;
      curtop += obj.offsetTop;
    } while (obj = obj.offsetParent);
    return {
      x: curleft,
      y: curtop
    };
  }
  return undefined;
}