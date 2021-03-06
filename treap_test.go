package treap

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func TestNewTreap(t *testing.T) {
	tests := []struct {
		name string
		want *Treap
	}{
		{
			name: "new treap",
			want: &Treap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewTreap())
		})
	}
}

func TestTreap_Search(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "search for value in empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value: "h",
			},
			want: false,
		},
		{
			name: "search for non-existing value in single value treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
				},
			},
			args: args{
				value: "h",
			},
			want: false,
		},
		{
			name: "search for existing value in single value treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
				},
			},
			args: args{
				value: "f",
			},
			want: true,
		},
		{
			name: "search for existing value in populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "h",
			},
			want: true,
		},
		{
			name: "search for nonexistent value in populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			trp := &Treap{
				root: tt.fields.root,
			}
			assert.Equal(t, tt.want, trp.Search(tt.args.value))
		})
	}
}

func TestTreap_Insert(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "insert value into an empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value: "c",
			},
		},
		{
			name: "insert value into left of treap with higher priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 1,
				},
			},
			args: args{
				value: "a",
			},
		},
		{
			name: "insert value into left of treap with lower priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 2,
				},
			},
			args: args{
				value: "a",
			},
		},
		{
			name: "insert value into right of treap with higher priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 1,
				},
			},
			args: args{
				value: "d",
			},
		},
		{
			name: "insert value into right of treap with lower priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 2,
				},
			},
			args: args{
				value: "d",
			},
		},
		{
			name: "insert value into populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "k",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trp := NewTreap()
			trp.root = tt.fields.root
			trp.Insert(tt.args.value)
			assert.True(t, trp.Search(tt.args.value))
		})
	}
}

func TestTreap_insert(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value    string
		priority int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		{
			name: "insert value into an empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value:    "c",
				priority: 1,
			},
			want: &node{
				value:    "c",
				priority: 1,
			},
		},
		{
			name: "insert value into left of treap with higher priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 1,
				},
			},
			args: args{
				value:    "a",
				priority: 2,
			},
			want: &node{
				value:    "a",
				priority: 2,
				right: &node{
					value:    "c",
					priority: 1,
				},
			},
		},
		{
			name: "insert value into left of treap with lower priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 2,
				},
			},
			args: args{
				value:    "a",
				priority: 1,
			},
			want: &node{
				value:    "c",
				priority: 2,
				left: &node{
					value:    "a",
					priority: 1,
				},
			},
		},
		{
			name: "insert value into right of treap with higher priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 1,
				},
			},
			args: args{
				value:    "d",
				priority: 2,
			},
			want: &node{
				value:    "d",
				priority: 2,
				left: &node{
					value:    "c",
					priority: 1,
				},
			},
		},
		{
			name: "insert value into right of treap with lower priority",
			fields: fields{
				root: &node{
					value:    "c",
					priority: 2,
				},
			},
			args: args{
				value:    "d",
				priority: 1,
			},
			want: &node{
				value:    "c",
				priority: 2,
				right: &node{
					value:    "d",
					priority: 1,
				},
			},
		},
		{
			name: "insert value into populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value:    "k",
				priority: 5,
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "k",
						priority: 5,
						left: &node{
							value:    "h",
							priority: 3,
						},
					},
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trp := &Treap{
				root: tt.fields.root,
			}
			trp.root = insert(trp.root, tt.args.value, tt.args.priority)
			assert.Equal(t, tt.want, trp.root)
			assert.True(t, trp.Search(tt.args.value))
		})
	}
}

func BenchmarkInsert(b *testing.B) {
	trp := NewTreap()
	for i := 0; i < b.N; i++ {
		trp.Insert(string(alpha[i%len(alpha)]))
	}
}

