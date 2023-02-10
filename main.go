package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

/*
関数の外部からサーバーのプロセスを中断可能にする
動的にポート番号を変更可能する

	1.*http.Server.ListenAndServeメソッドを実行してHTTPリクエストを受け付ける
	2.引数で渡されたcontext.Contextを通じて処理の中断命令を検知した時は
	  *http.Server.ShutdownメソッドでHTTPサーバーの機能を終了する
	3.run関数の戻り値として*http.Server.ListenAndServeメソッドの戻り値のエラーを返す
*/
func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		// 引数で受け取ったnet.Listenerを利用するので、Addrフィールドは指定しない
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		// ListenAndServeメソッドではなく、Serveメソッドに変更する
		// http.ErrServerClosedは
		// http.Server.Shutdown()が正常に終了したことを示すので異常ではない
		if err := s.Serve(l); err != nil &&
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
