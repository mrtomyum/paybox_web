<!DOCTYPE html>
<html>
<head>
    <title>terminal order</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="css/bootstrap.min.css">
    <script src="js/jquery-1.11.3.min.js"></script>
    <script src="js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="css/select_main.css">
</head>
<body>

<nav class="navbar navbar-default" style="margin-bottom: 0.2%; height: 60px; background: #262626;">
    <div class="container-fluid" style="padding-left: 0;">
        <div class="navbar-header" style=" width: 20%; text-align: left; padding-left: 1%; padding-top: 0.5%; ">
            <img src="img/logo.png" class="logo">
        </div>
        <div class="nav-r">
            version 0.1<br>เวลา 08.00 น.
        </div>
    </div>
</nav>

<div class="container">
    <div id="menu_main">
        <div id="img_bt"></div>
        <hr/>
    </div>
    <div id="language">
        <a href="javascript:onsayeng(1)"><img src="img/uk-flag.png" id="1" class="lang active_img"></a>
        <a href="javascript:onsaythai(2)"><img src="img/thaiflag.png" id="2" class="lang"></a>
        <a href="javascript:onsaychina(3)"><img src="img/China.png" id="3" class="lang"></a>
        <h1>language </h1>
    </div>
</div>

<script type="text/javascript" src="js/index.js"></script>
<script src='js/voice.js'></script>
</body>
</html>