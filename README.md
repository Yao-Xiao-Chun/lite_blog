# lite_blog

beegou + layui博客 
---
### 环境
+ golang 1.8+
+ port:8080
---


### 配置环境
- conf 目录中的db.conf是配置在数据库，配置你当前的数据库账号地址密码。
- session使用的是文件保存，默认存储在doc环境下，你可以修改配置，调整目录，你也可以使用redis活memcache等作为session存放，具体参考[beego配置redis](https://beego.me/docs/module/session.md)
- 代码默认开启debug模式，如果需要关闭，可以在 app.conf 中的进行调整，包括端口号，项目名称等

### 启动
+ go mod tidy 
+ go build ./main.go
### 登录  
+ 地址:http://ip:8080/admin  后台管理页面
+ 账户: 1234567@qq.com
+ 密码: 1234@abcd 
### 数据库

- mysql数据库 使用的是自动迁移，在首次运行环境过程中会自动生成数据表

### FAQ
 + 如果更新包下载错误，可以使用国内源
 + ```shell
     #配置 GOPROXY 环境变量 linux
    export GOPROXY=https://proxy.golang.com.cn,direct
   ```
 + ```shell
    # 配置 GOPROXY 环境变量 windows
    $env:GOPROXY = "https://proxy.golang.com.cn,direct"
    ```
