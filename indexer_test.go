package tinysearch

import (
	"reflect"
	"strings"
	"testing"
)

func TestUpdate(t *testing.T) {
	collection := []string{
		"Do you quarrel, sir?",
		"Quarrel sir! no, sir!",
		"No better.",
		"Well, sir",
	}
	indexer := NewIndexer(NewTokenizer())
	for i, doc := range collection {
		indexer.update(DocumentID(i), strings.NewReader(doc))
	}
	actual := indexer.index
	expected := &Index{
		Dictionary: map[string]PostingsList{
			"better": NewPostingsList(NewPosting(2, 1)),
		},
		TotalDocsCount: 4,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("want: %v\n got:%v\n", expected, actual)
	}

}
