# Code Contribution Guide 代码贡献指南 #

## 1、How to write code 如何写代码 ##

1. 每一个文件需要放在被归类好的文件夹下。文件夹以函数功能功别命名,并以 fun 开头，如 in_array 属于 php 的 array 函数，则其文件夹名为 funArray
1. 每一个 php 函数对应一个文件。如 in_array，的文件名为 fun/funArray/in_array.go

## 2、How to write unit tests 如何写单元测试 ##

<hr>

# 1、How to write code 如何写代码 #

## 代码编写要点 ##

# 2、How to write unit tests 如何写单元测试 #

## Installation 安装方法 ##

提交的代码必须通过 goconvey 单元测试。

    go get github.com/smartystreets/goconvey/convey 

goconvey 会被安装在 $GOPATH/src/github.com/smartystreets/goconvey

## Test method 测试方法 ## 

以 php 的 in_array() 函数为例，测试它在本项目下的单元测试这样做：

    cd $GOPATH/src
    cd github.com/yioio/fun/test/funArray
    go test

## Pay attention 注意要点 ##

1. 所有的测试代码写在 fun/tests 文件夹下的子文件夹。
1. 子文件夹名需要与 go 函数的 package 一一对应。例如 In_array 的 package 为 funArray，则其文件夹名为 funArray
1. 每一个测试代码对应测试一个 go 文件
1. 文件名必须以 _test.go 结尾。例如 In_array 的 测试文件名为 in_array_test.go
1. 测试文件名格式为：Test_开头 + 函数名。例如 In_array 的 测试 方法名为 func Test_In_array_string(t *testing.T)
1. goconvey 不仅仅提供了 ShouldEqual ，还有很多其它断言的方法。大家可以上网搜索学习