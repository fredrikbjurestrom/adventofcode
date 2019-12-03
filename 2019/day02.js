'use strict'

const fs = require('fs');
const IntcodeProgram = require('./utils/IntcodeProgram');

const input = fs.readFileSync("day02input.txt").toString().split(",").map(x => parseInt(x));

const operandResolver = (expected) => {
    let intcodeInstance = new IntcodeProgram(input);

    for (let noun = 0; noun < 100; noun++) {
        for (let verb = 0; verb < 100; verb++) {
            intcodeInstance.init(input);
            intcodeInstance.set(1, noun);
            intcodeInstance.set(2, verb);
            
            intcodeInstance.run(0);
            
            if (intcodeInstance.get(0) === expected) {
                return { noun, verb };
            }
        }   
    }
};

const partOne = (input) => {
    let intcodeInstance = new IntcodeProgram(input);
    
    // Revert to "1202 program alarm" state
    intcodeInstance.set(1, 12);
    intcodeInstance.set(2, 2);

    intcodeInstance.run(0);

    console.log("Part 1: " + intcodeInstance.get(0));
};

const partTwo = (expected) => {
    let result = operandResolver(expected);
    let reportValue = result.noun*100 + result.verb;
    console.log("Part 2: " + reportValue)
};


partOne(input);

partTwo(19690720);