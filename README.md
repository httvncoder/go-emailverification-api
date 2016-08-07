# go-emailverification-api
API for doing a smtp callback verification.

Check: https://en.wikipedia.org/wiki/Callback_verification

Todo:

- [ ] Mailserver Detection
- [ ] Graylisting detection
- [ ] More SMTP responses

Usage:
```
USAGE :
  -host string
    	Set the server host (default "127.0.0.1")
  -port string
    	Set the server port (default "3003")
```

For example:
```
$ ./go-emailverification-api -host=localhost -port=3003
```

Default values are:
* Host: 127.0.0.1
* Port: 3003

When it's started, test it with curl like the command below.

```
curl http://localhost:3003/v1/email/verify/email@example.com
```

Explain:
* Host and port: http://localhost:3003
* Version of test: /v1
* What to test: /email
* What to do with it: verify
* What email: email@example.com

# Dependencies

 * Go 1.6.x tested https://golang.org
 * httprouter https://github.com/julienschmidt/httprouter
 * govalidator https://github.com/asaskevich/govalidator

```
go get github.com/julienschmidt/httprouter
go get github.com/asaskevich/govalidator
```

# The MIT License (MIT)

Copyright (c) 2016 Sebastian Broekhoven
~~~
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
~~~