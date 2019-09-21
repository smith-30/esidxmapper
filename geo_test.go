package esidxmapper_test

import (
	"log"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/smith-30/esidxmapper"
)

var es7Factory = func() *es7.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

func TestSetGeoShape(t *testing.T) {
	type args struct {
		es  *es7.Client
		idx string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				es:  es7Factory(),
				idx: "target-idx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.args.es.Indices.Create(tt.args.idx)
			if err != nil {
				t.Fatalf("Indices.Create %s", err)
			}
			t.Logf("%v", res.String())
			if err := esidxmapper.SetGeoShape(tt.args.es, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("SetGeoShape() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
