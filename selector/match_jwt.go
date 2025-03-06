package selector

// matchJWT evaluates a JWT-based Rule against a Context's JWT audience.
func matchJWT(rule *Rule, jwtAud string) bool {
	if jwtAud == "" {
		return rule.Operand == "not" // No aud matches only if "not"
	}

	switch rule.Key {
	case "aud":
		switch rule.Operand {
		case "and":
			for _, rv := range rule.Values {
				if rv != jwtAud {
					return false
				}
			}
			return true
		case "not":
			for _, rv := range rule.Values {
				if rv == jwtAud {
					return false
				}
			}
			return true
		case "or":
			for _, rv := range rule.Values {
				if rv == jwtAud {
					return true
				}
			}
			return false
		default:
			return false
		}
	default:
		// Only "aud" supported for now, per your impl
		return false
	}
}
