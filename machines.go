package main

import (
	"fmt"
	"sync"
	"time"
)

type Machine struct {
	ID      string
	Host    string
	Running bool
	Started time.Time
}

type MachineStore struct {
	machines map[string]*Machine
	mu       sync.Mutex
}

func NewMachineStore() *MachineStore {
	return &MachineStore{
		machines: make(map[string]*Machine),
	}
}

func (ms *MachineStore) StartMachine(id, host string) *Machine {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	m := &Machine{ID: id, Host: host, Running: true, Started: time.Now()}
	ms.machines[id] = m
	return m
}

func (ms *MachineStore) StopMachine(id string) bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if m, ok := ms.machines[id]; ok {
		m.Running = false
		return true
	}
	return false
}

func (ms *MachineStore) CloneMachine(oldID, newID, newHost string) *Machine {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	fmt.Println("Available machine IDs:", ms.machines)
	fmt.Println("Trying to clone machine:", oldID)
	if old, ok := ms.machines[oldID]; ok && old.Running {
		clone := &Machine{ID: newID, Host: newHost, Running: true, Started: time.Now()}
		ms.machines[newID] = clone
		return clone
	}
	return nil
}

func (ms *MachineStore) ListMachines() []*Machine {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	var list []*Machine
	for _, m := range ms.machines {
		list = append(list, m)
	}
	return list
}
