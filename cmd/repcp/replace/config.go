package replace

import (
	"gopkg.in/yaml.v3"
)

// Config コマンドの設定ファイルを管理する構造体
type Config struct {
	Pkgs []*ReplaceString
}

// ReadConfig コマンドの設定YAMLを読み込む
func ReadConfig(config string) (*Config, error) {
	t := Config{}
	err := yaml.Unmarshal([]byte(config), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
