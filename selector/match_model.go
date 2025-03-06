package selector

// matchModel evaluates an input.model-based Rule against a Context's model.
func matchModel(rule *Rule, model string) bool {
	if model == "" {
		return rule.Operand == "not"
	}
	if fn, exists := operandDispatch[rule.Operand]; exists {
		return fn(rule.Values, []string{model})
	}
	return false
}
