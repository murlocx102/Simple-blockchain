pragma solidity ^0.4.24;

// 基于ERC20编写的代币合约
contract CoinInterface{  
    uint256 public totalSupply;  

    function balanceOf(address _owner) public constant returns (uint256 balance);  
    function transfer(address _to, uint256 _value) public returns (bool success);  
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);  
    function approve(address _spender, uint256 _value) public returns (bool success); 
    function allowance(address _owner, address _spender) public constant returns (uint256 remaining);  
    
    event Transfer(address indexed _from, address indexed _to, uint256 _value);  
    event Approval(address indexed _owner, address indexed _spender, uint256 _value); 
}  

contract TestDemo is CoinInterface {
    string public name;     // 币的名称  
    uint8 public decimals;  // 币使用的小数点后几位,精度.  
    string public symbol;   // 币的简称

    // 设置初始参数(币总量,币名称,币精度,币简称)
    function TestDemo(uint256 _initialAmount, string _coinName, uint8 _decimalUnits, string _coinSymbol) public {
        // totalSupply 默认不会超过最大值 (2^256 - 1).
        totalSupply = _initialAmount * 10 ** uint256(_decimalUnits);
        balances[msg.sender] = totalSupply;   // 将初始币数量给予合约创建者.

        name = _coinName;                     
        decimals = _decimalUnits;            
        symbol = _coinSymbol;  
    }  

    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(balances[msg.sender] >= _value && balances[_to] + _value > balances[_to]); // 支出地址的币要大于或等于转出的币数量.避免溢出的异常
        require(_to != 0x0);                // 禁止转到销毁币的地址
        balances[msg.sender] -= _value;     // 从合约创建者账户中减去数量为_value的币.(初始化已将总币给予了合约创建者) 
        balances[_to] += _value;            // 往接收账户增加数量为_value的币.  
        Transfer(msg.sender, _to, _value);  // 触发转币交易事件

        return true;  
    }  


    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {  
        require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value);  
        balances[_to] += _value;                //接收账户_to增加_value量  
        balances[_from] -= _value;              //支出账户_from减去_value量  
        allowed[_from][msg.sender] -= _value;   //合约创建者可以从账户_from中转出减少_value数量
        Transfer(_from, _to, _value);           //触发转币交易事件

        return true;  
    }

    function balanceOf(address _owner) public constant returns (uint256 balance) {  
        return balances[_owner];  
    }  


    function approve(address _spender, uint256 _value) public returns (bool success) {   
        allowed[msg.sender][_spender] = _value;  
        Approval(msg.sender, _spender, _value);

        return true;  
    }  

    function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {  
        return allowed[_owner][_spender]; // 允许_spender从_owner中转出的币数量
    }

    mapping (address => uint256) balances;  
    mapping (address => mapping (address => uint256)) allowed;
}