package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"strings"
)

var server = []string{
	"192.168.1.1",
	"192.168.2.2",
	"192.168.3.3",
	"192.168.4.4",
}

type HashRing struct {
	replicateCount int               // 每台服务所对应的节点数量（实际节点   虚拟节点）
	nodes          map[uint32]string // 键：节点哈希值 ， 值：服务器地址
	sortedNodes    []uint32          // 从小到大排序后的所有节点哈希值切片，可以认为这个就是 哈希环
}

/*
添加节点
*/
func (hr *HashRing) addNode(masterNode string) {
	for i := 0; i < hr.replicateCount; i++ {
		key := hr.hashKey(strconv.Itoa(i) + masterNode)
		hr.nodes[key] = masterNode
		hr.sortedNodes = append(hr.sortedNodes, key)
	}

	// 按照值从大到小的排序函数
	sort.Slice(hr.sortedNodes, func(i, j int) bool {
		return hr.sortedNodes[i] < hr.sortedNodes[j]
	})
}

/*
*
批量添加多个服务器节点（包含虚拟节点）的方法
*/
func (hr *HashRing) addNodes(masterNodes []string) {
	if len(masterNodes) > 0 {
		for _, node := range masterNodes {
			hr.addNode(node)
		}
	}
}

/*
*
移除节点
*/
func (hr *HashRing) removeNode(masterNode string) {
	for i := 0; i < hr.replicateCount; i++ {
		key := hr.hashKey(strconv.Itoa(i) + masterNode)

		// 移除映射关系
		delete(hr.nodes, key)

		if success, index := hr.getIndexForKey(key); success {
			// 从哈希环上移除实际节点和虚拟节点
			hr.sortedNodes = append(hr.sortedNodes[:index], hr.sortedNodes[index+1:]...)
		}
	}
}

/*
*
获取节点
*/
func (hr *HashRing) getNode(key string) string {

	if len(hr.nodes) == 0 {
		return ""
	}

	hashKey := hr.hashKey(key)
	nodes := hr.sortedNodes

	masterNode := hr.nodes[nodes[0]]

	for _, node := range nodes {
		// 如果客户端地址的哈希值小于当前节点的哈希值
		// 说明客户端的请求应当由该节点所对应的服务器来进行处理（逆时针）
		if hashKey < node {
			masterNode = hr.nodes[node]
			break
		}
	}

	return masterNode
}

func (hr *HashRing) getIndexForKey(key uint32) (bool, int) {

	index := -1
	success := false

	for i, v := range hr.sortedNodes {
		if v == key {
			index = i
			success = true
			break
		}
	}

	return success, index
}

/*
*
哈希函数（这里使用 crc32 算法来实现，返回的是一个 uint32 整型)
*/
func (hr *HashRing) hashKey(key string) uint32 {
	scratch := []byte(key)
	return crc32.ChecksumIEEE(scratch)
}

func New(nodes []string, replicateCount int) *HashRing {
	hr := new(HashRing)
	hr.replicateCount = replicateCount
	hr.nodes = make(map[uint32]string)
	hr.sortedNodes = []uint32{}
	hr.addNodes(nodes)

	return hr
}

func main() {
	hr := New(server, 10)
	// fmt.Println(hr.nodes)

	// hr.addNode("192.168.0.1")

	fifth := 0
	first, second, third, four := 0, 0, 0, 0

	for i := 0; i < 10; i++ {
		str := hr.getNode(strconv.Itoa(i))

		if strings.Compare(str, "192.168.1.1") == 0 {
			fmt.Printf("192.168.1.1：%v \n", i)
			first++
		} else if strings.Compare(str, "192.168.2.2") == 0 {
			fmt.Printf("192.168.2.2：%v \n", i)
			second++
		} else if strings.Compare(str, "192.168.3.3") == 0 {
			fmt.Printf("192.168.3.3：%v \n", i)
			third++
		} else if strings.Compare(str, "192.168.4.4") == 0 {
			fmt.Printf("192.168.4.4：%v \n", i)
			four++
		} else if strings.Compare(str, "192.168.5.5") == 0 {
			fmt.Printf("192.168.5.5：%v \n", i)
			fifth++
		}
	}

	fmt.Printf("%v %v %v %v %v", first, second, third, four, fifth)

}
