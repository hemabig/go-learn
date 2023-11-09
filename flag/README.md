1. -flag (bool)
   -flag=123
   -flag 123

第一种形式只支持布尔类型的选项，出现即为true，不出现为默认值。 第三种形式不支持布尔类型的选项。因为这种形式的布尔选项在类 Unix 系统中可能会出现意想不到的行为。看下面的命令：

cmd -x *
其中，*是 shell 通配符。如果有名字为 0、false的文件，布尔选项-x将会取false。反之，布尔选项-x将会取true。而且这个选项消耗了一个参数。 如果要显示设置一个布尔选项为false，只能使用-flag=false这种形式。

2. 遇到第一个非选项参数（即不是以-和--开头的）或终止符--，解析停止。运行下面程序, 不解析参数：

$ ./main.exe noflag -intflag 12

将会输出：
int flag: 0
bool flag: false
string flag: default

运行下面的程序：

$ ./main.exe -intflag 12 -- -boolflag=true
将会输出：

int flag: 12
bool flag: false
string flag: default

3. 参数格式
整数选项值可以接受 1234（十进制）、0664（八进制）和 0x1234（十六进制）的形式，并且可以是负数。实际上flag在内部使用strconv.ParseInt方法将字符串解析成int。 所以理论上，ParseInt接受的格式都可以。

布尔类型的选项值可以为：

取值为true的：1、t、T、true、TRUE、True；
取值为false的：0、f、F、false、FALSE、False。

4. 时间间隔
除了能使用基本类型作为选项，flag库还支持time.Duration类型，即时间间隔。时间间隔支持的格式非常之多，例如"300ms"、"-1.5h"、“2h45m"等等等等。 时间单位可以是 ns/us/ms/s/m/h/day 等。实际上flag内部会调用time.ParseDuration。具体支持的格式可以参见time（需fq）库的文档。