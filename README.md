Aron
--
__Aron__ is a simple GO script for finding hidden GET &amp; POST parameters with bruteforce.

![screen](https://i.imgur.com/lryusFo.png)

Installation
--

```sh

$ go get -u github.com/m4ll0k/Aron
$ ./Aron
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
$ go run aron.go -u http://www.test.com/index.php -g 
$ go run aron.go -u http://www.test.com/index.php<[?|id=1|id=1&]> -g
$ go run aron.go -u http://www.test.com/index.php<[?|id=1|id=1&]> -g -w my_wordlist.txt
```

_<[?|id=1|id=1&]>_ **=> Possible end URL**

**OR** __Note:__ in this case aron need the wordlist path 
```sh
$ aron -u http://www.test.com/index.php -g -w path/wordlist.txt
$ aron -u http://www.test.com/index.php<[?|id=1|id=1&]> -g -w path/wordlist.txt
```

__POST BRUTEFORCE:__

```sh
$ go run aron.go -u http://www.test.com/index.php -p
$ go run aron.go -u http://www.test.com/index.php<[?id=1]> -p
$ go run aron.go -u http://www.test.com/index.php<[?id=1]> -p -d "user=1"
$ go run aron.go -u http://www.test.com/index.php<[?id=1]> -p -d "user=1" -wordlist my_wordlist
```

**OR** __Note:__ in this case aron need the wordlist path 

```sh
$ aron -u http://www.test.com/index.php -p -w path/wordlist.txt
$ aron -u http://www.test.com/index.php<[?id=1]> -p -d "user=1" -w path/wordlist.txt
```
