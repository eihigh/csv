package csv_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/eihigh/csv"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		text    string
		heading bool
		header  []string
		records [][]string
		err     error
	}{
		{
			text:    "a,b,c\n1,2,3",
			heading: true,
			header:  []string{"a", "b", "c"},
			records: [][]string{{"1", "2", "3"}},
			err:     nil,
		},
		{
			text:    "a,b,c\n1,2",
			heading: true,
			header:  []string{"a", "b", "c"},
			records: [][]string{},
			err:     csv.ErrFieldCount,
		},
	}

	for _, tt := range tests {
		s := csv.NewScanner(csv.NewReader(strings.NewReader(tt.text)), tt.heading)
		records := [][]string{}
		for s.Scan() {
			records = append(records, s.Record())
		}
		if !reflect.DeepEqual(tt.header, s.Header()) {
			t.Errorf("scanner.Header(): want %+v\ngot %+v", tt.header, s.Header())
		}
		if !reflect.DeepEqual(tt.records, records) {
			t.Errorf("records: want %+v\ngot %+v", tt.records, records)
		}
		if tt.err == nil {
			if s.Err() != nil {
				t.Errorf("scanner.Err(): want %+v; got %+v", tt.err, s.Err())
			}
		} else if errors.Is(tt.err, s.Err()) {
			t.Errorf("scanner.Err(): want %+v; got %+v", tt.err, s.Err())
		}
	}
}
