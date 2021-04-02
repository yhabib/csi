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
	mov rax, 1
	imul rax, rdx
	imul rax, rcx
	add rax, r8
	imul rax, 4
	add rax, rdi
	mov rax, [rax]
	ret
