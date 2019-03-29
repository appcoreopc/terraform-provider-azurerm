package capi

import (
	"github.com/Azure/go-autorest/autorest"
)

type BaseResource struct {
	autorest.Response `json:"-"`

	ID *string `json:"id,omitempty"`

	Self *string `json:"_self,omitempty"`
	ETag *string `json:"_etag,omitempty"`
	RID  *string `json:"_rid,omitempty"`
	TS   *int    `json:"_ts,omitempty"`

	Path string `json:"-"` //path to object `dbs/{databaseName}/colls{collectionName}` ect
}

func (e BaseResource) PopulateCommon(r autorest.Response) {
	e.Response = r

}