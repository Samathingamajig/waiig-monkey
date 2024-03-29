package evaluator

import (
	"fmt"

	"github.com/Samathingamajig/waiig-monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}
				return arg.Elements[0]
			case *object.String:
				if len(arg.Value) == 0 {
					return NULL
				}
				return &object.String{Value: string(arg.Value[0])}
			default:
				return newError("argument to `first` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)
				if length == 0 {
					return NULL
				}
				return arg.Elements[length-1]
			case *object.String:
				length := len(arg.Value)
				if length == 0 {
					return NULL
				}
				return &object.String{Value: string(arg.Value[length-1])}
			default:
				return newError("argument to `last` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)
				if length == 0 {
					return NULL
				}
				newElements := make([]object.Object, length-1)
				copy(newElements, arg.Elements[1:length])
				return &object.Array{Elements: newElements}
			case *object.String:
				length := len(arg.Value)
				if length == 0 {
					return NULL
				}
				return &object.String{Value: arg.Value[1:]}
			default:
				return newError("argument to `rest` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			newValue := args[1]

			switch arg := args[0].(type) {
			case *object.Array:
				oldLength := len(arg.Elements)
				newElements := make([]object.Object, oldLength+1)
				copy(newElements, arg.Elements)
				newElements[oldLength] = newValue
				return &object.Array{Elements: newElements}
			default:
				return newError("argument to `push` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
