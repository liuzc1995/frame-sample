go web框架
===
## 此框架是本人基于本人导师fook的go web框架及BONFY的go-mege-master框架搭建而成

# 2019-09-23

### 修改点:

#### 关于模板添加
        经修改,可在template文件夹下自定义多个文件夹进行模板分类
        1. 模板添加在template根目录下的自定义文件夹中，例：template/content/index.html。
        2. 目前模板只识别template目录下的一级文件夹.html文件


# 2019-09-25
### 新增

        1. 静态资源目录配置
        2. 日志输出文件 生成目录 /run/

# 2019-09-26
### 新增

        1. 文件监控
        可配置 config.yml 配置项 server.watch_path 后追加虚监控的目录或者文件
            执行方式: 
                go build 生成可执行文件 
                运行可执行文件 ./server -auto=true  
                * 注:运行可执行文件后/run/目录下生成server.pid来记录进程
        1. 端口、日志、进程、监控文件路径可配置
                config.yml文件:
                    server.port:       端口配置
                    server.log:        日志配置
                    server.pid:        进程配置
                    server.watch_path: 端口配置

# 2019-09-28
### 修改点

        1. 路由调整
                分层配置路由 依赖库mux
        2. tree
                .
                ├── README.md
                ├── cmd
                │   └── db_init
                │       └── main.go
                ├── config
                │   └── g.go
                ├── config.yml
                ├── go.mod
                ├── go.sum
                └── server
                    ├── controller
                    │   ├── g.go
                    │   ├── home.go
                    │   └── utils.go
                    ├── model
                    │   ├── g.go
                    │   ├── user.go
                    │   └── utils.go
                    ├── run
                    ├── server.go
                    ├── static
                    │   ├── css
                    │   │   └── global.css
                    │   ├── img
                    │   └── js
                    ├── templates
                    │   ├── _base.html
                    │   └── index
                    │       └── index.html
                    ├── vm
                    │   ├── g.go
                    │   └── index.go
                    └── watch.go


### 撩我方式

                QQ: 52213314
                wx: a542213314