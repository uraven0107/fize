package controller

type Parameter struct {
	path string
}

type ParameterBuilder struct {
	param Parameter
}

func NewParameterBuilder() ParameterBuilder {
	return ParameterBuilder{
		param: Parameter{},
	}
}

func (builder ParameterBuilder) Path(path string) ParameterBuilder {
	builder.param.path = path
	return builder
}

func (builder ParameterBuilder) Build() Parameter {
	return builder.param
}
