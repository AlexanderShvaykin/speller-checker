package tree

import (
	"reflect"
	"testing"
)

var buff []string

func TestReadFiles(t *testing.T) {
	type args struct {
		root    string
		handler func(line string)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []string
	}{
		{
			name:    "Read files and call func with line",
			args:    args{root: "../../testdata", handler: testHandler},
			wantErr: false,
			want:    []string{"Собака", "Щука", "Щека", "Собака", "Щука"},
		},
		{
			name:    "Try read not existed files",
			args:    args{root: "../../testdata123", handler: testHandler},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		buff = []string{}
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadFiles(tt.args.root, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("ReadFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.want) > 0 && !reflect.DeepEqual(buff, tt.want) {
				t.Errorf("ReadFiles() call handler with %v, want %v", buff, tt.want)
			}
		})
	}
}

func testHandler(line string) {
	buff = append(buff, line)
}
