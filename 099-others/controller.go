package main

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

type FooReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

func main() {
	// フェイククライアントの作成
	scheme := runtime.NewScheme()
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	// シンプルなロガーの作成
	logger := funcr.New(func(prefix, args string) {
		fmt.Printf("%s: %s\n", prefix, args)
	}, funcr.Options{})

	// FooReconcilerのインスタンスを作成
	reconciler := FooReconciler{
		Client: client,
		Log:    logger,
		Scheme: scheme,
	}

	// 構造体の内容を出力
	fmt.Printf("FooReconciler: %+v\n", reconciler)
}
