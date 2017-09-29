main.go作为入口文件
go build 会把当前文件夹下的所有go文件都给编译了，因此文件不能有多个main函数

hotload
使用node child_process来启动go
node监听文件变化，如果发生变化，重启服务

