// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Script} from "forge-std/Script.sol";
import {OPSuccinctL2OutputOracle} from "../src/OPSuccinctL2OutputOracle.sol";
import {Utils} from "../test/helpers/Utils.sol";
import {Proxy} from "@optimism/src/universal/Proxy.sol";
import {console} from "forge-std/console.sol";

contract OPSuccinctUpgrader is Script, Utils {
    function run() public {
        vm.startBroadcast();

        address l2OutputOracleProxy = vm.envAddress("L2OO_ADDRESS");

        console.log("L2OO_ADDRESS:", l2OutputOracleProxy);

        bool executeUpgradeCall = vm.envOr("EXECUTE_UPGRADE_CALL", true);

        address OPSuccinctL2OutputOracleImpl = address(
            new OPSuccinctL2OutputOracle()
        );

        if (executeUpgradeCall) {
            Proxy existingProxy = Proxy(payable(l2OutputOracleProxy));
            existingProxy.upgradeTo(OPSuccinctL2OutputOracleImpl);
        } else {
            // Raw calldata for an upgrade call by a multisig.
            bytes memory multisigCalldata = abi.encodeWithSelector(
                Proxy.upgradeTo.selector,
                OPSuccinctL2OutputOracleImpl
            );
            console.log("Raw calldata for the upgrade call:");
            console.logBytes(multisigCalldata);
        }

        vm.stopBroadcast();
    }
}
