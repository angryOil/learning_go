package main

import "sync"

func main() {

}

type MutexScoreboardManger struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func newMutexScoreboardManager() *MutexScoreboardManger {
	return &MutexScoreboardManger{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManger) update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManger) read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
