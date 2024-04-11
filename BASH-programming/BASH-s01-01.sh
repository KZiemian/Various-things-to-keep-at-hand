#!/bin/bash

# echo "Jestem w skrypcie."

# arrayVar1=(jeden dwa trzy cztery piec)

# for t in ${arrayVar1[@]} ; do
#     echo "t = $t."
# done

# for index in ${!arrayVar1[@]} ; do
#     echo "index = $index"
# done

# echo "\${#arrayVar1[3]} = ${#arrayVar1[3]}."

# echo "\${arrayVar1[8]} = ${arrayVar1[8]}."

# sum=0

# arrayVar1=(1 2 3 4 5)

# for elem in ${arrayVar1[@]} ; do
#     sum=$(($sum+$elem))
# done

# echo "elem = $elem."

# echo "sum = $sum."

# arrayVar1=(jeden dwa trzy cztery piec)

# for index in ${!arrayVar1[@]} ; do
#     echo "arrayVar1[$index] = ${arrayVar1[$index]}."
# done

# echo "index = $index."

# if [ -e someFile.txt ] ; then
#     echo "Plik someFile.txt istnieje."
# else
#     echo "Plik someFile.txt nie istnieje."
# fi

# VAR1=1
# VAR2=2

# if [ VAR1 -lt VAR2 ] ; then
#     echo "Prawda."
# else
#     echo "Fałsz."
# fi

# if [ 2 -lt 1 ] ; then
#     echo "Prawda."
# else
#     echo "Fałsz."
# fi

# echo "Jestem po instrukcji if."

VAR1=1
VAR2=2

# if [ VAR1 -lt $VAR2 ] ; then
#     echo "Prawda"
# else
#     echo "Falsz."
# fi

# if [ $VAR1 -lt 3 ] ; then
#     echo "Prawda."
# else
#     echo "Falsz."
# fi

# if [ $VAR1 -lt 0 ] ; then
#     echo "Prawda."
# else
#     echo "Falsz."
# fi

# if [ VAR1 -lt VAR2 ] ; then
#     echo "Prawda."
# else
#     echo "Falsz."
# fi

# echo "Pierwsza zmienna wprowadzona z klawiatury: \$1 = $1."

if [ $1 -lt 100 ] ; then
    echo "Wprowadzona liczba jest mniejsza od 100."
else
    echo "Wprowadzona liczba jest wieksza lub rowna 100."
fi

echo "Jesteśmy za 'if'."
