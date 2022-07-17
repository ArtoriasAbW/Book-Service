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
	book, _ := NewBook("История западной философии", "Бертран Рассел", true)
	if err := Add(book); err != nil {
		log.Panic(err)
	}
}

func List() []*Book {
	res := make([]*Book, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func Add(u *Book) error {
	if _, ok := data[u.GetId()]; ok {
		return errors.Wrap(BookExists, strconv.FormatUint(uint64(u.GetId()), 10))
	}
	data[u.GetId()] = u
	return nil
}

func Update(u *Book) error {
	if _, ok := data[u.GetId()]; !ok {
		return errors.Wrap(BookNotExists, strconv.FormatUint(uint64(u.GetId()), 10))
	}
	data[u.GetId()] = u
	return nil
}

func Remove(id uint) error {
	if _, ok := data[id]; ok {
		delete(data, id)
		return nil
	}
	return errors.Wrap(BookNotExists, strconv.FormatUint(uint64(id), 10))
}
