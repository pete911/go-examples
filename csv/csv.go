package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func ToCsv[T any](in []T) ([]byte, error) {
	data := toTable(in)
	if len(data) == 0 {
		return []byte{}, nil
	}

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	if err := w.WriteAll(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func toTable[T any](in []T) [][]string {
	if len(in) == 0 {
		return nil
	}

	var data [][]string
	data = append(data, toHeader(in[0]))
	for _, v := range in {
		data = append(data, toRow(v))
	}
	return data
}

func toHeader(in any) []string {
	var header []string
	v := reflect.TypeOf(in)
	for i := 0; i < v.NumField(); i++ {
		if csvName := v.Field(i).Tag.Get("csv"); csvName != "" {
			header = append(header, csvName)
		}
	}
	return header
}

func toRow(in any) []string {
	var row []string
	v := reflect.ValueOf(in)
	for i := 0; i < v.NumField(); i++ {
		if csvName := v.Type().Field(i).Tag.Get("csv"); csvName != "" {
			field := v.Field(i)
			// bool
			if field.Type().Kind() == reflect.Bool {
				row = append(row, fmt.Sprintf("%t", field.Bool()))
				continue
			}
			// string slice
			if slice, ok := field.Interface().([]string); ok {
				row = append(row, strings.Join(slice, " "))
				continue
			}
			// time
			if field.Type() == reflect.TypeOf(time.Time{}) {
				row = append(row, field.Interface().(time.Time).Format(time.RFC3339))
				continue
			}
			// everything else (only string in our case)
			row = append(row, v.Field(i).String())
		}
	}
	return row
}
