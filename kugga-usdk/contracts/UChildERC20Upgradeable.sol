// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "./IChildTokenUpgradeable.sol";
import "./AccessControlMixinUpgradeable.sol";
import "./NativeMetaTransactionUpgradeable.sol";
import "./ChainConstants.sol";
import "./ContextMixinUpgradeable.sol";
 
contract UChildERC20Upgradeable is
    ERC20Upgradeable,
    IChildTokenUpgradeable,
    AccessControlMixinUpgradeable,
    NativeMetaTransactionUpgradeable,
    ChainConstants,
    ContextMixinUpgradeable
{
    bytes32 public constant DEPOSITOR_ROLE = keccak256("DEPOSITOR_ROLE");
    uint8 private _decimals;
    //constructor() public ERC20("", "") {}

    /**
     * @notice Initialize the contract after it has been proxified
     * @dev meant to be called once immediately after deployment
     */
    function initialize(
        string calldata name_,
        string calldata symbol_,
        uint8 decimals_,
        address childChainManager
    )
        external
        initializer
    {
      __ERC20_init(name_,symbol_);
      _setDecimals(decimals_);
      _setupContractId(string(abi.encodePacked("Child", symbol_)));
      _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
      _setRoleAdmin(DEPOSITOR_ROLE, DEFAULT_ADMIN_ROLE);
      _setupRole(DEPOSITOR_ROLE, childChainManager);
      _initializeEIP712(name_, ERC712_VERSION);
    }

    // This is to support Native meta transactions
    // never use msg.sender directly, use _msgSender() instead
    function _msgSender()
        internal
        override
        view
        returns (address sender)
    {
        return ContextMixinUpgradeable.msgSender();
    }
/*
    function changeName(string calldata name_) external only(DEFAULT_ADMIN_ROLE) {
        //_setName(name_);
        _setDomainSeperator(name_, ERC712_VERSION);
    }
    */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
    function setDecimals(uint8 newDecimals) external only(DEFAULT_ADMIN_ROLE){
      _setDecimals(newDecimals);
    }
    function _setDecimals(uint8 newDecimals) internal {
      _decimals = newDecimals;
    }
    /**
     * @notice called when token is deposited on root chain
     * @dev Should be callable only by ChildChainManager
     * Should handle deposit by minting the required amount for user
     * Make sure minting is done only by this function
     * @param user user address for whom deposit is being done
     * @param depositData abi encoded amount
     */
    function deposit(address user, bytes calldata depositData)
        external
        override
        only(DEPOSITOR_ROLE)
    {
        uint256 amount = abi.decode(depositData, (uint256));
        _mint(user, amount);
    }

  function deposit2(address user, uint256 amount2)
        external
        only(DEPOSITOR_ROLE)
    {
       // bytes  depositData = ;
        uint256 amount = abi.decode(abi.encode(amount2), (uint256));
        _mint(user, amount);
    }

    /**
     * @notice called when user wants to withdraw tokens back to root chain
     * @dev Should burn user's tokens. This transaction will be verified when exiting on root chain
     * @param amount amount of tokens to withdraw
     */
    function withdraw(uint256 amount) external {
        _burn(_msgSender(), amount);
    }
}