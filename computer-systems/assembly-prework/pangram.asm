
section .text
global pangram
pangram:
		xor 		ecx, 		ecx						;zeroes-out ecx to receive the result
		mov 		r9d, 		0x3ffffff

.loop:
		cmp     [rdi],  byte 0				;compares current char in the buffer of strings to 0 looking for NULL==end of string
		je			.end			
		movzx		edx, 		byte [rdi]		;movs zero extend next char from string
		or			edx, 		32						;sets bit 32 to true so uppercase becomes lowercase
		sub			edx, 		'a'						;substracts 'a' to normalize every single letter from 0-25
		bts			ecx,		edx						;sets bit in eax(bit base) defined by ecx(bit offset) to 1
  	inc 		rdi										;moves rdi pointer to next char
		jmp			.loop

.end:
		xor 		eax, 		eax						;zeroes-out eax to hold the result
		and 		ecx, 		r9d						;masks eax to get only the last 26bits
		cmp 		ecx, 		r9d						;compares eax to expeted solution by substracting eax from r9d and if substraction is zero sets ZF to 1
		sete 		al										;sets byte if ZF=1
		ret