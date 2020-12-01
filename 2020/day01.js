const fs = require('fs');

var input = fs.readFileSync("./input/day01input.txt").toString().split("\n").map(x => parseInt(x));

const findAddends = (input, noOfItems, sum) => {
  if (noOfItems < 1 || noOfItems > 3) {
    throw Error('Search depth not implemented')
  }

  for (let i = 0; i < input.length; i++) {
    if (noOfItems === 1 && input[i] === sum) {
      return [input[i]];
    }

    for (let j = 0; j < input.length; j++) {
      if (noOfItems === 2 && i !== j && input[i] + input[j] === sum) {
        return [input[i], input[j]];
      }

      for (let k = 0; k < input.length; k++) {
        if (noOfItems === 3 && j !== k && input[i] + input[j] + input[k] === sum) {
          return [input[i], input[j], input[k]];
        }
      }
    }
  }
}

const partOne = findAddends(input, 2, 2020);
const partTwo = findAddends(input, 3, 2020);

console.log(`Part 1: ${partOne.join('*')} = ${partOne.reduce((a, b) => a * b, 1)}`)
console.log(`Part 2: ${partTwo.join('*')} = ${partTwo.reduce((a, b) => a * b, 1)}`)