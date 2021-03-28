package sus

import (
	"errors"
	"sus/cmp"
	"sus/interp/bcinterp"
	"sus/interp/bcinterp/bytecode"
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
	mytesting.AssertEqF(parsing.DefaultParser().Parse("A = B /\\ True() \\/ C = D /\\ E = F"), propexpr.NewDisjunction(propexpr.NewConjunction(propexpr.NewUnification(valexpr.NewGetVar("A"), valexpr.NewGetVar("B")), propexpr.NewRuleCall("True", []valexpr.ValExpr{})), propexpr.NewConjunction(propexpr.NewUnification(valexpr.NewGetVar("C"), valexpr.NewGetVar("D")), propexpr.NewUnification(valexpr.NewGetVar("E"), valexpr.NewGetVar("F")))), cmp.Cmp, t)
}

func TestParser9(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("Test(1, 2, 3)"), propexpr.NewRuleCall("Test", []valexpr.ValExpr{valexpr.NewIntLit(1), valexpr.NewIntLit(2), valexpr.NewIntLit(3)}), cmp.Cmp, t)
}

func TestParser10(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("Test(())"), propexpr.NewRuleCall("Test", []valexpr.ValExpr{}), cmp.Cmp, t)
}

func TestParser11(t *testing.T) {
	mytesting.AssertEqF(parsing.DefaultParser().Parse("Test(1 + 2, 3 * 4 * 5)"), propexpr.NewRuleCall("Test", []valexpr.ValExpr{valexpr.NewPlus(valexpr.NewIntLit(1), valexpr.NewIntLit(2)), valexpr.NewMul(valexpr.NewMul(valexpr.NewIntLit(3), valexpr.NewIntLit(4)), valexpr.NewIntLit(5))}), cmp.Cmp, t)
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
	input := map[string]val.Val{"A": val.NewInt(100), "B": nil}
	expectedOutput := map[string]val.Val{"A": val.NewInt(100), "B": val.NewInt(100)}
	testBytecodeInterpreter(input, expectedOutput, "A = B", t)
}

func TestInterpreter2(t *testing.T) {
	input := map[string]val.Val{"A": nil}
	expectedOutput := map[string]val.Val{"A": val.NewInt(124)}
	testBytecodeInterpreter(input, expectedOutput, "A = 124", t)
}

func TestInterpreter3(t *testing.T) {
	input := map[string]val.Val{"A": val.NewInt(100), "B": nil, "C": nil}
	expectedOutput := map[string]val.Val{"A": val.NewInt(100), "B": val.NewInt(100), "C": val.NewInt(100)}
	testBytecodeInterpreter(input, expectedOutput, "A = B /\\ B = C", t)
}

func TestInterpreter4(t *testing.T) {
	input := map[string]val.Val{"A": nil}
	expectedOutput := map[string]val.Val{"A": val.NewInt(50)}
	testBytecodeInterpreter(input, expectedOutput, "1 = 2 \\/ 50 = A", t)
}

func TestInterpreter5(t *testing.T) {
	input := map[string]val.Val{"A": nil}
	expectedOutput := map[string]val.Val{"A": val.NewInt(3)}
	testBytecodeInterpreter(input, expectedOutput, "A = 1 + 2", t)
}

//func testInterpreter(vals interp.Solution, expectedOutput interp.Solution, s string, t *testing.T) {
//	expr := parsing.DefaultParser().Parse(s).(propexpr.PropExpr)
//	solution := astinterp.Query(expr, vals)
//
//	mytesting.AssertEqF(solution, expectedOutput, interp.ArrayCmp, t)
//}

func testBytecodeInterpreter(input map[string]val.Val, expectedOutput map[string]val.Val, s string, t *testing.T) {
	expr := parsing.DefaultParser().Parse(s).(propexpr.PropExpr)
	bc := bcinterp.CompileBody(expr, 1)
	solution := bcinterp.Solve(bc, input)

	for k, v := range expectedOutput {
		mytesting.AssertEq(solution[bc.VarTable[k]], v, t)
	}
}

func TestBytecodeCompiler1(t *testing.T) {
	expr := parsing.DefaultParser().Parse("1 + 2 = A \\/ True()").(propexpr.PropExpr)
	bc := bcinterp.CompileBody(expr, 0)
	expectedBc := bcinterp.RuleBody{Ops: []bytecode.Op{
		bytecode.CloneSolution(0, 2),
		bytecode.PutInt(0, 1, 0),
		bytecode.PutInt(0, 2, 1),
		bytecode.Add(0, 0, 1, 2),
		bytecode.PutVarNum(0, 3, 4),
		bytecode.Unify(0, 2, 4, 1),
		bytecode.DisjunctionPart1(1, 8),
		bytecode.RuleCall(2, 1, []bytecode.VarNum{}, 3),
		bytecode.DisjunctionPart2(1, 3, 4),
	}, LastSlot: 4, VarTable: map[string]bytecode.VarNum{"A": 3}}

	mytesting.AssertEqF(bc, expectedBc, bcinterp.RuleBodyEq, t)

	input := map[string]val.Val{"A": nil}
	solution := bcinterp.Solve(bc, input)

	for k, v := range map[string]val.Val{"A": val.NewInt(3)} {
		mytesting.AssertEq(solution[bc.VarTable[k]], v, t)
	}
}

func TestPrefixRunes(t *testing.T) {
	mytesting.Assert(lexing.PrefixRunes([]rune("abc"), []rune("abcdef")), t)
	mytesting.Assert(lexing.PrefixRunes([]rune("abc"), []rune("abc")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("abcdef"), []rune("abc")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("123"), []rune("abcdef")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("123"), []rune("abc")), t)
	mytesting.Assert(!lexing.PrefixRunes([]rune("123456"), []rune("abc")), t)
}
