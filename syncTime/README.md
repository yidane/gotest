作用是同步windows时间

## 交叉编译

Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序，最近使用了一下，非常好用，这里备忘一下。

#### Mac 下编译 Linux 和 Windows 64位可执行程序
~~~ code
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
~~~

#### Linux 下编译 Mac 和 Windows 64位可执行程序
~~~ code
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
~~~

#### Windows 下编译 Mac 和 Linux 64位可执行程序
~~~ code
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
~~~

GOOS：目标平台的操作系统（darwin、freebsd、linux、windows） 
GOARCH：目标平台的体系架构（386、amd64、arm） 
交叉编译不支持 CGO 所以要禁用它

上面的命令编译 64 位可执行程序，你当然应该也会使用 386 编译 32 位可执行程序 
很多博客都提到要先增加对其它平台的支持，但是我跳过那一步，上面所列的命令也都能成功，且得到我想要的结果，可见那一步应该是非必须的，或是我所使用的 Go 版本已默认支持所有平台。

## 编译

在当前目录执行脚本　
~~~ code
bash build.sh
~~~

## 发布
拷贝　syncTime.exe、syncTime.vbs、synvTime.bat到windows文件夹下即可。

## 运行方式
1. 可以以管理员方式执行syncTime.vbs或synvTime.bat。
2. 在Windows定时任务中配置定时任务，任务内容执行脚本 syncTime.vbs。
* ps:如果任务没有反应，需要将syncTime.vbs、synvTime.bat中的路径修改为绝对路径。
