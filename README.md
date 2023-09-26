# ttCache
通过装饰器模式实现了多种缓存模式的支持，同时解决了缓存异常问题

支持 OS 内存 缓存
OSCache 基于Map实现
 一、处理过期时间的三种策略
      1.每个key开一个goroutine盯着执行删除策略
         创建方法 ：
                    Set func (b *BuildInMapCache) Set(ctx , key , value，expiration）
                   expiration:不设置过期时间默认永不过期。
      2.开一个goroutine定时轮训
          创建方法 ：NewBuildInMapCacheOneGo
      3.类似于sql连接，下一次使用时候在检查。
          创建方法 ：Set

    


以及 Redis 实现缓存模式
read-through、write-through、cache-aside、write-back 

以及 都可加 singleflight 机制