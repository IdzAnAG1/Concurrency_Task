package variables

const (
	USER_STRUCT      = "USER_STRUCT"
	FUNC_INIT        = "FUNC_INIT"
	IMPLEMENTED_FUNC = "IMPLEMENTED_FUNC"
)

var (
	RagExpressions = map[string]string{
		USER_STRUCT:      `(?)type .*\b struct\b`,
		FUNC_INIT:        `func init\(\) \{`,
		IMPLEMENTED_FUNC: `func\s*\(\w+\s+\*?\w+\)\s+Launch\s*\(\s*\)\s*\{`,
	}
)
