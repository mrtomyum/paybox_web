{{define "index.tpl"}}
{{template "header.tpl" .}}
<div class="container">

    {{range .}}
        <a href="list.tpl">
        <div class="block-2 menu">
            <img src="{{.Image}}" width="70%" style="margin-bottom: 2%;">
            <h3 style="margin-top: 0;"><b>{{.Name}}</b></h3>
        </div>
        </a>
    {{end}}

        <div style="clear: both;"></div>

</div>

{{end}}