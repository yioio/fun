# Introduction 简介 #
1. Use Golang implement the commonly used built-in PHP functions. 使用Golang实现PHP的内置函数。
2. Similar to php2go | php2golang | go2php | golang2php and other projects. 类似于 php2go | php2golang | go2php | golang2php 等项目

# Wiki #
All document entries for this project 本项目所有文档入口

https://github.com/yioio/fun/blob/master/docs/wiki.md

# Instructions 使用说明 

## Download and install & 下载安装 

    go get -u github.com/yioio/fun
    go get -u github.com/smartystreets/goconvey/convey 

## Example 示例 

    package main

    import (
        "github.com/yioio/fun/funArray"
        "log"
    )

    func main() {

        ret1 := funArray.In_array("cat", []interface{}{"cat", "dog"})
        ret2 := funArray.In_array_string("cat", []string{"cat", "dog"})
        log.Println(ret1)
        log.Println(ret2) 
    }

# Other 其它

## Code Contribution Guide 代码贡献指南 
https://github.com/yioio/fun/blob/master/docs/guide-code-contribution.md

## Functional dictionary & development plan &  功能字典 & 开发计划
https://github.com/yioio/fun/blob/master/docs/function-dictionary.md

## Feedback  反馈 

If you have questions, an email to yumoop@163.com is welcome

如果您有任何问题，欢迎发邮件至 yumoop@163.com 或者加入 QQ 群讨论：655685367
