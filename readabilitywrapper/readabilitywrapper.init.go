package readabilitywrapper

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"readability-wrapper.ReadabilityProps",
		reflect.TypeOf((*ReadabilityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"readability-wrapper.ReadabilityResult",
		reflect.TypeOf((*ReadabilityResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"readability-wrapper.ReadabilityWrapper",
		reflect.TypeOf((*ReadabilityWrapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "parse", GoMethod: "Parse"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_ReadabilityWrapper{}
		},
	)
}
