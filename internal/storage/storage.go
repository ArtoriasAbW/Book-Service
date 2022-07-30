package storage

import (
	"log"
	"strconv"

	"github.com/pkg/errors"
)

var data map[uint]*Book

var BookNotExists = errors.New("book doesn't exist")
var BookExists = errors.New("book already exists")

func init() {
	log.Println("init storage")
	data = make(map[uint]*Book)
}

func List() []*Book {
	res := make([]*Book, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func Add(b *Book) error {
	if _, ok := data[b.GetId()]; ok {
		return errors.Wrap(BookExists, strconv.FormatUint(uint64(b.GetId()), 10))
	}
	data[b.GetId()] = b
	return nil
}

func Update(b *Book) error {
	if _, ok := data[b.GetId()]; !ok {
		return errors.Wrap(BookNotExists, strconv.FormatUint(uint64(b.GetId()), 10))
	}
	data[b.GetId()] = b
	return nil
}

func Remove(id uint) error {
	if _, ok := data[id]; ok {
		delete(data, id)
		return nil
	}
	return errors.Wrap(BookNotExists, strconv.FormatUint(uint64(id), 10))
}

func MarkRead(id uint) error {
	if _, ok := data[id]; ok {
		data[id].unread = false
		return nil
	}
	return errors.Wrap(BookNotExists, strconv.FormatUint(uint64(id), 10))
}

