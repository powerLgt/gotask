// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract NFT {
    string name;
    string symbol;
    uint counter;

    mapping (address => uint) userBalance;
    mapping (uint => address) userApprove;
    mapping (uint => address) nftOwners;
    mapping (uint => string) tokenUris;

    constructor(string memory _name, string memory _symbol)  {
        name = _name;
        symbol = _symbol;
    }

    function balanceOf(address owner) public view returns (uint256 balance) {
        return userBalance[owner];
    }

    function ownerOf(uint256 tokenId) public view returns (address owner) {
        return nftOwners[tokenId];
    }

    function approve(address to, uint256 tokenId) public {
        require(ownerOf(tokenId) == msg.sender, "the nft don't belong to you");
        userApprove[tokenId] = to;
    }

    function mintNFT(address recipient, string memory _tokenUri) public  {
        uint tokenId = ++counter;
        tokenUris[tokenId] = _tokenUri;
        userBalance[recipient]++;
        nftOwners[tokenId] = recipient;
    }

    function getTokenUri(uint tokenId) public view returns(string memory) {
        return tokenUris[tokenId];
    }
}