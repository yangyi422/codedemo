// 用于动态路由判断的树
package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node // 根节点
}

type node struct {
	islast  bool              // 代表该节点是否是最终的路由规则，是否能成为一个独立uri，时候自身就是最后的节点
	segment string            // uri中的字符串，代表此节点是否是路由中的某个段的字符串
	handler ControllerHandler // 代表这个节点中的控制器，用于最终的加载调用
	childs  []*node           // 代表此节点的子节点
}

func NewTree() *Tree {
	return &Tree{
		root: newNode(),
	}
}

func newNode() *node {
	return &node{}
}

// 判断一个segment是否是通用符号开头，例如：
func isUniversalSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果segment是通用符号开头的，则下层节点全部都满足要求
	if isUniversalSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤下层所有节点
	for _, cnode := range n.childs {
		if isUniversalSegment(cnode.segment) {
			// 如果下层节点是通用符号开头，则说明下层节点的uri满足筛选要求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// 如果下层节点不是通用符号开头，但是与入参文本完全匹配，那么uri也满足要求
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

// 判断路由是否已经在节点的所有子节点中存在了
func (n *node) matchNode(uri string) *node {
	// 使用分隔符切分uri
	segments := strings.SplitN(uri, "/", 2)
	// 第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isUniversalSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层节点
	cnodes := n.filterChildNodes(segment)
	// 如果当前的子节点没有一个符合条件，那么uri一定是不存在的
	if len(cnodes) == 0 {
		return nil
	}

	// 如果只有一个segment，那么就给这个segment打上最后的标记
	if len(segments) == 1 {
		// 如果这个segment已经是最后一个节点，就判断这个节点是否有islast标记
		for _, tn := range cnodes {
			if tn.islast {
				return tn
			}
		}
		return nil
	}

	// 如果有多个segment，则递归对应的子节点继续查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil

}

// 增加路由节点
func (t *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := t.root
	// 判断路由是否冲突
	if n.matchNode(uri) != nil {
		return errors.New("route exist" + uri)
	}

	segments := strings.Split(uri, "/")
	// 对每个segment检查
	for index, segment := range segments {
		// 最终进入Node segment的字段
		if !isUniversalSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		// 标记是否有而合适的子节点
		var objNode *node
		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.islast = true
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}

func (t *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := t.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
