package utils

type Listener interface {
	OnUpdate(instance interface{})
}

type ListenerFunc func(instance interface{})

func (f ListenerFunc) OnUpdate(instance interface{}) {
	f(instance)
}
