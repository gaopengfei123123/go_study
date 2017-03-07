这里设置 Liteide 的配置(环境macOS)

>Liteide 是一款专门为go开发的一款ide，省去了类似atom编辑器的一些麻烦的环境配置导致atom越来越臃肿，是不少go的开发者在推荐的一款ide

1.安装go
```
brew install go
```

2. 下载Liteide
>[下载地址](http://golangtc.com/download/liteide)

选择一个适合系统的版本，推荐qt5的，版本高的不容易过时


3. 打开编辑器 `查看->编辑当前环境`

列出我现在的环境
```
# native compiler darwin amd64

GOROOT=/usr/local/Cellar/go/1.8/libexec
GOPATH=$HOME/Documents/go
GOBIN=
GOARCH=amd64
GOOS=darwin
CGO_ENABLED=1

PATH=$GOROOT/bin:$PATH

LITEIDE_GDB=/usr/local/bin/gdb
LITEIDE_MAKE=make
LITEIDE_TERM=/usr/bin/open
LITEIDE_TERMARGS=-a Terminal
LITEIDE_EXEC=/usr/X11R6/bin/xterm
LITEIDE_EXECOPT=-e
```

上面这一块刚开始是注释掉的，需要我们自己填写，这个是每次运行编辑器的适合才会用到的配置，和安装时的系统配置不冲突

需要注意区分的就是 `GOROOT` 和 `GOPATH` 这两个变量

`GOROOT` 是go的执行目录，说白了就是go源码的安装目录，
比如 `import “fmt”`这条命令就是从 `GOROOT` 给的路径中找加载的包

`GOPATH` 是工作目录，类似apache的www目录，你的项目代码都是放到 `$GOPATH/src` 目录下的
同时我们通过 `go get` 获取的第三方package也是放在这里的为了能执行package里的文件也是需要把这个路径添加到环境变量当中，
因此才会出现 `PATH=$GOROOT/bin:$PATH` 这样的配置


4. 配置完成后运行个文件试一下
```
cd $HOME/Documents/go/src
mkdir hello && cd hello
```
通过liteide编辑器打开目录，创建一个文件  main.go

文件内容如下
```
package main

import "fmt"

func main(){
    fmt.Println("Hello world!")
}
```
然后点击编辑器右上角的 `BR(build and run)` 就能在编辑器下面的信息框中看到输出内容了

ready go！


[参考1](https://segmentfault.com/a/1190000004008942)

[参考2](http://www.jianshu.com/p/4e699ff478a5)
