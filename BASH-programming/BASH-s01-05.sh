#!/bin/bash

# COUNTER=1

# while [[ $COUNTER -le 10 ]] ; do
#     echo "COUNTER = $COUNTER."
#     ((COUNTER++))
# done

# echo "Pętla ukończona."

# COUNTER=1
# touch testFile.txt

# while [[ -e testFile.txt ]] ; do
#     echo "Plik test.txt istnieje."
#     echo "COUNTER = $COUNTER"

#     if [[ $COUNTER -eq 5 ]] ; then
# 	rm testFile.txt
# 	echo -e "\nPlik \"testFile.txt\" przestał istnieć."
#     fi

#     ((COUNTER++))
# done

# if [[ 1 -lt 2 ]] ; then
#     echo "Warunek \"-lt\" ma wartość logiczną \"Prawda\"."

#     echo "Teraz trochę arytmetyki."

#     VAR1=1
#     VAR2=2

#     echo "VAR1 = $VAR1, VAR2 = $VAR2."
#     echo "VAR1 + VAR2 = $(($VAR1 + $VAR2))"
# else
#     echo "Warunek \"-lt\" ma wartość logiczną \"Fałsz\"."
# fi

if [[ 2 -lt 1 ]] ; then
    echo "Warunek \"-lt\" ma wartość logiczną \"Prawda\"."
else
    echo "Warunek \"-lt\" ma wartość logiczną \"Fałsz\"."
    echo "Teraz trochę programowani w gałęzi 'else'."

    VAR3=3
    VAR4=4

    echo "VAR4 - VAR3 = $(($VAR4 - $VAR3))."
fi
