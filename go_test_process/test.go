package main
import (
    "fmt"
    "runtime"
)

var (   //定义变量
    v1 int 
    v2 string
    v3 int
    v4 string
) 

//常量的多重赋值
const (
    a,b int16 = 1,2
)

//枚举
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    numberOfDays  //这个常量没有导出  为小写，为包内私有
)

//多返回值
func GetName()(firstName,lastName,nickName string){
    return "May","Chan","Chibi Maruko"
}

func main(){
    //多重赋值
    //v1,v2 = v3,v4
    var (
        str string
    )
    str = "hello world"
    chone := str[0]

    // str[0] = 'X'//错误，字符串赋值之后，不能再修改

    fmt.Printf("The lenth of \"%s\" is %d \n" , str , len(str))   //Printf  Print  Println
    fmt.Printf("The first character of \"%s\" is \"%c\" \n" , str , chone)

    _,_,nickName := GetName()
    fmt.Println(nickName)  //Chibi Maruko
    fmt.Println(a,b)  //1 2
    fmt.Println(Sunday , Monday , numberOfDays)  //0 1 7

    fmt.Printf("My go version is \"%s\" \n" , runtime.Version())
}
