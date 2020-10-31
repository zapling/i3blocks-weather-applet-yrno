package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type CacheRecord struct {
	ValidUntil int64
	SSID       string
	Value      string
}

type CacheManager struct {
	cacheFolder string
}

func NewCacheManager(cacheFolder string) *CacheManager {
	return &CacheManager{cacheFolder: cacheFolder}
}

func (c *CacheManager) GetValue(ssid string) string {
	var cache []CacheRecord
	bytes, _ := ioutil.ReadFile(c.getPath() + "/cache.json")
	json.Unmarshal(bytes, &cache)

	for index, record := range cache {
		if record.SSID != ssid {
			continue
		}

		if time.Now().Unix() < record.ValidUntil {
			return record.Value
		}

		// cache invalid
		cache = append(cache[:index], cache[index+1:]...)
		c.write(cache)

		return ""
	}

	return ""
}

func (c *CacheManager) WriteCache(ssid string, value string) {
	cache := c.getRecords()

	until := time.Now().Add(time.Hour * 1).Unix()

	record := CacheRecord{ValidUntil: until, SSID: ssid, Value: value}
	cache = append(cache, record)

	file, err := os.OpenFile(c.getPath()+"/cache.json", os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		return // we do not care
	}

	bytes, err := json.Marshal(&cache)
	_, err = file.Write(bytes)
	if err != nil {
		return // we do not care
	}
}

func (c *CacheManager) getRecords() []CacheRecord {
	var cache []CacheRecord
	bytes, err := ioutil.ReadFile(c.getPath() + "/cache.json")
	if err != nil {
		c.createEmptyCacheFile()
		bytes = []byte("[]")
	}

	json.Unmarshal(bytes, &cache)

	return cache
}

func (c *CacheManager) write(cache []CacheRecord) {
	file, err := os.OpenFile(c.getPath()+"/cache.json", os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		return // do we care?
	}

	bytes, err := json.Marshal(&cache)
	if err != nil {
		return // do we care
	}

	_, err = file.Write(bytes)
	if err != nil {
		return // do we care
	}
}

func (c *CacheManager) createEmptyCacheFile() {
	err := os.MkdirAll(c.getPath(), 0755)
	if err != nil {
		return // we do not care
	}

	file, err := os.Create(c.getPath() + "/cache.json")
	if err != nil {
		return // we do not care
	}
	file.Write([]byte("[]"))
	file.Close()
}

func (c *CacheManager) getPath() string {
	return c.cacheFolder + "/weather-applet"
}
