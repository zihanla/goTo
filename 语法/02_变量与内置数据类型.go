package main

func main() {
	/*
		变量：Go是静态类型，变量声明必须确定变量类型
			var a int
			var a int = 1
			var a = 1
			a := 1
			msg := "Hello World!"
	*/

	/*
		简单类型：
			空值：nil
			整型：int(取决于操作系统),int8,int16,int32,int64,uint8,uint16,...
			浮点数：float32,float64
			字节：byte(== uint8)
			字符串：string
			布尔值：boolean,(true、false)
	*/

	/*
		字符串：
			Go语言中，字符串使用UTF8编码，英文每个字符占1 byte,中文一般每个字符占3 byte,于ASCII编码一样
			str1 := "Golang"
			str2 := "Go语言"
			fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8； reflect.TypeOf().Kind() 用于获取变量类型
			fmt.Println(str1[2], string(str1[2]))       // 108 l；uint8 == byte 占一字节，打印需要进行类型给转换，否则打印的是编码值
			fmt.Printf("%d %c\n", str2[2], str2[2])     // 232 è；字符串是以byte数组的形式存储
			fmt.Println("len(str2):", len(str2))        // len(str2): 8；Go占2 byte,语言占6 byte

	*/
	/*
		// 转换成[]rune类型后，字符串中的字符无论占多少个字节都用int32表示
		str2 := "Go语言"
		runeArr := []rune(str2)
		fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
		fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
		fmt.Println("len(runeArr):", len(runeArr))     // len(runeArr): 4
	*/
}
