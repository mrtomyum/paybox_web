{{define "menus.tpl"}}
{{template "header.tpl" .}}
<div class="container">

    {{range .}}
    <a href="{{.Link}}">
        <div class="block-2 menu">
            <img src="public/img/{{.Image}}" width="70%" style="margin-bottom: 2%;">
            <h3 style="margin-top: 0; font-size:16px;"><b>{{.Name}}</b></h3>
        </div>
    </a>
    {{end}}

    <div style="clear: both;"></div>

</div>
<div id="app">
    {{message}}
</div>

<script>
    new Vue({
        el: '#app',
        data: {
            message: 'Hello Vue!'
        }
    });

</script>
{{template "footer.tpl" .}}
{{end}}