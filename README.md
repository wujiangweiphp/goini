> 本函数主要是读取ini配置类文件

文件格式类似：

```
[Mysql]
dbhost = 127.0.0.1
dbname = test
dbuser = zhangsan
dbpwd = 123456
;dbuser = root
;dbpwd = 
```

### 1. 安装
```
go get -v -u github.com/wujiangweiphp/goin
```

### 2. 使用

```go
if dbhost,err:= GetValue("/etc/php/7.2/php.ini","Mysql->dbhost") ; err != nil {
  fmt.Println(err)
} else {
  fmt.Println("dbhost is ",val)
}
```

> 注意：keyname 是由: `区块名称->区块key` 组成



