1. 
“go generate”命令的作用是在编译前自动化生成某类代码；它常用于自动生成代码，
它可以在代码编译之前根据源代码生成代码。当运行“go generate”命令时，它将扫描与当前包相关的源代码文件，
找出所有包含“//go:generate”的特殊注释，提取并执行该特殊注释后面的命令

2. 安装 stringer
go get golang.org/x/tools/cmd/stringer
go install golang.org/x/tools/cmd/stringer