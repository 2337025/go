//程序所属包
package main
//导入依赖包
import "fmt"
//常量定义
const NAME string = "imooc"
//全局变量的声明与复制
var mainName string = "main name"
//一般类型声明
type imoocInt int
//结构的声明
type Learn struct{
}
//声明接口
type Ilearn interface{
}
//函数定义
func learnImooc(){
	fmt.Print("learnImooc\n")
}
func main(){
	learnImooc()
	fmt.Println("hello,world\n")

}
