#!/bin/bash

# STRING1=0001
# STRING2=1
# STRING3=""
# STRING4=

# echo "STRING1 = $STRING1."
# echo "STRING2 = $STRING2."
# echo "STRING3 = $STRING3."
# echo "STRING4 = $STRING4."
# echo "STRING5 = $STRING5."

# if [ -n $STRING1 ] ; then
#     echo "Warunek \"-n \$STRING1\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-n \$STRING1\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -n $STRING2 ] ; then
#     echo "Warunek \"-n \$STRING2\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-n \$STRING2\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -n $STRING3 ] ; then
#     echo "Warunek \"-n \$STRING3\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-n \$STRING3\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -n $STRING4 ] ; then
#     echo "Warunek \"-n \$STRING4\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-n \$STRING4\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -n $STRING5 ] ; then
#     echo "Warunek \"-n \$STRING5\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-n \$STRING%\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -z $STRING1 ] ; then
#     echo "Warunek \"-z \$STRING1\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING1\" ma wartość logiczną \"Prawda\"."
# fi

# if [ -z $STRING2 ] ; then
#     echo "Warunek \"-z \$STRING2\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING2\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -z $STRING3 ] ; then
#     echo "Warunek \"-z \$STRING3\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING4\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -z $STRING4 ] ; then
#     echo "Warunek \"-z \$STRING3\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING3\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ -z $STRING5 ] ; then
#     echo "Warunek \"-z \$STRING3\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING3\" ma wartość logiczną \"Fałsz\"."
# fi

# echo -e "\nNowa seria."

# if [[ -z $STRING1 ]] ; then
#     echo "Warunek \"-z \$STRING1\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING1\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ -z $STRING2 ]] ; then
#     echo "Warunek \"-z \$STRING2\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING2\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ -z $STRING3 ]] ; then
#     echo "Warunek \"-z \$STRING3\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING4\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ -z $STRING4 ]] ; then
#     echo "Warunek \"-z \$STRING4\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING4\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ -z $STRING5 ]] ; then
#     echo "Warunek \"-z \$STRING5\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-z \$STRING5\" ma wartość logiczną \"Fałsz\"."
# fi

VAR1=-5

if [[ 0 -lt $VAR1 ]] && [[ $VAR1 -lt 10 ]] ; then
    echo "Liczba całkowita $VAR1 zawiera się w przedziałem [0, 10]."
else
    echo "Liczba całkowita $VAR1 leży poza przedziałem [0, 10]."
fi
