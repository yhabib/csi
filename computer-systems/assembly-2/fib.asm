section .text
global fib
fib:
	mov eax, edi
	cmp eax, 1
	jle .end				;returns if less or equal to 1
	push rbx				;if proc uses rbx first thing is to save prev value	
	mov ebx, edi		;mov argument to save register so its kept per stack frame
	lea	edi, [ebx-2];using lea to calculate the value but not access the memory as it would have happened with mov
	call fib
	lea	edi, [ebx-1]
	mov ebx, eax;by moving this instruction after the calculation of ebx-1 we don't need to keep its value so it can be reused
	call fib
	add eax, ebx		;rax contains return of fib(n-1) and we add fib(n-2)
	pop rbx					;pops callee-save register

.end:
	ret
