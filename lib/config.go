package lib

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"sort"
)

type WafConfig struct {
	Upstream       string
	ListenAddress  string
	IpFilterMode   string
	IpAddresses    []string
	DenyExtensions []string
}

func LoadConfig(filename string) WafConfig {
	config := WafConfig{}
	contents, _ := ioutil.ReadFile(filename)
	err := toml.Unmarshal(contents, &config)
	if err != nil {
		panic(err)
	}

	// Sort for faster searching.
	sort.Strings(config.IpAddresses)

	return config
}
