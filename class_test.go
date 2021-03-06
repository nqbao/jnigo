package jnigo

import (
	"strings"
	"testing"
)

func TestJClass(t *testing.T) {
	jvm := CreateJVM()

	val, err := jvm.newJPrimitive(int32(1))
	if err != nil {
		t.Fatal(err)
	}

	testArray := [][]interface{}{
		[]interface{}{"java/lang/String", []JObject{}, ""},
		[]interface{}{"TestSubClass", []JObject{}, "TestSubClass"},
		[]interface{}{"TestClass", []JObject{val}, "TestClass"},
	}

	for _, test := range testArray {
		value, err := jvm.NewJClass(test[0].(string), test[1].([]JObject))
		if err != nil {
			t.Fatal(err)
		}

		if value.Signature() != "L"+test[0].(string)+";" {
			t.Fatal(value.GoValue())
		}

		if strings.HasPrefix(value.String(), test[2].(string)) == false {
			t.Fatal(value.String())
		}
	}
}

func TestJClassMethod(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"mvboolean", "()Z"},
		[]string{"mvbyte", "()B"},
		[]string{"mvchar", "()C"},
		[]string{"mvshort", "()S"},
		[]string{"mvint", "()I"},
		[]string{"mvlong", "()J"},
		[]string{"mvfloat", "()F"},
		[]string{"mvdouble", "()D"},
		[]string{"mvclass", "()LTestSubClass;"},

		[]string{"maboolean", "()[Z"},
		[]string{"mabyte", "()[B"},
		[]string{"machar", "()[C"},
		[]string{"mashort", "()[S"},
		[]string{"maint", "()[I"},
		[]string{"malong", "()[J"},
		[]string{"mafloat", "()[F"},
		[]string{"madouble", "()[D"},
		[]string{"maclass", "()[LTestSubClass;"},
	}

	value, err := jvm.NewJClass(clazz, []JObject{})
	if err != nil {
		t.Fatal(err)
	}

	if value.Signature() != "L"+clazz+";" {
		t.Fatal(value.GoValue())
	}

	for _, test := range testArray {
		_, err := value.CallFunction(test[0], test[1], []JObject{})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestJClassStaticMethod(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"smvboolean", "()Z"},
		[]string{"smvbyte", "()B"},
		[]string{"smvchar", "()C"},
		[]string{"smvshort", "()S"},
		[]string{"smvint", "()I"},
		[]string{"smvlong", "()J"},
		[]string{"smvfloat", "()F"},
		[]string{"smvdouble", "()D"},
		[]string{"smvclass", "()LTestSubClass;"},

		[]string{"smaboolean", "()[Z"},
		[]string{"smabyte", "()[B"},
		[]string{"smachar", "()[C"},
		[]string{"smashort", "()[S"},
		[]string{"smaint", "()[I"},
		[]string{"smalong", "()[J"},
		[]string{"smafloat", "()[F"},
		[]string{"smadouble", "()[D"},
		[]string{"smaclass", "()[LTestSubClass;"},
	}

	for _, test := range testArray {
		_, err := jvm.CallStaticFunction(clazz, test[0], test[1], []JObject{})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestJClassGetField(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"vboolean", "Z"},
		[]string{"vbyte", "B"},
		[]string{"vchar", "C"},
		[]string{"vshort", "S"},
		[]string{"vint", "I"},
		[]string{"vlong", "J"},
		[]string{"vfloat", "F"},
		[]string{"vdouble", "D"},
		[]string{"vclass", "LTestSubClass;"},

		[]string{"aboolean", "[Z"},
		[]string{"abyte", "[B"},
		[]string{"achar", "[C"},
		[]string{"ashort", "[S"},
		[]string{"aint", "[I"},
		[]string{"along", "[J"},
		[]string{"afloat", "[F"},
		[]string{"adouble", "[D"},
		[]string{"aclass", "[LTestSubClass;"},
	}

	value, err := jvm.NewJClass(clazz, []JObject{})
	if err != nil {
		t.Fatal(err)
	}

	if value.Signature() != "L"+clazz+";" {
		t.Fatal(value.GoValue())
	}

	for _, test := range testArray {
		_, err := value.GetField(test[0], test[1])
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestJClassGetStaticField(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"svboolean", "Z"},
		[]string{"svbyte", "B"},
		[]string{"svchar", "C"},
		[]string{"svshort", "S"},
		[]string{"svint", "I"},
		[]string{"svlong", "J"},
		[]string{"svfloat", "F"},
		[]string{"svdouble", "D"},
		[]string{"svclass", "LTestSubClass;"},
		[]string{"saboolean", "[Z"},
		[]string{"sabyte", "[B"},
		[]string{"sachar", "[C"},
		[]string{"sashort", "[S"},
		[]string{"saint", "[I"},
		[]string{"salong", "[J"},
		[]string{"safloat", "[F"},
		[]string{"sadouble", "[D"},
		[]string{"saclass", "[LTestSubClass;"},
	}

	for _, test := range testArray {
		t.Run(test[0], func(t *testing.T) {
			_, err := jvm.GetStaticField(clazz, test[0], test[1])
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestJClassSetField(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestSubClass"
	gobool, _ := jvm.newJPrimitive(false)
	gobyte, _ := jvm.newJPrimitive(byte(20))
	gochar, _ := jvm.newJPrimitive(uint16(20))
	goshort, _ := jvm.newJPrimitive(int16(20))
	goint, _ := jvm.newJPrimitive(int32(20))
	golong, _ := jvm.newJPrimitive(int64(20))
	gofloat, _ := jvm.newJPrimitive(float32(20))
	godouble, _ := jvm.newJPrimitive(float64(20))
	gojclass, _ := jvm.NewJClass("TestSubClass", []JObject{})

	goabool, _ := jvm.newJArray([]bool{true, false})
	goabyte, _ := jvm.newJArray([]byte{100, 100})
	goachar, _ := jvm.newJArray([]uint16{10000, 10000})
	goashort, _ := jvm.newJArray([]int16{10000, 10000})
	goaint, _ := jvm.newJArray([]int32{10000, 10000})
	goalong, _ := jvm.newJArray([]int64{10000, 10000})
	goafloat, _ := jvm.newJArray([]float32{1000.0, 1000.0})
	goadouble, _ := jvm.newJArray([]float64{1000.0, 1000.0})
	//goajclass, _ := jvm.newJArray([]*JClass{gojclass, gojclass})

	testArray := [][]interface{}{
		[]interface{}{"vboolean", gobool},
		[]interface{}{"vbyte", gobyte},
		[]interface{}{"vchar", gochar},
		[]interface{}{"vshort", goshort},
		[]interface{}{"vint", goint},
		[]interface{}{"vlong", golong},
		[]interface{}{"vfloat", gofloat},
		[]interface{}{"vdouble", godouble},
		[]interface{}{"vclass", gojclass},

		[]interface{}{"aboolean", goabool},
		[]interface{}{"abyte", goabyte},
		[]interface{}{"achar", goachar},
		[]interface{}{"ashort", goashort},
		[]interface{}{"aint", goaint},
		[]interface{}{"along", goalong},
		[]interface{}{"afloat", goafloat},
		[]interface{}{"adouble", goadouble},
		//[]interface{}{"aclass", goajclass},
	}

	value, err := jvm.NewJClass(clazz, []JObject{})
	if err != nil {
		t.Fatal(err)
	}

	if value.Signature() != "L"+clazz+";" {
		t.Fatal(value.GoValue())
	}

	for _, test := range testArray {
		err := value.SetField(test[0].(string), test[1].(JObject))
		if err != nil {
			t.Fatal(err)
		}
	}
}
