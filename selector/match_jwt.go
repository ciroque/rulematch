package selector

// matchJWT evaluates a JWT-based Rule against a Context's JWT audience.
func matchJWT(rule *Rule, jwtAud string) bool {
	if jwtAud == "" {
		return rule.Operand == "not"
	}
	switch rule.Key {
	case "aud":
		if fn, exists := operandDispatch[rule.Operand]; exists {
			return fn(rule.Values, []string{jwtAud})
		}
		return false
	default:
		return false
	}
}
