package lib

import (
	"reflect"
	"testing"
)

func TestGenerateQueryParamsFromStruct(t *testing.T) {
	type args struct {
		queryParams interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []QueryParam
		wantErr bool
	}{
		{
			name: "should return empty array",
			args: args{
				queryParams: struct {
					Name  string
					Place string
				}{},
			},
			want:    []QueryParam{},
			wantErr: false,
		},
		{
			name: "should return array with 1 element",
			args: args{
				queryParams: struct {
					Name  string
					Place string
				}{
					Name:  "John",
					Place: "",
				},
			},
			want: []QueryParam{
				{
					Key:   "Name",
					Value: "John",
				},
			},
			wantErr: false,
		},
		{
			name: "should return array with 2 elements",
			args: args{
				queryParams: struct {
					Name  string
					Place string
				}{
					Name:  "John",
					Place: "Jakarta",
				},
			},
			want: []QueryParam{
				{
					Key:   "Name",
					Value: "John",
				},
				{
					Key:   "Place",
					Value: "Jakarta",
				},
			},
			wantErr: false,
		},
		{
			name: "should return error",
			args: args{
				queryParams: "John",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return array with 1 element with queryKey",
			args: args{
				queryParams: struct {
					Name    string `queryKey:"name"`
					Address string
				}{
					Name:    "John",
					Address: "Jakarta",
				},
			},
			want: []QueryParam{
				{
					Key:   "name",
					Value: "John",
				},
				{
					Key:   "Address",
					Value: "Jakarta",
				},
			},
			wantErr: false,
		},
		{
			name: "should throw error if nil is passed",
			args: args{
				queryParams: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateQueryParamsFromStruct(tt.args.queryParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateQueryParamsFromStruct() got = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateQueryParamsFromStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
