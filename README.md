## httpHello

This is a small utility that listens to one or more ports for a http request and writes back a small html document
containing current time and whence you came

Usage: `httpHello 1111 1112 1113 1114 `\
This will listen on the ports 1111 to 1114 (*if they are available*) until the application is terminated

Do a `curl http://localhost:1111 `  and you get
```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
        <title>httpHello</title>
  </head>
  <body>
    <p>
       Hello!<br>

Time is 2021-03-25 23:11:28 and you are calling from localhost:1111

    </p>
  </body>
</html>
```

Since I use this for testing load balancing I have added the possibility to request a few http errorcodes.

The codes are 400, 403, 408, 410, 500 and 503 you can use them like this:

    curl http://localhost:1111/403

Then you get a 403 response like so:

    403 Forbidden
    Time is 2021-03-25 23:21:29 and you are calling from localhost:1111


