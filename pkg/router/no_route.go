package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

// 压缩字典树（Radix Tree）
// 参考 http://c.biancheng.net/view/5403.html
type node struct {
	Path  string
	Func  gin.HandlerFunc
	Trees map[string]*node // 并发读写 加锁
}

var rootNode = NewNode()
var rwMut = new(sync.RWMutex)

func NewNode() *node {
	return &node{
		Path:  "",
		Func:  nil,
		Trees: make(map[string]*node, 10),
	}
}

func (n *node) Set(path string, f gin.HandlerFunc) {
	n.Path = path
	n.Func = f
}

func (n *node) HasTress(name string) (ok bool) {
	rwMut.RLock()
	_, ok = n.Trees[name]
	rwMut.RUnlock()
	return
}

func (n *node) GetTree(name string) (tree *node, ok bool) {
	rwMut.RLock()
	tree, ok = n.Trees[name]
	rwMut.RUnlock()
	return
}

func (n *node) Add(name string, no *node) {
	rwMut.Lock()
	if n.Trees == nil {
		n.Trees = make(map[string]*node, 10)
	}
	n.Trees[name] = no
	rwMut.Unlock()
}

func (n *node) SetPathAndFunc(basePath string, f gin.HandlerFunc) error {
	if basePath == "/" {
		n.Func = f
		return nil
	}

	// 过滤前后的/
	basePath = strings.Trim(basePath, "/")
	// 分割uri
	s := strings.Split(basePath, "/")

	var (
		no    *node = n
		index       = 1
		l           = len(s)
	)

	for _, name := range s {
		// 去除零值
		if name == "" {
			continue
		}

		node, ok := no.GetTree(name)
		if !ok {
			node = NewNode()
			node.Path = name
			no.Add(name, node)
		}
		no = node

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

func (n *node) GetPathFunc(basePath string) gin.HandlerFunc {
	// 过滤前后的/
	basePath = strings.Trim(basePath, "/")
	// 分割uri
	s := strings.Split(basePath, "/")

	var (
		ok bool
		no *node           = n
		f  gin.HandlerFunc = n.Func
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
