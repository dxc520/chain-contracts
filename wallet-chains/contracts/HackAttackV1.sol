// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

 

contract HackAttackV1 {
     

    event receiveEvent(address indexed sender,uint256 indexed value );
    event fallbackEvent(address indexed sender,uint256 indexed value );


    constructor(address _erc1155Addr, address _erc721Addr) {
        
    }
 
    receive() external payable  {
       emit receiveEvent(msg.sender,msg.value);
    }
    
    fallback()external payable{
        emit fallbackEvent(msg.sender, msg.value);
    }
}
