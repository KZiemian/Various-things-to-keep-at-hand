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

for elem in 2, 3, 5, 7, 11, 13, 17, 19 ; do
    echo "prime number -> $elem."

    echo "elem + 1 = $(($elem + 1))."
done

VAR1="Ty"
echo "\$VAR1 = $((VAR1))."
