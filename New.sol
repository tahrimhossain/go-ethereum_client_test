// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.2 <0.9.0;

contract MyContract{

    uint256 a;
    uint256 b;
    uint256 c;
    event DataStored(uint256 val);

    function compute() public {
        
        c = a-b;
        emit DataStored(c);
        
    }
}