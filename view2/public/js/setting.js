$(document).ready(function () {

    sumtotal();

});
function AddM(id) {
    var add = document.getElementsByClassName('addM');
    var remove = document.getElementsByClassName('removeM');
    var cnt = document.getElementsByClassName('C');
    var price = document.getElementsByClassName('tc');

    add[id].style.display = 'block';
    remove[id].style.display = 'block';

    var data = parseInt(cnt[id].innerText) + 1;
    var ct = data + ",000";

    ///console.log('add = '+id+" ,cnt = "+data);

    cnt[id].innerHTML = data;
    price[id].innerHTML = ct;
    sumadd();
}

function RemoveM(id, text) {
    var add = document.getElementsByClassName('addM');
    var remove = document.getElementsByClassName('removeM');
    var cnt = document.getElementsByClassName('C');
    var price = document.getElementsByClassName('tc');

    if (parseInt(cnt[id].innerText) < 2) {
        add[id].style.display = 'none';
        remove[id].style.display = 'none';
        var data = 0;
        var ct = 0;
    } else {
        var data = parseInt(cnt[id].innerText) - 1;
        var ct = data + ",000";
    }
    cnt[id].innerHTML = data;
    price[id].innerHTML = ct;

    sumadd();
    //console.log('add = '+id+" ,cnt = "+data);
}

function sumadd() {
    var price = document.getElementsByClassName('tc');

    var total = 0;
    for (var i = 0; i < price.length; i++) {
        total += parseInt(price[i].innerText);
    }

    if (total == 0) {
        addMN.innerHTML = "฿ 0";
    } else {
        addMN.innerHTML = "฿ " + total + ",000";
    }
    sumtotal();

}

function sumtotal() {
    var addMN = document.getElementById('addMN');
    var scMN = document.getElementById('scMN');
    var ttMN = document.getElementById('ttMN');

    var pri = addMN.innerText.split(" ");
    var sc = scMN.innerText.split(" ");

    var totaltext = parseInt(sc[1]) + parseInt(pri[1]);
    console.log(parseInt(pri[1]));

    ttMN.innerHTML = "฿ " + totaltext + ",000";

}