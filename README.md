Aron
--
__Aron__ is a simple GO script for finding hidden GET &amp; POST parameters with bruteforce.

![screen](https://i.imgur.com/lryusFo.png)

Installation
--

```sh
$ git clone https://github.com/m4ll0k/Aron.git aron
$ cd aron 
$ go get github.com/m4ll0k/printer
# now check if $GOPATH is set
$ go env | grep -i gopath
# if $GOPATH not set, try with:
$ export GOPATH=$HOME/go
$ go run aron.go
# OR 
$ go build aron.go
$ cp aron /usr/bin/
$ aron
```

Usage
--

```go
Usage of ./aron:
  -H string
    			Set headers ('name:value,name:value')
  -d string
    			Set post data
  -g			Set get method
  -h			Show this Help
  -p			Set post method
  -u string
    			Set target URL
  -w string
    			Set your wordlist (default "dict.txt")
```

__GET BRUTEFORCE:__

```sh
$ go run aron.go -url http://www.test.com/index.php -get 
$ go run aron.go -url http://www.test.com/index.php<[?|id=1|id=1&]> -get
$ go run aron.go -url http://www.test.com/index.php<[?|id=1|id=1&]> -get -wordlist my_wordlist.txt
```

_<[?|id=1|id=1&]>_ **=> Possible end URL**

**OR** __Note:__ in this case aron need the wordlist path 
```sh
$ aron -url http://www.test.com/index.php -get -wordlist path/wordlist.txt
$ aron -url http://www.test.com/index.php<[?|id=1|id=1&]> -get -wordlist path/wordlist.txt
```

__POST BRUTEFORCE:__

```sh
$ go run aron.go -url http://www.test.com/index.php -post 
$ go run aron.go -url http://www.test.com/index.php<[?id=1]> -post
$ go run aron.go -url http://www.test.com/index.php<[?id=1]> -post -data "user=1"
$ go run aron.go -url http://www.test.com/index.php<[?id=1]> -post -data "user=1" -wordlist my_wordlist
```

**OR** __Note:__ in this case aron need the wordlist path 

```sh
$ aron -url http://www.test.com/index.php -post -wordlist path/wordlist.txt
$ aron -url http://www.test.com/index.php<[?id=1]> -post -data "user=1" -wordlist path/wordlist.txt
```
