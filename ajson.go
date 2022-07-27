package ajson

func (t *Result) Set(key string, value interface{}) {
	raw := t.Raw

	s, err := Set(raw, key, value)
	if err != nil {
		return
	}

	t.Raw = s
}

func (t *Result) Create(key string, value interface{}) {
	raw := t.Raw

	s, err := SetBytes([]byte(raw), key, value)
	if err != nil {
		return
	}

	t.Raw = string(s)
}

func (t *Result) Delete(key string) {
	raw := t.Raw

	s, err := Delete(raw, key)
	if err != nil {
		return
	}

	t.Raw = s
}
