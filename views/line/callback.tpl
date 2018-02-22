<!DOCTYPE html>

<html>

<head>
    <title>example-golang-oauth</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
    <div class="container">
        <br />
        <br />
        <br />
        <div class="alert alert-success" role="alert">LINEアカウント認証成功！</div>
        <div class="panel panel-success">
            <div class="panel-heading">
                情報
            </div>
            <div class="panel-body">
                <br />
                <ul class="list-group">
                    <li class="list-group-item">
                        <img style="width: 100%; max-width: 200px; margin: 0 auto;" src="{{.PictureUrl}}" alt="{{.PictureUrl}}"/>
                    </li>
                    <li class="list-group-item">
                        ID：{{.ID}}
                    </li>
                    <li class="list-group-item">
                        DisplayName：{{.DisplayName}}
                    </li>
                    <li class="list-group-item">
                        StatusMessage：{{.StatusMessage}}
                    </li>
                    <li class="list-group-item">
                        Token：{{.AccessToken}}
                    </li>
                    <li class="list-group-item">
                        Code：{{.Code}}
                    </li>
                </ul>
            </div>
        </div>
        <a href="http://localhost:8080/">戻る</a>
    </div>
</body>

</html>
