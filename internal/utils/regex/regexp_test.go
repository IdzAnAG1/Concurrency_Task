package regex

import "testing"

func TestContains(t *testing.T) {
	type test struct {
		regularExpression string
		content           []string
		ExpectedIndex     int
		ExpectedBoolVar   bool
	}
	tests := []test{
		{
			regularExpression: `[a-zA-Z_0-9]+`,
			content: []string{
				"()---()",
				"Test_123234_12334",
			},
			ExpectedIndex:   1,
			ExpectedBoolVar: true,
		},
		{
			regularExpression: `[a-zA-Z_0-9]+`,
			content: []string{
				"()---()",
				"()---()",
				"()---()",
				" test",
			},
			ExpectedIndex:   3,
			ExpectedBoolVar: true,
		},
		{
			regularExpression: `type\s+[a-zA-Z_0-9]+\s+struct\s+\{`,
			content: []string{
				"import fmt ",
				" type User_123543_dsfsa_1234 struct {",
				" name  string ",
				"}",
			},
			ExpectedIndex:   1,
			ExpectedBoolVar: true,
		},
		{
			regularExpression: `func\s*init\(\s*\)\s*\{`,
			content: []string{
				"} ",
				"func     init(     )       {",
				"()---()",
				"}",
			},
			ExpectedIndex:   1,
			ExpectedBoolVar: true,
		},
		{
			regularExpression: "func\\s+\\([a-zA-Z_0-9]+\\s+\\*?[a-zA-Z_0-9]+\\s*\\)\\s*Launch\\s*\\(\\s*\\)\\s*{",
			content: []string{
				"package test",
				"func    (t     *Test)    Launch(    )     {",
				"()---()",
				"}",
			},
			ExpectedIndex:   1,
			ExpectedBoolVar: true,
		},
		{
			regularExpression: "func\\s+\\([a-zA-Z_0-9]+\\s+\\*?[a-zA-Z_0-9]+\\s*\\)\\s*Launch\\s*\\(\\s*\\)\\s*{",
			content: []string{
				"package main",
				"()---()",
				"func (t Test) Launch() {",
				"()---()",
				"}",
			},
			ExpectedIndex:   2,
			ExpectedBoolVar: true,
		},
		{
			regularExpression: "func\\s+\\([a-zA-Z_0-9]+\\s+\\*?[a-zA-Z_0-9]+\\s*\\)\\s*Launch\\s*\\(\\s*\\)\\s*{",
			content: []string{
				"package main ",
				"()---()",
				"func Launch() {",
				"()---()",
				"}",
			},
			ExpectedIndex:   -1,
			ExpectedBoolVar: false,
		},
		{
			regularExpression: `func\s*init\(\s*\)\s*\{`,
			content: []string{
				"package main ",
				" type UserStruct struct {",
				"()---()",
				"}",
				"func(us *UserStruct) Launch(){",
				"()---()",
				"}",
				"()---()",
				"//testings",
			},
			ExpectedIndex:   -1,
			ExpectedBoolVar: false,
		},
		{
			regularExpression: `type\s+[a-zA-Z_0-9]+\s+struct\s+\{`,
			content: []string{
				"package main ",
				"func(us *UserStruct) Launch(){",
				"()---()",
				"}",
				"()---()",
				"//testings",
			},
			ExpectedIndex:   -1,
			ExpectedBoolVar: false,
		},
	}

	for _, el := range tests {
		testInt, testBool := Contains(el.regularExpression, el.content)
		if testInt != el.ExpectedIndex {
			t.Errorf("[UB] - The (output index :%d) does not match what is (expected :%d)",
				testInt, el.ExpectedIndex)
		}
		if testBool != el.ExpectedBoolVar {
			t.Errorf("[UB] - The (output Boolean value :%t) does not match the (expected value : %t)",
				testBool, el.ExpectedBoolVar)
		}
	}
}
