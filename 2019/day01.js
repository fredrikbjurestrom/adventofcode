const fs = require('fs');

var array = fs.readFileSync("./inputs/day01input.txt").toString().split("\n");;
console.log(array.filter(x => x !== '').reduce((acc, cur) => {return acc+(Math.floor(cur/3)-2)}, 0));