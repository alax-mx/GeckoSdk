package baseutils

import (
	"fmt"
	"reflect"
	"strconv"
)

// ShowObjectValue 打印对象完整信息
func ShowObjectValue(object interface{}) {
	dataType := reflect.TypeOf(object)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}

	showStructValue(object, dataType.Name())
}

// ShowStructValue ShowStructValue
func showStructValue(object interface{}, tag string) {
	dataValue := reflect.ValueOf(object)
	dataType := reflect.TypeOf(object)

	if dataType.Kind() == reflect.Ptr {
		if dataValue.IsNil() {
			fmt.Println(tag, " = nil")
			return
		}

		dataValue = dataValue.Elem()
		dataType = dataType.Elem()
	}

	num := dataType.NumField()
	for i := 0; i < num; i++ {
		field := dataType.Field(i)
		fieldName := field.Name
		fieldValue := dataValue.FieldByName(fieldName)

		if fieldName == "XXX_unrecognized" ||
			fieldName == "XXX_sizecache" {
			continue
		}

		showFieldValue(fieldValue, tag+"."+fieldName)
	}
}

// ShowFieldValue 打印某个属性的值
func showFieldValue(fieldValue reflect.Value, tag string) {
	realValue := fieldValue
	fieldType := fieldValue.Type()

	// 判断是否有效
	if !fieldValue.IsValid() {
		fmt.Println(tag, "= nil")
		return
	}

	// 如果是指针
	if fieldType.Kind() == reflect.Ptr {
		realValue = fieldValue.Elem()
		fieldType = fieldType.Elem()
	}

	// 结构数组
	if fieldType.Kind() == reflect.Slice {
		for i := 0; i < fieldValue.Len(); i++ {
			subValue := fieldValue.Index(i)
			subType := subValue.Type()
			if subType.Kind() == reflect.Ptr {
				subType = subType.Elem()
				subValue = subValue.Elem()
			}

			if subType.Kind() == reflect.Struct {
				showFieldValue(subValue, tag+"["+strconv.Itoa(i)+"]")
			} else {
				fmt.Println(tag+" =", realValue.Interface())
				break
			}
		}

		if fieldValue.Len() == 0 {
			fmt.Println(tag + " = empty")
		}
		return
	}

	// 判断是否是结构体
	if fieldType.Kind() == reflect.Struct {
		showStructValue(fieldValue.Interface(), tag)
		return
	}

	// 字符串
	if fieldType.Kind() == reflect.String {
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				fmt.Println(tag + " = nil")
			} else {
				fmt.Println(tag+" =", realValue.Interface())
			}
		} else {
			if fieldValue.Len() > 0 {
				fmt.Println(tag+" =", realValue)
			} else {
				fmt.Println(tag + " = nil")
			}
		}
		return
	}

	if !realValue.IsValid() {
		fmt.Println(tag + " = nil")
	} else {
		fmt.Println(tag+" =", realValue)
	}
}

//PrintBytes PrintBytes
func PrintBytes(data []byte, tag string) {
	fmt.Println("-----------", tag, "------------")
	length := len(data)
	for i := 0; i < length; i++ {
		fmt.Print(data[i], ", ")
		if i%16 == 0 && i > 0 {
			fmt.Println()
		}
	}
	fmt.Printf("\n")
}

// PrintBytesHex PrintBytes
func PrintBytesHex(data []byte, tag string) {
	length := len(data)
	fmt.Println("-----------", tag, "-", length, "------------")
	for i := 0; i < length; i++ {
		if i%16 == 0 && i > 0 {
			fmt.Println()
		}
		fmt.Printf("0x%02x, ", data[i])
	}
	fmt.Printf("\n")
}
