; NOTES:
;   movsv op uses specific registers: rdi for destinations, rsi for source
;   rep op uses rcx register to know amount of reps
;   repeats movsb operationr: desti == rdi, source == rsi, the amount of times defined in rxx
;   [byte] to take a part of a register for the copmarison

default rel

section .rodata
    prefix: db "One for "
    prefix_len: equ 8
    ; adds '\0' for C string
    suffix: db ", one for me.", 0
    suffix_len: equ 14  ;size of string plus '\0'
    name: db "you", 0   ;c terminating string to check when to stop
section .text
global two_fer
two_fer:
    mov rax, rdi    ;name to rax
    mov rdi, rsi    ;buffer to rdi
    
    lea rsi, [prefix]
    mov rcx, prefix_len
    rep movsb 

    cmp rax, 0
    jne .move_name
    lea rax, [name]
.move_name:
    mov rsi, rax
.add_name_char:
    cmp byte[rsi], 0
    je .add_suffix
    movsb
    jmp .add_name_char

.add_suffix:
    lea rsi, [suffix]
    mov rcx, suffix_len
    rep movsb 

    ret
