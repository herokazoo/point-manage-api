package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

/*
関数の外部からサーバーのプロセスを中断可能にする
 1.*http.Server.ListenAndServeメソッドを実行してHTTPリクエストを受け付ける
 2.引数で渡されたcontext.Contextを通じて処理の中断命令を検知した時は
   *http.Server.ShutdownメソッドでHTTPサーバーの機能を終了する
 3.run関数の戻り値として*http.Server.ListenAndServeメソッドの戻り値のエラーを返す
*/
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		// http.ErrServerClosedは
		// http.Server.Shutdown()が正常に終了したことを示すので異常ではない
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})
	// チャネルからの通知（終了通知）を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Goメソッドで起動した別ゴルーチンの終了を待つ
	return eg.Wait()
}

