rm -rf ./build
rm -rf ./go
mkdir -p ./go/boker_contract
solc --abi --bin -o ./build BokerManager.sol BokerFinance.sol BokerLogData.sol BokerLog.sol BokerDapp.sol BokerDappData.sol BokerFile.sol BokerFileData.sol BokerTokenPower.sol BokerTokenPowerData.sol BokerUser.sol BokerUserData.sol BokerNode.sol BokerNodeData.sol BokerInterface.sol BokerInterfaceBase.sol

ABICMD=/root/go/src/github.com/Bokerchain/Boker/chain/build/bin/abigen
$ABICMD -abi ./build/BokerManager.abi -bin ./build/BokerManager.bin -pkg boker_contract -out ./go/boker_contract/boker_manager.go -type BokerManager
$ABICMD -abi ./build/BokerInterface.abi -bin ./build/BokerInterface.bin --pkg boker_contract -out ./go/boker_contract/boker_interface.go -type BokerInterface
$ABICMD -abi ./build/BokerInterfaceBase.abi -bin ./build/BokerInterfaceBase.bin --pkg boker_contract -out ./go/boker_contract/boker_interface_base.go -type BokerInterfaceBase
$ABICMD -abi ./build/BokerDapp.abi -bin ./build/BokerDapp.bin --pkg boker_contract -out ./go/boker_contract/boker_dapp.go -type BokerDapp
$ABICMD -abi ./build/BokerDappData.abi -bin ./build/BokerDappData.bin --pkg boker_contract -out ./go/boker_contract/boker_dapp_data.go -type BokerDappData
$ABICMD -abi ./build/BokerFile.abi -bin ./build/BokerFile.bin --pkg boker_contract -out ./go/boker_contract/boker_file.go -type BokerFile
$ABICMD -abi ./build/BokerFileData.abi -bin ./build/BokerFileData.bin --pkg boker_contract -out ./go/boker_contract/boker_file_data.go -type BokerFileData
$ABICMD -abi ./build/BokerFinance.abi -bin ./build/BokerFinance.bin --pkg boker_contract -out ./go/boker_contract/boker_finance.go -type BokerFinance
$ABICMD -abi ./build/BokerLog.abi -bin ./build/BokerLog.bin --pkg boker_contract -out ./go/boker_contract/boker_log.go -type BokerLog
$ABICMD -abi ./build/BokerLogData.abi -bin ./build/BokerLogData.bin --pkg boker_contract -out ./go/boker_contract/boker_log_data.go -type BokerLogData
$ABICMD -abi ./build/BokerNode.abi -bin ./build/BokerNode.bin --pkg boker_contract -out ./go/boker_contract/boker_node.go -type BokerNode
$ABICMD -abi ./build/BokerNodeData.abi -bin ./build/BokerNodeData.bin --pkg boker_contract -out ./go/boker_contract/boker_node_data.go -type BokerNodeData
$ABICMD -abi ./build/BokerTokenPower.abi -bin ./build/BokerTokenPower.bin --pkg boker_contract -out ./go/boker_contract/boker_token_power.go -type BokerTokenPower
$ABICMD -abi ./build/BokerTokenPowerData.abi -bin ./build/BokerTokenPowerData.bin --pkg boker_contract -out ./go/boker_contract/boker_token_power_data.go -type BokerTokenPowerData
$ABICMD -abi ./build/BokerUser.abi -bin ./build/BokerUser.bin --pkg boker_contract -out ./go/boker_contract/boker_user.go -type BokerUser
$ABICMD -abi ./build/BokerUserData.abi -bin ./build/BokerUserData.bin --pkg boker_contract -out ./go/boker_contract/boker_user_data.go -type BokerUserData
rm -rf ./build