package esidxmapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strings"

	es7 "github.com/elastic/go-elasticsearch/v7"
	pkg_err "github.com/pkg/errors"
)

var geoShapeMappingBody []byte

var (
	errInvalidRequest = errors.New("invalid request")
	errInternal       = errors.New("internal err")
)

func init() {
	var geoShapeMappingSrc geoShapeMapping
	geoShapeMappingSrc.Properties.Location.Type = "geo_shape"

	b, _ := json.Marshal(geoShapeMappingSrc)
	geoShapeMappingBody = b
}

type geoShapeMapping struct {
	Properties struct {
		Location struct {
			Type string `json:"type,omitempty"`
		} `json:"location,omitempty"`
	} `json:"properties,omitempty"`
}

func SetGeoShape(es *es7.Client, idx string) error {
	res, err := es.Indices.PutMapping(bytes.NewReader(geoShapeMappingBody), es.Indices.PutMapping.WithIndex(idx))
	if err != nil {
		log.Fatalf("PutMapping response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		switch {
		case strings.HasPrefix(res.Status(), "4"):
			return pkg_err.Wrapf(errInvalidRequest, "resBody: %s", res.String())
		case strings.HasPrefix(res.Status(), "5"):
			return pkg_err.Wrapf(errInternal, "resBody: %s", res.String())
		default:
		}
		return pkg_err.Wrapf(errors.New("invalid status."), "code: %s", res.Status())
	}
	return nil
}
