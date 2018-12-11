package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

const projectID = "Loadme!"

type datastoreDB struct {
	client *datastore.Client
}

// Amiibo Todo remove
type Amiibo struct {
	ID   int64
	name string
}

func notmain() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)

	if err != nil {
		return
	}

	dataStore, _ := newDatastoreDB(client)

	key, err := dataStore.AddAmiibo(&Amiibo{name: "Test"})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("key: %v\n", key)
	amii, err := dataStore.GetAmiibo(key)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("key: %v\n", amii)
}

func newDatastoreDB(client *datastore.Client) (*datastoreDB, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &datastoreDB{
		client: client,
	}, nil
}

func (db *datastoreDB) AddAmiibo(a *Amiibo) (id int64, err error) {
	fmt.Printf("Putiing amiibo: %v\n", a)
	ctx := context.Background()
	k := datastore.IncompleteKey("amiibo", nil)
	k, err = db.client.Put(ctx, k, a)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put amiibo: %v", err)
	}

	return k.ID, nil
}

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
