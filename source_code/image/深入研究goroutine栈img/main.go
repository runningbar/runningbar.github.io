// main.go
func main() {
    foobar(1, 2)
}

// asm: $ GOOS=linux GOARCH=amd64 go tool compile -S -N -l main.go

0x0000 00000 (main.go:7)	TEXT	"".main(SB), ABIInternal, $40-0
// ... omit some lines
0x001d 00029 (main.go:8)	MOVQ	$1, (SP)
0x0025 00037 (main.go:8)	MOVQ	$2, 8(SP)
0x002e 00046 (main.go:8)	CALL	"".foobar(SB)
// ... omit some lines


package main

func add(a, b int64) (int64, int64) {
	var tmp int64 = 1
	tmp = tmp + a
	return a + b, a - b
}

func main() {
	var c int64 = 10
	var d int64 = 12
	add(c, d)
}

// asm: $ GOOS=linux GOARCH=amd64 go tool compile -S -N -l main.go

0x0000 00000 (main.go:3)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $16-32
0x0000 00000 (main.go:3)	SUBQ	$16, SP
0x0004 00004 (main.go:3)	MOVQ	BP, 8(SP)
0x0009 00009 (main.go:3)	LEAQ	8(SP), BP
// ... omit some lines：GC相关
0x000e 00014 (main.go:3)	MOVQ	$0, "".~r2+40(SP)
0x0017 00023 (main.go:3)	MOVQ	$0, "".~r3+48(SP)
0x0020 00032 (main.go:4)	MOVQ	$1, "".tmp(SP)
0x0028 00040 (main.go:5)	MOVQ	"".a+24(SP), AX
0x002d 00045 (main.go:5)	INCQ	AX
0x0030 00048 (main.go:5)	MOVQ	AX, "".tmp(SP)
0x0034 00052 (main.go:6)	MOVQ	"".a+24(SP), AX
0x0039 00057 (main.go:6)	ADDQ	"".b+32(SP), AX
0x003e 00062 (main.go:6)	MOVQ	AX, "".~r2+40(SP)
0x0043 00067 (main.go:6)	MOVQ	"".a+24(SP), AX
0x0048 00072 (main.go:6)	SUBQ	"".b+32(SP), AX
0x004d 00077 (main.go:6)	MOVQ	AX, "".~r3+48(SP)
0x0052 00082 (main.go:6)	MOVQ	8(SP), BP
0x0057 00087 (main.go:6)	ADDQ	$16, SP
0x005b 00091 (main.go:6)	RET