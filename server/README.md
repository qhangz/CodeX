# CodeX
## 目录结构
codex now
```text
├── server
    ├── api             (api层)
    │   └── v1          (v1版本接口具体实现)
    │   └── controllers (接口功能函数管理)                                    
    ├── middleware      (中间件层)                        
    ├── model           (模型层)                    
    │   ├── request     (入参结构体)                        
    │   └── response    (出参结构体)                                                 
    ├── resource        (静态资源文件夹)                        
    │   ├── json        (json文件)                                               
    ├── router          (路由层)     
    │   └── router      (router管理)       
    ├── files           (文件)                      
```
codex
```text
├── server
    ├── api             (api层)
    │   └── v1          (v1版本接口具体实现)
    │   └── controllers (接口功能函数管理)
    ├── docs            (swagger文档目录)
    ├── global          (全局对象)                                        
    ├── middleware      (中间件层)                        
    ├── model           (模型层)                    
    │   ├── request     (入参结构体)                        
    │   └── response    (出参结构体)                                                 
    ├── resource        (静态资源文件夹)                        
    │   ├── json        (json文件)                                               
    ├── router          (路由层)     
    │   └── router      (router管理)                       
    ├── source          (source层)                    
    └── utils           (工具包)                    
        ├── timer       (定时器接口封装)                        
        └── upload      (oss接口封装)      
```
aiapp
```text
├── controllers       (api层)
│   └── controllers   (接口功能函数管理)
│   └── router        (路由管理)
├── data              (json文件)
├── database          (核心文件)
│   └── util          (ConvertIDtoHexID)
│   └── .go           (各种功能函数与数据库实现)
├── docs              (swagger文档目录)
├── middleware        (中间件层)                     
├── models            (结构体)                        
```

```text
├── server
    ├── api             (api层)
    │   └── v1          (v1版本接口)
    ├── config          (配置包)
    ├── core            (核心文件)
    ├── docs            (swagger文档目录)
    ├── global          (全局对象)                    
    ├── initialize      (初始化)                        
    │   └── internal    (初始化内部函数)                            
    ├── middleware      (中间件层)                        
    ├── model           (模型层)                    
    │   ├── request     (入参结构体)                        
    │   └── response    (出参结构体)                            
    ├── packfile        (静态文件打包)                        
    ├── resource        (静态资源文件夹)                        
    │   ├── json        (json文件)                        
    │   ├── page        (表单生成器)                        
    │   └── template    (模板)                            
    ├── router          (路由层)                    
    ├── service         (service层)                    
    ├── source          (source层)                    
    └── utils           (工具包)                    
        ├── timer       (定时器接口封装)                        
        └── upload      (oss接口封装)               
```
## api description