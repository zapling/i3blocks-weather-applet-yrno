package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type cacheRecord struct {
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

func (c *CacheManager) GetCache(ssid string) string {
	var cache []cacheRecord
	bytes, _ := ioutil.ReadFile(c.getPath(true))
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

func (c *CacheManager) SetCache(ssid string, value string) {
	until := time.Now().Add(time.Hour * 1).Unix()
	record := cacheRecord{ValidUntil: until, SSID: ssid, Value: value}

	cache := c.getCurrentCache()
	cache = append(cache, record)

	c.write(cache)
}

// getCurrentCache returns the currently cached records from the cache file
func (c *CacheManager) getCurrentCache() []cacheRecord {
	var cache []cacheRecord
	bytes, err := ioutil.ReadFile(c.getPath(true))
	if err != nil {
		c.createEmptyCacheFile()
		bytes = []byte("[]")
	}

	json.Unmarshal(bytes, &cache)

	return cache
}

// write the supplied cache to the cache file, truncating any previous data
func (c *CacheManager) write(cache []cacheRecord) {
	file, err := os.OpenFile(c.getPath(true), os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Could not open cache file")
		os.Exit(1)
	}

	bytes, err := json.Marshal(&cache)
	if err != nil {
		fmt.Println("Could not convert cache records to json")
		os.Exit(1)
	}

	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println("Could not write data to cache file")
		os.Exit(1)
	}
}

// createEmptyCacheFile creates an empty cache file
func (c *CacheManager) createEmptyCacheFile() {
	err := os.MkdirAll(c.getPath(false), 0755)
	if err != nil {
		fmt.Println("Could not create cache folder(s)")
		os.Exit(1)
	}

	file, err := os.Create(c.getPath(true))
	if err != nil {
		fmt.Println("Could not create cache file")
		os.Exit(1)
	}
	_, err = file.Write([]byte("[]"))
	if err != nil {
		fmt.Println("Could not write data to cache file")
		os.Exit(1)
	}

	file.Close()
}

// getPath get the cache folder path
func (c *CacheManager) getPath(fullpath bool) string {
	path := c.cacheFolder + "/weather-applet"
	if fullpath == true {
		path += "/cache.json"
		return path
	}

	return path
}
