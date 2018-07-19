package objects

// func TestEmptySet(t *testing.T) {
// 	s := NewSet()
// 	values := s.Values()
// 	if values == nil {
// 		t.Error("Values is nil")
// 	}

// 	if len(values) != 0 {
// 		t.Error("Values should be empty")
// 	}
// }

// func TestAddItem(t *testing.T) {
// 	expected := []string{"test1", "test2"}
// 	s := NewSet()
// 	s.Add("test1")
// 	s.Add("test2")
// 	values := s.Values()
// 	if !reflect.DeepEqual(expected, values) {
// 		t.Errorf("Expected %v Got %v", expected, values)
// 	}
// }

// func TestRemove(t *testing.T) {
// 	expected := []string{"test1"}
// 	s := NewSet()
// 	s.Add("test1")
// 	s.Add("test2")
// 	s.Remove("test2")
// 	values := s.Values()
// 	if !reflect.DeepEqual(expected, values) {
// 		t.Errorf("Expected %v Got %v", expected, values)
// 	}
// }

// func TestHasPresent(t *testing.T) {
// 	v := "test1"
// 	s := NewSet()
// 	s.Add("test1")
// 	s.Add("test2")
// 	if !s.Has(v) {
// 		t.Errorf("Expected %s in set", v)
// 	}
// }

// func TestHasAbsent(t *testing.T) {
// 	v := "?"
// 	s := NewSet()
// 	s.Add("test1")
// 	s.Add("test2")
// 	if s.Has(v) {
// 		t.Errorf("Expected %s in set", v)
// 	}
// }

// func TestUnionWith(t *testing.T) {
// 	expected := []string{"a", "b", "c", "d"}

// 	x := NewSet()
// 	y := NewSet()

// 	x.Add("a")
// 	x.Add("b")
// 	y.Add("c")
// 	y.Add("d")

// 	union := x.UnionWith(y)
// 	values := union.Values()
// 	sort.Strings(values)

// 	if !reflect.DeepEqual(expected, values) {
// 		t.Errorf("Expected %q Got %q", expected, values)
// 	}
// }
