{{define "index.tpl"}}
{{template "header.tpl" .}}
<div class = "container">
    <!--แสดงรายการเมนูหลักในหน้าแรก -->
    <h1>Main Menu</h1>
    <h2>กรุณาเลือกเมนู</h2>
    <p>Please choose menu. / 请选择菜单</p>

    <!--<table class="table table-hover table-condense">-->
        <!--<tr>-->
            <!--<th>ID</th>-->
            <!--<th>Thai</th>-->
            <!--<th>English</th>-->
            <!--<th>Chinese</th>-->
        <!--</tr>-->
        <!--{{ range . }}-->
        <!--<tr>-->
            <!--<td>{{ .Id }}</td>-->
            <!--<td>{{ .NameTh }}</td>-->
            <!--<td>{{ .NameEn }}</td>-->
            <!--<td>{{ .NameCn }}</td>-->
        <!--</tr>-->
        <!--{{end}}-->
    <!--</table>-->

    <div class="container">
        <div class="row">
            {{range .}}
            <div class="col-sm-4 ">
                {{.Id}}
                {{.NameTh}}
                {{.NameEn}}
                {{.NameCn}}
            </div>
            {{end}}
        </div>
    </div>

    </div>
</div>

{{template "footer.tpl" .}}
{{end}}