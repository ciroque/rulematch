package selector

// matchTags evaluates a tag-based Rule against a Context's tags.
func matchTags(rule *Rule, ctxTags map[string]string) bool {
	if len(rule.Tags) == 0 {
		return true
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
		return false
	}
}
