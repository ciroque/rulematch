package selector

var operandDispatch = map[string]func(left, right []string) bool{
	"and": and,
	"not": not,
	"or":  or,
}

// Match evaluates a Rule against a Context, returning true if it matches.
func Match(rule *Rule, ctx *Context) bool {
	if rule == nil {
		return true // No rule means always match, per your docs
	}

	switch rule.Type {
	case "header":
		return matchHeader(rule, ctx.Headers)

	case "jwt":
		// Only aud for now, extend later if needed
		return matchJWT(rule, ctx.JWTAud)

	case "input.model":
		return matchModel(rule, ctx.Model)

	case "":
		return matchTags(rule, ctx.Tags)

	default:
		// Unknown types fail silently (could log or error if desired)
		return false
	}
}

// Evaluate selects the first Config from a list that matches the Context.
func Evaluate(configs []*Config, ctx *Context) (*Config, error) {
	for _, cfg := range configs {
		if Match(cfg.Rule, ctx) {
			return cfg, nil
		}
	}
	return nil, nil // No match, no error (caller decides what’s next)
}
