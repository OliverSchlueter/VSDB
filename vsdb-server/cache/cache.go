package cache

import (
	"time"
)

type Cache struct {
	data   map[string]string
	expire map[string]int64
}

func CreateCache() Cache {
	c := Cache{
		data:   map[string]string{},
		expire: map[string]int64{},
	}

	go c.startCacheRoutine()

	return c
}

func (receiver Cache) startCacheRoutine() {
	for true {
		receiver.checkForExpire()
		time.Sleep(1000 * time.Millisecond)
	}
}

func (receiver Cache) checkForExpire() {
	for key := range receiver.expire {
		contains, expire := receiver.Contains(key)

		if !contains {
			delete(receiver.expire, key)
			continue
		}

		now := time.Now().UnixMilli()

		if now >= expire {
			delete(receiver.data, key)
			delete(receiver.expire, key)
		}
	}
}

func (receiver Cache) InsertOrUpdate(key string, value string) {
	receiver.data[key] = value
}

// InsertAndExpire
// argument expire - time in milliseconds
func (receiver Cache) InsertAndExpire(key string, value string, expire int64) {
	receiver.InsertOrUpdate(key, value)
	receiver.expire[key] = time.Now().UnixMilli() + expire
}

func (receiver Cache) Delete(key string) {
	delete(receiver.data, key)
	delete(receiver.expire, key)
}

func (receiver Cache) Contains(key string) (bool, int64) {
	_, prs := receiver.data[key]

	expire, _ := receiver.expire[key]

	return prs, expire
}

func (receiver Cache) Get(key string) string {
	prs, _ := receiver.Contains(key)

	if prs {
		return receiver.data[key]
	}

	return ""
}
