package reflect

import (
	"reflect"
	"strings"
)

type FieldTag struct {
	Type    string
	Name    string
	Options []string
}

func IsNilOrEmpty(value interface{}) bool {
	typeName := reflect.ValueOf(value).Kind().String()

	switch typeName {
	case "string":
		return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
	case "struct":
		return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
	case "invalid":
		return true
	default:
		return false
	}
}

func GetFieldTag(field reflect.StructField, key string) *FieldTag {
	fieldTag := string(field.Tag)
	tags := strings.Split(fieldTag, " ")
	for _, tag := range tags {
		tagDetails := strings.Split(tag, ":")
		if len(tagDetails) == 2 {
			tagKey := tagDetails[0]
			if strings.EqualFold(tagKey, key) {
				result := FieldTag{
					Type: tagKey,
				}
				tagValue := tagDetails[1]
				tagValueDetails := strings.Split(tagValue, ",")
				if len(tagValueDetails) == 1 {
					result.Name = tagValueDetails[0]
					result.Options = []string{}
				}
				if len(tagDetails) >= 2 {
					result.Name = tagValueDetails[0]
					result.Options = tagValueDetails[1:]
				}
				result.Name = strings.TrimLeft(result.Name, "\"")
				result.Name = strings.TrimRight(result.Name, "\"")
				for i := 0; i < len(result.Options); i++ {
					result.Options[i] = strings.TrimLeft(result.Options[i], "\"")
					result.Options[i] = strings.TrimRight(result.Options[i], "\"")
				}
				return &result
			}
		}
	}

	return nil
}

func RemoveField(obj interface{}, fields ...string) map[string]interface{} {
	result := make(map[string]interface{})
	rt, rv := reflect.TypeOf(obj), reflect.ValueOf(obj)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		toRemove := false
		for e := 0; e < len(fields); e++ {
			if strings.EqualFold(field.Name, fields[e]) {
				toRemove = true
				break
			}
		}

		if !toRemove {
			switch rv.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				result[field.Name] = rv.Field(i).Int()
			case reflect.String:
				result[field.Name] = rv.Field(i).String()
			case reflect.Bool:
				result[field.Name] = rv.Field(i).Bool()
			case reflect.Struct:
				result[field.Name] = rv.Field(i)
			}
		}
	}

	return result
}
