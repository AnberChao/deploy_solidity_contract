package event

var event map[string]*ContractEvent

func GetEventOption() map[string]*ContractEvent {
	return event
}

func init() {
	event = make(map[string]*ContractEvent)
	event[_MyTokenProxyAddress] = NewMyTokenContractEvent(_MyTokenProxyAddress)
}

type EventOption interface {
	Name() string
	DataStruct() interface{}
	ShowData() string
}
