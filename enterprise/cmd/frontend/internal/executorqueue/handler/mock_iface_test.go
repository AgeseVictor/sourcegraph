// Code generated by go-mockgen 1.3.1; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the metadata.yaml file in the root of this repository.

package handler

import (
	"context"
	"sync"
	"time"

	store "github.com/sourcegraph/sourcegraph/internal/services/executors/store"
	types "github.com/sourcegraph/sourcegraph/internal/types"
)

// MockStore is a mock implementation of the Store interface (from the
// package
// github.com/sourcegraph/sourcegraph/internal/services/executors/store)
// used for unit testing.
type MockStore struct {
	// DeleteInactiveHeartbeatsFunc is an instance of a mock function object
	// controlling the behavior of the method DeleteInactiveHeartbeats.
	DeleteInactiveHeartbeatsFunc *StoreDeleteInactiveHeartbeatsFunc
	// GetByHostnameFunc is an instance of a mock function object
	// controlling the behavior of the method GetByHostname.
	GetByHostnameFunc *StoreGetByHostnameFunc
	// GetByIDFunc is an instance of a mock function object controlling the
	// behavior of the method GetByID.
	GetByIDFunc *StoreGetByIDFunc
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *StoreListFunc
	// UpsertHeartbeatFunc is an instance of a mock function object
	// controlling the behavior of the method UpsertHeartbeat.
	UpsertHeartbeatFunc *StoreUpsertHeartbeatFunc
}

// NewMockStore creates a new mock of the Store interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStore() *MockStore {
	return &MockStore{
		DeleteInactiveHeartbeatsFunc: &StoreDeleteInactiveHeartbeatsFunc{
			defaultHook: func(context.Context, time.Duration) (r0 error) {
				return
			},
		},
		GetByHostnameFunc: &StoreGetByHostnameFunc{
			defaultHook: func(context.Context, string) (r0 types.Executor, r1 bool, r2 error) {
				return
			},
		},
		GetByIDFunc: &StoreGetByIDFunc{
			defaultHook: func(context.Context, int) (r0 types.Executor, r1 bool, r2 error) {
				return
			},
		},
		ListFunc: &StoreListFunc{
			defaultHook: func(context.Context, store.ExecutorStoreListOptions) (r0 []types.Executor, r1 int, r2 error) {
				return
			},
		},
		UpsertHeartbeatFunc: &StoreUpsertHeartbeatFunc{
			defaultHook: func(context.Context, types.Executor) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockStore creates a new mock of the Store interface. All methods
// panic on invocation, unless overwritten.
func NewStrictMockStore() *MockStore {
	return &MockStore{
		DeleteInactiveHeartbeatsFunc: &StoreDeleteInactiveHeartbeatsFunc{
			defaultHook: func(context.Context, time.Duration) error {
				panic("unexpected invocation of MockStore.DeleteInactiveHeartbeats")
			},
		},
		GetByHostnameFunc: &StoreGetByHostnameFunc{
			defaultHook: func(context.Context, string) (types.Executor, bool, error) {
				panic("unexpected invocation of MockStore.GetByHostname")
			},
		},
		GetByIDFunc: &StoreGetByIDFunc{
			defaultHook: func(context.Context, int) (types.Executor, bool, error) {
				panic("unexpected invocation of MockStore.GetByID")
			},
		},
		ListFunc: &StoreListFunc{
			defaultHook: func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error) {
				panic("unexpected invocation of MockStore.List")
			},
		},
		UpsertHeartbeatFunc: &StoreUpsertHeartbeatFunc{
			defaultHook: func(context.Context, types.Executor) error {
				panic("unexpected invocation of MockStore.UpsertHeartbeat")
			},
		},
	}
}

// NewMockStoreFrom creates a new mock of the MockStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreFrom(i store.Store) *MockStore {
	return &MockStore{
		DeleteInactiveHeartbeatsFunc: &StoreDeleteInactiveHeartbeatsFunc{
			defaultHook: i.DeleteInactiveHeartbeats,
		},
		GetByHostnameFunc: &StoreGetByHostnameFunc{
			defaultHook: i.GetByHostname,
		},
		GetByIDFunc: &StoreGetByIDFunc{
			defaultHook: i.GetByID,
		},
		ListFunc: &StoreListFunc{
			defaultHook: i.List,
		},
		UpsertHeartbeatFunc: &StoreUpsertHeartbeatFunc{
			defaultHook: i.UpsertHeartbeat,
		},
	}
}

