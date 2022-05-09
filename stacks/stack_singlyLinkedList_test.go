package stacks

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateNewStack(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test create an empty stack",
			args: args{
				values: []interface{}{},
			},
			want: 0,
		},
		{
			name: "Test create a stack with 1 element",
			args: args{
				values: []interface{}{"1"},
			},
			want: 1,
		},
		{
			name: "Test create a stack with multiple element and types",
			args: args{
				values: []interface{}{"1", "abc", "cdf", 10, 0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateNewStack(tt.args.values...); got.Size() != tt.want {
				t.Errorf("CreateNewStack(): got = %v, want %v", got.Size(), tt.want)
			}
			// else {
			// 	t.Logf("CreateNewStack(): got = %v, want %v", got.Size(), tt.want)
			// }
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want int
	}{
		{
			name: "Test Size with empty stack",
			s:    CreateNewStack(),
			want: 0,
		},
		{
			name: "Test Size with 1 element stack",
			s:    CreateNewStack("abc"),
			want: 1,
		},
		{
			name: "Test Size with multiple element stack",
			s:    CreateNewStack("1", "abc", "cdf", 10, 0),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Stack.Size(): got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want bool
	}{
		{
			name: "Test stack empty with 0 elements",
			s:    CreateNewStack(),
			want: true,
		},
		{
			name: "Test stack empty with 1 elements",
			s:    CreateNewStack("abc"),
			want: false,
		},
		{
			name: "Test stack empty with multiple elements",
			s:    CreateNewStack("abc", "1", 10, "34", 45),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty(): got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		element []interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
		want int
	}{
		{
			name: "Push element in empty stack",
			s:    CreateNewStack(),
			args: args{
				element: nil,
			},
			want: 0,
		},
		{
			name: "Push element in stack with 1 element",
			s:    CreateNewStack(),
			args: args{
				element: []interface{}{"abc"},
			},
		},
		{
			name: "Push element in stack with multiple element",
			s:    CreateNewStack(),
			args: args{
				element: []interface{}{"abc", 1, 2, 3, "c", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Stack.Add(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name         string
		s            *Stack
		want         interface{}
		wantErr      error
		expectedSize int
	}{
		{
			name:         "Pop element in empty stack",
			s:            CreateNewStack(),
			want:         nil,
			wantErr:      errors.New("cannot remove element: empty stack"),
			expectedSize: 0,
		},
		{
			name:         "Pop element in a stack with 1 element",
			s:            CreateNewStack("a"),
			want:         "a",
			wantErr:      nil,
			expectedSize: 0, //expected size is 0 after pop of only 1 element
		},
		{
			name:         "Pop element in a stack with multiple element",
			s:            CreateNewStack("admin", "b", "c", 1, 87, 65, "45025"),
			want:         "45025", // expected Behavior LIFO: Last In First Out
			wantErr:      nil,
			expectedSize: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			size := tt.s.Size()
			if (err != nil) && (tt.wantErr.Error() != err.Error()) {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop(): got = %v, want %v", got, tt.want)
			}
			if size != tt.expectedSize {
				t.Errorf("Stack.Pop(): got size: %v, expected Size : %v", size, tt.expectedSize)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stack
		want    interface{}
		wantErr error
	}{
		{
			name:    "Peek test using 0 elements",
			s:       CreateNewStack(),
			want:    nil,
			wantErr: errors.New("cannot read element: empty stack"),
		},
		{
			name:    "Peek test using 1 elements",
			s:       CreateNewStack("abcde"),
			want:    "abcde",
			wantErr: nil,
		},
		{
			name:    "Peek test using multiple elements",
			s:       CreateNewStack("a", "b", 454, 87, "67", 45),
			want:    45,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Peek()
			if (err != nil) && (tt.wantErr.Error() != err.Error()) {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek(): got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Contains(t *testing.T) {
	type args struct {
		element interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
		want bool
	}{
		{
			name: "Test contains using empty stack",
			s:    CreateNewStack(),
			args: args{
				element: "a",
			},
			want: false,
		},
		{
			name: "Test contains using stack with some element",
			s:    CreateNewStack("a", "b", "v", 23, 24),
			args: args{
				element: "a",
			},
			want: true,
		},
		{
			name: "Test contains using stack with some element not presnt",
			s:    CreateNewStack("a", "b", "v", 23, 24),
			args: args{
				element: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.element); got != tt.want {
				t.Errorf("Stack.Contains(): got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_PrintStack(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
	}{
		{
			name: "Print Empty stack",
			s:    CreateNewStack(),
		},
		{
			name: "Print stack",
			s:    CreateNewStack("a", "b", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.PrintStack()
		})
	}
}

func BenchMarkPush(b *testing.B, s *Stack, values ...interface{}) {
	for i := 0; i < b.N; i++ {
		s.Push(values...)
	}
}

func BenchmarkPush(b *testing.B) {
	b.StopTimer()
	var values []interface{}
	for i := 0; i < 100; i++ {
		values = append(values, i)
	}
	b.StartTimer()
	BenchMarkPush(b, CreateNewStack(), values...)
}

func BenchMarkPop(b *testing.B, s *Stack) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < s.Size(); j++ {
			s.Pop()
		}
	}
}

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	s := CreateNewStack()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	b.StartTimer()
	BenchMarkPush(b, s)
}
