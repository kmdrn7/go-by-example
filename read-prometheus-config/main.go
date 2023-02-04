package readprometheusconfig

import (
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/prometheus/prometheus/config"
	_ "github.com/prometheus/prometheus/discovery/install"
)

func main() {

	logger := log.NewNopLogger()

	// finalPromeConfig := &config.Config{}
	// fmt.Println(finalPromeConfig)

	filename := "prometheus.yaml"

	fmt.Println("read content")
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	// cfg := &config.Config{}
	// *cfg = config.DefaultConfig

	cfg, err := config.Load(string(content), false, logger)
	if err != nil {
		fmt.Println(err.Error())
	}
	// cfg.SetDirectory(filepath.Dir(filename))
	// fmt.Println("set default config struct")

	fmt.Println(cfg)

	// _, err = config.LoadFile("prometheus.yaml", false, false, logger)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(thethe)

	// if finalPromeConfig, err := config.LoadFile(cfg.configFile, agentMode, false, log.NewNopLogger()); err != nil {
	// 	absPath, pathErr := filepath.Abs(cfg.configFile)
	// 	if pathErr != nil {
	// 		absPath = cfg.configFile
	// 	}
	// 	level.Error(logger).Log("msg", fmt.Sprintf("Error loading config (--config.file=%s)", cfg.configFile), "file", absPath, "err", err)
	// 	os.Exit(2)
	// }
}

// func yamlEncode(v interface{}) (*bytes.Buffer, error) {

// 	b := &bytes.Buffer{}
// 	yamlEncoder := yaml.NewEncoder(b)
// 	yamlEncoder.SetIndent(2)

// 	if err := yamlEncoder.Encode(v); err != nil {
// 		return nil, err
// 	}

// 	return b, nil
// }
