// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	SSL    SSL           `config:"ssl"`
}

type SSL struct {
	CAfile string `config:"cafile"`
}

var DefaultConfig = Config{
	Period: 5 * time.Minute,
	SSL: SSL{
		CAfile: "",
	},
}
