package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zilliztech/milvus-backup/internal/proto/commonpb"
)

// kvPairToMap largely copied from internal/proxy/task.go#parseIndexParams
func KVPairToMap(m []*commonpb.KeyValuePair) (map[string]string, error) {
	params := make(map[string]string)
	for _, kv := range m {
		if kv.Key == "params" {
			params, err := parseParamsMap(kv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range params {
				params[k] = v
			}
		} else {
			params[kv.Key] = kv.Value
		}
	}
	return params, nil
}

// parseParamsMap parse the jsonic index parameters to map
func parseParamsMap(mStr string) (map[string]string, error) {
	buffer := make(map[string]interface{})
	err := json.Unmarshal([]byte(mStr), &buffer)
	if err != nil {
		return nil, errors.New("Unmarshal params failed")
	}
	ret := make(map[string]string)
	for key, value := range buffer {
		valueStr := fmt.Sprintf("%v", value)
		ret[key] = valueStr
	}
	return ret, nil
}

func MapToKVPair(dict map[string]string) []*commonpb.KeyValuePair {
	kvs := make([]*commonpb.KeyValuePair, 0)

	for key, value := range dict {
		kvs = append(kvs, &commonpb.KeyValuePair{
			Key:   key,
			Value: value,
		})
	}
	return kvs
}

// KvPairsMap converts common.KeyValuePair slices into map
func KvPairsMap(kvps []*commonpb.KeyValuePair) map[string]string {
	m := make(map[string]string)
	for _, kvp := range kvps {
		m[kvp.Key] = kvp.Value
	}
	return m
}
