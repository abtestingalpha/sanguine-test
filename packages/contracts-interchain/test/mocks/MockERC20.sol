// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// @notice Simple ERC20 token for testing purposes with a public mint function
contract MockERC20 is ERC20 {
    constructor(string memory name_) ERC20(name_, name_) {}

    /// @notice For testing purposes, mints tokens to the given account
    function mintPublic(address account, uint256 amount) public {
        _mint(account, amount);
    }
}
