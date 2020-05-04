package parse

import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func ReadYamlFile(cfg interface{}, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	return err
}


func LoadConfig(cfg interface{}) error {
	cfgfile := flag.String("cfg", "", "config file")
	flag.Parse()
	err := ReadYamlFile(cfg, *cfgfile)
	if err != nil {
		log.Fatal("fail to parse the config file:", err)
		return err
	}
	return nil
}

//func loadConfig(fileName string) (Config, error) {
//	if err := config.Load(file.NewSource(
//		file.WithPath(fileName),
//	)); err != nil {
//		log.Fatal(err)
//		return Config{}, err
//	}
//	var cfg Config
//	if err := config.Get().Scan(&cfg); err != nil {
//		return Config{}, err
//	}
//
//	return cfg, nil
//}
