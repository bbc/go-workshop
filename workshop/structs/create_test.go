package structs

// func toMap(bundle Bundle) map[string]string {
// 	ele := reflect.ValueOf(&bundle).Elem()
// 	actual := map[string]string{}
// 	typeT := ele.Type()

// 	for i := 0; i < ele.NumField(); i++ {
// 		f := ele.Field(i)
// 		name := typeT.Field(i).Name
// 		nameType := fmt.Sprintf("%s", f.Type())
// 		actual[name] = nameType
// 	}

// 	return actual
// }

// func TestCreateBundle(t *testing.T) {
// 	expected := map[string]string{
// 		"ID":              "string",
// 		"ShouldBeRemoved": "bool",
// 		"IsExpected":      "bool",
// 		"Score":           "int",
// 		"Error":           "error",
// 	}

// 	bundle := CreateBundle()
// 	actual := toMap(bundle)
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Fatalf("Expected %v. Got %v", expected, actual)
// 	}
// }
