<html>
    <head>
    <title>Demo</title>
    </head>
    <body>
	    <form action="/" method="post">
	        Username:<input type="text" name="username">
	        Password:<input type="password" name="password">

	        <input type="submit" value="Login">
	    </form>
	    <p>{{.Notice}} <br>Username: {{.Name}} <br>Password: {{.Password}}</p>
	</body>
</html>