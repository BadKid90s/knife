package middleware

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var nodes = []*ServiceNode{
	{
		Address: "127.0.0.1:8080",
		Weight:  1,
	},
	{
		Address: "127.0.0.2:8080",
		Weight:  1,
	},
	{
		Address: "127.0.0.3:8080",
		Weight:  1,
	},
}

func TestLoadBalanceProxyHash(t *testing.T) {
	lb := &LoadBalancer{
		nodes:        nodes,
		currentIndex: 0,
	}

	addr := "127.0.0.1:8080"
	targetUrl1 := lb.hash(addr)
	targetUrl2 := lb.hash(addr)
	targetUrl3 := lb.hash(addr)
	msg := "compute hash value no the same"
	assert.Equal(t, targetUrl1, targetUrl2, msg)
	assert.Equal(t, targetUrl1, targetUrl3, msg)
	assert.Equal(t, targetUrl2, targetUrl3, msg)
}

func TestLoadBalanceProxyRandom(t *testing.T) {
	lb := &LoadBalancer{
		nodes:        nodes,
		currentIndex: 0,
	}

	adders := []string{
		"127.0.0.1:8080",
		"127.0.0.2:8080",
		"127.0.0.3:8080",
	}

	msg := "compute random index error"
	for i := 0; i < 100; i++ {
		targetUrl := lb.random()
		assert.Contains(t, adders, targetUrl, msg)
	}

}

func TestLoadBalanceProxyRoundRobin(t *testing.T) {
	lb := &LoadBalancer{
		nodes:        nodes,
		currentIndex: 0,
	}

	adders := []string{
		"127.0.0.1:8080",
		"127.0.0.2:8080",
		"127.0.0.3:8080",
	}

	msg := "compute roundRobin index error"
	for i := 0; i < 3; i++ {
		targetUrl := lb.roundRobin()
		println(targetUrl)
		index := (i + 1) % len(lb.nodes)
		assert.Equal(t, adders[index], targetUrl, msg)
	}
}

func TestLoadBalanceProxyWeightRoundRobin(t *testing.T) {
	lb := &LoadBalancer{
		nodes: []*ServiceNode{
			{
				Address: "127.0.0.1:8080",
				Weight:  5,
			},
			{
				Address: "127.0.0.2:8080",
				Weight:  3,
			},
			{
				Address: "127.0.0.3:8080",
				Weight:  2,
			},
		},
		currentIndex: 0,
	}

	addressNumbers := map[string]int{
		"127.0.0.1:8080": 0,
		"127.0.0.2:8080": 0,
		"127.0.0.3:8080": 0,
	}

	msg := "compute weightRoundRobin index error"
	for i := 0; i < 100; i++ {
		targetUrl := lb.weightRoundRobin()
		addressNumbers[targetUrl]++
	}
	assert.Equal(t, addressNumbers["127.0.0.1:8080"], 50, msg)
	assert.Equal(t, addressNumbers["127.0.0.2:8080"], 30, msg)
	assert.Equal(t, addressNumbers["127.0.0.3:8080"], 20, msg)

}
