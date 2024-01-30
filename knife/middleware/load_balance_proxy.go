package middleware

import (
	"fmt"
	"hash/fnv"
	"knife"
	"math/rand"
	"sync/atomic"
	"time"
)

type LoadBalanceType int

const (
	LoadBalanceRandom           = iota //随机
	LoadBalanceRoundRobin              //轮询
	LoadBalanceWeightRoundRobin        //加权轮询
	LoadBalanceHash                    //哈希
)

type ServiceNode struct {
	Address string
	Weight  int
}

type LoadBalancer struct {
	nodes        []*ServiceNode
	currentIndex int32
}

func LoadBalanceProxy(method LoadBalanceType, serviceNodes []*ServiceNode) knife.MiddlewareFunc {
	lb := &LoadBalancer{
		nodes:        append([]*ServiceNode{}, serviceNodes...),
		currentIndex: 0,
	}
	switch method {
	case LoadBalanceRandom:
		return proxy(lb.random())
	case LoadBalanceRoundRobin:
		return proxy(lb.roundRobin())
	case LoadBalanceWeightRoundRobin:
		return proxy(lb.weightRoundRobin())
	case LoadBalanceHash:
		return proxyHash(lb)
	default:
		return proxy(lb.random())
	}
}

// 随机算法的实现
func (lb *LoadBalancer) random() string {
	// 创建一个新的随机数种子
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	// 生成0到3之间的随机整数
	randomNumber := random.Intn(len(lb.nodes))
	//获取最终的服务地址
	return lb.nodes[randomNumber].Address
}

// 轮询算法的实现
func (lb *LoadBalancer) roundRobin() string {
	index := atomic.AddInt32(&lb.currentIndex, 1)

	i := index % int32(len(lb.nodes))

	//获取最终的服务地址
	return lb.nodes[i].Address

}

// 权重轮询算法的实现
// 通过计算总权重并进行取模运算得到最终的服务节点下标
func (lb *LoadBalancer) weightRoundRobin() string {
	totalWeight := lb.getTotalWeight()

	n := len(lb.nodes)
	index := n - 1
	hit := atomic.AddInt32(&lb.currentIndex, 1) % totalWeight

	for i := 0; i < n; i++ {
		weight := int32(lb.nodes[i].Weight)
		hit = (hit + weight) % totalWeight
		if hit < weight {
			return lb.nodes[i].Address
		}
	}

	//获取最终的服务地址
	return lb.nodes[index].Address

}

// Hash算法实现
// 通过请求地址计算hash值，让每次的请求都访问同一节点
func (lb *LoadBalancer) hash(addr string) string {
	// 创建一个 32 位的 FNV-1 哈希对象
	hashed := fnv.New32()

	// 对 int 类型的值 123 进行哈希计算
	_, err := hashed.Write([]byte(addr))
	if err != nil {
		panic(fmt.Sprintf("compute hash value error：%s", err))
	}
	hashValue := hashed.Sum32()
	// 输出哈希值
	fmt.Println("哈希值为：", hashValue)

	i := hashValue % uint32(len(lb.nodes))

	//获取最终的服务地址
	return lb.nodes[i].Address

}

// 获取所用的节点的权重
func (lb *LoadBalancer) getTotalWeight() int32 {
	totalWeight := 0
	for _, node := range lb.nodes {
		totalWeight += node.Weight
	}
	return int32(totalWeight)
}

func proxy(targetUrl string) knife.MiddlewareFunc {
	return func(context *knife.Context) {
		//代理请求
		Proxy(targetUrl)
	}
}

func proxyHash(balancer *LoadBalancer) knife.MiddlewareFunc {
	return func(context *knife.Context) {
		//通过ip计算得到最终的url
		targetUrl := balancer.hash(context.Req.RemoteAddr)
		//代理请求
		Proxy(targetUrl)
	}
}
