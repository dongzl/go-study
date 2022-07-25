package test_cover

import "testing"

// https://qiubo.ink/2019/06/27/Golang%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95%E5%8F%8A%E6%B5%8B%E8%AF%95%E8%A6%86%E7%9B%96%E7%8E%87/

// 普通测试,只展示测试结果及时间
// go test goshop/app/lib

// 展示测试明细, 此时仍没有覆盖率
// go test goshop/app/lib -v

// go test goshop/app/lib -v -covermode=count
// go test goshop/app/lib -v -coverprofile=count.out
// go tool cover -func=count.out
// go tool cover -html=count.out

func TestCheckPhoneNumber(t *testing.T) {
	type suite struct {
		data   string
		result bool
	}

	phoneExample := []suite{
		suite{"", false},
		suite{"12345678901", false},
		suite{"1020", false},
		suite{"19456782901", true},
	}

	for _, v := range phoneExample {
		if result := CheckPhoneNumber(v.data); result != v.result {
			t.Errorf("测试用例 %s : 测试结构 %t , 与期望不符合 %t", v.data, result, v.result)
		}
	}
}

func TestStringLength(t *testing.T) {
	type suite struct {
		data   string
		result int
	}

	strExample := []suite{
		suite{"1", 1},
		suite{"", 0},
		suite{"1678", 4},
		suite{"13456789301", 11},
		suite{"abcd", 4},
		suite{"hello world", 11},
		suite{"我在这", 3},
		suite{"你是谁?Who are you?", 16},
	}

	for _, v := range strExample {
		if result := StringLength(v.data); result != v.result {
			t.Errorf("测试用例 %s : 测试结构 %d , 与期望不符合 %d", v.data, result, v.result)
		}
	}
}
