# MYDICTIONARY 缓存

[English Version](./README.md)

### 1. 简介

这是一个用于缓存[MYDICTIONARY](https://github.com/zzc-tongji/mydictionary/blob/master/README.zh-Hans.md)在线查询结果的模块。

### 2. 缓存项

``` go
type ItemStruct struct {
	QueryString  string   `json:"queryString"`
	Word         string   `json:"word"`
	Definition   []string `json:"definition"`
	Status       string   `json:"status"`
	CreationTime int64    `json:"creationTime"`
}
```

`ItemStruct`是一个结构体，包含下列成员：

- `QueryString`指示了查询字符串。
- `Word`指示了词汇。
- `Definition`指示了释义。
- `Status`指示了状态。
- `CreationTime`是一个UNIX时间戳，它指示了本缓存项的创建时间。

### 3. 缓存

``` go
type CacheStruct struct {
	path         string
	shelfLifeDay int64
	Content      []ItemStruct `json:"content"`
}
```

`CacheStruct`是一个结构体。

它具有一个公开的成员`Content`用于存放所有的缓存项。

``` go
func (cache *CacheStruct) Read(path string, shelfLifeDay int64) (err error)
```

这个函数用于从被`path`指定的JSON文件中读取缓存，同时将`shelfLifeDay`设定为所有缓存项的保存期限（天）。

- 如果`path`指定的文件不存在，那么创建之。
- 如果`shelfLifeDay`为0，那么缓存永不过期。

在读取缓存之后，每个缓存项会被检查以确定是否过期（由缓存项的`CreationTime`、缓存的`shelfLifeDay`和当前时间共同决定）。**随后，所有过期的缓存项将被移除。**

``` go
func (cache *CacheStruct) Query(queryString string) (item ItemStruct, err error)
```

这个函数用于在缓存中搜索`queryString`。

``` go
func (cache *CacheStruct) Add(item ItemStruct)
```

这个函数用于添加`item`到缓存。

``` go
func (cache *CacheStruct) Write()
```

这个函数用于将缓存写回JSON文件。

### 3. 其他

- 所以代码文件是用[Atom](https://atom.io/)编写的。
- 所有".md"文件是用[Typora](http://typora.io)编写的。
- 所有".md"文件的风格是[Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown)。
- 各行以LF（Linux）结尾。
