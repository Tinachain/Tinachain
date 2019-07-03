rm -rf build
rm -rf go
mkdir -p go
mkdir -p build
./solc --abi --bin -o build BokerInterface.sol BokerInterfaceBase.sol BokerLogData.sol BokerLog.sol BokerManager.sol BokerNode.sol BokerNodeData.sol Chainware.sol

ABICMD=/root/go/src/github.com/Tinachain/Tina/chain/build/bin/abigen
$ABICMD -abi build/BokerManager.abi -bin build/BokerManager.bin -pkg boker_contract -out ../../boker_interface/chainware_manager.go -type BokerManager
$ABICMD -abi build/BokerInterface.abi -bin build/BokerInterface.bin --pkg boker_contract -out ../../boker_interface/chainware_interface.go -type BokerInterface
$ABICMD -abi build/BokerInterfaceBase.abi -bin build/BokerInterfaceBase.bin --pkg boker_contract -out ../../boker_interface/chainware_interface_base.go -type BokerInterfaceBase
$ABICMD -abi build/BokerLog.abi -bin build/BokerLog.bin --pkg boker_contract -out ../../boker_interface/chainware_log.go -type BokerLog
$ABICMD -abi build/BokerLogData.abi -bin build/BokerLogData.bin --pkg boker_contract -out ../../boker_interface/chainware_log_data.go -type BokerLogData
$ABICMD -abi build/BokerNode.abi -bin build/BokerNode.bin --pkg boker_contract -out ../../boker_interface/chainware_node.go -type BokerNode
$ABICMD -abi build/BokerNodeData.abi -bin build/BokerNodeData.bin --pkg boker_contract -out ../../boker_interface/chainware_node_data.go -type BokerNodeData
$ABICMD -abi build/Chainware.abi -bin build/Chainware.bin --pkg boker_contract -out ../../boker_interface/chainware.go -type Chainware

rm -rf build
