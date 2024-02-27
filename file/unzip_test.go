package file

import "testing"

func TestUnzip(t *testing.T) {
	type args struct {
		zipFile string
		destDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_unzip_1",
			args: args{
				zipFile: "C:\\Users\\leig\\Developer\\github\\wow-helper\\Interface.zip",
				destDir: "C:\\Users\\leig\\Developer\\github\\wow-helper",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unzip(tt.args.zipFile, tt.args.destDir); (err != nil) != tt.wantErr {
				t.Errorf("Unzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
