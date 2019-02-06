# zhihu
仿知乎网站
> 在原项目的基础上进行了调整优化，集成代码功能，支持oj。

## 截图
![file](screenshots/question.png)
![file](screenshots/comment.png)

## 项目结构
```
-zhihu
    |-config        配置文件
    |-controllers   控制器
    |-middleware    中间件
    |-models        数据库访问
    |-router        路由转发
    |-static        静态资源
    |-utils         实用工具
    |-vendor        项目依赖
    |-views         模板文件
    |-main.go       程序执行入口
```

## 安装
**软件依赖**
- Golang
- MySQL
- Redis

本项目使用govendor管理依赖包
```
go get -u -v github.com/kardianos/govendor
```

```
go get -u - v github.com/novioleo/zhihu
cd $GOPATH/src/github.com/novioleo/zhihu
govendor sync
go run main.go
```

## Docker
直接运行`docker-compose -d up`，能够一键启动。
如果未安装docker-compose，可以使用`pip install docker-compose`，也可以直接在官网下载。

## TODO List
- [x] 修改原项目代码，保证其他人能直接clone下来运行
- [ ] 在提问页面上增加算法问题选项
- [ ] 打通Judge Server，实现后端功能
- [ ] 增加回答页面的oj功能
