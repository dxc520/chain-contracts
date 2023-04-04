// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./AccessControlUpgradeable.sol";

contract AccessControlMixinUpgradeable is AccessControlUpgradeable {
    function __AccessControlMixin_init() internal onlyInitializing {
    }

    function __AccessControlMixin_init_unchained() internal onlyInitializing {
    }
    string private _revertMsg;
    function _setupContractId(string memory contractId) internal {
        _revertMsg = string(abi.encodePacked(contractId, ": INSUFFICIENT_PERMISSIONS"));
    }

    modifier only(bytes32 role) {
        require(
            hasRole(role, _msgSender()),
            _revertMsg
        );
        _;
    }
}
