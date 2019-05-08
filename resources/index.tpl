<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>idgen</title>
    <meta name="viewport" content="initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no" draggable="false"/>
    <meta name="format-detection" content="telephone=no"/>
    <meta name="apple-mobile-web-app-capable" content="yes"/>
    <meta name="keywords" content=""/>
    <meta name="description" content=""/>
    <style>
        .list-li{
            display: block;
            width: 100%;
            border-width: 0px;
            outline: none;
            background: #fff;
        }
        .list-li::selection{
            background: rgb(30, 202, 194);
        }
    </style>

</head>
<body>
    <input class="list-li" value="{{ .Name }}" />
    <input class="list-li" value="{{ .Mobile }}" />
    <input class="list-li" value="{{ .IdNo }}" />
    <input class="list-li" value="{{ .Bank }}" />
    <input class="list-li" value="{{ .Email }}" />
    <input class="list-li" value="{{ .Addr }}" />
    <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
    <script>
        $(".list-li").on('select', function(){
            document.execCommand('copy', true);
        });
    </script>
</body>
</html>