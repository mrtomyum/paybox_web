{{define "index.tpl"}}
{{template "header.tpl" .}}

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
{{template "footer.tpl" .}}
{{end}}