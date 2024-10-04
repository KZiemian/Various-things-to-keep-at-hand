#!/usr/bin/python3

auxiliaryList1 = []
auxiliaryList2 = []

# for i in range(10):
#     auxiliaryList.append(str(i))

for i in range(1, 10):
    for j in range(1, 10):
        auxiliaryList1.append(str(i) + str(j))

# print(auxiliaryList)

for i in range(1, 10):
    for elem in auxiliaryList1:
        auxiliaryList2.append(str(i) + elem)

with open("PINy-lista.txt", "w") as file1, \
     open("PINy-for-Exel.txt", "w") as file2:
    for i in range(1, 10):
        for elem in auxiliaryList2:
            file1.write(str(i) + elem + "\n")
            file2.write(str(i) + elem + "\r\n")
