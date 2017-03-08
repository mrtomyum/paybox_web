# How to deployment Paybox Terminal to terminal box.
วิธีติดตั้งซอฟท์แวร์ลงไปในตู้

## Setup SSH
ให้สร้าง SSH public key ของเครื่องเพื่อให้เข้าใช้งานกับตู้ได้สะดวก และตั้งชื่อ Host ว่า pb
หมายถึง root@192.168.10.64 อ้างอิงวิธีการที่นี่ https://www.digitalocean.com/community/tutorials/how-to-set-up-ssh-keys--2

## Hardware Service (Device)
```
$ ssh pb
$ killall HW_SERVICE
$ exit
```

copy file HW_SERVICE จากเครื่องเราไปที่ตู้
```
$ scp HW_SERVICE pb:/opt/paybox/hw_service
$ ssh pb
$ ./opt/paybox/hw_service/HW_SERVICE
```

## Web Service (Go + JS - HTML)
Cross Compile โค้ดของ go project ด้วย xgo https://github.com/karalabe/xgo
```
$ xgo --targets=linux/arm-7 github.com/mrtomyum/paybox_web
```
จะได้ไฟล์ชื่อ paybox_web-linux-arm-7

```
$ ssh pb
$ killall WEB_SERVICE
$ exit
```
copy file WEB_SERVICE จากเครื่องเราไปที่ตู้ปลายทางชื่อ WEB_SERVICE
```
$ scp paybox_web-linux-arm-7 pb:/opt/paybox/web_service/WEB_SERVICE
$ ssh pb
$ ./opt/paybox/web_service/WEB_SERVICE
```