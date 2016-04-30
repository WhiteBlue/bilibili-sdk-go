package lib

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type JSON struct {
	data interface{}
}

//Build JSON from bytes
func ParseJSON(data []byte) (*JSON, error) {
	j := new(JSON)
	if err := json.Unmarshal(data, &j.data); err == nil {
		return j, nil
	}else {
		return nil, err
	}
}


//Get JSON value
func (this *JSON) Get(key string) *JSON {
	if m, ok := (this.data).(map[string]interface{}); ok {
		if val, ok := m[key]; ok {
			return &JSON{val}
		}
	}
	return &JSON{nil}
}

func (this *JSON) String() (string, bool) {
	if val, ok := this.data.(string); ok {
		return val, true
	}
	return "", false
}

func (this *JSON) Int() (int, bool) {
	switch this.data.(type) {
	case string:
		if val, err := strconv.Atoi(reflect.ValueOf(this.data).String()); err == nil {
			return val, true
		}
	case float32, float64:
		return int(reflect.ValueOf(this.data).Float()), true
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(this.data).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(this.data).Uint()), true
	}

	return 0, false
}

func (this *JSON) Float32() (float32, bool) {
	if val, ok := this.data.(float32); ok {
		return val, true
	}
	return 0, false
}

func (this *JSON) Bool() (bool, bool) {
	if val, ok := this.data.(bool); ok {
		return val, true
	}
	return false, false
}

func (this *JSON) Map() (map[string]interface{}, bool) {
	if val, ok := this.data.(map[string]interface{}); ok {
		return val, true
	}
	return nil, false
}

func (this *JSON) Array() ([]interface{}, bool) {
	if val, ok := this.data.([]interface{}); ok {
		return val, true
	}
	return nil, false
}

func (this *JSON) JSONMap() (map[string]*JSON, bool) {
	if val, ok := this.data.(map[string]interface{}); ok {
		rMap := make(map[string]*JSON)
		for k, v := range val {
			rMap[k] = &JSON{v}
		}
		return rMap, true
	}
	return nil, false
}

func (this *JSON) JSONArray() ([]*JSON, bool) {
	if val, ok := this.data.([]interface{}); ok {
		jArray := make([]*JSON, len(val))
		for i, v := range val {
			jArray[i] = &JSON{v}
		}
		return jArray, true
	}
	return nil, false
}

func (this *JSON) ToString() string {
	if bytes, err := json.Marshal(this.data); err == nil {
		return string(bytes)
	}
	return ""
}