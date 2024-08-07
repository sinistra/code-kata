// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fetchermock

import (
	"code-kata/internal/fetcher"
	"sync"
)

// Ensure, that FetcherMock does implement fetcher.Fetcher.
// If this is not the case, regenerate this file with moq.
var _ fetcher.Fetcher = &FetcherMock{}

// FetcherMock is a mock implementation of fetcher.Fetcher.
//
//	func TestSomethingThatUsesFetcher(t *testing.T) {
//
//		// make and configure a mocked fetcher.Fetcher
//		mockedFetcher := &FetcherMock{
//			FetchFunc: func() ([]fetcher.Todo, error) {
//				panic("mock out the Fetch method")
//			},
//			GetFunc: func(id int)  {
//				panic("mock out the Get method")
//			},
//			ValidateFunc: func() error {
//				panic("mock out the Validate method")
//			},
//		}
//
//		// use mockedFetcher in code that requires fetcher.Fetcher
//		// and then make assertions.
//
//	}
type FetcherMock struct {
	// FetchFunc mocks the Fetch method.
	FetchFunc func() ([]fetcher.Todo, error)

	// GetFunc mocks the Get method.
	GetFunc func(id int)

	// ValidateFunc mocks the Validate method.
	ValidateFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Fetch holds details about calls to the Fetch method.
		Fetch []struct {
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// ID is the id argument value.
			ID int
		}
		// Validate holds details about calls to the Validate method.
		Validate []struct {
		}
	}
	lockFetch    sync.RWMutex
	lockGet      sync.RWMutex
	lockValidate sync.RWMutex
}

// Fetch calls FetchFunc.
func (mock *FetcherMock) Fetch() ([]fetcher.Todo, error) {
	if mock.FetchFunc == nil {
		panic("FetcherMock.FetchFunc: method is nil but Fetcher.Fetch was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFetch.Lock()
	mock.calls.Fetch = append(mock.calls.Fetch, callInfo)
	mock.lockFetch.Unlock()
	return mock.FetchFunc()
}

// FetchCalls gets all the calls that were made to Fetch.
// Check the length with:
//
//	len(mockedFetcher.FetchCalls())
func (mock *FetcherMock) FetchCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFetch.RLock()
	calls = mock.calls.Fetch
	mock.lockFetch.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *FetcherMock) Get(id int) {
	if mock.GetFunc == nil {
		panic("FetcherMock.GetFunc: method is nil but Fetcher.Get was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	mock.GetFunc(id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedFetcher.GetCalls())
func (mock *FetcherMock) GetCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Validate calls ValidateFunc.
func (mock *FetcherMock) Validate() error {
	if mock.ValidateFunc == nil {
		panic("FetcherMock.ValidateFunc: method is nil but Fetcher.Validate was just called")
	}
	callInfo := struct {
	}{}
	mock.lockValidate.Lock()
	mock.calls.Validate = append(mock.calls.Validate, callInfo)
	mock.lockValidate.Unlock()
	return mock.ValidateFunc()
}

// ValidateCalls gets all the calls that were made to Validate.
// Check the length with:
//
//	len(mockedFetcher.ValidateCalls())
func (mock *FetcherMock) ValidateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockValidate.RLock()
	calls = mock.calls.Validate
	mock.lockValidate.RUnlock()
	return calls
}
