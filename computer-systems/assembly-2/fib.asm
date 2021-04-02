section .text
global fib
fib:
	push rbx				;if proc uses rbx first thing is to save prev value	
	mov ebx, edi		;mov argument to save register so its kept per stack frame
	mov eax, ebx
	cmp eax, 1
	jle .end				;returns if less or equal to 1
	lea	edi, [ebx-2]
	call fib
	push rax				;local variable holding return value of fib(n-2)
	lea	edi, [ebx-1]
	call fib
	pop r9					;recovers value of fib(n-1) into random register
	add eax, r9d		;rax contains return of fib(n-1) and we add fib(n-2)

.end:
	pop rbx					;pops callee-save register
	ret
