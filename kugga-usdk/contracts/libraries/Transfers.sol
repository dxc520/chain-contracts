// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";

import "./Errors.sol";

library Transfers {
    using SafeERC20Upgradeable for IERC20Upgradeable;
    using AddressUpgradeable for address payable;

    /// @notice Transfers (or checks sent value) given asset from sender to running contract
    /// @param asset Asset to transfer (address(0) to check native sent value)
    /// @param amount Amount to transfer
    /// @return extraValue Extra amount of native token passed
    function transferIn(address asset, uint256 amount)
        internal
        returns (uint256 extraValue)
    {
        if (isNative(asset)) {
            require(msg.value >= amount, Errors.INVALID_MESSAGE_VALUE);
            return msg.value - amount;
        } else {
            uint256 balanceBefore = IERC20Upgradeable(asset).balanceOf(
                address(this)
            );
            IERC20Upgradeable(asset).safeTransferFrom(
                msg.sender,
                address(this),
                amount
            );
            (bool isSuccess, uint256 newBalance) = SafeMathUpgradeable.trySub(
                IERC20Upgradeable(asset).balanceOf(address(this)),
                balanceBefore
            );
            require(isSuccess, Errors.INVALID_RECEIVED_AMOUNT);
            require(newBalance == amount, Errors.INVALID_RECEIVED_AMOUNT);
            return msg.value;
        }
    }

    /// @notice Transfers given token from running contract to given address
    /// @param asset Asset to transfer (address(0) to transfer native token)
    /// @param to Address to transfer to
    /// @param amount Amount to transfer
    function transferOut(
        address asset,
        address to,
        uint256 amount
    ) internal {
        if (isNative(asset)) {
            payable(to).sendValue(amount);
        } else {
            IERC20Upgradeable(asset).safeTransfer(to, amount);
        }
    }

    /// @notice Approves given token to given spender (with checks for address(0) as native)
    /// @param asset Token to approve
    /// @param spender Spender address
    /// @param amount Amount to approve
    function approve(
        address asset,
        address spender,
        uint256 amount
    ) internal {
        if (isNative(asset)) {
            return;
        }

        uint256 allowance = IERC20Upgradeable(asset).allowance(
            address(this),
            spender
        );
        if (allowance > 0) {
            // https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
            IERC20Upgradeable(asset).safeApprove(spender, 0);
        }
        IERC20Upgradeable(asset).safeIncreaseAllowance(spender, amount);
    }

    /// @notice Gets balance of given token
    /// @param asset Token to get balance of (address(0) for native token)
    function getBalance(address asset) internal view returns (uint256) {
        if (isNative(asset)) {
            return address(this).balance;
        } else {
            return IERC20Upgradeable(asset).balanceOf(address(this));
        }
    }

 function getBalanceWithOwner(address asset,address owner) internal view returns (uint256) {
        if (isNative(asset)) {
            return owner.balance;
        } else {
            return IERC20Upgradeable(asset).balanceOf(owner);
        }
    }


    function getAllowance(address asset, address owner)
        internal
        view
        returns (uint256)
    {
        if (isNative(asset)) {
            return address(this).balance;
        } else {
            return IERC20Upgradeable(asset).allowance(owner, address(this));
        }
    }

    function isNative(address asset) internal pure returns (bool) {
        return asset == address(0);
    }
}
