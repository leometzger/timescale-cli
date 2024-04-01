package printer

import (
	"log/slog"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"
)

type Printable interface {
	Headers() []string
	Values() []string
}

type Printer interface {
	Print(ref any, values []any) error
}

type TabWriterPrinter struct {
	writer *tabwriter.Writer
}

func NewTabwriterPrinter() *TabWriterPrinter {
	return &TabWriterPrinter{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

func (p *TabWriterPrinter) Print(ref any, values []any) error {
	// @TODO add error handling
	// @TODO refactor a little bit
	r := reflect.TypeOf(ref)
	for i := 0; i < r.NumField(); i++ {
		var header string

		field := r.Field(i)
		header, ok := field.Tag.Lookup("header")
		if !ok {
			header = field.Name
		}

		p.writer.Write([]byte(header))
		p.writer.Write([]byte("\t"))
	}

	p.writer.Write([]byte("\n"))

	for _, value := range values {
		r := reflect.TypeOf(value)

		for i := 0; i < r.NumField(); i++ {
			k := r.Field(i).Type
			v := reflect.ValueOf(value).Field(i)

			switch k.Kind() {
			case reflect.String:
				p.writer.Write([]byte(v.String()))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				p.writer.Write([]byte(strconv.FormatInt(v.Int(), 10)))
			case reflect.Bool:
				p.writer.Write([]byte(strconv.FormatBool(v.Bool())))
			default:
				slog.Error("invalid kind", "kind", k.Kind())
			}
			p.writer.Write([]byte("\t"))
		}
		p.writer.Write([]byte("\n"))
	}
	p.writer.Flush()
	return nil
}
