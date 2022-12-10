package main

import (
	"sentimentalgirl/pocs"
	"sentimentalgirl/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var checkurls [] string
var thread int

func runpoc(checkurl string) int{
	fmt.Println( "[ " + checkurl + " ]" + " 正在检测")
	if(utils.Checks == "all"){
		if(pocs.CheckCustomReadFile(checkurl) == -1){
			return -1
		}
		if(pocs.CheckCustomRce(checkurl) == -1){
			return -1
		}
		if(pocs.CheckTreexmlRce(checkurl) == -1){
			return -1
		}
	}else {
		checkvuls := strings.Split((utils.Checks),",")
		for _,checkvul := range checkvuls{
			switch checkvul {
			case "custom-readfile":
				pocs.CheckCustomReadFile(checkurl)
				return -1
			case "custom-rce":
				pocs.CheckCustomRce(checkurl)
				return -1
			case "treexml-rce":
				pocs.CheckTreexmlRce(checkurl)
				return -1
			default:
				utils.Usage()
				os.Exit(-1)
			}
		}
	}
	return 0
}

func main(){
	url := flag.String("url","","")
	file := flag.String("file","","")
	check := flag.String("check","","")
	thread := flag.Int("t",10,"")
	output := flag.String("o","results.txt","")
	flag.Parse()
	if(((*file == "" && *url == "") || (*file != "" && *url != "")) || (*check) == ""){
		utils.Usage()
		os.Exit(-1)
	}
	utils.Init(runpoc,*check,*thread,*output)
	if(*file != ""){
		fileread,err := os.OpenFile(*file,os.O_RDWR,0666)
		if(err != nil){
			fmt.Print("打开 [" + (*file) + "] 文件失败")
			os.Exit(-1)
		}
		defer fileread.Close()
		buf := bufio.NewReader(fileread)
		for{
			line, err := buf.ReadString('\n')
			if (err != nil) {
				break
			}
			line = strings.Replace(line,"\r","",1)
			line = strings.Replace(line,"\n","",1)
			checkurl := utils.StandardUrl(line)
			checkurls = append(checkurls, checkurl)
		}
		for i := 0;i < len(checkurls);i++{
			utils.Wg.Add(1)
			go utils.Run(checkurls[i])
		}
		utils.Wg.Wait()
		utils.FilePtr.Close()
	}else {
		checkurl := utils.StandardUrl(*url)
		runpoc(checkurl)
	}
}
