﻿调试编译   nim c rml.nim
正式编译   nim c -d:release rml.nim

实现Rebol语法首先需要实现代码的token化

types.nim里定义了token的类型，目前仅实现了少量主要类型
token.nim里定义了Token类型，Rebol代码会统一解析为一个Token序列。Token类型主要为两部分，Token.tp为类型标识，Token.val为实际值的存储位置。其中Token.val为union类型，长度为16字节，用来存放各种可能出现的值类型。

为了避开Nim的GC，自行实现了List和BindMap类型，其中List类型等同于seq，是一个动态的数组。BindMap等同于Rebol中语境的概念，其本质上是HashMap，增加了一个father字段，用来实现语境的层级关系。

Rebol代码转化为Token序列，需要对代码文本进行处理，strTool.nim中定义了一些字符串处理函数，其中最主要的函数是切割字符串的strCut函数，目前主要把代码切割为通用标识符、字符串、方块和圆块。toToken.nim则实现了讲这些切割后的字符串序列转化为不同类型Token序列的功能，在单个Token转化的过程中，方块和圆块类型内部会进一步转化为Token序列。

转化生成的Token序列需要进行求值运算，并将最后一个表达式的运算结果在控制台打印出来。evalStack.nim实现了Rebol代码求值运算逻辑，其中evalExp对应处理一个Rebol表达式，eval函数用于处理输入的整段代码。

/native封装了解释器的内置函数，中缀运算符在/op内进行初始化，实际调用的是native中的函数。

rml.nim是控制台入口，启动过程首先构建了一个lib语境，并将内置函数和中缀运算符绑定到lib语境中，之后构建一个user语境，父语境为lib语境，然后实例化一个evalStack，当控制台收入输入的代码时，由evalStack解释执行，并将结果输出到控制台。