// StoreDeleteInactiveHeartbeatsFunc describes the behavior when the
// DeleteInactiveHeartbeats method of the parent MockStore instance is
// invoked.
type StoreDeleteInactiveHeartbeatsFunc struct {
	defaultHook func(context.Context, time.Duration) error
	hooks       []func(context.Context, time.Duration) error
	history     []StoreDeleteInactiveHeartbeatsFuncCall
	mutex       sync.Mutex
}

// DeleteInactiveHeartbeats delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockStore) DeleteInactiveHeartbeats(v0 context.Context, v1 time.Duration) error {
	r0 := m.DeleteInactiveHeartbeatsFunc.nextHook()(v0, v1)
	m.DeleteInactiveHeartbeatsFunc.appendCall(StoreDeleteInactiveHeartbeatsFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// DeleteInactiveHeartbeats method of the parent MockStore instance is
// invoked and the hook queue is empty.
func (f *StoreDeleteInactiveHeartbeatsFunc) SetDefaultHook(hook func(context.Context, time.Duration) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteInactiveHeartbeats method of the parent MockStore instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *StoreDeleteInactiveHeartbeatsFunc) PushHook(hook func(context.Context, time.Duration) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreDeleteInactiveHeartbeatsFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, time.Duration) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreDeleteInactiveHeartbeatsFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, time.Duration) error {
		return r0
	})
}

