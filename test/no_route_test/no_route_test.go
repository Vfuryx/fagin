package no_route_test

import (
	"fagin/pkg/router/no_router"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

var rootNode = no_router.NewNode()

func TestMain(m *testing.M) {
	var err error
	err = rootNode.SetPathAndFunc("/", func(ctx *gin.Context) {
		fmt.Println("/")
	})
	if err != nil {
		panic(err)
	}

	err = rootNode.SetPathAndFunc("/a/b/c", func(ctx *gin.Context) {
		fmt.Println("/a/b/c")
	})
	if err != nil {
		panic(err)
	}

	err = rootNode.SetPathAndFunc("/a/c", func(ctx *gin.Context) {
		fmt.Println("/a/c")
	})
	if err != nil {
		panic(err)
	}

	err = rootNode.SetPathAndFunc("/a/b/cc", func(ctx *gin.Context) {
		fmt.Println("/a/b/cc")
	})
	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestNode(t *testing.T) {
	f := rootNode.GetPathFunc("/a/b/c//b")
	if f == nil {
		t.Fatal("失败")
	}
	f(nil)
}

func BenchmarkNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode.GetPathFunc("/a/b/c//b")
	}
}
