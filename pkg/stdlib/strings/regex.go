package strings

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"regexp"
)

/*
 * Returns the matches in the given string text, using the regex.
 * @param text (String) - The string to search in.
 * @param regex (String) - A regular expression to use for matching the text.
 * @param caseInsensitive (Boolean) - If set to true, the matching will be case-insensitive. The default is false.
 * @return (Array) - An array of strings containing the matches.
 */
func RegexMatch(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 3)

	if err != nil {
		return values.None, err
	}

	text := args[0].String()
	exp := args[1].String()

	if len(args) > 2 {
		if args[2] == values.True {
			exp = "(?i)" + exp
		}
	}

	reg, err := regexp.Compile(exp)

	if err != nil {
		return values.None, err
	}

	matches := reg.FindAllStringSubmatch(text, -1)
	res := values.NewArray(10)

	if len(matches) == 0 {
		return res, nil
	}

	for _, m := range matches[0] {
		res.Push(values.NewString(m))
	}

	return res, nil
}

/*
 * Splits the given string text into a list of strings, using the separator.
 * @param text (String) - The string to split.
 * @param regex (String) - A regular expression to use for splitting the text.
 * @param caseInsensitive (Boolean) - If set to true, the matching will be case-insensitive. The default is false.
 * @param limit (Int) - Limit the number of split values in the result. If no limit is given, the number of splits returned is not bounded.
 * @return (Array) - An array of strings splited by teh expression.
 */
func RegexSplit(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 4)

	if err != nil {
		return values.None, err
	}

	text := args[0].String()
	exp := args[1].String()
	limit := -1

	if len(args) > 2 {
		if args[2].Type() == core.IntType {
			limit = int(args[2].(values.Int))
		}
	}

	reg, err := regexp.Compile(exp)

	if err != nil {
		return values.None, err
	}

	matches := reg.Split(text, limit)
	res := values.NewArray(10)

	if len(matches) == 0 {
		return res, nil
	}

	for _, m := range matches {
		res.Push(values.NewString(m))
	}

	return res, nil
}

/*
 * Splits the given string text into a list of strings, using the separator.
 * @param text (String) - The string to split.
 * @param regex (String) - A regular expression to use for splitting the text.
 * @param caseInsensitive (Boolean) - If set to true, the matching will be case-insensitive. The default is false.
 * @return (Boolean) - Returns true if the pattern is contained in text, and false otherwise.
 */
func RegexTest(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 3)

	if err != nil {
		return values.None, err
	}

	text := args[0].String()
	exp := args[1].String()

	if len(args) > 2 {
		if args[2] == values.True {
			exp = "(?i)" + exp
		}
	}

	reg, err := regexp.Compile(exp)

	if err != nil {
		return values.None, err
	}

	matches := reg.MatchString(text)

	return values.NewBoolean(matches), nil
}

/*
 * Splits the given string text into a list of strings, using the separator.
 * @param text (String) - The string to split.
 * @param regex (String) - A regular expression search pattern.
 * @param replacement (String) - The string to replace the search pattern with
 * @param caseInsensitive (Boolean) - If set to true, the matching will be case-insensitive. The default is false.
 * @return (String) - Returns the string text with the search regex pattern replaced with the replacement string wherever the pattern exists in text
 */
func RegexReplace(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 3, 4)

	if err != nil {
		return values.EmptyString, err
	}

	text := args[0].String()
	exp := args[1].String()
	repl := args[2].String()

	if len(args) > 3 {
		if args[3] == values.True {
			exp = "(?i)" + exp
		}
	}

	reg, err := regexp.Compile(exp)

	if err != nil {
		return values.None, err
	}

	out := reg.ReplaceAllString(text, repl)

	return values.NewString(out), nil
}
