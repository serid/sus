package sus

import (
	"errors"
	"sus/cmp"
	"sus/interp"
	"sus/interp/astinterp"
	"sus/interp/val"
	mytesting "sus/stuff/testing"
	"sus/syntax/lexing"
	"sus/syntax/lexing/lexeme"
	"sus/syntax/parsing"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
	"testing"
)

func TestLexer1(t *testing.T) {
	mytesting.Assert(lexeme.CompareLexemeSlices(lexing.Lexate("1 + 2"), []lexeme.Lexeme{
		lexeme.Int(1),
		lexeme.Plus(),
		lexeme.Int(2),
	}), t)
}

func TestLexer2F(t *testing.T) {
	r, err := lexing.LexateE("1 # 2")

	mytesting.AssertEq(r, nil, t)
	mytesting.Assert(errors.Is(err, lexing.NewUnrecognizedCharacterError('#')), t)
}

func TestParser1(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("1 + 2 = 0"), propexpr.NewUnification(valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewIntLit(2)), valexpr.NewIntLit(0)), t)
}

func TestParser2(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("1 + 2 * 4 = 0"), propexpr.NewUnification(valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewMul(valexpr.NewIntLit(2), valexpr.NewIntLit(4))), valexpr.NewIntLit(0)), t)
}

func TestParser3(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("1 * 2 + 4 = 0"), propexpr.NewUnification(valexpr.NewPlus(valexpr.NewMul(valexpr.NewIntLit(1), valexpr.NewIntLit(2)), valexpr.NewIntLit(4)), valexpr.NewIntLit(0)), t)
}

func TestParser4(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("(1 + 2) * 4 = 0"), propexpr.NewUnification(valexpr.NewMul(valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewIntLit(2)), valexpr.NewIntLit(4)), valexpr.NewIntLit(0)), t)
}

func TestParser5(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("1 + (2 * 4) = 0"), propexpr.NewUnification(valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewMul(valexpr.NewIntLit(2), valexpr.NewIntLit(4))), valexpr.NewIntLit(0)), t)
}

func TestParser6(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("1 + 2 + 3 + 10 = 0"), propexpr.NewUnification(valexpr.NewPlus(valexpr.NewPlus(valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewIntLit(2)), valexpr.NewIntLit(3)), valexpr.NewIntLit(10)), valexpr.NewIntLit(0)), t)
}

func TestParser7(t *testing.T) {
	mytesting.AssertEq(parsing.DefaultParser().Parse("((1) + ((2) * 4)) = 0"), propexpr.NewUnification(valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewMul(valexpr.NewIntLit(2), valexpr.NewIntLit(4))), valexpr.NewIntLit(0)), t)
}

func TestParser8(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("@1 = @2 /\\ True() \\/ @5 = @6 /\\ @7 = @8"), propexpr.NewDisjunction(propexpr.NewConjunction(propexpr.NewUnification(valexpr.NewGetVar(1), valexpr.NewGetVar(2)), propexpr.NewRuleCall(1, []valexpr.ValExpr{})), propexpr.NewConjunction(propexpr.NewUnification(valexpr.NewGetVar(5), valexpr.NewGetVar(6)), propexpr.NewUnification(valexpr.NewGetVar(7), valexpr.NewGetVar(8)))), cmp.Cmp, t)
}

func TestParser9(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("Test(1, 2, 3)"), propexpr.NewRuleCall(1, []valexpr.ValExpr{valexpr.NewIntLit(1), valexpr.NewIntLit(2), valexpr.NewIntLit(3)}), cmp.Cmp, t)
}

func TestParser10(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("Test(())"), propexpr.NewRuleCall(1, []valexpr.ValExpr{}), cmp.Cmp, t)
}

func TestParser11(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("Test(1 + 2, 3 * 4 * 5)"), propexpr.NewRuleCall(1, []valexpr.ValExpr{valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewIntLit(2)), valexpr.NewMul(valexpr.NewMul(valexpr.NewIntLit(3), valexpr.NewIntLit(4)), valexpr.NewIntLit(5))}), cmp.Cmp, t)
}

func TestParser1F(t *testing.T) {
	r, err := parsing.DefaultParser().ParseE("1 +")
	mytesting.AssertEq(r, nil, t)
	mytesting.Assert(err != nil, t)
	mytesting.Assert(errors.Is(err, parsing.NewEOFError(lexeme.Any())), t)
}

//func TestParser2F(t *testing.T) {
//	r, err := parsing.DefaultParser().ParseE("1 + 2 3 4")
//	mytesting.AssertEq(r, nil, t)
//	mytesting.Assert(err != nil, t)
//	mytesting.Assert(errors.Is(err, parsing.TrailingLexemesError{}), t)
//}

func TestInterpreter1(t *testing.T) {
	input := []val.Val{val.NewInt(100), nil}
	expectedOutput := []val.Val{val.NewInt(100), val.NewInt(100)}
	testInterpreter(input, expectedOutput, "@0 = @1", t)
}

func TestInterpreter2(t *testing.T) {
	input := []val.Val{nil}
	expectedOutput := []val.Val{val.NewInt(124)}
	testInterpreter(input, expectedOutput, "@0 = 124", t)
}

func TestInterpreter3(t *testing.T) {
	input := []val.Val{val.NewInt(100), nil, nil}
	expectedOutput := []val.Val{val.NewInt(100), val.NewInt(100), val.NewInt(100)}
	testInterpreter(input, expectedOutput, "@0 = @1 /\\ @1 = @2", t)
}

func TestInterpreter4(t *testing.T) {
	input := []val.Val{nil}
	expectedOutput := []val.Val{val.NewInt(50)}
	testInterpreter(input, expectedOutput, "1 = 2 \\/ 50 = @0", t)
}

func TestInterpreter5(t *testing.T) {
	input := []val.Val{nil}
	expectedOutput := []val.Val{val.NewInt(3)}
	testInterpreter(input, expectedOutput, "@0 = 1 + 2", t)
}

func testInterpreter(vals interp.Solution, expectedOutput interp.Solution, s string, t *testing.T) {
	expr := parsing.DefaultParser().Parse(s).(propexpr.PropExpr)
	solution := astinterp.Query(expr, vals)

	mytesting.AssertEqF(solution, expectedOutput, interp.ArrayCmp, t)
}

func TestPrefixRunes(t *testing.T) {
	mytesting.Assert(lexing.PrefixRunes([]rune("abc"), []rune("abcdef")), t)
	mytesting.Assert(lexing.PrefixRunes([]rune("abc"), []rune("abc")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("abcdef"), []rune("abc")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("123"), []rune("abcdef")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("123"), []rune("abc")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("123456"), []rune("abc")), t)
}
