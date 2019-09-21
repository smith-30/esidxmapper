package esidxmapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	es7 "github.com/elastic/go-elasticsearch/v7"
)

var geoShapeMappingBody []byte

func init() {
	var geoShapeMappingSrc geoShapeMapping
	geoShapeMappingSrc.Properties.Location.Type = "geo_shape"

	b, _ := json.Marshal(geoShapeMappingSrc)
	geoShapeMappingBody = b
}

type geoShapeMapping struct {
	// Mappings struct {
	Properties struct {
		Location struct {
			Type string `json:"type,omitempty"`
		} `json:"location,omitempty"`
	} `json:"properties,omitempty"`
	// } `json:"mappings,omitempty"`
}

func SetGeoShape(es *es7.Client, idx string) error {
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Printf("%v\n", res.String())

	fmt.Printf("%v\n", string(geoShapeMappingBody))
	res, err = es.Indices.PutMapping(bytes.NewReader(geoShapeMappingBody), es.Indices.PutMapping.WithIndex(idx))
	if err != nil {
		log.Fatalf("PutMapping response: %s", err)
	}
	fmt.Printf("%v\n", res.String())
	return nil
}
