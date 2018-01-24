# Cache for MYDICTIONARY

[简体中文版](./README.zh-Hans.md)

### 1. Introduction

It is a module used for caching online query result for [MYDICTIONARY](https://github.com/zzc-tongji/mydictionary).

### 2. Cache Item

``` go
type ItemStruct struct {
	QueryString  string   `json:"queryString"`
	Word         string   `json:"word"`
	Definition   []string `json:"definition"`
	Status       string   `json:"status"`
	CreationTime int64    `json:"creationTime"`
}
```

`ItemStruct` is a structure and has got these members:

- `QueryString` indicates the string of query.
- `Word` indicates the word.
- `Definition` indicates definitions.
- `Status` indicates the status.
- `CreationTime` is an unix timestamp which indicates when this item is created.

### 3. Cache

``` go
type CacheStruct struct {
	path         string
	shelfLifeDay int64
	Content      []ItemStruct `json:"content"`
}
```

`CacheStruct` is a structure.

It has got a public member `Content` which stores all cache items.

``` go
func (cache *CacheStruct) Read(path string, shelfLifeDay int64) (err error)
```

The function is used for reading cache from a JSON file indicates by `path`, and set the life period in days for all cache items by `shelfLifeDay`.

- If the file indicated by `path` is not existent, create it.
- If `shelfLifeDay` is 0, cache will never expire.

After reading cache, each item will be checked whether it is expired (determined by its `CreationTime`, cache's `shelfLifeDay` and the current time). **Then, all expired items will be removed.**

``` go
func (cache *CacheStruct) Query(queryString string) (item ItemStruct, err error)
```

The function is used for searching `queryString` from cache.

``` go
func (cache *CacheStruct) Add(item ItemStruct)
```

The function is used for adding `item` to cache.

``` go
func (cache *CacheStruct) Write()
```

The function is used for writing cache to the JSON file which it comes from.

### 3. Others

- All code files are edited by [Atom](https://atom.io/).
- All ".md" files are edited by [Typora](http://typora.io).
- The style of all ".md" files is [Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown).
- There is a LF (Linux) at the end of each line.
