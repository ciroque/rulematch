package selector

// matchModel evaluates an input.model-based Rule against a Context's model.
func matchModel(rule *Rule, model string) bool {
	if model == "" {
		return rule.Operand == "not" // No model matches only if "not"
	}

	switch rule.Operand {
	case "and":
		for _, rv := range rule.Values {
			if rv != model {
				return false
			}
		}
		return true
	case "not":
		for _, rv := range rule.Values {
			if rv == model {
				return false
			}
		}
		return true
	case "or":
		for _, rv := range rule.Values {
			if rv == model {
				return true
			}
		}
		return false
	default:
		// Shouldn’t hit this (validated in ParseRule), but fallback to false
		return false
	}
}
