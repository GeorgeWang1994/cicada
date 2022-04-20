# cicada

## 架构

整体流程为

1. 探针通过RPC向event发送数据
2. event在接收到事件数据以后，将数据发向kafka，接着启动多个协程从kafka中读取数据，接着数据分发到数据库的redis中（近期都只对redis进行操作），
   同时也将数据写入到告警的缓存中，同时启动多个协程将在redis中超过一定分钟范围内的事件数据刷进clickhouse中
3. 告警组件从后台拉取告警策略信息，同时从缓存中读取数据，同时开启合并操作（将多个事件合并到一起），将数据经过告警策略后，发给对应的provider
4. portal后台组件从clickhouse或者redis中读取事件数据展示出来
5. 用户可以在后端配置告警策略信息，具体数据存储在mysql中

## 组件

* agentd: 专门用来上报数据用的探针
* judge：根据告警策略判断是否需要报警
* ev：具备事件转发的功能，并且将数据存储到缓存中，定期将缓存中的数据同步到CK
* alarm：负责将数据告警出去，会提供多个provider
* portal：面向用户的查询界面，提供数据的查询，以及对告警策略的配置

## Todo

1. 实现可以根据指定部分的事件名称来告警
2. 实现可以在指定时间内告警
3. 实现可以根据各种业务字段来告警（比如根据协议、端口等等）
4. ~~在ev中直接将事件数据存储到CK中不太合理，有可能存在短时间内频繁更改事件的需求~~
   * ✅已经将最近10分钟的数据存储到Redis中
5. 某些特殊情况下没有考虑到数据被取出来但是未处理成功的情况应该怎么办（比如在读取最近30分钟的事件数据）
6. 确认目前的func能否支持现有的业务，如果不能建议重新思考实现方案
7. 支持k8s
8. 事件的条件告警怎么办
9. 组件之间心跳怎么办，是否需要
10. 对于近期存在redis中的数据如何做统计
11. 所有事件是否共用一套处理逻辑还是单独处理
