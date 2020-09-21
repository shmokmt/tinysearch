package tinysearch

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func setup() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/tinysearch")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("TRUNCATE TABLE documents")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.RemoveAll("_index_data"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("_index_data", 0777); err != nil {
		log.Fatal(err)
	}
	return db
}

func Testmain(m *testing.M) {
	testDB = setup()
	defer testDB.Close()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateIndex(t *testing.T) {

}

func TestSearch(t *testing.T)  {
	engine := NewSearchEngine(testDB)
	query := "Quarrel, sir."
	actual, err := engine.Search(query, 5)
	if err != nil {
		t.Fatalf("failed searchTopK: %v", err)
	}
	expected := []*SearchResult{
		{3, 1.754887502163469, "test3"},
		{1, 1.1699250014423126, "test1"},
	}

	for !reflect.DeepEqual(actual, expected) {
		t.Fatal("\ngot:\n%v\nwant:\n%v\n", actual, expected)
	}
	
}
