package cache

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Cache struct {
	ReadMessages int `yaml:"readMessages"`
}

var defaultCache = Cache{0}

func GetCache() *Cache {
	buf, err := ioutil.ReadFile("./cache.yaml")
	if err != nil {
		log.Println("Could not load cache file. Creating one now.", err)
		SetCache(defaultCache)
		return GetCache()
	}

	c := &Cache{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		log.Fatal("Could not load config.", err)
	}

	return c
}

func SetCache(newCache Cache) {
	d, err := yaml.Marshal(&newCache)

	if err != nil {
		log.Fatal("Could Encode Values for cache file.", err)
	}

	err = ioutil.WriteFile("./cache.yaml", d, 0644)

	if err != nil {
		log.Fatal("Error while writing to cache file.", err)
	}

}
