# micro-mall-comments

#### 介绍
评论服务

#### 软件架构
gRPC应用

#### 框架，库依赖
kelvins框架支持（gRPC，cron，queue，web支持）：https://gitee.com/kelvins-io/kelvins   
g2cache缓存库支持（两级缓存）：https://gitee.com/kelvins-io/g2cache   

#### 安装教程

1.仅构建  sh build.sh   
2 运行  sh build-run.sh   
3 停止 sh stop.sh   

#### 使用说明
参考配置
```toml
[kelvins-server]
Environment = "dev"

[kelvins-logger]
RootPath = "./logs"
Level = "debug"

[kelvins-auth]
Token = "c9VW6ForlmzdeDkZE2i8"
TransportSecurity = false
ExpireSecond = 100

[kelvins-mysql]
Host = "127.0.0.1:3306"
UserName = "root"
Password = "xxx"
DBName = "micro_mall_comments"
Charset = "utf8mb4"
PoolNum =  10
MaxIdleConns = 5
ConnMaxLifeSecond = 3600
MultiStatements = true
ParseTime = true

[kelvins-redis]
Host = "127.0.0.1:6379"
Password = "xxx"
DB = 1
PoolNum = 10

[kelvins-queue-amqp]
Broker = "amqp://micro-mall:szJ9aePR@localhost:5672/micro-mall"
DefaultQueue = "trade_pay_notice"
ResultBackend = "redis://xxx@127.0.0.1:6379/10"
ResultsExpireIn = 36000
Exchange = "trade_pay_notice"
ExchangeType = "direct"
BindingKey = "trade_pay_notice"
PrefetchCount = 5
TaskRetryCount = 3
TaskRetryTimeout = 3600

[email-config]
Enable = false
User = "xxx@qq.com"
Password = "xxxx"
Host = "smtp.qq.com"
Port = "465"

```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request