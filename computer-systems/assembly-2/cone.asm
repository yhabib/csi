; V=1/3*pi*r^2*h
default rel
section .rodata
	pi_thirds dd 1.04719

section .text
global volume
volume:
	; ymm0 is radius
	; ymm1 is height
	mulss xmm0, xmm0
	mulss xmm0, xmm1
	mulss xmm0, [pi_thirds]
 	ret
