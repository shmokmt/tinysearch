# tinysearch

![reviewdog](https://github.com/shmokmt/tinysearch/workflows/reviewdog/badge.svg)

Tiny full-text search engine for learning

- Go
- inverted index
- tf-idf
- MySQL8.x

## How to dev

```sh
docker-compose up -d
make install
# create index
tinysearch create testdata/document
# search by full-text search engine
tinysearch search "qurrel sir"
```

## Comparison

## See Also

- [Go で検索エンジンに入門する本](https://booth.pm/ja/items/1576277)
