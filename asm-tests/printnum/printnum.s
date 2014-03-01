    addi $8, $0, 3721
    j printnum

; print the uint32 in $8 to output
printnum:
    bne $8, $0, .nonzero
    addi $1, $0, 0x30
    addi $30, $31, 4
    j printchar
    j .end
.nonzero:
    addi $2, $0, 10
    addi $3, $0, 1    ; base 1
.find:
    mulu $3, $3, $2
    slt $4, $8, $3 
    beq $4, $0, .find ; $8 >= $3
.start:
    divu $3, $3, $2
.loop:
    divu $1, $8, $3
    addi $1, $1, 0x30
    addi $30, $31, 4    ; save the return point
    j printchar
    modu $8, $8, $3
    divu $3, $3, $2
    bne $3, $0, .loop   
.end:
    addi $1, $0, 0xa
    addi $30, $31, 4
    j printchar
    sb $0, 0x4($0)
    
; print the digit in $1 to output
printchar:
.wait:
    lbu $20, 5           ; is output ready?
    bne $20, $0, .wait   
    sb $1, 5
    add $31, $0, $30    ; return
