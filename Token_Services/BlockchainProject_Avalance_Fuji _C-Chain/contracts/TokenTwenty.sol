// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TokenTwenty is ERC20 {
    constructor() ERC20("TokenTwenty", "TT") {
        _mint(msg.sender, 1000000000 * 10 ** uint(decimals()));
    }
}
