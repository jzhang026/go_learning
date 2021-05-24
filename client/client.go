package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)




func main() {
	h := &ComplexHello{}

	PrepareComplexRPC(h)
	if resp, err := h.SayHello(&Input{Msg: "Adrian"}); err !=nil {
		log.Fatal(err)
	} else {
		fmt.Println(resp.Msg)
	}
}


type ComplexHello struct {
	SayHello func (in *Input) (*Output, error)
}

func (c ComplexHello) ServiceName() string {
	return "hello"
}

type Input struct {
	Msg string
}

type Output struct {
	Msg string
}

func PrepareComplexRPC (val interface{}) {
	value := reflect.ValueOf(val)
	ele := value.Elem()
	t := ele.Type()
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := ele.Field(i)

		if fieldValue.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				input := args[0].Interface()
				output := reflect.New(fieldType.Type.Out(0).Elem()).Interface()
				inData, err := json.Marshal(input)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
				}

				client := http.Client{}
				serviceName := val.(Service).ServiceName()
				resp, err := client.Post(CfgMap[serviceName].Endpoint, "application/json",
					bytes.NewReader(inData))

				if err != nil {
					return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
				}

				err = json.Unmarshal(data, output)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
				}

				return []reflect.Value{reflect.ValueOf(output), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fieldValue.Set(reflect.MakeFunc(fieldType.Type,fn))
		}
	}
}

