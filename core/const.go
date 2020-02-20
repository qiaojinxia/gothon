package core

type PyType = int

const (
	PY_NONE       = iota - 1 //NOne类型
	PY_INT                   //python int类型 64位
	PY_Float                 //python int类型
	PY_Bool                  //python int类型
	PY_COMPLEX               //python复数
	PY_STRING                //python string类型
	PY_LIST                  //python list类型
	PY_TUPLE                 //python 原表类型
	PY_SET                   //python set类型
	PY_DICTIONARY            //python 字典类型

)
