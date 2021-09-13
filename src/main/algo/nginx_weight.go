package main

import "fmt"
/**
Nginx 平滑的基于权重的轮训算法
 */

func main() {
	nodes := []*Node{
		&Node{"a", 0, 5},
		&Node{"b", 0, 1},
		&Node{"c", 0, 1},
	}

	for i := 0; i < 7; i++ {
		best := SmoothWrr(nodes)
		if best != nil {
			fmt.Println(best.Name)
		}
	}
}

type Node struct {
	Name    string  //节点名字
	Current int     //节点当前权重
	Weight  int     //节点权重
}

//具有命名的函数返回参数, return 可以不带参数，自动返回函数签名中的命名参数
func SmoothWrr(nodes []*Node) (best *Node) {
	if len(nodes) == 0 {
		return
	}
	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}
		total += node.Weight  //权重总和
		node.Current += node.Weight //每个节点，用它们的当前值加上它们自己的权重
		if best == nil || node.Current > best.Current {//选出第1个最大值
			best = node
		}
	}
	if best == nil {
		return
	}
	best.Current -= total  //选择当前值最大的节点为选中节点，并把它的当前值减去所有节点的权重总和
	return
}



