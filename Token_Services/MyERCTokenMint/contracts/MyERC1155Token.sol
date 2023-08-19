// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

import '@openzeppelin/contracts/token/ERC1155/ERC1155.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract MyERC1155Token is ERC1155, Ownable {
    constructor() ERC1155('https://example.com/tokens/{id}.json') {
        _mint(msg.sender, 1, 100, ""); // Başlangıçta 100 adet token mint
    }

    function mint(address to, uint256 id, uint256 amount, bytes memory data) public onlyOwner {
        _mint(to, id, amount, data);
    }

    function burn(address from, uint256 id, uint256 amount) public onlyOwner {
        _burn(from, id, amount);
    }
}
