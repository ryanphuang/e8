    addi $8, $0, 3927

; and the procedure starts here
    bne $8, $0, .nonzero
.zero:
    addi $1, $0, 0x30
    addi $30, $31, 4
    j printChar
    j .end

.nonzero:
    addi $10, $0, 10
    addi $3, $0, 1    ; base 1
.find:
    mulu $3, $3, $10
    slt $4, $8, $3 
    beq $4, $0, .find ; $8 >= $3

    divu $3, $3, $10
.loop:
    divu $1, $8, $3

    addi $1, $1, 0x30   ; convert digit to char
    ; call printChar to print the digit
    addi $30, $31, 4    ; save the return point
    j printChar

    modu $8, $8, $3     ; remove that printed digit
    divu $3, $3, $10    ; and shift the base
    bne $3, $0, .loop   

.end:
    ; print an end line
    addi $1, $0, 0xa
    addi $30, $31, 4    ; save the return point
    j printChar
    
    sb $0, 0x4($0)   ; halt
    
; print the digit in $1 to output
printChar:
.loop:
    lbu $29, 5           ; is output ready?
    bne $29, $0, .loop
    sb $1, 5
    add $31, $0, $30    ; return
