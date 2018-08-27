## 使用方法

### 1.获取包
```
go get github.com/enjoyass/AliDysmsapi
```
### 2.使用方法
```go
    package main

    import (
        "github.com/enjoyass/AliDysmsapi"
        "fmt"
    )

    func main () {
        ssr,err:=dysmsapi.SendSms(phone, accessKeyId, accessKeySecret, TemplateParam,TemplateCode,SignName)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(ssr)
    }
```
