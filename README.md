# cicada

## 架构

整体流程为

1. 探针通过RPC向event发送数据
2. event在接收到事件数据以后，将数据发向kafka，接着启动多个协程从kafka中读取数据，接着数据分发到数据库的redis中（近期都只对redis进行操作），
   同时也将数据写入到告警的缓存中，同时启动多个协程将在redis中超过一定分钟范围内的事件数据刷进clickhouse中
3. 告警组件从后台拉取告警策略信息，同时从缓存中读取数据，同时开启合并操作（将多个事件合并到一起），将数据经过告警策略后，发给对应的provider
4. portal后台组件从clickhouse或者redis中读取事件数据展示出来
5. 用户可以在后端配置告警策略信息，具体数据存储在mysql中

可能存在的问题

1. 条件告警怎么办
2. 组件心跳怎么办
3. 对于近期存在redis中的数据如何做统计
4. 所有事件是否共用一套处理逻辑还是单独处理

## 组件

* agentd: 专门用来上报数据用的探针
* judge：根据告警策略判断是否需要报警
* ev：具备事件转发的功能，并且将数据存储到数据库中
* alarm：负责将数据告警出去，会提供多个provider
* portal：面向用户的查询界面，提供数据的查询，以及对告警策略的配置
