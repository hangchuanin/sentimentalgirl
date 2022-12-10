# sentimentalgirl

蓝凌OA漏洞扫描工具，扫描以下三个漏洞

- custom-rce
- custom-readfile
- treexml-rce

# 编译

```
set GO111MODULE=off
go get github.com/fatih/color
set GOOS=windows
go build main.go
```

# 使用

```
G:\>main.exe

 __                           __
(_  _ .__|_o._ _  _ .__|_ _.|/__._o|
__)(/_| ||_|| | |(/_| ||_(_||\_|| ||

    main.exe -url https://127.0.0.1 -check=all -t 15 -o result.txt
    main.exe -file urls.txt -check=custom-readfile,custom-rce,treexml-rce -t 15 -o result.txt
                               author hangchuanin
                               github https://github.com/hangchuanin
```

# 免责声明

本工具仅面向合法授权的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。请勿对非授权目标进行扫描。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。

在安装并使用本工具前，请您务必审慎阅读、充分理解各条款内容，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要安装并使用本工具。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。
