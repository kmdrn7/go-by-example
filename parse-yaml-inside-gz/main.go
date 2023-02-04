package parseyamlinsidegz

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var firstGzipBytes = []byte{0x1f, 0x8b, 0x08}

func Run() {

	// create empty object with type interface
	var promeConfig interface{}

	// read gziped prometheus yaml
	file, err := os.ReadFile("/home/a.ramadhan/workspace/rise/github/prometheus-config-merger/example/prometheus_operator.yaml.gz")
	if err != nil {
		log.Panic(err.Error())
	}

	// check if file is gzip file
	if bytes.Equal(file[0:3], firstGzipBytes) {
		fmt.Println("this is gzip file")
	} else {
		fmt.Println("this is not gzip file")
	}

	// create gzip io reader
	zr, err := gzip.NewReader(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer zr.Close()

	// read actual file inside gziped file
	file, err = io.ReadAll(zr)
	if err != nil {
		log.Fatal(err.Error())
	}

	// unmarshal bytes to interface object
	if err := yaml.Unmarshal(file, &promeConfig); err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(promeConfig)

}
