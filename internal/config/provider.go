package config

func NewProvider() (*Config, error) {
	cfgReader := NewReader()

	err := cfgReader.Read()
	if err != nil {
		return nil, err
	}

	return cfgReader.Get(), nil
}
