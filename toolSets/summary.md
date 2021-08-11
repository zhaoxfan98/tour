# 命令行工具

## 1.1 打开工具之旅

![](./flag.png)

`flag.Parse`方法，它总是在所有命令行参数注册的最后进行调用，函数功能是解析并绑定命令行参数

`FlagSet.Parse`，其主要承担了 parse 方法的异常分流处理

最后会流转到命令行解析的核心方法 `FlagSet.parseOne` 下进行处理

flag的命令行参数类型是可以自定义的，也就是Value.Set方法，只需要实现其对应的Value相关的两个接口就可以了

```
type Value interface {
	String() string
	Set(string) error
}
```

### 小结

基于第三方开源库 Cobra 和标准库 strings、unicode 实现了多种模式的单词转换，非常简单，也是在日常的工作中较实用的一环，因为我们经常会需要对输入、输出数据进行各类型的转换和拼装。