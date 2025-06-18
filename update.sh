#!/bin/bash

# This script will do the following.
# - Generate proto files
# - Add to git, commit, update version, create tag, push both tag and commit.

version=88
old_version=$version

echo -e '\033[4;32;1mGenerating Proto Code\033[m'
./generate.sh

echo -e '\033[4;32;1mAdding To Git\033[m'
git add .
echo -e '\033[4;32;1mGive a message: \033[m'
read message
git commit -m "$message"
version=$((version+1))
echo "Creating a tag - v1.0.$version"
git tag "v1.0.$version"
echo -e '\033[4;32;1mPushing Code To Github\033[m'
git push origin develop
git push --tags
echo -e '\033[4;32;1mUpdating Go Repository\033[m'

sed -i "s/version=$old_version/version=$version/g" /home/cipher/GoProjects/robotix-proto/update.sh
echo -e '\033[4;32;1mDone!\033[m'

//goat@gtlap001:~/Documents/gtstudio-proto$ ./update.sh
Generating Proto Code
./generate.sh: line 8: protoc: command not found
Processing ./protos/rcc/v1/rcc.proto file..
./generate.sh: line 12: protoc: command not found
Processing ./protos/agent/v1/agent.proto file..
./generate.sh: line 16: protoc: command not found
Processing ./protos/fleet_manager/v1/fleet_manager.proto file..
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/commands.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/config.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/datapoint.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/file.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/health.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/math.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/media.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/navigation.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/ros.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/specifications.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/text.proto file..
./generate.sh: line 26: protoc: command not found
/usr/bin/python3: Error while finding module specification for 'grpc_tools.protoc' (ModuleNotFoundError: No module named 'grpc_tools')
Processing ./protos/model/v1/webrtc.proto file..
### Protobuf files generated successfully ###
Adding To Git
Give a message: 
Update streamDataResponse proto
[develop 8cd6d29] Update streamDataResponse proto
 2 files changed, 7 insertions(+), 2 deletions(-)
Creating a tag - v1.0.89
fatal: tag 'v1.0.89' already exists
Pushing Code To Github
warning: url contains a newline in its username component: https://sineka-goatrobotics%0A@github.com/GOAT-Robotics/gtstudio-proto.git/
fatal: credential url cannot be parsed: https://sineka-goatrobotics%0A@github.com/GOAT-Robotics/gtstudio-proto.git/
warning: url contains a newline in its username component: https://sineka-goatrobotics%0A@github.com/GOAT-Robotics/gtstudio-proto.git/
fatal: credential url cannot be parsed: https://sineka-goatrobotics%0A@github.com/GOAT-Robotics/gtstudio-proto.git/
Updating Go Repository
sed: can't read /home/cipher/GoProjects/robotix-proto/update.sh: No such file or directory
Done!