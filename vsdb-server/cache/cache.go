package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Cache struct {
	savePersistent bool
	savePath       string
	data           map[string]string
	expire         map[string]int64
}

func CreateCache(savePersistent bool, savePath string) Cache {
	c := Cache{
		savePersistent: savePersistent,
		savePath:       savePath,
		data:           map[string]string{},
		expire:         map[string]int64{},
	}

	if savePersistent {
		c.loadDataFromPath()
		go c.startSaveRoutine()
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

func (receiver Cache) startSaveRoutine() {
	for true {
		receiver.Save()
		time.Sleep(10 * time.Second)
	}
}

func (receiver Cache) Save() {
	path := receiver.savePath + "data.json"

	fmt.Println("Saving data to " + path)

	data, _ := json.Marshal(receiver.data)

	file, err := os.Create(path)
	if err != nil {
		return
	}

	defer file.Close()

	file.Write(data)
}

func (receiver Cache) loadDataFromPath() {
	_, err := os.OpenFile(receiver.savePath, os.O_RDONLY, os.ModeTemporary)

	if os.IsNotExist(err) {
		return
	}

	d, err := os.ReadFile(receiver.savePath)
	if err != nil {
		return
	}

	json.Unmarshal(d, &receiver.data)
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

func (receiver Cache) GetAllEntries() map[string]string {
	newMap := make(map[string]string)

	for k, v := range receiver.data {
		newMap[k] = v
	}
	return newMap
}

func (receiver Cache) GetAllKeys() []string {
	var keys []string

	for k, _ := range receiver.data {
		keys = append(keys, k)
	}

	return keys
}
