pragma solidity ^0.4.19;

library MiniSafe {
    function add(uint a, uint b) internal pure returns (uint) {
      uint c = a + b;
      assert(c >= a);
      return c;
    }
  }

contract DogsOfRome {
    using MiniSafe for uint256;

    uint256   public counter;
    uint256   public lastQBlock;
    address   public romulus;
    address   public owner = msg.sender;
    mapping   (address => uint256) public balance;
    uint256   public usage_fee;


    modifier onlyOwner() {
        require (msg.sender==owner);
        _;
    }

    modifier roman() {
        require (msg.sender==romulus);
        _;
    }

    function mustHaveEnough() internal {
        require(balance[msg.sender].add(msg.value) >= usage_fee);
        balance[msg.sender] += msg.value;
        balance[msg.sender] -= usage_fee;
    }

    event Romulus(address indexed asker,uint256 nonce,string question);
    event Scooby(address indexed asker,uint256 nonce,string question);
    event Remus(address indexed asker,uint256 nonce,string answer);

    function DogsOfRome(address _auth) public {
        romulus = _auth;
        usage_fee = 100 finney;
        lastQBlock = block.number;
    }

    function setUsageRate(uint256 newRate) onlyOwner public {
        usage_fee = newRate;
    }

    function askRomulus(string question) public payable {
        mustHaveEnough();
        Romulus(msg.sender,counter,question);
        counter++;
    }

    function askScoobyDoo(string question) public payable {
        mustHaveEnough();
        Scooby(msg.sender,counter,question);
        counter++;
    }

    function letRomulusReply(address asker, uint256 question, string answer, uint256 questionBlock ) public roman {
        Remus(asker,question,answer);
        lastQBlock = questionBlock > lastQBlock ? questionBlock : lastQBlock;
    }
    
    function withdraw() public {
        uint256 bal = balance[msg.sender];
        balance[msg.sender] = 0;
        msg.sender.transfer(bal);
    }
}