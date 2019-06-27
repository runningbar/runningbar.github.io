// main.c
int main() {
    foobar(1, 2, 3, 4, 5, 6, 7, 8)
}

// asm
main:
	// ... omit some lines
	movl	$8, 8(%rsp)
	movl	$7, (%rsp)
	movl	$6, %r9d
	movl	$5, %r8d
	movl	$4, %ecx
	movl	$3, %edx
	movl	$2, %esi
	movl	$1, %edi
	call	foobar
    // ... omit some lines

