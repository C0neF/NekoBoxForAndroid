package libcore

import (
	"sync"

	tun "github.com/sagernet/sing-tun"
	"github.com/sagernet/sing/common/control"
	"github.com/sagernet/sing/common/x/list"
)

type interfaceMonitorStub struct {
	access       sync.Mutex
	myInterfaces []string
}

func (s *interfaceMonitorStub) Start() error {
	return nil
}

func (s *interfaceMonitorStub) Close() error {
	return nil
}

func (s *interfaceMonitorStub) DefaultInterface() *control.Interface {
	return nil
}

func (s *interfaceMonitorStub) OverrideAndroidVPN() bool {
	return false
}

func (s *interfaceMonitorStub) AndroidVPNEnabled() bool {
	return false
}

func (s *interfaceMonitorStub) RegisterCallback(callback tun.DefaultInterfaceUpdateCallback) *list.Element[tun.DefaultInterfaceUpdateCallback] {
	return nil
}

func (s *interfaceMonitorStub) UnregisterCallback(element *list.Element[tun.DefaultInterfaceUpdateCallback]) {
}

func (s *interfaceMonitorStub) RegisterMyInterface(interfaceName string) {
	s.access.Lock()
	defer s.access.Unlock()
	s.myInterfaces = append(s.myInterfaces, interfaceName)
}

func (s *interfaceMonitorStub) MyInterface() string {
	s.access.Lock()
	defer s.access.Unlock()
	if len(s.myInterfaces) > 0 {
		return s.myInterfaces[0]
	}
	return ""
}

func (s *interfaceMonitorStub) MyInterfaces() []string {
	s.access.Lock()
	defer s.access.Unlock()
	return s.myInterfaces
}
