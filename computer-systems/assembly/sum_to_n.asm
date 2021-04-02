; NOTES:
; 	rdi contains the first argument
; 	rax contains the return value
; 	ADD, or INC or LEAQ could be used
; 	TIME

global	sum_to_n
section	.text

sum_to_n:
		mov			rsi, 0 		;Initialize both counter and sum to 0	
		mov			rax, 0
loop:
		cmp 		rsi, rdi  ;compare counter with given n
		je 			exit			;return if equal
		inc 		rsi				;increment counter
		add			rax, rsi	;add current value to rax
		jmp 		loop			;iterate again
exit:
		ret

; Alternative
sum_to_n_constant:
		; N * (N+1) / 2
		mov rax, rdi
		inc rax
		mul rdi
		mov rsi, 2
		div rsi
		ret