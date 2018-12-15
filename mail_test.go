package mimer

import (
	"testing"
	"bytes"
)

func TestMail(t *testing.T) {

	m := New()
	m.FromName("Vitold S.")
	m.From("vit1251@gmail.com")
	m.To("support@gmail.com")
	m.Subject("Test")

	expected := "123"

	var b bytes.Buffer
	_, err := m.WriteTo(&b)
	if err != nil {
		t.Errorf("WriteTo error %v", err)
	}
	actual := b.String()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}
