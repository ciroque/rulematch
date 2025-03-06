package selector

// matchHeader evaluates a header-based Rule against a Context's headers.
func matchHeader(rule *Rule, headers map[string][]string) bool {
	vals, ok := headers[rule.Key]
	if !ok || len(vals) == 0 {
		return rule.Operand == "not"
	}
	if fn, exists := operandDispatch[rule.Operand]; exists {
		return fn(rule.Values, vals)
	}
	return false
}
