package main

import (
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Stack"
	"fmt"
	"math"
	"strconv"
)

func main() {
	//res := baseConversation(2013, 16)
	//fmt.Println(res)
	//31023

	//exp := "{2+[1+1-(2+3)]-3}"
	//fmt.Println(paren(exp, 0, len(exp)))
	//true

	//exp := "22$"
	//22

	exp := "(1+2^3!-4)*(5!-(6-(7-(89-0!))))$"
	//exp := "(2^3)$"
	//var RPN string
	res, RPN := evaluate(exp)
	fmt.Printf("result = %.0f\n", res)
	fmt.Printf("RPN = %s\n", RPN)
	//result = 2013
	//RPN =  1 2 3 ! ^ + 4 - 5 ! 6 7 89 0 ! - - - - *
}

/* ----------- 数值转换 ----------- */
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

/* ----------- 括号匹配 ----------- */
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

/* ----------- 中缀表达式 Infix ----------- */
func evaluate(S string) (res float64, RPN string) {
	var opnd, optr Stack.Stack
	optr.Push('$')
	RPN = ""
	for !optr.Empty() {
		//fmt.Printf("%c\n", S[0])
		if '0' <= S[0] && S[0] <= '9' {
			S = readNumber(S, &opnd)
			//opnd.Push(S[0]-48)
			//S = S[1:]
			RPN = RPNappend(RPN, strconv.Itoa(int(opnd.Top().(uint8))))
		} else {
			//fmt.Printf("S[0]= %c\n", S[0])
			//fmt.Printf("S[0]= %T\n", S[0])

			//fmt.Printf("Top = %c\n", optr.Top())
			//fmt.Printf("Top = %T\n", optr.Top())
			var top int32
			if c, ok := optr.Top().(uint8); ok {
				top = int32(c)
			}
			if c, ok := optr.Top().(int32); ok {
				top = c
			}
			//fmt.Printf("%c\n", orderBetween(top, int32(S[0])))
			switch orderBetween(top, int32(S[0])) {
			case '<':
				optr.Push(S[0])
				S = S[1:]
			case '=':
				optr.Pop()
				S = S[1:]
			case '>':
				op := optr.Pop()
				RPN = RPNappend(RPN, string(int(op.(uint8))))
				if op.(uint8) == '!' {
					pOpnd := f64Assert(opnd.Pop())
					opnd.Push(calcu1opnd(int(op.(uint8)), pOpnd))
				} else {
					pOpnd2 := f64Assert(opnd.Pop())
					pOpnd1 := f64Assert(opnd.Pop())
					opnd.Push(calcu2opnd(pOpnd1, int(op.(uint8)), pOpnd2))
				}
			}
		}
	}
	if num, ok := opnd.Top().(uint8); ok {
		opnd.Pop()
		return float64(num), RPN
	}
	return opnd.Pop().(float64), RPN
}

func readNumber(S string, opnd *Stack.Stack) string {
	res := S[0] - 48
	i := 1
	for '0' <= S[i] && S[i] <= '9' {
		res = 10*res + (S[i] - 48)
		i++
	}
	opnd.Push(res)
	return S[i:]
}

func f64Assert(in interface{}) float64 {
	var res float64
	if c, ok := in.(uint8); ok {
		res = float64(c)
	}
	if c, ok := in.(float64); ok {
		res = c
	}
	return res
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

func calcu1opnd(op int, num float64) float64 {
	//num -= 48
	switch op {
	//case '!': return factorial(num)
	default:
		return factorial(num)
	}
}

func factorial(num float64) float64 {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// {'+', '-'} < {'*', '/'} < {'^'} < {'!'}
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

func orderBetween(topOp int32, newOp int32) int {
	rowIdx, colIdx := optrIdx(topOp), optrIdx(newOp)
	return pri[rowIdx][colIdx]
}

func optrIdx(op int32) int {
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

/* ----------- 逆波兰表达式表达式 Reverse Polish Notation ----------- */
func RPNappend(rpn string, op string) string {
	rpn = rpn + " " + op
	return rpn
}
