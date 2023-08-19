// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MyERC721Token is ERC721 {
    constructor() ERC721("My ERC721 Token", "M721") {
        _mint(msg.sender, 1);
    }
}
