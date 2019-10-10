# 说明
flyutils是一个golang项目小工具助手，封装了对字符串、切片、时间等相关的函数，方便、简单、易用。

# 下载安装
go get github.com/will110/flyutils

# 实例
1、截取字符串

	str := "helloabc你好123"
	//从前面开始截取
	s := flyutils.Substr(str, 2, 3)
	fmt.Println(s)
	
	//从后面开始截取
	s = flyutils.Substr(str, -3, 5)
	fmt.Println(s)
	
2、生成长度为8位的随机字符串

    s := flyutils.GenerateRandomStr(8)
    fmt.Println(s)

3、删除切片中的元素

	s := []string{"a", "b", "c", "d"}
	s1 := flyutils.SliceDeleteString("b", s)
	fmt.Println(s1)

	n1 := []int64{22, 321, 89, 1102}
	n2 := flyutils.SliceDeleteInt64(89, n1)
	fmt.Println(n2)
	
4、获取今天是几号

	w := flyutils.GetMonthDay()
	fmt.Println(w)

5、获取今天星期几

	d := flyutils.GetWeekday()
	fmt.Println(d)
	

# 文档：
[https://godoc.org/github.com/will110/flyutils](https://godoc.org/github.com/will110/flyutils "文档")