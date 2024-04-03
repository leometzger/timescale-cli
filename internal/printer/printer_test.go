package printer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPrintable struct {
	Name  string
	Value string
}

func (p TestPrintable) Headers() []string {
	return []string{"Name", "Value"}
}

func (p TestPrintable) Values() []string {
	return []string{
		p.Name,
		p.Value,
	}
}

func TestPrinterPrintCorrectly(t *testing.T) {
	var buffer bytes.Buffer

	printer := NewTabwriterPrinter(&buffer)

	ref := &TestPrintable{}
	values := []Printable{
		&TestPrintable{Name: "name1", Value: "value1"},
		&TestPrintable{Name: "name2", Value: "value2"},
		&TestPrintable{Name: "name3", Value: "value3"},
	}

	err := printer.Print(ref, values)

	assert.Nil(t, err)
	assert.NotEmpty(t, buffer.String())

	for _, value := range values {
		assert.Contains(t, buffer.String(), value.Values()[0])
		assert.Contains(t, buffer.String(), value.Values()[1])
	}
}
