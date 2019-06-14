var string = "";

for (i=1 ; i <= 150 ; i++) {
    if ((i % 3) == 0) {
	string += "fizz"
    }
    if ((i % 5) == 0) {
	string += "buzz"
    }
    if ((i % 7) == 0) {
	string += "baz"
    }
}

console.log(string)
