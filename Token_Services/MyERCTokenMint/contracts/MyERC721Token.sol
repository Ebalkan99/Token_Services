// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

import '@openzeppelin/contracts/token/ERC721/ERC721.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract MyERC721Token is ERC721, Ownable {
    uint256 public mintPrice;
    uint256 public totalSupply;
    uint256 public maxSupply;
    uint8 public maxMint;
    bool public isMyERC721TokenEnabled;
    string internal baseTokenUri;
    mapping(address => uint8) public walletMints;

    constructor() payable ERC721('MyERC721Token', 'ME7T') {
        mintPrice = 0.01 ether;
        totalSupply = 0;
        maxSupply = 1000;
        maxMint = 3;
        isMyERC721TokenEnabled = true; // Varsayılan olarak mint işlemi etkin
        baseTokenUri = "https://example.com/tokens/"; // Varsayılan token URI
    }

    function setIsMyERC721TokenEnabled(bool isMyERC721TokenEnabled_) external onlyOwner {
        isMyERC721TokenEnabled = isMyERC721TokenEnabled_;
    }

    function setBaseTokenUri(string calldata newBaseTokenUri) external onlyOwner {
        baseTokenUri = newBaseTokenUri;
    }

    function tokenURI(uint256 tokenId_) public view override returns (string memory) {
        require(_exists(tokenId_), 'Token does not exist');
        return string(abi.encodePacked(baseTokenUri, Strings.toString(tokenId_)));
    }


    function mint(uint8 quantity_) public payable {
        require(isMyERC721TokenEnabled, 'Minting is not enabled');
        require(msg.value == quantity_ * mintPrice, 'Insufficient minting fee');
        require(totalSupply + quantity_ <= maxSupply, 'Exceeds max supply');
        require(walletMints[msg.sender] + quantity_ <= maxMint, 'Exceeds max mint per wallet');

        for (uint8 i = 0; i < quantity_; i++) {
            uint256 newTokenId = totalSupply + 1;
            totalSupply++;
            _safeMint(msg.sender, newTokenId);
            walletMints[msg.sender]++;
        }
    }

    function withdraw() external onlyOwner {
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, 'Withdrawal failed');
    }
}
