string = ''

for i in range(1, 151):
    if (i % 3) == 0:
        string += "fizz"
    if (i % 5) == 0:
        string += "buzz"
    if (i % 7) == 0:
        string += "baz"

print (string)
