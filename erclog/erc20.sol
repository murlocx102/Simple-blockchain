pragma solidity ^0.4.24;

// 创建ERC-20智能合约的事件日志的interface文件
contract ERC20 {
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}