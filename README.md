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
$ ./aron -u http://www.test.com/index.php -g 
$ ./aron -u http://www.test.com/index.php<[?|id=1|id=1&]> -g
$ ./aron -u http://www.test.com/index.php<[?|id=1|id=1&]> -g -w my_wordlist.txt
```

__POST BRUTEFORCE:__

```sh
$ ./aron -u http://www.test.com/index.php -p #basic fuzz with post method 
$ ./aron -u http://www.test.com/index.php<[?id=1]> -p -d "user=1" #set post data
$ ./aron -u http://www.test.com/index.php<[?id=1]> -p -d "user=1" -H 'cookie: test=20,x-header: 10' #set headers name:value,name:value,...
$ ./aron -u http://www.test.com/index.php<[?|id=1|id=1&]> -p -d "user=1" -w my_wordlist.txt #set wordlist
```