func TestTreap_Delete(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete value from empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value: "a",
			},
		},
		{
			name: "delete non-existing value from treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
				},
			},
			args: args{
				value: "a",
			},
		},
		{
			name: "delete non-existing value from treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
				},
			},
			args: args{
				value: "f",
			},
		},
		{
			name: "delete existing value from treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
				},
			},
			args: args{
				value: "d",
			},
		},
		{
			name: "delete existing value from left side of treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
					left: &node{
						value:    "b",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 3,
					},
				},
			},
			args: args{
				value: "b",
			},
		},
		{
			name: "delete existing value from right side of treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
					left: &node{
						value:    "b",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 3,
					},
				},
			},
			args: args{
				value: "e",
			},
		},
		{
			name: "delete subtree value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "t",
			},
		},
		{
			name: "delete right leaf value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "e",
			},
		},
		{
			name: "delete left leaf value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "h",
			},
		},
		{
			name: "delete root value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "f",
			},
		},
		{
			name: "delete non-existent value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "b",
			},
		},
		{
			name: "delete non-existent value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "z",
			},
		},
		{
			name: "delete non-existent value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "i",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trp := NewTreap()
			trp.root = tt.fields.root
			trp.Delete(tt.args.value)
			assert.False(t, trp.Search(tt.args.value))
		})
	}
}

func TestTreap_delete(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		{
			name: "delete value from empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value: "a",
			},
			want: nil,
		},
		{
			name: "delete non-existing value from treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
				},
			},
			args: args{
				value: "a",
			},
			want: &node{
				value:    "d",
				priority: 4,
			},
		},
		{
			name: "delete non-existing value from treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
				},
			},
			args: args{
				value: "f",
			},
			want: &node{
				value:    "d",
				priority: 4,
			},
		},
		{
			name: "delete existing value from treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
				},
			},
			args: args{
				value: "d",
			},
			want: nil,
		},
		{
			name: "delete existing value from left side of treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
					left: &node{
						value:    "b",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 3,
					},
				},
			},
			args: args{
				value: "b",
			},
			want: &node{
				value:    "d",
				priority: 4,
				right: &node{
					value:    "e",
					priority: 3,
				},
			},
		},
		{
			name: "delete existing value from right side of treap",
			fields: fields{
				root: &node{
					value:    "d",
					priority: 4,
					left: &node{
						value:    "b",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 3,
					},
				},
			},
			args: args{
				value: "e",
			},
			want: &node{
				value:    "d",
				priority: 4,
				left: &node{
					value:    "b",
					priority: 2,
				},
			},
		},
		{
			name: "delete subtree value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "t",
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "x",
					priority: 6,
					left: &node{
						value:    "h",
						priority: 3,
					},
				},
			},
		},
		{
			name: "delete right leaf value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "e",
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "h",
						priority: 3,
					},
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
		{
			name: "delete left leaf value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "h",
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
		{
			name: "delete root value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "f",
			},
			want: &node{
				value:    "d",
				priority: 8,
				left: &node{
					value:    "c",
					priority: 2,
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "h",
						priority: 3,
						left: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
		{
			name: "delete non-existent value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "b",
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "h",
						priority: 3,
					},
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
		{
			name: "delete non-existent value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "z",
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "h",
						priority: 3,
					},
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
		{
			name: "delete non-existent value from populated treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
						},
					},
				},
			},
			args: args{
				value: "i",
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "h",
						priority: 3,
					},
					right: &node{
						value:    "x",
						priority: 6,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trp := NewTreap()
			trp.root = delete(tt.fields.root, tt.args.value)
			assert.Equal(t, tt.want, trp.root)
			assert.False(t, trp.Search(tt.args.value))
		})
	}
}

func Test_rotateRight(t *testing.T) {
	type args struct {
		root *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "rotate tree right",
			args: args{
				root: &node{
					value:    "5",
					priority: 10,
					right: &node{
						value:    "7",
						priority: 9,
					},
					left: &node{
						value:    "3",
						priority: 8,
						right: &node{
							value:    "4",
							priority: 5,
						},
						left: &node{
							value:    "2",
							priority: 3,
						},
					},
				},
			},
			want: &node{
				value:    "3",
				priority: 8,
				right: &node{
					value:    "5",
					priority: 10,
					right: &node{
						value:    "7",
						priority: 9,
					},
					left: &node{
						value:    "4",
						priority: 5,
					},
				},
				left: &node{
					value:    "2",
					priority: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert returned value is the new root after the rotation
			assert.Equal(t, tt.want, rotateRight(tt.args.root, tt.args.root.left))
		})
	}
}

