package order

import (
	"errors"
	"fmt"
	"sync"
)

type State int                             // 状态
type Event string                          // 事件
type Handler func(opt *Opt) (State, error) // 处理方法，并返回新的状态

// FSM 有限状态机
type FSM struct {
	mu       sync.Mutex                  // 排他锁
	state    State                       // 当前状态
	handlers map[State]map[Event]Handler // 当前状态可触发的有限个事件
}

// 获取当前状态
func (f *FSM) getState() State {
	return f.state
}

// 设置当前状态
func (f *FSM) setState(newState State) {
	f.state = newState
}

// addHandlers 添加事件和处理方法
func (f *FSM) addHandlers() (*FSM, error) {
	if statusEvent == nil || len(statusEvent) <= 0 {
		return nil, errors.New("[警告] 未定义 statusEvent")
	}

	for state, events := range statusEvent {
		if len(events) <= 0 {
			return nil, errors.New(fmt.Sprintf("[警告] 状态(%s)未定义事件", StatusText(state)))
		}

		for _, event := range events {
			handler := eventHandler[event]
			if handler == nil {
				return nil, errors.New(fmt.Sprintf("[警告] 事件(%s)未定义处理方法", event))
			}

			if _, ok := f.handlers[state]; !ok {
				f.handlers[state] = make(map[Event]Handler)
			}

			if _, ok := f.handlers[state][event]; ok {
				return nil, errors.New(fmt.Sprintf("[警告] 状态(%s)事件(%s)已定义过", StatusText(state), event))
			}

			f.handlers[state][event] = handler
		}
	}

	return f, nil
}

// Call 事件处理
func (f *FSM) Call(event Event, opts ...Option) (State, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	events := f.handlers[f.getState()]
	if events == nil {
		return 0, errors.New(fmt.Sprintf("[警告] 状态(%s)未定义任何事件", StatusText(f.getState())))
	}

	opt := new(Opt)
	for _, f := range opts {
		f(opt)
	}

	fn, ok := events[event]
	if !ok {
		return 0, errors.New(fmt.Sprintf("[警告] 状态(%s)不允许操作(%s)", StatusText(f.getState()), event))
	}

	status, err := fn(opt)
	if err != nil {
		return 0, err
	}

	oldState := f.getState()
	f.setState(status)
	newState := f.getState()

	fmt.Println(fmt.Sprintf("操作[%s]，状态从 [%s] 变成 [%s]", event, StatusText(oldState), StatusText(newState)))

	return f.getState(), nil
}

// NewFSM 实例化 FSM
func NewFSM(initState State) (fsm *FSM, err error) {
	fsm = new(FSM)
	fsm.state = initState
	fsm.handlers = make(map[State]map[Event]Handler)

	fsm, err = fsm.addHandlers()
	if err != nil {
		return
	}

	return
}
