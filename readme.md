# Newbee Mall 牛逼👍
```
1、enter这个词用的好、类似js/ts的index module化开发对外暴露接口
2、代码看起来简洁舒服

```

# Java飞速转Go
```
不求甚解
快速入门
关键理念
copy+实战+idea
```
# 饭一口一口吃 代码一行一行码
```
AI已经摸透大部分人/码农的思维模式和流程模式、企业为所减人力成本、应用级别的开发呈逐渐被替代的趋势
所以穷苦人家的孩子赶紧学AI吧、爱和创新总不能也被一台机器赶超吧
```
## 面向对象
```
1、间接实现多态与继承
组合式api概念
combination = is a+has a
2、init
3、包的可见性
4、函数一等公民
```
## 接收器
```
只接收一个值或引用类型*
```

# copy code by finger practice
```
拿来主义与逆向工程：码者的想法
优化创新：

- 直接用1.20版本 1.16 17 18版本 go mod 麻烦
- 名字、应用级别的非lib的项目没有引用的场景、没有拉代码的必要 niu-mall
- 从main.go开始
- 代码结构 从Java项目中也可知utils cfg model service controller(router) 必须的，component（middleware）成了一个习惯
- 基准压测 go test -test.bench=.*、 go test -v -bench=. benchmark_test.go

- mode entity、vo、param --- service  ---- api --- router

- todo
 搜索：商品 订单
 token redis session化
```