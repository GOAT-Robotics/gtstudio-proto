#!/bin/bash

RCC_PROTO="./protos/rcc/v1/rcc.proto"
NODE_SDK="./protos/node_sdk/v1/node_sdk.proto"
AGENT_PROTO="./protos/agent/v1/agent.proto"
FLEET_MANAGER_PROTO="./protos/fleet_manager/v1/fleet_manager.proto"

# Remove existing Go files
rm -rf ./protos/v1

# Create new directory
mkdir ./protos/v1

# Generates Go files for GT Studio from the rcc.proto file
protoc --go_out=. --go-grpc_out=. $RCC_PROTO
echo "Processing $RCC_PROTO file.."

# Generates Go files for GT Studio - Node SDK from the node_sdk.proto file
protoc --go_out=. --go-grpc_out=. $NODE_SDK
echo "Processing $NODE_SDK file.."

# Generates Go files for GT Studio - Agent from the agent.proto file
protoc --go_out=. --go-grpc_out=. $AGENT_PROTO
echo "Processing $AGENT_PROTO file.."

# Generates Go files for GT Studio - Fleet Manager from the fleet_manager.proto file
protoc --go_out=. --go-grpc_out=. $FLEET_MANAGER_PROTO
echo "Processing $FLEET_MANAGER_PROTO file.."

# Compiles Python files for the agent.proto file
##### NOTE: Change the local GT-Agent repo directory to save the generated Python pb files ####
python3 -m grpc_tools.protoc -I. --python_out=/home/amizhthan/RCC/robotix-agent/node --grpc_python_out=/home/amizhthan/RCC/robotix-agent/node ./protos/agent/v1/agent.proto

# Loops over all the model files and creates both Go and Python files
for f in ./protos/model/v1/*.proto
do
    protoc --go_out=. --go-grpc_out=. $f
    python3 -m grpc_tools.protoc -I. --python_out=/home/amizhthan/RCC/robotix-agent/node --grpc_python_out=/home/amizhthan/RCC/robotix-agent/node $f

    echo "Processing $f file.."
done

# Move generated Go files to ./protos/v1/
mv ./github.com/GOAT-Robotics/robotix-proto/protos/v1/* ./protos/v1/

# remove the github.com directory
rm -rf ./github.com

echo "### Protobuf files generated successfully ###"
