package reader

import (
	"context"
	"github.com/whosonfirst/go-reader"
	"io"
	"io/ioutil"
	"testing"
)

func TestHTTPReader(t *testing.T) {

	reader_uri := "https://data.whosonfirst.org"
	file_uri := "101/736/545/101736545.geojson"

	ctx := context.Background()

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatal(err)
	}

	fh, err := r.Read(ctx, file_uri)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatal(err)
	}

}
