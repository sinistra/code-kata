package fetcher

import (
	"reflect"
	"testing"
)

func TestNewTodoHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		collection string
		count      int
	}{
		{
			name:       "success",
			url:        "",
			collection: "",
			count:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTodoHandler(tt.url, tt.collection, tt.count)
			_, ok := interface{}(got).(Fetcher)
			if !ok {
				t.Error("NewTodoHandler() is not type Fetcher")
			}
		})
	}
}

func Test_todoHandler_Fetch(t1 *testing.T) {
	tests := []struct {
		name       string
		url        string
		count      int
		collection string
		want       []Todo
		wantErr    bool
	}{
		{
			name:       "success",
			url:        "https://jsonplaceholder.typicode.com/todos",
			count:      1,
			collection: "e",
			want: []Todo{
				{
					Title:     "quis ut nam facilis et officia qui",
					UserID:    1,
					ID:        2,
					Completed: false,
					Error:     "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &todoHandler{
				url:        tt.url,
				count:      tt.count,
				collection: tt.collection,
				resultChan: make(chan Todo, tt.count),
			}
			got, err := t.Fetch()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoHandler_Validate(t1 *testing.T) {
	tests := []struct {
		name       string
		url        string
		count      int
		collection string
		wantErr    bool
	}{
		{
			name:       "success",
			url:        "https://jsonplaceholder.typicode.com/todos",
			count:      1,
			collection: "e",
			wantErr:    false,
		}, {
			name:       "success_slash",
			url:        "https://jsonplaceholder.typicode.com/todos/",
			count:      1,
			collection: "e",
			wantErr:    false,
		},
		{
			name:       "count_fail",
			url:        "https://jsonplaceholder.typicode.com/todos",
			count:      0,
			collection: "e",
			wantErr:    true,
		},
		{
			name:       "collection_fail",
			url:        "https://jsonplaceholder.typicode.com/todos",
			count:      1,
			collection: "x",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &todoHandler{
				url:        tt.url,
				count:      tt.count,
				collection: tt.collection,
				resultChan: make(chan Todo, tt.count),
			}
			if err := t.Validate(); (err != nil) != tt.wantErr {
				t1.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
