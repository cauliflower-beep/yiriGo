package consistentHash

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
)

type Consistent struct {
	nodesReplicas int               //一个实节点对应虚拟节点的数量
	hashSortNodes []uint32          //所有节点的排序数组
	circle        map[uint32]string //所有虚拟节点对应的实node
	nodes         map[string]bool   //所有实节点
}

func NewConsistent() *Consistent {
	return &Consistent{
		nodesReplicas: 20,
		circle:        make(map[uint32]string),
		nodes:         make(map[string]bool),
	}
}

// Add 增加节点
func (c *Consistent) Add(node string) error {
	if _, ok := c.nodes[node]; ok { //判断新加节点是否存在
		return fmt.Errorf("%s already existed", node)
	}
	c.nodes[node] = true
	for i := 0; i < c.nodesReplicas; i++ { //添加虚拟节点
		replicasKey := getReplicasKey(i, node) //虚拟节点
		c.circle[replicasKey] = node
		c.hashSortNodes = append(c.hashSortNodes, replicasKey) //所有节点ID
	}
	sort.Slice(c.hashSortNodes, func(i, j int) bool { //排序
		return c.hashSortNodes[i] < c.hashSortNodes[j]
	})
	return nil
}

// Remove 删除节点
func (c *Consistent) Remove(node string) error {
	if _, ok := c.nodes[node]; !ok { //判断新加节点是否存在
		return fmt.Errorf("%s not existed", node)
	}
	delete(c.nodes, node)
	for i := 0; i < c.nodesReplicas; i++ {
		replicasKey := getReplicasKey(i, node)
		delete(c.circle, replicasKey) //删除虚拟节点
	}
	c.refreshHashSortNodes()
	return nil
}

func (c *Consistent) GetNode() (node []string) {
	for v := range c.nodes {
		node = append(node, v)
	}
	return node
}

// Get 获取具体映射到的node实节点
func (c *Consistent) Get(key string) (string, error) {
	if len(c.nodes) == 0 {
		return "", errors.New("not add node")
	}
	index := c.searchNearbyIndex(key)
	host := c.circle[c.hashSortNodes[index]]
	return host, nil
}

// refreshHashSortNodes 刷新所有虚拟节点的排序位置
func (c *Consistent) refreshHashSortNodes() {
	c.hashSortNodes = nil
	for v := range c.circle {
		c.hashSortNodes = append(c.hashSortNodes, v)
	}
	sort.Slice(c.hashSortNodes, func(i, j int) bool { //排序
		return c.hashSortNodes[i] < c.hashSortNodes[j]
	})
}

// searchNearbyIndex 查找映射到哪个虚拟节点上
func (c *Consistent) searchNearbyIndex(key string) int {
	hashKey := hashKey(key)
	index := sort.Search(len(c.hashSortNodes), func(i int) bool { //key算出的节点，距离最近的node节点下标  sort.Search数组需要升序排列好
		return c.hashSortNodes[i] >= hashKey
	})
	if index >= len(c.hashSortNodes) {
		index = 0
	}
	return index
}

func getReplicasKey(i int, node string) uint32 {
	return hashKey(fmt.Sprintf("%s#%d", node, i))
}

// hashKey hash加密
func hashKey(host string) uint32 {
	return crc32.ChecksumIEEE([]byte(host))
}
