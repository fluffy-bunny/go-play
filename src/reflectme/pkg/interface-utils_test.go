package pkg

import (
	"fmt"
	"reflect"
	"testing"
)

func printMethods(rt reflect.Type) {

	offset := 0

	prntFuncs := func() {
		methods := GetMethods(rt)

		for _, method := range methods {
			fmt.Printf("func %s(", method.Name)
			for i := offset; i < method.Type.NumIn(); i++ {
				fmt.Print(method.Type.In(i))
				if i < method.Type.NumIn()-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println(")")

		}
	}
	switch rt.Kind() {
	case reflect.Ptr:
		offset = 1
		prntFuncs()

	case reflect.Interface:
		prntFuncs()
	}
}

func TestHello(t *testing.T) {
	u, _ := NewReflectObjectInspect(RT_ISomething)
	fmt.Println(RT_ISomething.Name())
	fmt.Println(u.GetMethods())

	u, _ = NewReflectObjectInspect(RT_something)
	fmt.Println(RT_ISomething.Name())
	fmt.Println(u.GetMethods())

	u, _ = NewReflectObjectInspect(RT_other)
	fmt.Println(RT_ISomething.Name())
	fmt.Println(u.GetMethods())

}
