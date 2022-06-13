package test_utils

import (
	"reflect"
	"testing"
)

func TestParseContentBody(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name            string
		args            args
		wantContentBody string
		wantErr         bool
	}{
		{
			name:            "it should return an error when the content body is empty",
			args:            args{content: "---\n---\n"},
			wantContentBody: "",
			wantErr:         true,
		},
		{
			name:            "it should return the rest of the front matter deleted",
			args:            args{content: "---\n---\na"},
			wantContentBody: "a",
			wantErr:         false,
		},
		{
			name:            "it can parse the body even when the content includes multiple \"---\"",
			args:            args{content: "---\n---\na\n---\nb---"},
			wantContentBody: "a\n---\nb---",
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContentBody, err := ParseContentBody(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseContentBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotContentBody != tt.wantContentBody {
				t.Errorf("ParseContentBody() = %v, want %v", gotContentBody, tt.wantContentBody)
			}
		})
	}
}

func TestParseImageFilePaths(t *testing.T) {
	type args struct {
		contentBody string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "it can parse an image literal including a relative image path",
			args: args{contentBody: "![](relative.png)"},
			want: []string{"relative.png"},
		},
		{
			name: "it can parse an image literal including a absolute image path",
			args: args{contentBody: "![](/absolute.png)"},
			want: []string{"/absolute.png"},
		},
		{
			name: "it should ignore the \"[]()\" literal",
			args: args{contentBody: "[](/no-bang.png)"},
			want: []string{},
		},
		{
			name: "it should ignore urls",
			args: args{contentBody: "![](http://foo.com/a.png)"},
			want: []string{},
		},
		{
			name: "it can parse multiple image literals",
			args: args{contentBody: "![](/a.jpg) ![](foo/b.gif)"},
			want: []string{"/a.jpg", "foo/b.gif"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseImageFilePaths(tt.args.contentBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseImageFilePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateImageFilePath(t *testing.T) {
	type args struct {
		imageFilePath string
		postFilePath  string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "it should return an error when the post file path does not include a YYYY string",
			args:    args{imageFilePath: "", postFilePath: "anywhere/not_year/slug.md"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "it should return an error when the post file path has an invalid slug",
			args:    args{imageFilePath: "", postFilePath: "anywhere/2021/invalid_slug.md"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "it should return false when the image file path does not start with \"/external/posts\" dir",
			args:    args{imageFilePath: "/static/external/posts/2021/slug/a.png", postFilePath: "anywhere/2021/slug.md"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "it should return false when the year is mismatched",
			args:    args{imageFilePath: "/external/posts/2020/slug/a.png", postFilePath: "anywhere/2021/slug.md"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "it should return false when the slug is mismatched",
			args:    args{imageFilePath: "/external/posts/2021/slug2/a.png", postFilePath: "anywhere/2021/slug.md"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "it should return true when the image file path is probably an url",
			args:    args{imageFilePath: "http://foo.com/path/to/x.jpg", postFilePath: "anywhere/2021/slug.md"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "it should return true when the image and post file path match",
			args:    args{imageFilePath: "/external/posts/2021/slug/a.png", postFilePath: "anywhere/2021/slug.md"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateImageFilePath(tt.args.imageFilePath, tt.args.postFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateImageFilePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateImageFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
