{{define "index.tpl"}}
{{template "header.tpl" .}}
<h1>Main Menu</h1>
<h2>กรุณาเลือกเมนู</h2>
<p>Please choose menu. / 请选择菜单</p>
<div class = "container">
    <h3>User</h3>
    <table class="table table-hover table-condense">
        <tr>
            <th>ID</th>
            <th>Thai</th>
            <th>English</th>
            <th>Chinese</th>
        </tr>
        {{ range . }}
        <tr>
            <td>{{ .Id }}</td>
            <td>{{ .NameTh }}</td>
            <td>{{ .NameEn }}</td>
            <td>{{ .NameCn }}</td>
        </tr>
        {{end}}
    </table>

</div>

{{template "footer.tpl" .}}
{{end}}
