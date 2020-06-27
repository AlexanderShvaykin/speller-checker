package tree

import (
	"reflect"
	"testing"
)

func TestReadFiles(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Read dir",
			args:    args{root: "../../testdata"},
			want:    []string{"../../testdata/1.txt", "../../testdata/2.txt", "../../testdata/testdir/foo.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFiles(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}
