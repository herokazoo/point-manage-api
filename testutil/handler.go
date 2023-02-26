package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// HTTPハンドラーのテストで利用する検証手順
// JSON文字列をUnmarshalし、cmpパッケージを使って差分比較する
func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannot unmarshal got %q: %v", got, err)
	}
	if diff := cmp.Diff(jg, jw); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}

	// 期待としても実体としてもレスポンスボディがない場合はAssertJSONを呼ぶ必要はない
	if len(gb) == 0 && len(body) == 0 {
		return
	}
	AssertJSON(t, body, gb)
}

// ゴールデンテストの入力値と期待値をファイルから取得する
func LoodFile(t *testing.T, path string) []byte {
	t.Helper()
	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read from %q: %v", path, err)
	}
	return bt
}
