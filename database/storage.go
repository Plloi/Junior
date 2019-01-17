package database

import (
	"context"
	"fmt"
	"reflect"

	"cloud.google.com/go/datastore"
)

type datastoreDB struct {
	client *datastore.Client
	ctx    context.Context
}

// NewDatastore .
func NewDatastore(projectID string) *datastoreDB {
	DS := &datastoreDB{
		ctx: context.Background(),
	}
	DS.getClient(projectID)
	return DS
}

func (db *datastoreDB) getClient(projectID string) {
	client, err := datastore.NewClient(db.ctx, projectID)
	if err != nil {
		fmt.Printf("Error: Unable to create datastore client: %v", err)
		return
	}
	t, err := client.NewTransaction(db.ctx)
	if err != nil {
		fmt.Printf("datastoredb: could not connect: %v", err)
		return
	}
	if err := t.Rollback(); err != nil {
		fmt.Printf("datastoredb: could not connect: %v", err)
		return
	}
	db.client = client
	return
}

func (db *datastoreDB) Add(a interface{}) (id int64, err error) {
	k := datastore.IncompleteKey(reflect.TypeOf(a).String(), nil)
	k, err = db.client.Put(db.ctx, k, a)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put amiibo: %v", err)
	}

	return k.ID, nil
}

/*
// GetBook retrieves a book by its ID.
func (db *datastoreDB) GetAmiibo(id int64) (*Amiibo, error) {
	ctx := context.Background()
	k := db.datastoreKey(id)
	ami := &Amiibo{}
	if err := db.client.Get(ctx, k, ami); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Book: %v", err)
	}
	ami.ID = id
	fmt.Printf("key: %v\n", ami)
	return ami, nil
}

func (db *datastoreDB) datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey("amiibo", id, nil)
}
*/
