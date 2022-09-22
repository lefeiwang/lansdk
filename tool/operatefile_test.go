package tool

import (
	"testing"
)

func init() {
	RemoveFile("/app/test/temp/frame")
}

func TestWriteFile(t *testing.T) {
	type args struct {
		srcFile string
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "WriteFile",
			args:    args{srcFile: "/app/test/temp/frame/1.txt", content: []byte{1}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.srcFile, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	type args struct {
		srcFile  string
		destFile string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "CopyFile",
			args:    args{srcFile: "../../test/file/1.jpeg", destFile: "/app/test/temp/frame/1.jpeg"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CopyFile(tt.args.srcFile, tt.args.destFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
