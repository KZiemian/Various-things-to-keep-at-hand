#!/bin/bash

# for elem in 2 3 5 7 11 13 17 19 ; do
#     echo "prime number -> $elem"
# done

# echo "Kod po pętli 'for'."

# for elem in 2, 3, 5, 7, 11, 13, 17, 19 ; do
#     echo "prime number -> $elem"
# done

# echo "Kod po pętli 'for'."

# for elem in 2 3 5 7 11 13 17 19 ; do
#     echo "prime number -> $elem"

#     varTemp=$(($elem + 1))
#     echo "varTemp = $varTemp."
# done

# for elem in 2, 3, 5, 7, 11, 13, 17, 19 ; do
#     echo "prime number -> $elem."

#     echo "elem + 1 = $(($elem + 1))."
# done

# VAR1="Ty"
# echo "\$VAR1 = $((VAR1))."

# funkcja_1 () {
#     echo "Witam was z wnętrza funkcji 1."
# }

# funkcja_1

# funkcja_2 () { echo "Witam was z wnętrza inaczej zdefiniowanej funkcji" ; }

# funkcja_2

# function funkcja_3 {
#     echo "Witam was z wnętrza funkcji 3."
# }

# funkcja_3

# VAR1="A"
# VAR2="B"

# jakas_funkcja() {
#     VAR1="A1"
#     VAR2="B1"

#     echo "Wewnątrz funkcji, VAR1: $VAR1, VAR2: $VAR2."
# }

# echo "Przed wykonaniem 'jakas_funkcja', VAR1=$VAR1, VAR2=$VAR2."

# jakas_funkcja

# echo "Po wykonaniu 'jakas_funkcja, VAR1=$VAR1, VAR2=$VAR2."

# VAR1="A"
# VAR2="B"

# jakas_inna_funkcja() {
#     local VAR1="A1"
#     VAR2="B1"
#     VAR3="C1"

#     echo "Wewnątrz 'jakas_inna_funkcja', VAR1: $VAR1, VAR2: $VAR2, VAR3: $VAR3"
# }

# echo "Przed wywowalaniem 'jakas_inna_funkcja', VAR1: $VAR2, VAR2: $VAR2, VAR3: $VAR3."

# jakas_inna_funkcja

# echo "Po wywolaniu 'jakas_inna_funkcja', VAR1: $VAR1, VAR2: $VAR2, VAR3: $VAR3."

# VAR1="A"

# funkcja_BASH_01() {
#     local VAR2="B"

#     echo "Wewnątrz funkcji VAR1: $VAR1, VAR2: $VAR2"
# }

# echo "Przed wywołaniem funkcji, VAR1: $VAR1, VAR2: $VAR2."

# funkcja_BASH_01

# echo "Po wywołaniu funkcji, VAR1: $VAR1, VAR2: $VAR2."

add() {
    local VAR_LOC_1=$(($1 + $2))

    echo "VAR_LOC_1: $1 + $2 = $VAR_LOC_1"

    return 0
}

add 1 2
add 2 2
add 3 2
add 2 3

echo "$?"
