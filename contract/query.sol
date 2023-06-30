pragma solidity 0.6.12;

interface IVault {
    struct Position {
        uint256 size;
        uint256 collateral;
        uint256 averagePrice;
        uint256 entryFundingRate;
        uint256 reserveAmount;
        int256 realisedPnl;
        uint256 lastIncreasedTime;
    }

    function cumulativeFundingRates(address collateralToken) external view returns (uint256);
    
    function positions(bytes32 positionKey) external view returns (uint256, uint256, uint256, uint256, uint256, int256, uint256);

    function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) external pure returns (bytes32);

    function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) external view returns (bool, uint256);

    function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) external view returns (uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256);
}

contract Query {

    uint256 public constant FUNDING_RATE_PRECISION = 1000000;

    IVault public vault;

    constructor(address vaultAddress) public {
        vault = IVault(vaultAddress);
    }

    /* 
    return:
        size, 
        collateral, 
        averagePrice, 
        entryFundingRate, 
        reserveAmount, 
        realisedPnl, 
        lastIncreasedTime, 
        floatingPnl, 
        uncollectedFundingFee
    */
    function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) public view 
    returns (uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime, int256 floatingPnl, uint256 uncollectedFundingFee) {
        return abi.decode(_getPosition(_account, _collateralToken, _indexToken, _isLong), (uint256, uint256, uint256, uint256, uint256, int256, uint256, int256, uint256));
    }

    // margin fee is not inclueded in realisedPnl & floatingPnl
    // return value = abi.encode(size, collateral, averagePrice, entryFundingRate, reserveAmount, realisedPnl, lastIncreasedTime, floatingPnl, uncollectedFundingFee)
    function _getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) internal view returns (bytes memory) { 
        bytes memory returnValue;
        {
        bytes32 positionKey = vault.getPositionKey(_account, _collateralToken, _indexToken, _isLong);
        (bool success, bytes memory res) = address(vault).staticcall(abi.encodeWithSignature("positions(bytes32)", positionKey));
        require(success, "fail to get position");
        returnValue = res;
        }
        returnValue = _addFloatingPnlToRes(returnValue, _indexToken, _isLong);
        returnValue = _addUncollectedFundingFeeToRes(returnValue, _collateralToken);
        return returnValue;
    }


    function _addFloatingPnlToRes(bytes memory resBefore, address _indexToken, bool _isLong) internal view returns (bytes memory) {
        (uint size, ,uint averagePrice, , , ,uint lastIncreasedTime) = abi.decode(resBefore, (uint256, uint256, uint256, uint256, uint256, int256, uint256));
        (bool hasProfit, uint256 delta) = vault.getDelta(_indexToken, size, averagePrice, _isLong, lastIncreasedTime);
        int256 floatingPnl = int256(delta) * (hasProfit ? int256(1) : int256(-1));
        return abi.encodePacked(resBefore, abi.encode(floatingPnl));
    }

    function _addUncollectedFundingFeeToRes(bytes memory resBefore, address _collateralToken) internal view returns (bytes memory) {
        (uint size, , , uint entryFundingRate, , , , ) = abi.decode(resBefore, (uint256, uint256, uint256, uint256, uint256, int256, uint256, int256));
        uint256 fundingRate = vault.cumulativeFundingRates(_collateralToken) - entryFundingRate;
        uint256 uncollectedFundingFee = (size * fundingRate) / FUNDING_RATE_PRECISION;
        return abi.encodePacked(resBefore, abi.encode(uncollectedFundingFee)); 
    }
}