; V=1/3*pi*r^2*h
default rel
section .data
	PI dd 3.141592
	three dd 3.0

section .text
global volume
volume:
	; ymm0 is radius
	; ymm1 is height
	mulss xmm0, xmm0
	mulss xmm0, xmm1
	mulss xmm0, [PI]
	divss xmm0, [three]
 	ret
