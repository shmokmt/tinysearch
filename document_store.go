package tinysearch

import "database/sql"

// DocumentStore is a connector of MySQL.
type DocumentStore struct {
	db *sql.DB
}

func NewDocumentStore(db *sql.DB) *DocumentStore {
	return &DocumentStore{db: db}
}

func (ds *DocumentStore) save(title string) (DocumentID, error) {
	query := "INSERT INTO documents (document_title) VALUES (?)"
	ret, err := ds.db.Exec(query, title)
	if err != nil {
	}
	id, err := ret.LastInsertId()
	return DocumentID(id), err
}
