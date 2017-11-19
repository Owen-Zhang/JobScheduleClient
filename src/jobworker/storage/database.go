package storage

import (
	"sync"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataStorageArgs struct {
	Hosts    string,
	DBName   string,
	Auth     map[string]string
}

type DataStorage struct {
	mutex  *sync.Mutex,
	mdb   *mgo.Database
	mss   *mgo.Session
	cc map[string]*mgo.Collection
}

func NewDataStorage(arg *DataStorageArgs) (*DataStorage, error) {
	
}
