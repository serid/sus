package bcinterp

import (
	"fmt"
	"sus/interp"
	"sus/interp/bcinterp/bytecode"
	"sus/interp/val"
)

func Solve(body RuleBody, input map[string]val.Val) interp.Solution {
	solutions := make([]interp.Solution, body.Result + 1)

	// TODO: compute len(vals) from source code, don't hard-code 20
	vals := make(interp.Solution, 20)

	// Copy inputs to actual val array (solution) that will be solved
	for k, v := range input {
		vals[body.VarTable[k]] = v
	}

	solutions[0] = vals

	var i bytecode.BodyAddress = 0
	for int(i) < len(body.Ops) {
		op := body.Ops[i]
		switch op.OpCode {
		case bytecode.OpPutInt:
			data := op.Data.(bytecode.PutIntData)
			solutions[data.Context][data.Output] = val.NewInt(data.Data)
		case bytecode.OpPutVarNum:
			data := op.Data.(bytecode.PutVarNumData)
			solutions[data.Context][data.Output] = val.Ref{Value: data.Data}
		case bytecode.OpAdd:
			data := op.Data.(bytecode.ValBinOpData)
			solutions[data.Context][data.Output] = val.NewInt(solutions[data.Context][data.Input1].(val.Int).Value + solutions[data.Context][data.Input2].(val.Int).Value)
		case bytecode.OpAssign:
			data := op.Data.(bytecode.AssignData)
			solutions[data.Context][data.Output] = solutions[data.Context][data.Input]
		case bytecode.OpCloneSolution:
			data := op.Data.(bytecode.SolCloneSolutionData)
			solutions[data.Output] = solutions[data.Input].Clone()
		case bytecode.OpUnify:
			data := op.Data.(bytecode.SolUnifyData)

			// Pointer to first value in solutions[data.Context][] array
			var val1 *val.Val
			// Pointer to first value in solutions[data.Context][] array
			var val2 *val.Val

			// If a value is nil, it's an empty variable
			// If a value is a Ref, dereference it
			// Otherwise, it's some value in a variable
			val1 = maybeDereference(solutions[data.Context], data.Input1)
			val2 = maybeDereference(solutions[data.Context], data.Input2)

			// If one of args to Unification is nil, copy value overwriting None

			if *val1 == nil {
				*val1 = *val2
			}

			if *val2 == nil {
				*val2 = *val1
			}

			// If Values in expressions are equal, return true, otherwise false
			if *val1 == *val2 {
				solutions[data.Output] = solutions[data.Context]
			} else {
				solutions[data.Output] = nil
			}
		case bytecode.OpConjunctionPart1:
			data := op.Data.(bytecode.SolMaybeSkipData)

			// If solution of first operand is nil, skip calculation of second operand
			if solutions[data.Input1] == nil {
				i = data.SkipAddress
				continue
			}
		case bytecode.OpConjunctionPart2:
			data := op.Data.(bytecode.SolBinOpData)

			// If solution of first operand is nil, return nil
			if solutions[data.Input1] == nil {
				solutions[data.Output] = nil
			} else {
				solutions[data.Output] = solutions[data.Input2]
			}
		case bytecode.OpDisjunctionPart1:
			data := op.Data.(bytecode.SolMaybeSkipData)

			// If solution of first operand is not nil, skip calculation of second operand
			// TODO: implement returning multiple solutions
			if solutions[data.Input1] != nil {
				i = data.SkipAddress
				continue
			}
		case bytecode.OpDisjunctionPart2:
			data := op.Data.(bytecode.SolBinOpData)

			// If solution of first operand is not nil, return it
			if solutions[data.Input1] != nil {
				solutions[data.Output] = solutions[data.Input1]
			} else {
				solutions[data.Output] = solutions[data.Input2]
			}
		case bytecode.OpRuleCall:
			data := op.Data.(bytecode.SolRuleCallData)
			// TODO: implement RuleCall

			solutions[data.Output] = solutions[data.Context]
		default:
			panic(fmt.Sprintf("unhandled opcode: %#v", op.OpCode))
		}
		i++
	}

	result := solutions[body.Result]

	return result
}

func maybeDereference(solution interp.Solution, where bytecode.VarNum) *val.Val {
	if solution[where] == nil {
		return &solution[where]
	} else if r, ok := solution[where].(val.Ref); ok {
		return &solution[r.Value]
	} else {
		return &solution[where]
	}
}