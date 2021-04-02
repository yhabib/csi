;fib(n) {
;  if (n<2) return n;
;  fib = 1, aux=0;
;  for(2->n) 
;	   temp = fib
;	   fib += aux
;	   aux = temp
;  return fib 
; } 
section .text
global fib_iter
fib_iter:
	mov rax, rdi		;rax contains the fib sum
	cmp rdi, 2
	jl .end
	mov rax, 1
	mov r9, 0				;r9 auxiliar variable to track fib(n-1)
	mov r10, 2			;r10 used as a counter for the loop
.loop:
	push rax
	add rax, r9
	pop r9
	inc r10
	cmp r10, rdi
	jle	.loop

.end:
	ret
