'use strict'

const intcodeProgram = () => {
    let _memory = [];

    const init = (memory) => {
        _memory = memory.slice();
    }

    const run = (address) => {
        let instruction = _memory[address];
        
        if (instruction === 99 || instruction == null)
        {
            return;
        }
        else if (instruction !== 1 && instruction !== 2) {
            run(address+1);
        }

        let para1address = _memory[address+1];
        let para2address = _memory[address+2];
        let para3address = _memory[address+3];

        if (instruction === 1) {
            _memory[para3address] = _memory[para1address] + _memory[para2address];
        }
        else if (instruction === 2) {
            _memory[para3address] = _memory[para1address] * _memory[para2address];
        }

        run(address + 4);
    }

    const get = (address) => {
        return _memory[address];
    }

    const set = (address, value) => {        
        _memory[address] = value;
    }

    return {
        init,
        run,
        get,
        set
    }
};

module.exports = intcodeProgram();