package lru

import (
	"sync"
	"time"
)

type EntryWithTimeout struct {
	key    string
	value  any
	expire int64
	pre    *EntryWithTimeout
	next   *EntryWithTimeout
}

type CacheWithTimeout struct {
	mu       sync.RWMutex
	timeout  time.Duration
	cap      int
	entryMap map[string]*EntryWithTimeout
	head     *EntryWithTimeout
	tail     *EntryWithTimeout
}

func NewCacheWithTimeout(cap int, timeout time.Duration) *CacheWithTimeout {
	mc := &CacheWithTimeout{
		timeout:  timeout,
		cap:      cap,
		entryMap: make(map[string]*EntryWithTimeout),
	}
	return mc
}

func (c *CacheWithTimeout) Put(id string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().UnixNano()
	if v, ok := c.entryMap[id]; ok {
		// 已存在则更新
		v.value = value
		v.expire = now + c.timeout.Nanoseconds()
		// 重新移动到head
		c.moveEntryToHead(v)
		return
	}
	// 不存在则新建
	v := &EntryWithTimeout{
		key:    id,
		value:  value,
		expire: now + c.timeout.Nanoseconds(),
		next:   c.head,
	}
	c.addNewEntryToHead(v)
}

func (c *CacheWithTimeout) Get(id string) (value any) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if v, ok := c.entryMap[id]; ok {
		// 判断是否过期
		if time.Now().UnixNano() > v.expire {
			// 删除此值，并返回nil
			delete(c.entryMap, v.key)
			if v.pre != nil {
				v.pre.next = v.next
			} else {
				c.head = v.next
			}
			if v.next != nil {
				v.next.pre = v.pre
			} else {
				c.tail = v.pre
			}
			return nil
		}
		// 没过期，返回值，并移动到head
		c.moveEntryToHead(v)
		return v.value
	}
	return nil
}

func (c *CacheWithTimeout) addNewEntryToHead(e *EntryWithTimeout) {
	if c.head != nil {
		c.head.pre = e
	}
	c.head = e
	if c.tail == nil {
		c.tail = e
	}
	c.entryMap[e.key] = e
	if len(c.entryMap) > c.cap {
		// out of cap, remove tail
		c.removeTail()
	}
}

func (c *CacheWithTimeout) removeTail() {
	if c.tail == nil {
		return
	}
	removeKey := c.tail.key
	c.tail = c.tail.pre
	if c.tail != nil {
		c.tail.next = nil
	}
	delete(c.entryMap, removeKey)
}

func (c *CacheWithTimeout) moveEntryToHead(e *EntryWithTimeout) {
	if c.head == e {
		return
	}
	// 摘除当前节点
	if e.pre != nil {
		e.pre.next = e.next
	}
	if e.next != nil {
		e.next.pre = e.pre
	}

	if e == c.tail {
		c.tail = e.pre
	}

	e.pre = nil
	e.next = c.head
	if c.head != nil {
		c.head.pre = e
	}
	c.head = e
	if c.tail == nil {
		c.tail = e
	}
}
