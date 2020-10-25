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
# set env
export INDEX_DIR_PATH="./testdata/index"
# create index
go run ./cmd/tinysearch/main.go create ./testdata/document
# search by full-text search engine
go run ./cmd/tinysearch/main.go search "qurrel sir"
```

## See Also

- [Go で検索エンジンに入門する本](https://booth.pm/ja/items/1576277)