func (f *StoreDeleteInactiveHeartbeatsFunc) nextHook() func(context.Context, time.Duration) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreDeleteInactiveHeartbeatsFunc) appendCall(r0 StoreDeleteInactiveHeartbeatsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreDeleteInactiveHeartbeatsFuncCall
// objects describing the invocations of this function.
func (f *StoreDeleteInactiveHeartbeatsFunc) History() []StoreDeleteInactiveHeartbeatsFuncCall {
	f.mutex.Lock()
	history := make([]StoreDeleteInactiveHeartbeatsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreDeleteInactiveHeartbeatsFuncCall is an object that describes an
// invocation of method DeleteInactiveHeartbeats on an instance of
// MockStore.
type StoreDeleteInactiveHeartbeatsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 time.Duration
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreDeleteInactiveHeartbeatsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreDeleteInactiveHeartbeatsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StoreGetByHostnameFunc describes the behavior when the GetByHostname
// method of the parent MockStore instance is invoked.
type StoreGetByHostnameFunc struct {
	defaultHook func(context.Context, string) (types.Executor, bool, error)
	hooks       []func(context.Context, string) (types.Executor, bool, error)
	history     []StoreGetByHostnameFuncCall
	mutex       sync.Mutex
}

// GetByHostname delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockStore) GetByHostname(v0 context.Context, v1 string) (types.Executor, bool, error) {
	r0, r1, r2 := m.GetByHostnameFunc.nextHook()(v0, v1)
	m.GetByHostnameFunc.appendCall(StoreGetByHostnameFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the GetByHostname method
// of the parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreGetByHostnameFunc) SetDefaultHook(hook func(context.Context, string) (types.Executor, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByHostname method of the parent MockStore instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreGetByHostnameFunc) PushHook(hook func(context.Context, string) (types.Executor, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetByHostnameFunc) SetDefaultReturn(r0 types.Executor, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, string) (types.Executor, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetByHostnameFunc) PushReturn(r0 types.Executor, r1 bool, r2 error) {
	f.PushHook(func(context.Context, string) (types.Executor, bool, error) {
		return r0, r1, r2
	})
}

func (f *StoreGetByHostnameFunc) nextHook() func(context.Context, string) (types.Executor, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetByHostnameFunc) appendCall(r0 StoreGetByHostnameFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetByHostnameFuncCall objects
// describing the invocations of this function.
func (f *StoreGetByHostnameFunc) History() []StoreGetByHostnameFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetByHostnameFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetByHostnameFuncCall is an object that describes an invocation of
// method GetByHostname on an instance of MockStore.
type StoreGetByHostnameFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 types.Executor
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetByHostnameFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetByHostnameFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// StoreGetByIDFunc describes the behavior when the GetByID method of the
// parent MockStore instance is invoked.
type StoreGetByIDFunc struct {
	defaultHook func(context.Context, int) (types.Executor, bool, error)
	hooks       []func(context.Context, int) (types.Executor, bool, error)
	history     []StoreGetByIDFuncCall
	mutex       sync.Mutex
}

// GetByID delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) GetByID(v0 context.Context, v1 int) (types.Executor, bool, error) {
	r0, r1, r2 := m.GetByIDFunc.nextHook()(v0, v1)
	m.GetByIDFunc.appendCall(StoreGetByIDFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the GetByID method of
// the parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreGetByIDFunc) SetDefaultHook(hook func(context.Context, int) (types.Executor, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByID method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreGetByIDFunc) PushHook(hook func(context.Context, int) (types.Executor, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetByIDFunc) SetDefaultReturn(r0 types.Executor, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, int) (types.Executor, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetByIDFunc) PushReturn(r0 types.Executor, r1 bool, r2 error) {
	f.PushHook(func(context.Context, int) (types.Executor, bool, error) {
		return r0, r1, r2
	})
}

func (f *StoreGetByIDFunc) nextHook() func(context.Context, int) (types.Executor, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetByIDFunc) appendCall(r0 StoreGetByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetByIDFuncCall objects describing the
// invocations of this function.
func (f *StoreGetByIDFunc) History() []StoreGetByIDFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetByIDFuncCall is an object that describes an invocation of method
// GetByID on an instance of MockStore.
type StoreGetByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 types.Executor
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// StoreListFunc describes the behavior when the List method of the parent
// MockStore instance is invoked.
type StoreListFunc struct {
	defaultHook func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error)
	hooks       []func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error)
	history     []StoreListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) List(v0 context.Context, v1 store.ExecutorStoreListOptions) ([]types.Executor, int, error) {
	r0, r1, r2 := m.ListFunc.nextHook()(v0, v1)
	m.ListFunc.appendCall(StoreListFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreListFunc) SetDefaultHook(hook func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreListFunc) PushHook(hook func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreListFunc) SetDefaultReturn(r0 []types.Executor, r1 int, r2 error) {
	f.SetDefaultHook(func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreListFunc) PushReturn(r0 []types.Executor, r1 int, r2 error) {
	f.PushHook(func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error) {
		return r0, r1, r2
	})
}

func (f *StoreListFunc) nextHook() func(context.Context, store.ExecutorStoreListOptions) ([]types.Executor, int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreListFunc) appendCall(r0 StoreListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreListFuncCall objects describing the
// invocations of this function.
func (f *StoreListFunc) History() []StoreListFuncCall {
	f.mutex.Lock()
	history := make([]StoreListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreListFuncCall is an object that describes an invocation of method
// List on an instance of MockStore.
type StoreListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.ExecutorStoreListOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []types.Executor
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 int
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// StoreUpsertHeartbeatFunc describes the behavior when the UpsertHeartbeat
// method of the parent MockStore instance is invoked.
type StoreUpsertHeartbeatFunc struct {
	defaultHook func(context.Context, types.Executor) error
	hooks       []func(context.Context, types.Executor) error
	history     []StoreUpsertHeartbeatFuncCall
	mutex       sync.Mutex
}

// UpsertHeartbeat delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) UpsertHeartbeat(v0 context.Context, v1 types.Executor) error {
	r0 := m.UpsertHeartbeatFunc.nextHook()(v0, v1)
	m.UpsertHeartbeatFunc.appendCall(StoreUpsertHeartbeatFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the UpsertHeartbeat
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreUpsertHeartbeatFunc) SetDefaultHook(hook func(context.Context, types.Executor) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UpsertHeartbeat method of the parent MockStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreUpsertHeartbeatFunc) PushHook(hook func(context.Context, types.Executor) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreUpsertHeartbeatFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, types.Executor) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreUpsertHeartbeatFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, types.Executor) error {
		return r0
	})
}

func (f *StoreUpsertHeartbeatFunc) nextHook() func(context.Context, types.Executor) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreUpsertHeartbeatFunc) appendCall(r0 StoreUpsertHeartbeatFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreUpsertHeartbeatFuncCall objects
// describing the invocations of this function.
func (f *StoreUpsertHeartbeatFunc) History() []StoreUpsertHeartbeatFuncCall {
	f.mutex.Lock()
	history := make([]StoreUpsertHeartbeatFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreUpsertHeartbeatFuncCall is an object that describes an invocation of
// method UpsertHeartbeat on an instance of MockStore.
type StoreUpsertHeartbeatFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 types.Executor
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreUpsertHeartbeatFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreUpsertHeartbeatFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
