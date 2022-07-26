package alljson

func (t *Result) Set(key string, value interface{}) {
	raw := t.Raw

	s, err := Set(raw, key, value)
	if err != nil {
		return
	}

	t.Raw = s
}
