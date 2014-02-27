; .string msg "Hello, world.\n\000"
; .const halt 4
; .const stdout 5

;.func main
    add $1, $0, $0      ; init counter
loop:
    lbu $2, 0x1000($1)  ; load byte
    ; lbu $2, msg($1)   ; load byte
    beq $2, $0, end     ; +5
wait:
    lbu $3, 5           ; is output ready?
    bne $3, $0, wait    ; -2
    sb $2, 5            ; output byte
    addi $1, $1, 1      ; update counter
    j loop              ; -7
end:
    sb $0, 0x4($0)

; Output:
; Hello, world.

; Regdump:
; $0 = 0
