// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract MyERC1155Token is ERC1155 {
    constructor() ERC1155("uri") {
        uint256 decimals = 18;  // Örnek bir ondalık basamak sayısı
        _mint(msg.sender, 1, 1000000000 * (10 ** decimals), "");
    }
}
