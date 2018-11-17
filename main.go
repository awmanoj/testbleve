package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("blacklist.index", mapping)
	if err != nil {
		log.Println("err", err)
		return
	}

	data1 := struct {
		ID     uint64
		Code   string
		Name   string
		ShopID uint64
	}{
		ID:     3801,
		Code:   "effect of this product is so amazing and not comparable to any other product",
		Name:   "test2",
		ShopID: 2884978,
	}

	data2 := struct {
		ID          uint64
		Description string
		Name        string
		ColID       uint64
	}{
		ID:          3802,
		Description: "this product is amazing in its quality and has great effect.",
		Name:        "test1",
		ColID:       323243,
	}

	// index some data
	err = index.Index(fmt.Sprintf("%d", data1.ID), data1)
	// TODO: check error

	err = index.Index(fmt.Sprintf("%d", data2.ID), data2)
	// TODO: check error

	// search for some text
	query := bleve.NewMatchQuery(os.Args[1])
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)

	log.Println(searchResults)
}
