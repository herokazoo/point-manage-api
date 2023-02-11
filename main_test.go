package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

/*
1 キャンセル可能なcontext.Contextオブジェクトを作成
2 別ゴルーチンでテスト対象のrun関数を実行してHTTPサーバーを起動する
3 エンドポイントに対してGETリクエストを送信する
4 cancel関数を実行する
5 *errgroup.Group.Waitメソッド経由でrun関数の戻り値を検証する
6 GETリクエストで取得したレスポンスボディが期待する文字列であることを検証する
*/
func TestRun(t *testing.T) {
	t.Skip("TODO: リファクタリング")

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// リッスンしているポート番号の確認
	t.Logf("try request to %q", url)
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	// HTTPサーバーの戻り値を検証する
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
	// run関数に終了通知を送信する
	cancel()
	// run関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
