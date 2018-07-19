package translations

import "testing"

var expected map[string]string

func init() {
	expected = make(map[string]string)
	expected["#{availability.mon}"] = "Mon"
	expected["#{availability.mon} text #{labels.category.nature} text"] = "Mon text Nature text"
	expected["#{availability.expires-tomorrow} 10#{availability.pm}#{availability.expires-tomorrow-suffix}"] = "Expires tomorrow 10pm"
}

func TestTraslateSingleKey(t *testing.T) {
	key := "#{availability.mon}"
	actual := Translate(key)

	if expected[key] != actual {
		t.Errorf("Expected %s got %s", expected[key], actual)
	}
}

func TestTraslateMultipleKeys(t *testing.T) {
	key := "#{availability.mon} text #{labels.category.nature} text"
	actual := Translate(key)

	if expected[key] != actual {
		t.Errorf("Expected %q got %q", expected[key], actual)
	}
}

func TestIgnoreAbsentKeys(t *testing.T) {
	key := "preceding text #{DOES_NOT_EXIST_IN_LOOKUP} subsequent text."
	actual := Translate(key)

	if key != actual {
		t.Errorf("Expected %s got %s", key, actual)
	}
}

func TestIgnoreEmptyValues(t *testing.T) {
	key := "#{availability.expires-tomorrow} 10#{availability.pm}#{availability.expires-tomorrow-suffix}"
	actual := Translate(key)

	if expected[key] != actual {
		t.Errorf("Expected %s got %s", expected[key], actual)
	}
}

func TestEmptyKey(t *testing.T) {
	key := ""
	actual := Translate(key)

	if key != actual {
		t.Errorf("Expected %q got %q", key, actual)
	}
}
