; &D[i][j] = Xd + L*(C*i + j) => rdi + 4*(rdx*rcx + r8)
; L==4
; Xd==rdi
; C==rdx
; i==rcx
; j==r8
section .text
global index
index:
	; rdi: matrix 
	; rsi: rows 
	; rdx: cols
	; rcx: rindex
	; r8 : cindex
	mov rax, rdx	
	mul rcx								;rdx*rcx
	add rax, r8						;rdx*rcx + r8
	lea rax, [4*rax + rdi];4(rdx*rcx + r8) + rdi
	mov rax, [rax]				;get value from address
	ret
