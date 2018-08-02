Aron
--
__Aron__ is a simple GO script for finding hidden GET &amp; POST parameters with bruteforce.

![screen](https://i.imgur.com/e1iEOXP.png)

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
```

Usage
--

```go
   ___                         
   /   |  _________  ___       
  / /| | / ___/ __ \/ __\   
 / ___ |/ /  / /_/ / / / /   
/_/  |_/_/   \____/_/ /_/ (v0.1.0 beta)
----------------------------
  Momo (M4ll0k) Outaadi 

Usage of aron:
  -data="": 
		Set post data
  -get=false: 
		Set get method
  -post=false: 
		Set post method
  -url="": 
		Set target URL
  -wordlist="dict.txt": 
		Set your wordlist
```

__GET BRUTEFORCE:__

```sh
$ go run aron.go -url http://www.test.com/index.php -get 
$ go run aron.go -url http://www.test.com/index.php? -get
$ go run aron.go -url http://www.test.com/index.php?id=1 -get
$ go run aron.go -url http://www.test.com/index.php?id=1& -get
$ go run aron.go -url http://www.test.com/index.php?id=1 -get -wordlist my_wordlist.txt
```

__POST BRUTEFORCE:__

```sh
$ go run aron.go -url http://www.test.com/index.php -post 
$ go run aron.go -url http://www.test.com/index.php?id=1 -post
$ go run aron.go -url http://www.test.com/index.php?id=1 -post -data "user=1"
$ go run aron.go -url http://www.test.com/index.php?id=1 -post -data "user=1" -wordlist my_wordlist
```
