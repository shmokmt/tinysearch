package tinysearch

import (
	"reflect"
	"testing"
)

func TestSearchTopK(t *testing.T) {
	s := NewSearcher("testdata/index")
	actual := s.SearchTopK([]string{"quarrel", "sir"}, 1)

	expected := &TopDocs{2,
		[]*ScoreDoc{
			{3, 1.754887502163469},
		}}
	for !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got: \n%v\nexpected:%v\n", actual, expected)
	}

}
