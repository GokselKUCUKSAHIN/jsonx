package jsonx

import "github.com/bytedance/sonic"

func Unmarshal(buf []byte, val any) error {
	return sonic.Unmarshal(buf, val)
}

func UnmarshalString(buf string, val any) error {
	return sonic.UnmarshalString(buf, val)
}

func Marshal(val any) ([]byte, error) {
	return sonic.Marshal(val)
}

func MarshalString(val any) (string, error) {
	return sonic.MarshalString(val)
}

func Cast[Model any](buf []byte, err error) (*Model, error) {
	if err != nil {
		return nil, err
	}
	var result Model
	if err = Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CastSlice[Model any](buff []byte, err error) ([]Model, error) {
	if err != nil {
		return nil, err
	}
	var result []Model
	if err = Unmarshal(buff, &result); err != nil {
		return nil, err
	}
	return result, nil
}
