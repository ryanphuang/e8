main:
    addi $1, 0, 4

    addi $14, $0, 4096
    sw $1, 0($14)
    addi $14, $14, 16
    j fab
    ; TODO

