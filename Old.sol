// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.2 <0.9.0;

contract MyContract{

    uint256 a;
    uint256 b;
    uint256 c;

    event DataStored(uint256 val);

    constructor(){
        a = 5;
        b = 3;
        c = 0;
    }

    function compute() public {
        
        c = a+b;
        emit DataStored(c);

    }

    function getResult() public view returns (uint256){

        return c;
    }
}