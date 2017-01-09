{{define "list"}}
<!DOCTYPE html>
<html>
<head>
    <title>terminal order</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="css/bootstrap.min.css">
    <script src="js/jquery-1.11.3.min.js"></script>
    <script src="js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="css/main.css">
</head>
<body>
<div id="cleft">
    <div id="contentleft">
        <div class="logo"><img src="img/logo.png"></div>
        <div class="Status"><h4>Status : take this</h4></div>
        <div class="order">
            <h2>รายการ</h2>
            <div class="scroll">
                <label class="orderlist">
                    <div class="ordername"> คาปูชิโน S</div>
                    <div class="orderqty">1 แก้ว</div>
                    <div class="orderprice">35 ฿</div>
                    <div class="ordercancel">
                        <button class="btn btn-danger btn-xs" style="padding-left: 22.5%; padding-right: 22.5%;">-
                        </button>
                    </div>
                </label>

                <label class="orderlist">
                    <div class="ordername"> คาปูชิโน S</div>
                    <div class="orderqty">1 แก้ว</div>
                    <div class="orderprice">35 ฿</div>
                    <div class="ordercancel">
                        <button class="btn btn-danger btn-xs" style="padding-left: 22.5%; padding-right: 22.5%;">-
                        </button>
                    </div>
                </label>

                <label class="orderlist">
                    <div class="ordername"> คาปูชิโน S</div>
                    <div class="orderqty">1 แก้ว</div>
                    <div class="orderprice">35 ฿</div>
                    <div class="ordercancel">
                        <button class="btn btn-danger btn-xs" style="padding-left: 22.5%; padding-right: 22.5%;">-
                        </button>
                    </div>
                </label>

                <label class="orderlist">
                    <div class="ordername"> คาปูชิโน S</div>
                    <div class="orderqty">1 แก้ว</div>
                    <div class="orderprice">35 ฿</div>
                    <div class="ordercancel">
                        <button class="btn btn-danger btn-xs" style="padding-left: 22.5%; padding-right: 22.5%;">-
                        </button>
                    </div>
                </label>

                <label class="orderlist">
                    <div class="ordername"> คาปูชิโน S</div>
                    <div class="orderqty">1 แก้ว</div>
                    <div class="orderprice">35 ฿</div>
                    <div class="ordercancel">
                        <button class="btn btn-danger btn-xs" style="padding-left: 22.5%; padding-right: 22.5%;">-
                        </button>
                    </div>
                </label>

                <label class="orderlist">
                    <div class="ordername"> คาปูชิโน S</div>
                    <div class="orderqty">1 แก้ว</div>
                    <div class="orderprice">35 ฿</div>
                    <div class="ordercancel">
                        <button class="btn btn-danger btn-xs" style="padding-left: 22.5%; padding-right: 22.5%;">-
                        </button>
                    </div>
                </label>

            </div> <!-- scroll -->

            <div id="price1"><h4>ราคาทั้งหมด
                <input type="number" name="pri1" id="pri1" readonly> บาท</h4>
            </div>
            <div id="price2"><h4 style="padding-left: 4%;">จำนวนเงิน
                <input type="number" name="pri2" id="pri2" readonly> บาท</h4>
            </div>

            <div id="sub">
                <button class="ok" style="width:45%; height: 50px; font-size: 20px;">
                    <b>พิมพ์ใบเสร็จ</b>
                </button>
                <a href="index.html">
                    <button class="cancel" style="width:45%; margin-left: 3%; height: 50px; font-size: 20px;">
                        <b>ยกเลิก</b></button>
                </a>
            </div>
        </div> <!-- order -->

    </div><!-- contentleft -->
</div><!-- cleft -->
<div id="chead">
    <div class="contenthead"></div>
    <div class="language">
        <div style="width:50%; float: left; line-height:14px; text-align: right; padding-right: 5%;">
            <font size="5">language : </font>
            <br style="padding-right: 2%;">เวลา 08.00 น.
            <br style="padding-right: 2%;">version 0.1
        </div>
        <div style="width:50%; float: left;">
            <img src="img/uk-flag.png" class="lang">
            <img src="img/thaiflag.png" class="lang">
            <img src="img/China.png" class="lang">
        </div>
    </div>
</div>

<div id="cent">
    <div id="centitem">
        <div id="itemlist">
            <div id="menu_scroll">
                {{range .}}
                <a href="{{.Link}}">
                    <div class="block-3" data-toggle="modal" data-target="#myModal" onclick="showmodal()">
                        <img src="img/{{.Image}}" onError="this.src = 'img/noimg.jpg'" class="block-img">
                        <h5 style="margin-top: 0;">
                            <div style="width: 80%; float: left;"><b>{{.Name}}</b></div>
                            <div style="width: 20%; float: left;"><b>xx฿</b></div>
                        </h5>
                    </div>
                </a>
                {{end}}
            </div>

        </div>
        <div id="bt_menu">
            <ul>
                <li onclick="active(0)"><a href="#" id="0" class="active"><img src="img/hot_list.png">HOT</a></li>
                <li onclick="active(1)"><a href="#" id="1"><img src="img/ice_list.png">Ice</a></li>
                <li onclick="active(2)"><a href="#" id="2"><img src="img/Smoothies.png">Frappe</a></li>
                <li onclick="active(3)"><a href="#" id="3"><img src="img/kids_list.png">kids</a></li>
                <li onclick="active(4)"><a href="#" id="4"><img src="img/snack_list.png">snack</a></li>
            </ul>
        </div>

    </div>
</div>

<!-- Modal -->
<div class="modal fade" id="myModal" role="dialog">
    <div class="modal-dialog">

        <!-- Modal content-->
        <div class="modal-content">
            <div class="modal-header">
                <div class="modal-name"><h1 class="modal-title">ItemCode : ItemName</h1></div>
            </div>
            <div class="modal-body">
                <div id="modal-left">
                    <div id="modal-img">
                        <div>
                            <img class="crop" src="img/capu.jpg" onError="this.src = 'img/noimg.jpg'" width="100%"/>
                        </div>
                    </div>
                    <div id="modal-price">
                        <input type="text" id="mo-pri" name="mo-pri" value="35 ฿" readonly>
                    </div>
                </div>
                <div id="modal-right">
                    <div id="modal-qty">
                        <img src="img/delete.png"><input type="text" id="mo_qty" value="1" readonly><img
                            src="img/add.png">
                    </div>
                    <div id="modal-size">
                        <a href="#"><h1 id="s" onclick="active_size('s')" class="acsize"><img src="img/s.png"
                                                                                              class="img-size"><b>Small </b>
                        </h1></a>
                        <a href="#"><h1 id="m" onclick="active_size('m')"><img src="img/m.png" class="img-size"><b>Medium </b>
                        </h1></a>
                        <a href="#"><h1 id="l" onclick="active_size('l')"><img src="img/l.png"
                                                                               class="img-size"><b>large </b></h1></a>
                    </div>
                    <div id="modal-submit">
                        <a href="#"><img src="img/ok.png" width="45%" style="margin-right: 8.5%;"></a>
                        <a href="#"><img src="img/no.png" width="45%" data-dismiss="modal"></a>
                    </div>
                </div>
                <br style="clear: both;">
            </div>

        </div>

    </div>
</div>
<script type="text/javascript" src="js/item.js"></script>
</body>
</html>
{{end}}