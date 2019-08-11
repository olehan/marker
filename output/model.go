package output

import "plugin"

type (
    Output  plugin.Symbol
    Outputs map[string]plugin.Symbol
)

func NewOutputs() Outputs {
    return Outputs{}
}
