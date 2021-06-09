package goconfigenv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var (
	ErrNotImplemented error = fmt.Errorf("not implemented")
)

func Load(config interface{}) {
	configRef := reflect.ValueOf(config).Elem()
	if configRef.Kind() != reflect.Struct {
		panic("config must be struct type")
	}

	numOfFields := configRef.NumField()

	for i := 0; i < numOfFields; i++ {
		currFieldRef := configRef.Field(i)
		envKey, ok := configRef.Type().Field(i).Tag.Lookup("configenv")
		if !ok || len(envKey) < 1 {
			panic(fmt.Sprintf("config field %s does not have `configenv` tag", configRef.Type().Field(i).Name))
		}
		envValue, err := getEnvVarValue(envKey)
		if err != nil {
			panic(err)
		}
		if err := assignvalue(currFieldRef, envValue); err != nil {
			panic(err)
		}
	}
}

// Helper Functions

func getEnvVarValue(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable `%s` does not exist", key)
	}
	return value, nil
}

func assignvalue(fieldRef reflect.Value, value string) error {
	switch fieldRef.Kind() {
	case reflect.String:
		fieldRef.SetString(value)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		fieldRef.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 0, fieldRef.Type().Bits())
		if err != nil {
			return err
		}
		fieldRef.SetInt(i)
	default:
		return ErrNotImplemented
	}

	return nil
}
