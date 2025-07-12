package variables

const (
	USER_STRUCT = "USER_STRUCT"

	FUNC_INIT = "FUNC_INIT"

	IMPLEMENTED_FUNC = "IMPLEMENTED_FUNC"

	UserStructStartName = "UserStruct_%s"

	UserStructText = "type %s struct {\n\n}\n"

	FuncInitText = "func init() {\n" +
		"\tStorage := task_storage.GetStorageInstance()\n" +
		"\tStorage.AddInStorage(\"Task_%s\",\n \t\t&%s{})\n" +
		"}\n"
	FunctionForImplementation = "func (t %s) Launch() {\n\n}\n"

	CommentForUserStruct = "// Come up with your own name for the structure and \n " +
		"//replace the current %s structure name"

	CommentForFunctionForImplementation = "//This is an interface implementation function. " +
		"\n//To run the program, it uses a receiver with a name created automatically. " +
		"\n//Change the name in the receiver to the one you created."

	CommentForFunctionInit = "//This function allows you to immediately add your task to the execution stream.\n" +
		"//Change it only if you change the name of the structure created automatically."
)

var (
	RegExpressions = map[string]string{
		USER_STRUCT:      `type\s+[a-zA-Z_0-9]+\s+struct\s+\{`,
		FUNC_INIT:        `func\s*init\(\s*\)\s*\{`,
		IMPLEMENTED_FUNC: `func\s+\([a-zA-Z_0-9]+\s+\*?[a-zA-Z_0-9]+\s*\)\s*Launch\s*\(\s*\)\s*{`,
	}
)
