package printer

import (
	"io"
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

func NewTabwriterPrinter(output io.Writer) *TabWriterPrinter {
	return &TabWriterPrinter{
		writer: tabwriter.NewWriter(output, 0, 0, 2, ' ', 0),
	}
}

func (p *TabWriterPrinter) Print(ref Printable, values []Printable) error {
	for i := range ref.Headers() {
		_, err := p.writer.Write([]byte(ref.Headers()[i] + "\t"))
		if err != nil {
			return err
		}
	}

	p.writer.Write([]byte("\n"))

	for i := range values {
		for j := range values[i].Values() {
			_, err := p.writer.Write([]byte(values[i].Values()[j] + "\t"))
			if err != nil {
				return err
			}
		}
		p.writer.Write([]byte("\n"))
	}

	return p.writer.Flush()
}
