// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qot

import (
	"reflect"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	testutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
	"testing"
)

// qotTestClient starts a mock server that answers protoID with a cloned empty
// typed response (the mock server fills required fields), and returns a client
// connected to it.
func qotTestClient(t *testing.T, protoID uint32, resp proto.Message) *futuapi.Client {
	t.Helper()
	srv := testutil.NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(srv.Stop)
	srv.RegisterHandler(protoID, func([]byte) (proto.Message, error) {
		return proto.Clone(resp), nil
	})
	cli, cleanup := testutil.NewTestClient(t, srv)
	t.Cleanup(cleanup)
	return cli
}

// fillNonZero recursively sets scalar fields, nil pointer fields and empty
// slice fields on a request. Scalars receive a non-zero sentinel so range and
// "must be set" validations, and proto2 required fields, all pass. It works for
// both generated protobuf requests (pointer scalars) and hand-written wrapper
// request structs (value scalars).
func fillNonZero(v any) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !f.CanSet() {
			continue
		}
		switch f.Kind() {
		case reflect.Ptr:
			if f.IsNil() {
				f.Set(reflect.New(f.Type().Elem()))
			}
			if f.Elem().Kind() == reflect.Struct {
				fillNonZero(f.Interface())
			} else {
				setScalar(f.Elem())
			}
		case reflect.Slice:
			if f.Len() == 0 {
				el := reflect.New(f.Type().Elem()).Elem()
				f.Set(reflect.Append(f, el))
			}
			for j := 0; j < f.Len(); j++ {
				e := f.Index(j)
				switch {
				case e.Kind() == reflect.Ptr:
					if e.IsNil() {
						e.Set(reflect.New(e.Type().Elem()))
					}
					if e.Elem().Kind() == reflect.Struct {
						fillNonZero(e.Interface())
					}
				case e.Kind() == reflect.Struct && e.CanAddr():
					fillNonZero(e.Addr().Interface())
				}
			}
		case reflect.Struct:
			if f.CanAddr() {
				fillNonZero(f.Addr().Interface())
			}
		default:
			setScalar(f)
		}
	}
}

func setScalar(f reflect.Value) {
	if !f.CanSet() {
		return
	}
	switch f.Kind() {
	case reflect.String:
		f.SetString("x")
	case reflect.Bool:
		f.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if f.Int() == 0 {
			f.SetInt(1)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if f.Uint() == 0 {
			f.SetUint(1)
		}
	case reflect.Float32, reflect.Float64:
		if f.Float() == 0 {
			f.SetFloat(1)
		}
	}
}
