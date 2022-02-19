package no_router

import (
	"errors"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// Node 压缩字典树（Radix Tree）
// 参考 http://c.biancheng.net/view/5403.html
type Node struct {
	Path  string
	Func  gin.HandlerFunc
	Trees map[string]*Node // 并发读写 加锁
}

var rootNode = NewNode()
var rwMut = new(sync.RWMutex)

func NewNode() *Node {
	return &Node{
		Path:  "",
		Func:  nil,
		Trees: make(map[string]*Node),
	}
}

func (n *Node) Set(path string, f gin.HandlerFunc) {
	n.Path = path
	n.Func = f
}

func (n *Node) HasTress(name string) (ok bool) {
	rwMut.RLock()
	_, ok = n.Trees[name]
	rwMut.RUnlock()

	return
}

func (n *Node) GetTree(name string) (tree *Node, ok bool) {
	rwMut.RLock()
	tree, ok = n.Trees[name]
	rwMut.RUnlock()

	return
}

func (n *Node) Add(name string, no *Node) {
	rwMut.Lock()
	if n.Trees == nil {
		n.Trees = make(map[string]*Node)
	}

	n.Trees[name] = no

	rwMut.Unlock()
}

func (n *Node) SetPathAndFunc(basePath string, f gin.HandlerFunc) error {
	if basePath == "/" {
		n.Func = f
		return nil
	}

	// 过滤前后的/
	basePath = strings.Trim(basePath, "/")
	// 分割uri
	s := strings.Split(basePath, "/")

	var (
		no    = n
		index = 1
		l     = len(s)
	)

	for _, name := range s {
		// 去除零值
		if name == "" {
			continue
		}

		Node, ok := no.GetTree(name)
		if !ok {
			Node = NewNode()
			Node.Path = name
			no.Add(name, Node)
		}

		no = Node

		if index == l {
			if no.Func != nil {
				return errors.New(basePath + " 设置方法出错，方法存在重叠")
			}

			no.Func = f

			return nil
		}

		index++
	}

	return nil
}

func (n *Node) GetPathFunc(basePath string) gin.HandlerFunc {
	// 过滤前后的/
	basePath = strings.Trim(basePath, "/")
	// 分割uri
	s := strings.Split(basePath, "/")

	var (
		ok bool
		no = n
		f  = n.Func
	)

	for _, name := range s {
		// 去除零值
		if name == "" {
			continue
		}

		no, ok = no.GetTree(name)
		if !ok {
			break
		}

		if no.Func != nil {
			f = no.Func
		}
	}

	return f
}

// NoRoute 存入 NoRouteMap
func NoRoute(basePath string, handle gin.HandlerFunc) {
	err := rootNode.SetPathAndFunc(basePath, handle)
	if err != nil {
		panic(err)
	}
}

// NoRouteHandle 404处理
func NoRouteHandle(ctx *gin.Context) {
	if fun := rootNode.GetPathFunc(ctx.Request.RequestURI); fun != nil {
		fun(ctx)
	}
}
