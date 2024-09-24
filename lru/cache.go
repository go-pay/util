package lru

import (
	"sync"
)

// Entry 缓存中的实体
type Entry struct {
	Key   string // 键
	Value any    // 值
	pre   *Entry // 前一个实体
	next  *Entry // 后一个实体
}

// Cache LRU 缓存
type Cache struct {
	mu    sync.RWMutex
	cache map[string]*Entry
	cap   int
	head  *Entry
	tail  *Entry
}

// NewCache 创建一个新的 LRU 缓存
func NewCache(cap int) *Cache {
	return &Cache{cache: make(map[string]*Entry), cap: cap}
}

// Put 将一个键值对放入缓存
func (c *Cache) Put(key string, v any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 如果键已存在，更新值并将其移动到头部
	if e, has := c.cache[key]; has {
		e.Value = v
		c.moveToHead(e)
		return
	}
	// 创建一个新实体
	entry := &Entry{Key: key, Value: v, pre: nil, next: c.head}
	// 将新实体放入头部
	c.putNewEntryHead(entry)
	// 如果缓存已满，移除尾部实体
	if len(c.cache) > c.cap {
		c.removeTail()
	}
}

// Get 获取键对应的值
func (c *Cache) Get(key string) (v any) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// 如果键存在，移动到头部并返回值
	if e, has := c.cache[key]; has {
		c.moveToHead(e)
		return e.Value
	}
	return nil
}

// moveToHead 将实体移动到链表头部
func (c *Cache) moveToHead(e *Entry) {
	// 如果实体已经在头部，直接返回
	if e == c.head {
		return
	}
	// 将实体从当前位置摘除
	if e.pre != nil {
		e.pre.next = e.next
	}
	if e.next != nil {
		e.next.pre = e.pre
	}
	// 如果e是最后一个实体，设置尾部为e的前一个实体
	if e == c.tail {
		c.tail = e.pre
	}
	// 将e实体放入链表头部
	e.next = c.head
	e.pre = nil
	if c.head != nil {
		c.head.pre = e
	}
	c.head = e
	// 如果链表尾部为空，设置尾部为当前实体e
	if c.tail == nil {
		c.tail = e
	}
}

func (c *Cache) putNewEntryHead(e *Entry) {
	// 如果头部不为空，将头部前指针指向新实体
	if c.head != nil {
		c.head.pre = e
	}
	// 设置新实体为头部
	c.head = e
	// 如果尾部为空，设置尾部为新实体
	if c.tail == nil {
		c.tail = e
	}
	// 将新实体放入缓存
	c.cache[e.Key] = e
}

func (c *Cache) removeTail() {
	// 如果尾部为空，直接返回
	if c.tail == nil {
		return
	}
	// 获取尾部实体
	removeEntryKey := c.tail.Key
	// 将尾部前移
	c.tail = c.tail.pre
	// 如果尾部不为空，将尾部的后指针置空
	if c.tail != nil {
		c.tail.next = nil
	}
	// 从缓存中删除尾部实体
	delete(c.cache, removeEntryKey)
}
