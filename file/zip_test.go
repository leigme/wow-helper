package file

import "testing"

func TestZip(t *testing.T) {
	type args struct {
		src, dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_zip_1",
			args: args{src: "C:\\Program Files (x86)\\World of Warcraft\\_classic_era_\\Interface",
				dest: "C:\\Users\\leig\\Developer\\github\\wow-helper"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Zip(tt.args.src, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
