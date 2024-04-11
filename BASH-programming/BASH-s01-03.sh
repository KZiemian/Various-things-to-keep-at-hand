#!/bin/bash

# if [[ 1 < 2 ]] ; then
#     echo "Prawda."
# else
#     echo "Fałsz."
# fi

# if [[ 2 < 1 ]] ; then
#     echo "Prawda."
# else
#     echo "Fałsz."
# fi

# if [[ 2 > 3 ]] ; then
#     echo "Prawda."
# else
#     echo "Fałsz."
# fi

# if [[ 1 < 2 ]] ; then
#     echo "Warunek logiczny \"<\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek logiczny \"<\" ma wartość logiczną \"Fałsz\"."
# fi

# VAR1=001
# VAR2=1
# VAR3=

# if [ -n $VAR4 ] ; then
#     echo "Warunek \"-n\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-n\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ $VAR1 = $VAR2 ] ; then
#     echo "Warunek \"=\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunke \"=\" ma wartość logcizną \"Fałsz\"."
# fi

# if [ $VAR1 -eq $VAR2 ] ; then
#     echo "Warunek \"-eq\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-eq\" ma wartość logiczną \"Fałsz\"."
# fi

# if [ 1 -lt 1 ] ; then
#     echo "Warunek \"-lt\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-lt\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ 1 -lt 2 ]] ; then
#     echo "Warunek \"-lt\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-lt\" ma wrtość logiczną \"Fałsz\"."
# fi

# if [[ 2 -lt 1 ]] ; then
#     echo "Warunek \"-lt\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-lt\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ 1 -lt 1 ]] ; then
#     echo "Warunke \"-lt\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-lt\" ma wartość logiczną \"Fałsz\"."
# fi

VAR1=001
VAR2=1

# if [[ $VAR1 = $VAR2 ]] ; then
#     echo "Warunek \"=\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"=\" ma wartość logiczną \"Fałsz\"."
# fi

# if [[ $VAR1 -eq $VAR2 ]] ; then
#     echo "WArunek \"-eq\" ma wartość logiczną \"Prawda\"."
# else
#     echo "Warunek \"-eq\" ma wartość logiczną \"Fałsz\"."
# fi

STRING1=001
STRING2=1
STRING3=""
STRING4=

if [[ -n $STRING1 ]] ; then
    echo "Warunek \"-n \$STRING1\" ma wartość logiczną \"Prawda\"."
else
    echo "Warunek \"-n \$STRING1\" ma wartość logiczną \"Fałsz\"."
fi

if [[ -n $STRING2 ]] ; then
    echo "Warunek \"-n \$STRING2\" ma wartość logiczną \"Prawda\"."
else
    echo "Warunek \"-n \$SRTING2\" ma wartość logiczną \"Fałsz\"."
fi

if [[ -n $STRING3 ]] ; then
    echo "Warunek \"-n \$STRING3\" ma wartość logiczną \"Prawda\"."
else
    echo "Warunek \"-n \$STRING3\" ma wartość logiczną \"Fałsz\"."
fi

if [[ -n $STRING4 ]] ; then
    echo "Warunek \"-n \$STRING4\" ma wartość logiczną \"Prawda\"."
else
    echo "Warunek \"-n \$STRING4\" ma wartość logiczną \"Fałsz\"."
fi

if [[ -n $STRING5 ]] ; then
    echo "Warunek \"-n \$STRING5\" ma wartość logiczną \"Prawda\"."
else
    echo "Warunek \"-n \$STRING5\" ma wartość logiczną \"Fałsz\"."
fi
