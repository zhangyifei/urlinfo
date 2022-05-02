package model

import (
	"context"
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	cachec "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mongoc"
)

var prefixUrlCacheKey = "cache:Url:"

type UrlModel interface {
	Insert(ctx context.Context, data *Url) error
	BatchInsert(ctx context.Context, data []*Url) error
	FindOne(ctx context.Context, query DBQueryString) (*Url, error)
	Upsert(ctx context.Context, data *Url) (*mgo.ChangeInfo, error)
	Delete(ctx context.Context, id string) error
}

type defaultUrlModel struct {
	*mongoc.Model
}

func NewUrlModel(url, collection string, c cachec.CacheConf) UrlModel {
	return &defaultUrlModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultUrlModel) Insert(ctx context.Context, data *Url) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultUrlModel) BatchInsert(ctx context.Context, data []*Url) error {

	dataInterface := make([]interface{}, len(data))
	for i := range data {

		if !data[i].ID.Valid() {
			data[i].ID = bson.NewObjectId()
		}

		dataInterface[i] = data[i]
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(dataInterface...)
}

func (m *defaultUrlModel) FindOne(ctx context.Context, query DBQueryString) (*Url, error) {

	fmt.Println(bson.M{"hostnameport": query.Hostnameport})

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data Url
	key := prefixUrlCacheKey + query.Hostnameport + query.Queryparamter
	err = m.GetCollection(session).FindOne(&data, key, bson.M{"hostnameport": query.Hostnameport, "queryparamter": query.Queryparamter})
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUrlModel) Upsert(ctx context.Context, data *Url) (*mgo.ChangeInfo, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)
	key := prefixUrlCacheKey + data.Hostnameport + data.Queryparamter
	return m.GetCollection(session).Upsert(bson.M{"hostnameport": data.Hostnameport, "queryparamter": data.Queryparamter}, data, key)
}

func (m *defaultUrlModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixUrlCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}
