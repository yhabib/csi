; NOTES:
; 	rdi contains the first argument
; 	rax contains the return value
; 	ADD, or INC or LEAQ could be used
; 	TIME

global	sum_to_n

section	.text

sum_to_n:
	; Initialize both counter and sum to 0	
	mov rsi, 0
	mov rax, 0

loop:
	; compare counter with given n
	cmp rsi, rdi
	; return if equal
	je exit
	; increment counter
	inc rsi
	; add current value to rax
	add	rax, rsi
	; iterate again
	jmp loop

sum_to_n_constant:
	; N * (N+1) / 2
	mov rax, rdi
	inc rax
	mul rdi
	mov rsi, 2
	div rsi

exit:
	ret