{{define "device"}}
{{template "header.tpl" .}}
<h1>{{.Title}}</h1>
<div class="container">
    Hello
    <h1>Hello, this is device index page.</h1>
    <a href="/ping" class="btn btn-info" role="button">Ping</a>
</div>
{{template "footer.tpl" .}}
{{end}}