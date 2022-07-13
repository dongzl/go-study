package mock

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

// https://geektutu.com/post/gee.html
// https://geektutu.com/post/quick-gomock.html

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
