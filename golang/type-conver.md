_**Golang的类型转换**_

  


```golang
package main


import (
"fmt"
"strconv"
)

func main(){
   
   //string
   str := "99"
   
   //int
   numInt := 99
   
   //int64
   var numInt64 int64 = 99

   //string转int
   strToInt , _ := strconv.Atoi(str)

   
   //string转int64
   strToInt64, _ := strconv.ParseInt(str, 10, 64)

   
   //int转string
   intToStr := strconv.Itoa(numInt)

   
   //int64转string
   int64ToStr :=strconv.FormatInt(numInt64,10)

   fmt.Println(strToInt , strToInt64 , intToStr, int64ToStr)

   numFloat64  := 99.99//float64

   str2 := "99.99"

   //float64转string
   float64ToStr := strconv.FormatFloat(numFloat64 , 'f', 2, 64)

   //string转float64
   strToFloat64 , _ := strconv.ParseFloat(str2 , 64)

   
   //string转float32
   strToFloat32 , _ := strconv.ParseFloat(str2 , 32)

   fmt.Println(float64ToStr,strToFloat64,strToFloat32)

   
   //接口类型转换
   var intInterface interface{} = 123
   
   //接口转int
   fmt.Println(intInterface.(int))

   
   var strInterface interface{} = "123"

   //接口转string
   fmt.Println(strInterface.(string))

   
   var float64Interface interface{} = 123.123
   //接口转float64
   fmt.Println(float64Interface.(float64))


   
   //强制转化
   compleInt := 12345
   compleStr := "12345"
   //int强转为int32

   fmt.Println(int32(compleInt))

   
   //int强转为int64
   fmt.Println(int64(compleInt))

   
   //int强转为float64
   fmt.Println(float64(compleInt))

   
   //int强转为float32
   fmt.Println(float32(compleInt))

}

```
### 强制转换结论
int , float 无法强转为string, string 无法强转为 int , float, 非同一类型不能强制转换

### 接口转换类型注意点
接口转换时是返回两个参数第二个参数是是否转换成功的flag





