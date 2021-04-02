section .text
global fib
fib:
	push rbx				;if proc uses rbx first thing is to save prev value	
	mov rbx, rdi		;mov argument to save register so its kept per stack frame
	mov rax, rbx
	cmp rax, 1
	jle .end				;returns if less or equal to 1
	lea	rdi, [rbx-2]
	call fib
	push rax				;local variable holding return value of fib(n-2)
	lea	rdi, [rbx-1]
	call fib
	pop r9					;recovers value of fib(n-1) into random register
	add rax, r9			;rax contains return of fib(n-1) and we add fib(n-2)

.end:
	pop rbx					;pops callee-save register
	ret
