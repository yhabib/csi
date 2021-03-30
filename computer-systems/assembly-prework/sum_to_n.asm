; NOTES:
; 	rdi contains the first argument
; 	rax contains the return value
; 	ADD, or INC or LEAQ could be used
; 	TIME

global	sum_to_n

section	.text

sum_to_n:
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

exit:
	ret
