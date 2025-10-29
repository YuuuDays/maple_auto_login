package service

import "testing"

func TestGetEnv(t *testing.T) {
	// テーブル駆動テスト&サブテストで書こう!

	okTests := []struct {
		id   string
		pass string
	}{
		{"idid", "paspas"},
		{"id", "pas"},
	}

	t.Run("成功", func(t *testing.T) {
		got := "これはサンプルです。"
		want := GetEnv()

		if got != want {
			t.Errorf("got %q want &q", got, want)
		}
	})

}
