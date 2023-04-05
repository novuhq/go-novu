package lib

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

func MustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return u
}

type QueryParam struct {
	Key   string
	Value string
}

func GenerateQueryParamsFromStruct[T interface{}](queryParamsStruct T) ([]QueryParam, error) {
	var queryParamList []QueryParam = make([]QueryParam, 0)

	values := reflect.ValueOf(queryParamsStruct)
	if values.Kind() != reflect.Struct {
		return nil, errors.New("queryParamsStruct must be a struct. Make sure you are not passing a pointer to struct")
	}

	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		if !field.IsZero() {
			// Check if the field has a queryKey tag and use that as the key for the query param
			key := types.Field(i).Tag.Get("queryKey")
			if key == "" {
				key = types.Field(i).Name
			}

			// Check if the field is a string, bool or int and convert it to string
			var value string
			switch field.Kind() {
			case reflect.String:
				value = field.Interface().(string)

			case reflect.Bool:
				value = strconv.FormatBool(field.Interface().(bool))

			case reflect.Int:
				value = strconv.FormatInt(int64(field.Interface().(int)), 10)

			default:
				return nil, errors.New("unsupported type in struct field. Supported types are: string, bool and int")
			}

			queryParamList = append(queryParamList, QueryParam{
				Key:   key,
				Value: value,
			})
		}
	}

	return queryParamList, nil
}
