package jsonx

import (
	stringutil "github.com/GokselKUCUKSAHIN/jsonx/internal/string-util"
	jsoniter "github.com/json-iterator/go"
)

func Unmarshal(buf []byte, val any) error {
	return jsoniter.Unmarshal(buf, val)
}

func UnmarshalString(buf string, val any) error {
	return jsoniter.Unmarshal(stringutil.Byte(buf), val)
}

func Marshal(val any) ([]byte, error) {
	return jsoniter.Marshal(val)
}

func MarshalString(val any) (string, error) {
	res, err := jsoniter.Marshal(val)
	if err != nil {
		return "", err
	}
	return stringutil.String(res), nil
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
