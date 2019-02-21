package config

type TemplateValue struct {
	KeyList   []string `toml:"keyList"`
	ValueList []string `toml:"valueList"`
}

func (t *TemplateValue) FormatConfig() {

}
