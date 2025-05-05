package lighthouse

import (
	"os"
	"reflect"
	"testing"
)

func TestUploadData(t *testing.T) {

	fileBytes, err := os.ReadFile("./Kvasir.webp")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	type args struct {
		apikey   string
		fileName string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				apikey:   "27f11aa2.ece6165afe694a718212ad611a08a3a2",
				fileName: "Kvasir.webp",
				data:     fileBytes,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UploadData(tt.args.apikey, tt.args.fileName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UploadData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadDataSimple(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				hash: "QmUNabFWZhQ7LuQ9tqkZrufB9Q5N5Yb2JGWb3RMuXhybaQ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DownloadDataSimple(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownloadData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DownloadData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetFileInfo(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    *FileInfo
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				hash: "QmUNabFWZhQ7LuQ9tqkZrufB9Q5N5Yb2JGWb3RMuXhybaQ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileInfo(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadChunkedData(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				hash: "QmPe96CRaXHrsYvCPxu9PHz5JPFp9sxnCrsFywVUj1tg58",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadChunkedData(tt.args.hash, "./")
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadChunkedData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadChunkedData() = %v, want %v", got, tt.want)
			}
		})
	}
}
