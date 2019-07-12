package unjson

import (
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	type myStruct struct {
		Str string    `json:"str"`
		I64 int64     `json:"i64"`
		I32 int32     `json:"i32"`
		I16 int16     `json:"i16"`
		I8  int8      `json:"i8"`
		I   int       `json:"i"`
		F32 float32   `json:"f32"`
		F64 float64   `json:"f64"`
		Tm  time.Time `json:"tm"`
	}

	data := myStruct{
		Str: "string",
		I64: 123,
		I32: 123,
		I16: 123,
		I8:  123,
		I:   123,
		F32: 123.32,
		F64: 123.32,
		Tm:  time.Unix(123, 123),
	}

	tests := []struct {
		name    string
		data    interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "empty",
			data:    myStruct{},
			want:    `{str: "", i64: 0, i32: 0, i16: 0, i8: 0, i: 0, f32: 0, f64: 0, tm: "0001-01-01 00:00:00 +0000 UTC"}`,
			wantErr: false,
		},
		{
			name:    "struct",
			data:    data,
			want:    `{str: "string", i64: 123, i32: 123, i16: 123, i8: 123, i: 123, f32: 123.32, f64: 123.32, tm: "1970-01-01 07:02:03.000000123 +0700 WIB"}`,
			wantErr: false,
		},
		{
			name:    "struct",
			data:    &data,
			want:    `{str: "string", i64: 123, i32: 123, i16: 123, i8: 123, i: 123, f32: 123.32, f64: 123.32, tm: "1970-01-01 07:02:03.000000123 +0700 WIB"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalWithoutQuotes(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarshalWithoutQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
