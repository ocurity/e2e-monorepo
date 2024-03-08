// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

package actor

import (
	iradix "github.com/hashicorp/go-immutable-radix"
)

type PathListStore iradix.Tree

func NewPathListStore(paths []string) *PathListStore {
	if len(paths) == 0 {
		return nil
	}

	txn := iradix.New().Txn()
	for _, path := range paths {
		txn.Insert([]byte(path), struct{}{})
	}

	return (*PathListStore)(txn.Commit())
}

func (s *PathListStore) unwrap() *iradix.Tree { return (*iradix.Tree)(s) }

func (s *PathListStore) Find(path string) (exists bool) {
	_, _, exists = s.unwrap().Root().LongestPrefix([]byte(path))
	return
}
