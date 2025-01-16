package dataloader

type NoopMutex struct{}

func (m NoopMutex) Lock()   {}
func (m NoopMutex) Unlock() {}