func Test_rotateLeft(t *testing.T) {
	type args struct {
		root  *node
		pivot *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "rotate tree left",
			args: args{
				root: &node{
					value:    "3",
					priority: 8,
					right: &node{
						value:    "5",
						priority: 10,
						right: &node{
							value:    "7",
							priority: 9,
						},
						left: &node{
							value:    "4",
							priority: 5,
						},
					},
					left: &node{
						value:    "2",
						priority: 3,
					},
				},
			},
			want: &node{
				value:    "5",
				priority: 10,
				right: &node{
					value:    "7",
					priority: 9,
				},
				left: &node{
					value:    "3",
					priority: 8,
					right: &node{
						value:    "4",
						priority: 5,
					},
					left: &node{
						value:    "2",
						priority: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert returned value is the new root after the rotation
			assert.Equal(t, tt.want, rotateLeft(tt.args.root, tt.args.root.right))
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		n     *node
		value string
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "binary search for value",
			args: args{
				n: &node{
					value: "abc",
				},
				value: "abc",
			},
			want: &node{
				value: "abc",
			},
		},
		{
			name: "binary search for value in root of tree",
			args: args{
				n: &node{
					value: "b",
					right: &node{
						value: "c",
					},
					left: &node{
						value: "a",
					},
				},
				value: "b",
			},
			want: &node{
				value: "b",
				right: &node{
					value: "c",
				},
				left: &node{
					value: "a",
				},
			},
		},
		{
			name: "binary search for value on right of tree",
			args: args{
				n: &node{
					value: "b",
					right: &node{
						value: "c",
					},
					left: &node{
						value: "a",
					},
				},
				value: "c",
			},
			want: &node{
				value: "c",
			},
		},
		{
			name: "binary search for value on left of tree",
			args: args{
				n: &node{
					value: "b",
					right: &node{
						value: "c",
					},
					left: &node{
						value: "a",
					},
				},
				value: "a",
			},
			want: &node{
				value: "a",
			},
		},
		{
			name: "binary search for value not in tree",
			args: args{
				n: &node{
					value: "abcd",
				},
				value: "abc",
			},
			want: nil,
		},
		{
			name: "binary search for value in empty tree",
			args: args{
				n:     nil,
				value: "z",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, binarySearch(tt.args.n, tt.args.value))
		})
	}
}

func TestTreapMixedOps(t *testing.T) {
	trp := NewTreap()

	// Fill the treap up with random strings
	inserted := fillTree(trp, 10000)
	assert.True(t, hasTreapProperties(trp.root))

	// For each random string inserted
	for k := range inserted {
		// Assert that it's in the treap
		assert.True(t, trp.Search(k))

		// Delete the value from the treap
		trp.Delete(k)

		// Assert that treap properties are still true and
		// the value is no longer in the treap
		assert.True(t, hasTreapProperties(trp.root))
		assert.False(t, trp.Search(k))
	}

	// Assert that all values that were inserted are now deleted
	assert.Nil(t, trp.root)
}

func fillTree(trp *Treap, count int) map[string]bool {
	inserted := make(map[string]bool, count)
	for i := 0; i < count; i++ {
		var b bytes.Buffer
		for j := 0; j < rand.Intn(len(alpha)); j++ {
			rand.Seed(time.Now().UnixNano())
			b.WriteByte(alpha[rand.Intn(len(alpha))])
		}

		randStr := b.String()
		inserted[randStr] = true
		trp.Insert(randStr)
	}

	return inserted
}

// hasTreapProperties returns true if the passed tree has both
// binary search tree properties and max heap properties.
func hasTreapProperties(root *node) bool {
	if root == nil {
		return true
	}

	isValid := true
	if root.left != nil &&
		(root.left.value > root.value || root.left.priority > root.priority) {
		isValid = false
	}
	if root.right != nil &&
		(root.right.value < root.value || root.right.priority > root.priority) {
		isValid = false
	}

	return isValid &&
		hasTreapProperties(root.left) &&
		hasTreapProperties(root.right)
}
