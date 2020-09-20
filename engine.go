package tinysearch

import (
	"database/sql"
	"io"
	"os"
	"path/filepath"
)

// Engine is type for full-text search engine.
// This has a directory name of indexes info.
type Engine struct {
	tokenizer     *Tokenizer
	indexer       *Indexer
	documentStore *DocumentStore
	indexDir      string
}

func NewSearchEngine(db *sql.DB) *Engine {
	tokenizer := NewTokenizer()
	indexer := NewIndexer(tokenizer)
	documentStore := NewDocumentStore(db)

	path, ok := os.LookupEnv("INDEX_DIR_PATH")
	if !ok {
		current, _ := os.Getwd()
		path = filepath.Join(current, "_index_data")
	}
	return &Engine{
		tokenizer:     tokenizer,
		indexer:       indexer,
		documentStore: documentStore,
		indexDir:      path,
	}

}

func (e *Engine) AddDocument(title string, reader io.Reader) error {
	id, err := e.documentStore.save(title)
	if err != nil {
		return err
	}
	e.indexer.update(id, reader)
	return nil
}

func (e *Engine) Flush() error {
	writer := NewIndexWriter(e.indexDir)
	return writer.Flush(e.indexer.index)
}
