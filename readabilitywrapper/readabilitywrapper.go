// Wrapper around readability lib
package readabilitywrapper

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/Neurostep/readability-wrapper-go/readabilitywrapper/jsii"
)

type ReadabilityProps struct {
	Name *string `json:"name"`
}

type ReadabilityResult struct {
	Byline *string `json:"byline"`
	Content *string `json:"content"`
	Dir *string `json:"dir"`
	Excerpt *string `json:"excerpt"`
	Length *float64 `json:"length"`
	SiteName *string `json:"siteName"`
	TextContent *string `json:"textContent"`
	Title *string `json:"title"`
}

type ReadabilityWrapper interface {
	Props() *ReadabilityProps
	SetProps(val *ReadabilityProps)
	Parse(article *string, url *string) *ReadabilityResult
}

// The jsii proxy struct for ReadabilityWrapper
type jsiiProxy_ReadabilityWrapper struct {
	_ byte // padding
}

func (j *jsiiProxy_ReadabilityWrapper) Props() *ReadabilityProps {
	var returns *ReadabilityProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewReadabilityWrapper(props *ReadabilityProps) ReadabilityWrapper {
	_init_.Initialize()

	j := jsiiProxy_ReadabilityWrapper{}

	_jsii_.Create(
		"readability-wrapper.ReadabilityWrapper",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewReadabilityWrapper_Override(r ReadabilityWrapper, props *ReadabilityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"readability-wrapper.ReadabilityWrapper",
		[]interface{}{props},
		r,
	)
}

func (j *jsiiProxy_ReadabilityWrapper) SetProps(val *ReadabilityProps) {
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (r *jsiiProxy_ReadabilityWrapper) Parse(article *string, url *string) *ReadabilityResult {
	var returns *ReadabilityResult

	_jsii_.Invoke(
		r,
		"parse",
		[]interface{}{article, url},
		&returns,
	)

	return returns
}

