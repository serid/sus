package bcinterp

import (
	"sus/cmp"
	"sus/interp/bcinterp/bytecode"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// Executable body of a rule
type RuleBody struct {
	Ops    []bytecode.Op
	Result bytecode.SlotNum
}

func CompileBody(expr propexpr.PropExpr, firstFreeVariable bytecode.VarNum) RuleBody {
	var body = RuleBody{Ops: make([]bytecode.Op, 0), Result: -1}

	cs := compilePropExpr(expr, &body.Ops, compilerState{
		NextSlotNum: 1,
		NextVarNum:  firstFreeVariable,
	})

	body.Result = cs.SlotResult()

	return body
}

func compilePropExpr(expr propexpr.PropExpr, body *[]bytecode.Op, cs compilerState) compilerState {
	switch pExpr := expr.(type) {
	case propexpr.Unification:
		context := cs.SlotResult()

		cs1 := compileValExpr(pExpr.E1, context, body, cs)
		input1 := cs1.VarResult()
		cs2 := compileValExpr(pExpr.E2, context, body, cs1)
		input2 := cs2.VarResult()
		output := cs2.NextSlotNum
		*body = append(*body, bytecode.Unify(context, input1, input2, output))
		return cs2.SkipSol()
	case propexpr.Conjunction:
		cs1 := compilePropExpr(pExpr.E1, body, cs)
		input1 := cs1.SlotResult()

		*body = append(*body, bytecode.ConjunctionPart1(-1, -1)) // Set dummy Op to replace it later
		part1Address := len(*body) - 1                           // Save address to know where to fix Op later

		cs2 := compilePropExpr(pExpr.E2, body, cs1)
		input2 := cs2.SlotResult()

		output := cs2.NextSlotNum
		*body = append(*body, bytecode.ConjunctionPart2(input1, input2, output))

		skipAddress := len(*body) - 1 // Address will point to `ConjunctionPart2` when it will be pushed
		(*body)[part1Address] = bytecode.ConjunctionPart1(input1, bytecode.BodyAddress(skipAddress))

		return cs2.SkipSol()
	case propexpr.Disjunction:

		// Push a clone instruction to clone solution before calculating first operand
		*body = append(*body, bytecode.CloneSolution(-1, -1)) // Set dummy Op to replace it later
		cloneOpAddress := len(*body) - 1                      // Save address to know where to fix Op later

		cs1 := compilePropExpr(pExpr.E1, body, cs)
		input1 := cs1.SlotResult()

		csc := cs1.SkipSol() // CompilerState after cloning

		(*body)[cloneOpAddress] = bytecode.CloneSolution(cs.SlotResult(), csc.SlotResult())

		*body = append(*body, bytecode.DisjunctionPart1(-1, -1)) // Set dummy Op to replace it later
		part1Address := len(*body) - 1                           // Save address to know where to fix Op later

		cs2 := compilePropExpr(pExpr.E2, body, csc)
		input2 := cs2.SlotResult()

		output := cs2.NextSlotNum
		*body = append(*body, bytecode.DisjunctionPart2(input1, input2, output))

		skipAddress := len(*body) - 1 // Address will point to `DisjunctionPart2` when it will be pushed
		(*body)[part1Address] = bytecode.DisjunctionPart1(input1, bytecode.BodyAddress(skipAddress))

		return cs2.SkipSol()
	case propexpr.RuleCall:
		var input = make([]bytecode.VarNum, len(pExpr.Args))
		var newCs = cs
		// Compile valExpressions and save their VarNum-s
		for i, exprArg := range pExpr.Args {
			newCs = compileValExpr(exprArg, newCs.SlotResult(), body, newCs)
			input[i] = newCs.VarResult()
		}

		// TODO: convert function name (symbol) to an RuleId
		_ = pExpr.Name

		*body = append(*body, bytecode.RuleCall(newCs.SlotResult(), 1, input, newCs.NextSlotNum))
		return cs.SkipSol()
	default:
		panic("unsupported propexpr")
	}
}

func compileValExpr(expr valexpr.ValExpr, context bytecode.SlotNum, body *[]bytecode.Op, cs compilerState) compilerState {
	switch pExpr := expr.(type) {
	case valexpr.IntLit:
		i := pExpr.Data
		*body = append(*body, bytecode.PutInt(context, i, cs.NextVarNum))
		return cs.SkipVar()
	case valexpr.Plus:
		cs1 := compileValExpr(pExpr.E1, context, body, cs)
		input1 := cs1.VarResult()
		cs2 := compileValExpr(pExpr.E2, context, body, cs1)
		input2 := cs2.VarResult()
		output := cs2.NextVarNum
		*body = append(*body, bytecode.Add(context, input1, input2, output))
		return cs2.SkipVar()
	case valexpr.GetVar:
		*body = append(*body, bytecode.PutVarNum(context, pExpr.VarNum, cs.NextVarNum))
		return cs.SkipVar()
	default:
		panic("unsupported valexpr")
	}
}

func RuleBodyEq(a, b interface{}) bool {
	ra := a.(RuleBody)
	rb := b.(RuleBody)

	return ra.Result == rb.Result && SliceOpEq(ra.Ops, rb.Ops)
}

func SliceOpEq(a, b interface{}) bool {
	ra := a.([]bytecode.Op)
	rb := b.([]bytecode.Op)

	if (ra == nil) || (rb == nil) {
		panic("Attempted to comapre nil slices for equivalence.")
	}

	if len(ra) != len(rb) {
		return false
	}

	for i := range ra {
		if !cmp.Cmp(ra[i], rb[i]) {
			return false
		}
	}
	return true
}
