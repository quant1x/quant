package cache

import "testing"

func TestCheckFilepath(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "t01",
			args: args{
				"a.csv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckFilepath(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("CheckFilepath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
