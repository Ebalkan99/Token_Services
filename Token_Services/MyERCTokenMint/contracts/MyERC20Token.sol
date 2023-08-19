// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract MyERC20Token is ERC20, Ownable {
    constructor() ERC20('MyERC20Token', 'ME2T') {
        _mint(msg.sender, 1000000 * 10 ** decimals()); // Başlangıçta 1.000.000 token mint
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function burn(address from, uint256 amount) public onlyOwner {
        _burn(from, amount);
    }
}
