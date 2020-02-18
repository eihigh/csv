package csv

import (
	"encoding/csv"
	"io"
)

var (
	ErrTrailingComma = csv.ErrTrailingComma
	ErrBareQuote     = csv.ErrBareQuote
	ErrQuote         = csv.ErrQuote
	ErrFieldCount    = csv.ErrFieldCount
)

type Reader = csv.Reader

func NewReader(r io.Reader) *Reader {
	return csv.NewReader(r)
}

type Writer = csv.Writer

type ParseError = csv.ParseError

type Scanner struct {
	r      *csv.Reader
	header []string
	record []string
	err    error
}

func NewScanner(r *csv.Reader, header bool) *Scanner {
	s := &Scanner{r: r}
	if !header {
		return s
	}
	if s.Scan() {
		s.header = s.record
		s.record = nil
	}
	return s
}

func (s *Scanner) Header() []string {
	return s.header
}

func (s *Scanner) Scan() bool {
	s.record, s.err = s.r.Read()
	if s.err != nil {
		return false
	}
	return true
}

func (s *Scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

func (s *Scanner) Record() []string {
	return s.record
}
