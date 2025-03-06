package selector

import "net/http"

// Rule represents a single selector configuration from the YAML.
type Rule struct {
	Type    string            `yaml:"type,omitempty"`    // e.g., "header", "jwt", "input.model"; empty for tags
	Operand string            `yaml:"operand,omitempty"` // "and", "or", "not"; defaults to "any"
	Key     string            `yaml:"key,omitempty"`     // Used for "header" (header name) or "jwt" (claim name)
	Values  []string          `yaml:"values,omitempty"`  // Values to match against (for header, jwt, model)
	Tags    map[string]string `yaml:"tags,omitempty"`    // Tags to match (for tag selectors), e.g., "language": "en"
}

// Config represents a configuration object (service, stage, profile) with an attached selector.
type Config struct {
	Name string `yaml:"name"`
	Rule *Rule  `yaml:"selector,omitempty"`
}

// Context holds request data for selector matching.
type Context struct {
	Headers map[string][]string // HTTP headers from the request
	JWTAud  string              // JWT "aud" claim (extend for other claims if needed)
	Model   string              // Input model name
	Tags    map[string]string   // Tags added during processing, e.g., "language": "en"
}

// NewContext creates a new Context from an HTTP request and additional data.
func NewContext(r *http.Request, jwtAud, model string, tags map[string]string) *Context {
	headers := make(map[string][]string)
	for k, v := range r.Header {
		headers[k] = v
	}
	return &Context{
		Headers: headers,
		JWTAud:  jwtAud,
		Model:   model,
		Tags:    tags,
	}
}
