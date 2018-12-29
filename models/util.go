package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func InputRequest(r *http.Request, ptr interface{}) error {
	body := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		//VPNGatewayCommonLibs.LOG(ctx, "InputRequest").WithFields(log.Fields{"request": r.Body}).WithError(err).Debug("json decode error")
		return fmt.Errorf("")
	}
	fmt.Println(body)
	data := make(map[string]string)
	for k, _v := range body {
		switch v := _v.(type) {
		case string:
			data[k] = v
		case float64:
			data[k] = strconv.FormatFloat(v, 'f', -1, 64)
		}
	}
	fmt.Println(data)
	//VPNGatewayCommonLibs.LOG(ctx, "InputRequest").WithFields(log.Fields{"request": data}).Debug("begin")

	dataArr := make(map[string][]string)
	re, err := regexp.Compile(`[A-Za-z0-9]\.[0-9]`)
	fmt.Println(re)
	if err != nil {
		//VPNGatewayCommonLibs.LOG(ctx, "InputRequest").WithError(err).Debug("regext error")
		return fmt.Errorf("")
	}
	for k, v := range data {
		if re.MatchString(k) {
			_k := strings.Split(k, ".")[0]
			if _, ok := dataArr[_k]; !ok {
				dataArr[_k] = []string{v}
			} else {
				dataArr[_k] = append(dataArr[_k], v)
			}
		}
	}
	fmt.Println(dataArr)
	v := reflect.ValueOf(ptr).Elem()
	fmt.Println(v.NumField())
	for i, l := 0, v.NumField(); i < l; i++ {
		fieldTag := v.Type().Field(i).Tag
		fmt.Println(fieldTag)
		fieldValue := v.Field(i)
		fmt.Println(fieldValue.Kind())
		if !fieldValue.CanSet() {
			fmt.Println("oooo")
			continue
		}
		key := fieldTag.Get("key")
		defaultValue := fieldTag.Get("default")
		required := fieldTag.Get("required")
		value, ok := data[key]
		if !ok {
			value = defaultValue
			if required != "" {
				if _, _ok := dataArr[key]; !_ok {
					return fmt.Errorf("%s", key)
				}
			}
		}
		fmt.Println(ptr)
		if fieldValue.Kind() == reflect.Slice {
			if value, ok := dataArr[key]; ok {
				for _, _v := range value {
					elem := reflect.New(fieldValue.Type().Elem()).Elem()
					if err := populate(elem, _v); err != nil {
						//VPNGatewayCommonLibs.LOG(ctx, "InputRequest").WithError(err).Debug("slice populate error")
						return fmt.Errorf("%s", key)
					}
					fieldValue.Set(reflect.Append(fieldValue, elem))
				}
			}
		} else {
			if err := populate(fieldValue, value); err != nil {
				//VPNGatewayCommonLibs.LOG(ctx, "InputRequest").WithError(err).Debug("populate error")
				return fmt.Errorf("%s", key)
			}
		}
	}

	//VPNGatewayCommonLibs.LOG(ctx, "InputRequest").Debug("parsing success")
	return nil
}
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Uint32:
		i, _ := strconv.ParseUint(value, 10, 32)
		v.SetUint(i)

	case reflect.Int:
		i, _ := strconv.ParseInt(value, 10, 64)
		v.SetInt(i)

	case reflect.Bool:
		b, _ := strconv.ParseBool(value)
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}


func OutputResponse( w http.ResponseWriter, r interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, err := json.Marshal(r)
	if err != nil {
		return
	}

	w.Write(b)
	//VPNGatewayCommonLibs.LOG(ctx, "OutputResponse").WithFields(log.Fields{"response": r}).Debug("finish")

	return
}