// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

interface IChildTokenUpgradeable {
    function deposit(address user, bytes calldata depositData) external;
}