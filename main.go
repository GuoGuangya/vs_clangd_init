package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const clangFormat = "---\nLanguage: Cpp\nBasedOnStyle: LLVM\nAccessModifierOffset: -4\nAlignConsecutiveAssignments: false\nAlignConsecutiveDeclarations: false\nAlignOperands: true\nAlignTrailingComments: false\nAlwaysBreakTemplateDeclarations: Yes\nBraceWrapping: \n  AfterCaseLabel: false\n  AfterClass: false\n  AfterControlStatement: false\n  AfterEnum: false\n  AfterFunction: false\n  AfterNamespace: false\n  AfterStruct: false\n  AfterUnion: false\n  AfterExternBlock: false\n  BeforeCatch: false\n  BeforeElse: false\n  BeforeLambdaBody: false\n  BeforeWhile: false\n  SplitEmptyFunction: true\n  SplitEmptyRecord: true\n  SplitEmptyNamespace: true\nBreakBeforeBraces: Custom\nBreakConstructorInitializers: AfterColon\nBreakConstructorInitializersBeforeComma: false\nColumnLimit: 120\nConstructorInitializerAllOnOneLineOrOnePerLine: false\nContinuationIndentWidth: 8\nIncludeCategories: \n  - Regex: '^<.*'\n    Priority: 1\n  - Regex: '^\".*'\n    Priority: 2\n  - Regex: '.*'\n    Priority: 3\nIncludeIsMainRegex: '([-_](test|unittest))?$'\nIndentCaseLabels: true\nIndentWidth: 4\nInsertNewlineAtEOF: true\nMacroBlockBegin: ''\nMacroBlockEnd: ''\nMaxEmptyLinesToKeep: 2\nNamespaceIndentation: All\nSpaceAfterCStyleCast: true\nSpaceAfterTemplateKeyword: false\nSpaceBeforeRangeBasedForLoopColon: false\nSpaceInEmptyParentheses: false\nSpacesInAngles: false\nSpacesInConditionalStatement: false\nSpacesInCStyleCastParentheses: false\nSpacesInParentheses: false\nTabWidth: 4\n...\n"

func initVscodeSettings() {
	file, err := os.OpenFile(".vscode/settings.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(fileContent))

	settings := make(map[string]interface{})
	err = json.Unmarshal(fileContent, &settings)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(settings)
	settings["editor.formatOnSave"] = true
	settings["clang-format.executable"] = "/usr/bin/clang-format-18"
	newJson, err := json.MarshalIndent(&settings, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(string(newJson))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initClangFormatFile() {
	err := os.WriteFile(".clang-format", []byte(clangFormat), 0666)
	if err != nil {
		return
	}
}

func main() {
	_, err := os.Stat("/usr/bin/clang-format-18")
	if err != nil {
		fmt.Println("Please execute command: sudo apt-get install clang-format-18")
		return
	}

	_, err = os.Stat(".vscode/settings.json")
	if err != nil {
		fmt.Println("Please init vscode project at first")
		return
	}

	initClangFormatFile()
	initVscodeSettings()
}
