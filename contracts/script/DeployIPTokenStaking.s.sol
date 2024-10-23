// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import { IPTokenStaking } from "../src/protocol/IPTokenStaking.sol";

// forge script script/DeployIPTokenStaking.s.sol --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast
// mininet
// forge script script/DeployIPTokenStaking.s.sol --rpc-url http://54.215.121.164:8545 --private-key <PRIVATE_KEY> --broadcast
contract DeployIPTokenStaking is Script {
    function run() public {
        // Retrieve constructor arguments
        uint256 stakingRounding = 1 gwei;
        uint256 defaultMinFee = 1 ether;

        // Deploy the contract
        vm.startBroadcast();
        IPTokenStaking staking = new IPTokenStaking(stakingRounding, defaultMinFee);
        console.log("IPTokenStaking deployed at:", address(staking));
        vm.stopBroadcast();
    }
}