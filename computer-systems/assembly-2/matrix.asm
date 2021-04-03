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
	imul rdx, rcx;rdx*rcx
	add rdx, r8;rdx*rcx + r8
	mov rax, [4*rdx + rdi];4(rdx*rcx + r8) + rdi
	ret
