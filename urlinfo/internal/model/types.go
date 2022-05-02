package model

//go:generate goctl model mongo -c -t Url
import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Url struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Hostnameport  string        `bson:"hostnameport"`
	Queryparamter string        `bson:"queryparamter"`
	UpdateDate    time.Time     `bson:"updatedate,omitempty"`
}

type DBQueryString struct {
	Hostnameport  string
	Queryparamter string
}
