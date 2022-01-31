package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type (
	MethodSet  map[string]string
	ISomething interface {
		Something() string
		Something2(v int64)
		Something3(v int64, v2 int64) (string, error)
	}
	something struct {
		something string
	}
	other struct {
		Value int64
	}

	ReflectObjectInspect struct {
		Type         reflect.Type
		MethodOffset int
	}
)

func (s *MethodSet) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *MethodSet) Remove(key string) {
	delete(*s, key)
}

func (s *MethodSet) Clear() {
	for k := range *s {
		delete(*s, k)
	}
}

func NewReflectObjectInspect(rt reflect.Type) (*ReflectObjectInspect, error) {
	switch rt.Kind() {
	case reflect.Ptr:
		return &ReflectObjectInspect{
			Type:         rt,
			MethodOffset: 1,
		}, nil

	case reflect.Interface:
		return &ReflectObjectInspect{
			Type:         rt,
			MethodOffset: 0,
		}, nil

	}
	return nil, errors.New("only pointer or interface type is supported")
}

func (s *ReflectObjectInspect) GetMethods() MethodSet {
	methods := GetMethods(s.Type)

	sset := make(MethodSet)
	for _, method := range methods {
		builderFull := strings.Builder{}
		sPtr := ""
		switch s.Type.Kind() {
		case reflect.Ptr:
			sPtr = "*"
		}

		builderFull.WriteString(fmt.Sprintf("func (s %s%s) ", sPtr, s.Type.Name()))
		builder := strings.Builder{}
		builder.WriteString(fmt.Sprintf("%s(", method.Name))

		for i := s.MethodOffset; i < method.Type.NumIn(); i++ {
			builder.WriteString(fmt.Sprint(method.Type.In(i)))

			if i < method.Type.NumIn()-1 {
				builder.WriteString(", ")

			}
		}
		builder.WriteString(")")
		numOut := method.Type.NumOut()
		if numOut > 0 {
			if numOut > 1 {
				builder.WriteString(" (")
			} else {
				builder.WriteString(" ")
			}
			for i := 0; i < numOut; i++ {
				builder.WriteString(fmt.Sprint(method.Type.Out(i)))
				if i < numOut-1 {
					builder.WriteString(", ")
				}
			}
		}
		if numOut > 1 {
			builder.WriteString(")")
		}
		sFuncN := builder.String()
		builderFull.WriteString(sFuncN)
		sset[builder.String()] = builderFull.String()

	}
	return sset
}

func (s *something) Something() {
	s.something = "something"
}
func (s *something) Something2(v int64) {
	s.something = "something2"
}

func (s *other) Set(v int64) {
	s.Value = v
}

var RT_ISomething = reflect.TypeOf((*ISomething)(nil)).Elem()

var RT_something = reflect.TypeOf(&something{})
var RT_other = reflect.TypeOf(&other{})

func GetMethods(rt reflect.Type) []reflect.Method {
	var methods []reflect.Method
	for i := 0; i < rt.NumMethod(); i++ {
		method := rt.Method(i)
		if method.PkgPath != "" {
			continue
		}
		methods = append(methods, method)
	}
	return methods
}
