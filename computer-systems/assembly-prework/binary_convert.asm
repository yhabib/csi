; NOTES:
; 	rdi is a pointer to the value, it holds the address

section .text
global binary_convert
binary_convert:
		xor			rax,			rax					;Initialieze rax to 0
		xor 		rcx,			rcx							
.loop:	
		movzx		rcx, 			byte [rdi]	;takes next byte from string buffer
		cmp 		rcx, 			0			 			;checks against end of string NULL
		je 			.end									;jumps to end
  	sub 		rcx, 			'0'					;normalize each number. '1' in ascii 49 -> 1
		sal			rax,			1						;for each iteration shifts left one pos == num * 2^i 
		add 		rax,			rcx					;adds either 0 or 1 based on current bit
		inc 		rdi										
		jmp 		.loop
.end:
	  ret