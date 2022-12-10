package utils

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var Checks string
var Wg sync.WaitGroup
var Fun func(checkurl string) int
var HttpClient http.Client
var Chan chan struct{}
var FilePtr *os.File
var WriteFlush *bufio.Writer

func OpenWriteFile(output string) *os.File {
	file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND, 0666)
	if(err != nil){
		fmt.Println("打开 [" + output + "] 文件失败")
		os.Exit(-1)
	}
	return file
}

func Init(funptr func(checkurl string) int,check string,thread int,output string){
	Checks = check
	Wg = sync.WaitGroup{}
	Fun = funptr
	Chan = make(chan struct{},thread)
	HttpClient = GetHttpClient()
	FilePtr = OpenWriteFile(output)
	WriteFlush = bufio.NewWriter(FilePtr)
}

func Usage(){
	logo := `
 __                           __     
(_  _ .__|_o._ _  _ .__|_ _.|/__._o| 
__)(/_| ||_|| | |(/_| ||_(_||\_|| ||

    main.exe -url https://127.0.0.1 -check=all -t 15 -o result.txt
    main.exe -file urls.txt -check=custom-readfile,custom-rce,treexml-rce -t 15 -o result.txt
`
	color.Red(logo)
	fmt.Println("                               author hangchuanin ")
	fmt.Println("                               github https://github.com/hangchuanin ")
}

func StandardUrl(url string) string{
	if(strings.HasSuffix(url,"/")){
		url = strings.TrimRight(url,"/")
	}
	return url
}

func GetHttpClient() http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: transport,
		Timeout: 5 * time.Second,
	}
	return client
}

func Run(url string) {
	Chan <- struct{}{}
	Fun(url)
	<- Chan
	Wg.Done()
}

