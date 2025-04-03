// process/process_manager.go
package process

import (
	"os/exec"
	"sync"
	"fmt"
)

type ProcessInfo struct {
	Cmd *exec.Cmd
}

var (
	processes = make(map[string]*ProcessInfo)
	mu        sync.RWMutex
)

func StartProcess(id string, command string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := processes[id]; exists {
		return fmt.Errorf("process %s already running", id)
	}

	cmd := exec.Command("cmd", "/C", command)
	err := cmd.Start()
	if err != nil {
		return err
	}

	processes[id] = &ProcessInfo{Cmd: cmd}
	go cmd.Wait()

	return nil
}

func StopProcess(id string) error {
	mu.Lock()
	defer mu.Unlock()

	proc, exists := processes[id]
	if !exists {
		return fmt.Errorf("process %s not running", id)
	}

	err := proc.Cmd.Process.Kill()
	if err != nil {
		return err
	}

	delete(processes, id)
	return nil
}

func IsRunning(id string) bool {
	mu.RLock()
	defer mu.RUnlock()
	_, exists := processes[id]
	return exists
}

func StopAll() {
	mu.Lock()
	defer mu.Unlock()
	for id, proc := range processes {
		proc.Cmd.Process.Kill()
		delete(processes, id)
	}
}

func GetAllRunning() []string {
	mu.RLock()
	defer mu.RUnlock()
	running := []string{}
	for id := range processes {
		running = append(running, id)
	}
	return running
}
