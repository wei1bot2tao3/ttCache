# ttCache
通过装饰器模式实现了多种缓存模式的支持，同时解决了缓存异常问题

OSCache是一个使用内置的 Map 实现的 Go 缓存包。它提供了各种缓存管理功能，例如设置带有过期时间的键值对、按键检索值以及从缓存中删除条目。
OSCache 基于Map实现
 使用方法：创建 BuildInMapCache
         请使用NewBuildInMapCache 函数 同时指定大小： 
         cache := OSCache.NewBuildInMapCache(size)
        
  一、处理过期时间的三种策略
      1.每个key开一个goroutine盯着执行删除策略
         创建方法 ：NewBuildInMapCacheGos（cache）
      2.开一个goroutine定时轮训
          创建方法 ：NewBuildInMapCacheOneGo（cache, expiration）轮训时间
      3.类似于sql连接，下一次使用时候在检查。
          创建方法 ：Set

    


支持 Redis 实现缓存模式
read-through、write-through、cache-aside、write-back 

以及 都可加 singleflight 机制