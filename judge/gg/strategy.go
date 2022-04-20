package gg

import (
	"cicada/pkg/model"
	"sync"
)

type SafeStrategy struct {
	sync.RWMutex
	M []model.AlarmStrategy
}

var (
	AlarmStrategy = &SafeStrategy{M: make([]model.AlarmStrategy, 0)} // 缓存终端和策略的映射关系
)

func (s *SafeStrategy) ReInit(m []model.AlarmStrategy) {
	s.Lock()
	defer s.Unlock()
	s.M = m
}

func (s *SafeStrategy) Get() []model.AlarmStrategy {
	s.RLock()
	defer s.RUnlock()
	return s.M
}

type SafeSubscribeStrategy struct {
	sync.RWMutex
	M []model.SubscribeStrategy
}

var (
	SubscribeStrategy = &SafeStrategy{M: make([]model.AlarmStrategy, 0)}
)

func (s *SafeSubscribeStrategy) ReInit(m []model.SubscribeStrategy) {
	s.Lock()
	defer s.Unlock()
	s.M = m
}

func (s *SafeSubscribeStrategy) Get() []model.SubscribeStrategy {
	s.RLock()
	defer s.RUnlock()
	return s.M
}
