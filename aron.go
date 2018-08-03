/*
** -----------------------------------
** Aron - Hidden Parameters Bruteforce 
** Coded by Momo (M4ll0k) Outaadi
** https://github.com/m4ll0k
** -----------------------------------
*/

package main

import (
        "os"
		"fmt"
		"flag"
		"time"
		"bufio"
		"bytes"
		"strings"
		"net/http"
		"io/ioutil"
		"github.com/m4ll0k/printer"
)

var (
	_url      = flag.String("url","","\n\t\tSet target URL")
	_get      = flag.Bool("get",false,"\n\t\tSet get method")
	_post     = flag.Bool("post",false,"\n\t\tSet post method")
	_data     = flag.String("data","","\n\t\tSet post data")
	_wordlist = flag.String("wordlist","dict.txt","\n\t\tSet your wordlist")
)

func Banner() {
	fmt.Printf("%s    ___                         %s\n",printer.GREEN,printer.RESET)                   
	fmt.Printf("%s   /   |  _________  ___       %s\n",printer.GREEN,printer.RESET)
	fmt.Printf("%s  / /| | / ___/ __ \\/ __\\   %s\n",printer.GREEN,printer.RESET)
	fmt.Printf("%s / ___ |/ /  / /_/ / / / /   %s\n",printer.GREEN,printer.RESET)
	fmt.Printf("%s/_/  |_/_/   \\____/_/ /_/ %s(v0.1.0 beta)%s\n",printer.GREEN,printer.YELLOW,printer.RESET)
	fmt.Printf("----------------------------\n")
	fmt.Printf("  Momo (M4ll0k) Outaadi \n\n") 
}

func CheckUrl(url string) string {
	// check url 
	if url != "" {
		if strings.Contains(url,"://"){
			return url
		} else {
			if strings.Contains(url,"."){
				return "http://" + url
			} else {
				printer.Warn("Please enter with valid URL",true)
				os.Exit(0)
				return ""
			}
		}
	} else {
		printer.Warn("Please enter with your URL!",true)
		os.Exit(0)
		return ""
	}
	return ""
}

func JoinPost(data string, param string) string {
	// Params Join
	var params = param + "=fuzz"
	if data == "" {
		return params
	} else {
		if strings.HasSuffix(data,"&") {
			return data + param + "=fuzz"
		} else {
			if strings.HasSuffix(data,"=") {
				return data + "fuzz&" + param + "=fuzz"
			} else {
				return data + "&" + param + "=fuzz"
			}
		}
	}
	return params
}

func JoinGet(url string, param string) string {
	// URL Join 
	var params = param + "=fuzz"
	if strings.HasSuffix(url,"?") {
		if strings.Contains(url,"=") {
			if strings.HasSuffix(url,"&") {
				return url + params
			} else {
				return url + "&" + params
			}
		} else {
			return url + params
		}
		return url + params
	} else {
		if strings.Contains(url,"=") {
			if strings.HasSuffix(url,"&") {
				return url + params
			} else {
				return url + "&" + params
			}
			return url + "&" + params
		} else {
			return url + "?" + params
		}
	}
	return url
}


func Request(url string, method string, data string, payload string) {
	// Do Request
	m := strings.ToLower(method)
	if m == "post" {
		// 1 req content
		f_content,f_url := Post(url,data)
		// 2 req content
		s_content,s_url := Post(url,JoinPost(data,payload))
		if s_content != f_content {
			if s_url == url {
				printer.Info("Found a potentially valid parameter: ",false)
				printer.Plus("URL  => "+s_url,false)
				printer.Plus("POST => "+JoinPost(data,payload),false)
			} else {
				if f_url == url {
					fmt.Printf("")
				}
			}
		}
	} else {
		// 1 req content
		f_content,f_url := Get(url)
		// 2 req content
		real_url := JoinGet(url,payload)
		s_content,s_url := Get(real_url)
		fmt.Printf("%s\n",s_url)
		if s_content != f_content {
			if s_url == real_url {
				printer.Info("Found a potentially valid parameter: ",false)
				printer.Plus("URL => " + s_url,false)
			} else {
				if s_url == f_url {
					fmt.Printf("")
				}
			}
		}
	}
}


func Get(url string) (string,string) {
	// create new request
	req,err := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent","Mozilla/5.0")
	// client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
			},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		return string(content),resp.Request.URL.String()
	}
	return "",""
}

func Post(url string, data string) (string,string) {
	post_data := []byte(data)
	// create new request
	req, err := http.NewRequest("POST",url,bytes.NewBuffer(post_data))
	req.Header.Set("User-Agent","Mozilla/5.0")
	// client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp,err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content),resp.Request.URL.String()
}

func main() {
	// Main 
	Banner()
	flag.Parse()
	// -- time 
	t := time.Now()
	time_now := t.Format("2006-01-02 15:04:05")
	printer.Plus("URL: "+*_url,false)
	printer.Plus("Starting: "+time_now,false)
	fmt.Printf("\n")
	//-- end time
	file,err := os.Open(*_wordlist)
	if err != nil {
		printer.Warn("Error: \""+*_wordlist+"\" not found!!",true)
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//
	uri := CheckUrl(*_url)
	if uri == "" {
		flag.Usage()
		os.Exit(0)
	}
	printer.Test("Bruteforcing...",true)
	if *_get {
		// read payload
		for scanner.Scan() {
			Request(uri,"GET","",scanner.Text())
		}
		// --
	} else {
		if *_post { 
			// read payload 
			for scanner.Scan() {
				Request(uri,"POST",*_data,scanner.Text())
			}
		}
	}
	printer.Test("Done!",true)
}
