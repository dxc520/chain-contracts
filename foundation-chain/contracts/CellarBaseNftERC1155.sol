// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "./ERC1155URIStorageUpgradeable.sol";

contract CellarBaseNftERC1155 is  OwnableUpgradeable, AccessControlUpgradeable, ERC1155URIStorageUpgradeable {
    bytes32 private constant MINTER_ROLE = keccak256("MINTER_ROLE"); //铸币角色

    mapping (uint256 => bool) private _hasMinted;

    
    event Gift(address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount);
    event GiftBatch(address indexed from, address indexed to, uint256 [] tokenIds, uint256 [] amount);

    function _init_base_unchanged()internal onlyInitializing{
                __Ownable_init();
                _setupRole(MINTER_ROLE, _msgSender());//owner，可以铸币

    }

     modifier isMinterRole {
        require(
            hasRole(MINTER_ROLE, _msgSender()),
            "User must have minter role to mint"
        );
        _;
    }

    modifier isCallerRoleOrOwner(address _operator) {
        require(
            hasRole(MINTER_ROLE, _msgSender()) || isApprovedForAll(_operator, _msgSender()),
            "User must have caller role or is the approvedOwner"
        );
        _;
    }
    /**
     * 查询是否拥有铸币权限
     */
    function hasMintRole(address _account) public view returns (bool) {
        return hasRole(MINTER_ROLE,_account);
    }

    /**
     *  授予用户铸币权限
     */
    function grantMintRole(address operator)  public onlyOwner {
        require(address(0)!=operator);
        _setupRole(MINTER_ROLE, operator);//owner，可以铸币
    }

    /**
     *  收回用户铸币权限
     */
    function renounceMintRole( address account) public onlyOwner {
        renounceRole(MINTER_ROLE, account);
    }

    function _checkedMinted(uint256  id)internal view returns (bool){
        return _hasMinted[id];
    }

    function _checkedBatchMinted(uint256 [] memory  ids)internal view returns (bool){
        for (uint256 i = 0; i < ids.length; ++i) {
            if (_hasMinted[ids[i]]){
                return true;
            }
        }   
        
        return false;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC1155Upgradeable, AccessControlUpgradeable)  returns (bool) {
        return interfaceId == type(IAccessControlUpgradeable).interfaceId
            || interfaceId == type(IERC1155Upgradeable).interfaceId
            || interfaceId == type(IERC1155MetadataURIUpgradeable).interfaceId
            || super.supportsInterface(interfaceId);
    }

    function _mintExt(
        address to,
        uint256 id,
        uint256 amount,
        string memory uri
    ) internal virtual {
        require(!_checkedMinted(id),"ERC1155: token already minted.");
         _mint(to, id, amount, "");
         _setTokenURI(id, uri);
         _hasMinted[id]=true;
    }

    /**
     * @dev xref:ROOT:erc1155.adoc#batch-operations[Batched] version of {_mint}.
     *
     * Requirements:
     *
     * - `ids` and `amounts` must have the same length.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155BatchReceived} and return the
     * acceptance magic value.
     */
    function _mintBatchExt(
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        string [] memory uris
    ) internal virtual {
        require(!_checkedBatchMinted(ids),"ERC1155:token already minted.");

        _mintBatch(to, ids, amounts, "");
        _setBatchTokenURI(ids, uris);
        _setBatchMinted(ids);

    }

    function _setBatchMinted(uint256 []memory tokenIds) internal virtual {
        
        for (uint256 i = 0; i < tokenIds.length; ++i) {
            _hasMinted[tokenIds[i]]=true;
        }
    }

    function gift(
        address from,
        address to,
        uint256 id,
        uint256 amount
    ) public virtual  {
        super.safeTransferFrom(from,to,id,amount,"");
        emit Gift(from,to,id,amount);
    }

    function giftBatch(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory amounts
    ) public virtual   {
        super.safeBatchTransferFrom(from,to,ids,amounts,"");
        emit GiftBatch(from,to,ids,amounts);
    }

}