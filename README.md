# conv
超简单的进制转换工具，十进制转换为不同进制（二进制、八进制、十六进制）的字符串格式

feizhiwu/conv is a conversion util, you can use it convert  decimal to different string format(binary, octal, hexadecimal) 

## 和strconv.FormatInt的区别
多了补位和负数支持


## 示例
#### 转二进制
```go
	fmt.Println(conv.FormatInt(4, 2, 8)) //00000100
	fmt.Println(conv.FormatInt(8, 2, 16)) //0000000000001000
	fmt.Println(conv.FormatInt(-16, 2, 8)) //11110000
```

#### 转十六进制
```go
	fmt.Println(conv.FormatInt(4, 16, 8)) //04
	fmt.Println(conv.FormatInt(8, 16, 16)) //0008
	fmt.Println(conv.FormatInt(-16, 16, 8)) //f0
```