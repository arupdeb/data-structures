package queue

import (
	"errors"
	"log"
	"reflect"
	"testing"
)

func TestCreateNewQueue(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test empty queue create",
			args: args{
				values: []interface{}{},
			},
			want: 0,
		},
		{name: "Test queue create with 1 element",
			args: args{
				values: []interface{}{"abcd"},
			},
			want: 1,
		},
		{name: "Test queue create with multiple element",
			args: args{
				values: []interface{}{"abcd", 1, 56, "45", "vc", 1092783635},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateNewQueue(tt.args.values...); got.Size() != tt.want {
				t.Errorf("CreateNewQueue(): got = %v, want %v", got.Size(), tt.want)
			}
		})
	}
}

func Test_queue_Size(t *testing.T) {
	tests := []struct {
		name string
		q    *queue
		want int
	}{
		{
			name: "Test Size with empty queue",
			q:    CreateNewQueue(),
			want: 0,
		},
		{
			name: "Test Size with 1 element queue",
			q:    CreateNewQueue("1"),
			want: 1,
		},
		{
			name: "Test Size with multi element queue",
			q:    CreateNewQueue("1", "avc", 23, 45, "df", 76),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Size(); got != tt.want {
				t.Errorf("queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Enqueue(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		q    *queue
		args args
		want int
	}{
		{
			name: "test Enqueue with 0 element queue",
			q:    CreateNewQueue(),
			args: args{
				values: nil,
			},
			want: 0,
		},
		{
			name: "test Enqueue with 1 element queue",
			q:    CreateNewQueue(),
			args: args{
				values: []interface{}{"a"},
			},
			want: 1,
		},
		{
			name: "test Enqueue with multi element queue",
			q:    CreateNewQueue(),
			args: args{
				values: []interface{}{"abc", "f", "23", 23, 54, "c", 67.546},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Enqueue(tt.args.values...)
			if got := tt.q.Size(); got != tt.want {
				t.Errorf("queue.Enqueue(): got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		q       *queue
		want    interface{}
		wantErr error
	}{
		{
			name:    "test dequeue with 0 element",
			q:       CreateNewQueue(),
			want:    nil,
			wantErr: errors.New("cannot dequeue elememt: empty queue"),
		},
		{
			name:    "test dequeue with 1 element",
			q:       CreateNewQueue("a"),
			want:    "a",
			wantErr: nil,
		},
		{
			name:    "test dequeue with multiple element",
			q:       CreateNewQueue("a45", "b", "c"),
			want:    "a45",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Peek(t *testing.T) {
	tests := []struct {
		name    string
		q       *queue
		want    interface{}
		wantErr error
	}{
		{
			name:    "Test Peek with 0 element",
			q:       CreateNewQueue(),
			want:    nil,
			wantErr: errors.New("cannot peek elememt: empty queue"),
		},
		{
			name:    "Test Peek with 1 element",
			q:       CreateNewQueue(234),
			want:    234,
			wantErr: nil,
		},
		{
			name:    "Test Peek with multi element",
			q:       CreateNewQueue("fgh", "werf", 3, "c", 56785),
			want:    "fgh",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Peek()
			if (err != nil) && (tt.wantErr.Error() != err.Error()) {
				t.Errorf("queue.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Remove(t *testing.T) {
	type args struct {
		element interface{}
	}
	tests := []struct {
		name    string
		q       *queue
		args    args
		wantErr error
	}{
		{
			name: "Test remove with 0 element",
			q:    CreateNewQueue(),
			args: args{
				element: "abc",
			},
			wantErr: errors.New("cannot remove elememt: empty queue"),
		},
		{
			name: "Test remove with 1 element and existing",
			q:    CreateNewQueue("abc"),
			args: args{
				element: "abc",
			},
			wantErr: nil,
		},
		{
			name: "Test remove with 1 element and non-existing",
			q:    CreateNewQueue("abc"),
			args: args{
				element: "cdf",
			},
			wantErr: errors.New("cannot find element to remove"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.PrintQueue()
			err := tt.q.Remove(tt.args.element)
			log.Println(err)
			if (err != nil) && (err.Error() != tt.wantErr.Error()) {
				t.Errorf("queue.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_queue_PrintQueue(t *testing.T) {
	tests := []struct {
		name string
		q    *queue
	}{
		{
			name: "Print Empty Queue",
			q: CreateNewQueue(),
		},
		{
			name: "Print Queue 1 element",
			q: CreateNewQueue("jsdhsp"),
		},
		{
			name: "Print Queue with multiple element",
			q: CreateNewQueue("a", "b","c", 1,3,7653, "abc"),
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.PrintQueue()
		})
	}
}
