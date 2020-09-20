package tinysearch

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Inverted index
type Index struct {
	Dictionary     map[string]PostingsList // 辞書
	TotalDocsCount int                     // ドキュメントの総数
}

// NewIndex create a new index
func NewIndex() *Index {
	dict := make(map[string]PostingsList)
	return &Index{
		Dictionary:     dict,
		TotalDocsCount: 0,
	}
}

// DocumentID is the same AUTO_INCREMENT of MySQL
type DocumentID int64

// Posting is an element of PostingsList
type Posting struct {
	DocID         DocumentID
	Positions     []int
	TermFrequency int
}

func NewPosting(docID DocumentID, positions ...int) *Posting {
	return &Posting{docID, positions, len(positions)}
}

// PostingsList has some Posting. This is obtained as value in Dictionary
type PostingsList struct {
	*list.List
}

func NewPostingsList(postings ...*Posting) PostingsList {
	l := list.New()
	for _, posting := range postings {
		l.PushBack(posting)
	}
	return PostingsList{l}
}

func (pl PostingsList) add(p *Posting) {
	pl.PushBack(p)
}

func (pl PostingsList) last() *Posting {
	e := pl.List.Back()
	if e == nil {
		return nil
	}
	return e.Value.(*Posting)
}

func (pl PostingsList) Add(new *Posting) {
	last := pl.last()
	if last == nil || last.DocID != new.DocID {
		pl.add(new)
		return
	}
	last.Positions = append(last.Positions, new.Positions...)
	last.TermFrequency++
}

func (idx Index) String() string {
	var padding int
	keys := make([]string, 0, len(idx.Dictionary))
	for k := range idx.Dictionary {
		l := utf8.RuneCountInString(k)
		if padding < l {
			padding = l
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	strs := make([]string, len(keys))
	format := " [%-" + strconv.Itoa(padding) + "s] -> %s"
	for i, k := range keys {
		if PostingsList, ok := idx.Dictionary[k]; ok {
			strs[i] = fmt.Sprintf(format, k, PostingsList.String())
		}
	}
	return fmt.Sprintf("total documents : %v\ndictionary:\n%v\n",
		idx.TotalDocsCount, strings.Join(strs, "\n"))
}

func (pl PostingsList) String() string {
	str := make([]string, 0, pl.Len())
	for e := pl.Front(); e != nil; e = e.Next() {
		str = append(str, e.Value.(*Posting).String())
	}
	return strings.Join(str, "=>")
}

func (p Posting) String() string {
	return fmt.Sprintf("(%v, %v, %v)",
		p.DocID, p.TermFrequency, p.Positions)
}
