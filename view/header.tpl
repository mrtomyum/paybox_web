{{define "header.tpl"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>PAYBOX</title>
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.5/css/bootstrap.min.css" integrity="sha384-AysaV+vQoT3kOAXZkl02PThvDr8HYKPZhNT5h/CXfBThSRXQ6jW5DO2ekP5ViFdi" crossorigin="anonymous">

</head>
<body>

<nav class="navbar navbar-fixed-top navbar-light bg-faded">
    <div class="container-fluid">
        <a class="navbar-brand" href="/">
            <!-- <img src="/assets/brand/bootstrap-solid.svg" width="30" height="30" class="d-inline-block align-top" alt=""> -->
            PAYBOX: Coffee Shop
        </a>
        <ul class="nav navbar-nav">
            <li class="nav-item active">
                <a class="nav-link" href="#">หน้าแรก <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#">แจ้งปัญหา</a>
            </li>
        </ul>
        <div class="btn-group float-lg-right">
            <button type="button" class="btn btn-primary">ไทย</button>
            <button type="button" class="btn btn-primary">English</button>
            <button type="button" class="btn btn-primary">中文</button>
        </div>
    </div>
</nav>
{{end}}
