package printer

import (
	"os"
	"text/tabwriter"
)

type Printable interface {
	Headers() []string
	Values() []string
}

type Printer interface {
	Print(ref Printable, values []Printable) error
}

type TabWriterPrinter struct {
	writer *tabwriter.Writer
}

func NewTabwriterPrinter() *TabWriterPrinter {
	return &TabWriterPrinter{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

func (p *TabWriterPrinter) Print(ref Printable, values []Printable) error {
	headers := ref.Headers()
	for _, header := range headers {
		p.writer.Write([]byte(header))
		p.writer.Write([]byte("\t"))
	}
	p.writer.Write([]byte("\n"))

	for _, value := range values {
		for _, v := range value.Values() {
			p.writer.Write([]byte(v))
			p.writer.Write([]byte("\t"))
		}
		p.writer.Write([]byte("\n"))
	}
	p.writer.Flush()
	return nil
}
