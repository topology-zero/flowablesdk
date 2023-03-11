# flowablesdk

使用 GO 语言调用 flowable 暴露出来的 rest 接口

[文档地址](https://topology-zero.github.io/flowablesdk/)

### 背景

在 GO 开源的项目中，没有什么使用的特别广泛的工作流项目，但是在 Java 项目中 有一个非常著名的开源工作流项目，那就是大名鼎鼎的 flowable，但是 flowable 没有提供对应的 GO sdk ，但是 flowable 提供了 rest 接口跨平台调用，本项目 就是将 rest 接口封装一层。提供 struct 返回，简化调用

### 使用

```shell
go get -v github.com/topology-zero/flowablesdk@v1.3.9
```

### 前置要求

- flowable
    - java8
    - spring-boot
    - spring-boot 集成 flowable
    - 扩展 flowable
- go
    - golang 1.18+
    - gin
    - gorm
    - gorm-gen
    - logrus
    - go-zero
- vue
    - vue-admin-template
    - bpmn-js

### 配套项目

- [spring-boot + flowable 项目](https://github.com/topology-zero/flowable-rest)
- [后台管理UI](https://github.com/topology-zero/go-flowable-vue)
- [后台管理接口](https://github.com/topology-zero/go-flowable-admin)

### 联系

##### 可直接在 issues 提出

##### 微信

<img decoding="async" src="https://tc.masterjoy.top/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20230216101038.jpg" width="30%" />

### License

MIT
