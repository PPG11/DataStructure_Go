package main

import (
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Stack"
	"math"
)

func main() {
	//res := baseConversation(2013, 16)
	//fmt.Println(res)
	//31023

	//exp := "{2+[1+1-(2+3)]-3}"
	//fmt.Println(paren(exp, 0, len(exp)))
	//true

	exp := "3+2$"
	//fmt.Println(evaluate(exp))
}

func baseConversation(num int, base int) string {
	var T Stack.Stack
	var digit = [...]string{`0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `A`, `B`, `C`,
		`D`, `E`, `F`}
	for num != 0 {
		T.Push(digit[num%base])
		num /= base
	}
	res := ""
	for T.Size() > 0 {
		res += T.Pop().(string)
	}
	return res
}

func paren(exp string, lo int, hi int) bool {
	var S Stack.Stack
	for i := lo; i < hi; i++ {
		//fmt.Println(S)
		switch {
		case exp[i] == '(' || exp[i] == '[' || exp[i] == '{':
			S.Push(exp[i])
		case exp[i] == ')' && !S.Empty():
			if S.Pop().(uint8) != '(' {
				return false
			}
		case exp[i] == ']' && !S.Empty():
			if S.Pop().(uint8) != '[' {
				return false
			}
		case exp[i] == '}' && !S.Empty():
			if S.Pop().(uint8) != '{' {
				return false
			}
		case (exp[i] == ')' || exp[i] == ']' || exp[i] == '}') && S.Empty():
			return false
		}
	}
	//fmt.Println(S.Size())
	return S.Empty()
}

func evaluate(S string) float64 {
	var opnd, optr Stack.Stack
	optr.Push('$')
	for !optr.Empty() {
		if '0' <= S[0] && S[0] <= '9' {
			//readNumber(S[0], opnd)
			opnd.Push(S[0])
		} else {
			switch orderBetween(optr.Top().(uint8), S[0]) {
			case '<':
				optr.Push(S[0])
				S = S[1:]
			case '=':
				optr.Pop()
				S = S[1:]
			case '>':
				op := optr.Pop()
				if op == '!' {
					opnd.Push(calcu1opnd(op.(int), opnd.Pop().(int)))
				} else {
					pOpnd2 := opnd.Pop().(float64)
					pOpnd1 := opnd.Pop().(float64)
					opnd.Push(calcu2opnd(pOpnd1, op.(int), pOpnd2))
				}
			}
		}
	}
	return opnd.Pop().(float64)
}

func calcu2opnd(opnd1 float64, op int, opnd2 float64) float64 {
	switch op {
	case '+':
		return opnd1 + opnd2
	case '-':
		return opnd1 - opnd2
	case '*':
		return opnd1 * opnd2
	case '/':
		return opnd1 / opnd2
	//case '^': return math.Pow(opnd1, opnd2)
	default:
		return math.Pow(opnd1, opnd2)
	}
}

func calcu1opnd(op int, num int) int {
	switch op {
	//case '!': return factorial(num)
	default:
		return factorial(num)
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

var pri = [][]int{
	//              |-------------- 当前运算符 --------------|
	//              +    -    *    /    ^    !    (    )    $
	/*  --  +  */ {'>', '>', '<', '<', '<', '<', '<', '>', '>'},
	/*  |   -  */ {'>', '>', '<', '<', '<', '<', '<', '>', '>'},
	/*  栈  *  */ {'>', '>', '>', '>', '<', '<', '<', '>', '>'},
	/*  顶  /  */ {'>', '>', '>', '>', '<', '<', '<', '>', '>'},
	/*  运  ^  */ {'>', '>', '>', '>', '>', '<', '<', '>', '>'},
	/*  算  !  */ {'>', '>', '>', '>', '>', '>', ' ', '>', '>'},
	/*  符  (  */ {'<', '<', '<', '<', '<', '<', '<', '=', ' '},
	/*  |   )  */ {' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	/*  --  $  */ {'<', '<', '<', '<', '<', '<', '<', ' ', '='},
}

func orderBetween(topOp uint8, newOp uint8) int {
	rowIdx, colIdx := optrIdx(topOp), optrIdx(newOp)
	return pri[rowIdx][colIdx]
}

func optrIdx(op uint8) int {
	switch op { // + - * / ^ ! ( ) $
	case '+':
		return 0
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 3
	case '^':
		return 4
	case '!':
		return 5
	case '(':
		return 6
	case ')':
		return 7
	//case '$': return 8
	default:
		return 8
	}
}
