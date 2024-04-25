package cache

import (
	"fmt"
	"knife"
	"net/http"
	"sync"
	"time"
)

// Cache  缓存中间件
// timeout 缓存过期时间（单位秒）
// duration 缓存监视器检查间隔时间（单位秒）
func Cache(expiration int, duration int) knife.MiddlewareFunc {
	cacheMap := &cache{
		cacheMap:   make(map[string]*cacheItem),
		mutex:      sync.RWMutex{},
		expiration: expiration,
	}

	//启动缓存监视器
	go cacheMap.startMonitor(time.Duration(duration) * time.Second)

	return func(context *knife.Context) {
		key := fmt.Sprintf("%s#%s", context.Req.Method, context.Req.URL.Path)

		if cacheItem, found := cacheMap.getFromCache(key); found {
			_, err := context.Writer.Write(cacheItem.data)
			if err != nil {
				knife.Logger.Printf("Cache middleware write data error %s ", err)
			}
			context.Abort(http.StatusOK)
			knife.Logger.Printf("Use cache middleware data")
		} else {
			cacheWriter := &cacheResponseWriter{
				HttpResponseWriter: context.Writer,
			}
			context.Writer = cacheWriter

			context.Next()

			cacheMap.saveToCache(key, cacheWriter.data)
		}
	}
}

// 缓存结构体
type cacheItem struct {
	data       []byte
	expiration time.Time
}

// 判断是否过期
func (item *cacheItem) isExpired() bool {
	return time.Now().After(item.expiration)
}

// 缓存对象
type cache struct {
	cacheMap   map[string]*cacheItem
	mutex      sync.RWMutex
	expiration int
}

// 从缓存中获取数据
func (c *cache) getFromCache(key string) (*cacheItem, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	cacheItem, found := c.cacheMap[key]
	if found && !cacheItem.isExpired() {
		return cacheItem, true
	}
	return nil, false
}

// 将数据存入缓存
func (c *cache) saveToCache(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cache := &cacheItem{
		data:       data,
		expiration: time.Now().Add(time.Duration(c.expiration) * time.Second), // 设置缓存过期时间，这里设置为5分钟
	}
	c.cacheMap[key] = cache
}

// 缓存监视器，判断缓存中的key是否过期，过期移除
// duration 监视器检查触发时间间隔
func (c *cache) startMonitor(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			knife.Logger.Printf("check cache is isExpired ")
			c.mutex.Lock()
			for key, item := range c.cacheMap {
				if item.isExpired() {
					knife.Logger.Printf("cache is isExpired,key:%s ", key)
					delete(c.cacheMap, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}

// 自定义的ResponseWriter，记录数据
type cacheResponseWriter struct {
	knife.HttpResponseWriter
	data []byte
}

// 重写Write方法，将数据暂存
func (w *cacheResponseWriter) Write(data []byte) (int, error) {
	w.data = data
	return w.HttpResponseWriter.Write(data)
}
