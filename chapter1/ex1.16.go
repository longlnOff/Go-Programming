package main

import (
	"fmt"
)

const GloabalLimit = 100

const MAxCacheSize int = 10 * GloabalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache) + 1 > MAxCacheSize {
		return
	}
	cache[key] = val
}

func GetBook(title string) string {
	return cacheGet(CacheKeyBook + title)
}

func SetBook(title, val string) {
	cacheSet(CacheKeyBook + title, val)
}

func getCD(title string) string {
	return cacheGet(CacheKeyCD + title)
}

func setCD(title, val string) {
	cacheSet(CacheKeyCD + title, val)
}

func main() {

	cache = make(map[string]string)

	// add a book to the cache
	SetBook("Alice in Wonderland", "Lewis Carroll")
	SetBook("Alice in Wonderland 2", "Lewis Carroll 2")
	fmt.Println(GetBook("Alice in Wonderland"))
	fmt.Println(GetBook("Alice in Wonderland 2"))

	// add a CD to the cache
	setCD("Greatest Hits", "Queen")
	setCD("Greatest Hits 2", "Queen 2")
	fmt.Println(getCD("Greatest Hits"))
	fmt.Println(getCD("Greatest Hits 2"))
}