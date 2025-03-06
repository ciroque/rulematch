package selector

// matchHeader evaluates a header-based Rule against a Context's headers.
func matchHeader(rule *Rule, headers map[string][]string) bool {
	vals, ok := headers[rule.Key]
	if !ok || len(vals) == 0 {
		return rule.Operand == "not" // No header matches only if "not"
	}

	switch rule.Operand {
	case "and":
		for _, rv := range rule.Values {
			found := false
			for _, hv := range vals {
				if hv == rv {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	case "not":
		for _, rv := range rule.Values {
			for _, hv := range vals {
				if hv == rv {
					return false
				}
			}
		}
		return true
	case "or":
		for _, rv := range rule.Values {
			for _, hv := range vals {
				if hv == rv {
					return true
				}
			}
		}
		return false
	default:
		// Shouldn’t hit this (validated in ParseRule), but fallback to false
		return false
	}
}
