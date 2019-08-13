package output

import "plugin"

type (
    Output  plugin.Symbol
    Outputs map[string]plugin.Symbol
)

func NewOutputs() Outputs {
    return Outputs{}
}

func (o Outputs) Has(key string) bool {
    _, ok := o[key]
    return ok
}

func (o Outputs) Get(key string) plugin.Symbol {
    return o[key]
}
