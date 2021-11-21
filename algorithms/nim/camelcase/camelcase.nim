let input = "oneTwoThreeFourFive"

var words = 0
for letter in input:
  let l = int(letter)
  if l <= 96:
    words += 1

echo "There are ", words + 1, " words in the input string."
