package selector

// matchTags evaluates a tag-based Rule against a Context's tags.
func matchTags(rule *Rule, ctxTags map[string]string) bool {
	if len(rule.Tags) == 0 {
		return true // No tags means match, per your impl
	}

	switch rule.Operand {
	case "and":
		for k, v := range rule.Tags {
			if ctxV, ok := ctxTags[k]; !ok || ctxV != v {
				return false
			}
		}
		return true
	case "not":
		for k, v := range rule.Tags {
			if ctxV, ok := ctxTags[k]; ok && ctxV == v {
				return false
			}
		}
		return true
	case "or":
		for k, v := range rule.Tags {
			if ctxV, ok := ctxTags[k]; ok && ctxV == v {
				return true
			}
		}
		return false
	default:
		// Shouldn’t hit this (validated in ParseRule), but fallback to false
		return false
	}
}
