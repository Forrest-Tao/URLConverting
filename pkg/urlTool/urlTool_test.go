package urlTool

import "testing"

// TestGetBasePath 使用 ide 自带的工具，填充具体数据，进行测试
func TestGetBasePath(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "基本示例", args: args{targetUrl: "https://www.runoob.com/git/git-basic-operations.html"},
			want: "git-basic-operations.html", wantErr: false},
		{name: "相对路径示例", args: args{targetUrl: "xxx/git/git-basic-operations.html"},
			want: "", wantErr: true},
		{name: "空字符串", args: args{targetUrl: ""},
			want: "", wantErr: true},
		{name: "带query的查询", args: args{targetUrl: "https://www.runoob.com/git/git-basic-operations.html/?a=1&b=10"},
			want: "git-basic-operations.html", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.targetUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
