# mywork

beegou框架基础+layui的轻博客前台
---
# 环境要求
golang 1.8+
---
beego框架
---
go get github.com/satori/go.uuid   -- 获取uuid的用户包

go get  github.com/jinzhu/gorm     -- 获取gorm中的orm

go get  github.com/go-sql-driver/mysql --获取mysql 驱动包

golang.org/x/net/html golang中国下载 放入src目录

go install golang.org/x/net/html     安装

go get github.com/PuerkitoBio/goquery  --文章截取包

go get github.com/dchest/captcha --验证码包


go get github.com/molizz/goip   -- 地址转换包  没用

go get github.com/tabalt/ipquery
# 配置环境
- conf 目录中的db.conf是配置在数据库，配置你当前的数据库账号地址密码。
- session使用的是文件保存，默认存储在doc环境下，你可以修改配置，调整目录，你也可以使用redis活memcache等作为session存放，具体参考[beego配置redis](https://beego.me/docs/module/session.md)
- 代码默认开启debug模式，如果需要关闭，可以在 app.conf 中的进行调整，包括端口号，项目名称等
# 运行方法
----
bee run blog // 参数是你的beego的项目名称

# 数据库
-------
- mysql数据库 使用的是自动迁移，在首次运行环境过程中会自动生成数据表

