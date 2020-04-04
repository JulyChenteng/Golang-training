package slice

import (
	"reflect"
	"unsafe"
)

type T interface{}

type MySlice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func (slice *MySlice) Set(index int, val interface{}) {
	if index < 0 || index >= ILen(*slice) {
		panic("input error : index")
	}

	arrTyp, arr := slice.getArray()

	if arrTyp.Elem() == reflect.TypeOf(val) {
		arr.Index(index).Set(reflect.ValueOf(val))
	}
}

func (slice *MySlice) Get(index int) interface{} {
	if index < 0 || index >= ILen(*slice) {
		panic("input error : index")
	}

	_, arr := slice.getArray()

	return arr.Index(index).Interface()
}

func (slice *MySlice) getArray() (reflect.Type, reflect.Value) {
	v := *(*reflect.Value)(slice.array)
	typ := v.Type().Elem()
	arr := v.Elem()

	return typ, arr
}

func IMake(et T, params ...int) MySlice {
	if len(params) > 2 || len(params) < 1 {
		panic("Params is not enough!")
	}

	if len(params) == 1 {
		if params[0] <= 0 {
			panic("error input param: len and cap <= 0")
		}

		typ := reflect.ArrayOf(params[0], reflect.TypeOf(et).Elem())
		value := reflect.New(typ)

		return MySlice{unsafe.Pointer(&value), params[0], params[0]}
	} else {
		if params[0] > params[1] && params[2] > 0 {
			panic("error input param: len > cap or cap <= 0")
		}

		tp := reflect.ArrayOf(params[1], reflect.TypeOf(et).Elem())
		v := reflect.New(tp)

		//fmt.Println(v.Type(),v.Kind())  // v.kind()为reflect.Ptr, v.Type()为*[8]int

		return MySlice{unsafe.Pointer(&v), params[0], params[1]}
	}
}

func ILen(slice MySlice) int {
	return slice.len
}

func ICap(slice MySlice) int {
	return slice.cap
}

func IAppend(slice MySlice, elems ...T) MySlice {
	var sli = slice

	if sli.len+len(elems) > slice.cap {
		sli = growSlice(slice, len(elems))
	}

	arrType, array := slice.getArray()
	typ := arrType.Elem()
	arr := array.Elem()

	//fmt.Println(v.Type(),v.Kind())		//*[8]int   reflect.Ptr
	//fmt.Println(v.Type().Elem(), typ) 	//[8]int	int
	//fmt.Println(arr.Type())   			//[8]int

	len := sli.len

	//类型检测
	for _, value := range elems {
		if reflect.TypeOf(value) == typ {
			arr.Index(len).Set(reflect.ValueOf(value))
			len++
		} else {
			panic("type is incompatible!")
			break
		}
	}

	return MySlice{sli.array, len, sli.cap}
}

//len表示IAppend参数elems的长度
func growSlice(old MySlice, len int) MySlice {
	newCap := old.cap
	doubleCap := 2 * old.cap

	if old.cap+len > doubleCap {
		newCap = old.cap + len
	} else {
		if old.len < 1024 {
			newCap = doubleCap
		} else {
			for 0 < newCap && newCap < old.len+len {
				newCap += newCap / 4
			}

			//newCap溢出
			if newCap <= 0 {
				newCap = old.len + len
			}
		}
	}

	value := *(*reflect.Value)(old.array)
	arrTyp := value.Type().Elem()
	//fmt.Println(arrTyp)    //[8]int

	tp := reflect.ArrayOf(newCap, arrTyp.Elem())
	v := reflect.New(tp)

	return MySlice{unsafe.Pointer(&v), old.len, newCap}
}
