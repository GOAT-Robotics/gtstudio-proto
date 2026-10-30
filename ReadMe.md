<h1 align="center">🤖 Robotix Proto</h1>

<p align="center">
  Protocol Buffer + gRPC definitions for the <b>Robotix</b> platform by
  <a href="https://github.com/GOAT-Robotics">GOAT Robotics</a>
</p>

---

## 🧩 What is this?

The shared contract between Robotix services. Define messages and gRPC services once in `.proto` files, generate Go code, and every service stays in sync.

| Proto | Purpose |
|-------|---------|
| 🛰️ `rcc/v1/rcc.proto` | Robotix Command Center (RCC) |
| 🚚 `fleet_manager/v1/fleet_manager.proto` | Fleet Manager |
| 📦 `model/v1/*.proto` | Shared data models |

> The Agent and Node SDK generators are marked deprecated in `generate.sh` and are no longer generated.

## 📁 Structure

```
protos/
├── rcc/v1/              → RCC service definitions
├── fleet_manager/v1/    → Fleet Manager service definitions
├── model/v1/            → shared models
└── v1/                  → ⚙️ GENERATED Go code (wiped on every run, do not edit)
generate.sh              → compile protos into Go code
update.sh                → generate + commit + tag + push
go.mod                   → Go module
```

## ⚙️ Setup (Ubuntu / Debian)

**1. Install Go and the protobuf compiler**

```bash
sudo apt update
sudo apt install -y golang-go protobuf-compiler
go version      # needs Go 1.26+ (go.mod: 1.26.6)
protoc --version
```

**2. Install the Go plugins**

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.2
```

**3. Add Go binaries to your PATH**

```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
protoc-gen-go --version
protoc-gen-go-grpc --version
```

## 🚀 Usage

**Generate code only** (no git changes):

```bash
chmod +x generate.sh update.sh
./generate.sh
```

**Generate and release:**

```bash
./update.sh
```

`update.sh` will:

1. 🔨 Run `generate.sh`
2. 🏷️ Pick the next version from the remote tags (`vMAJOR.MINOR.PATCH`, patch + 1)
3. ➕ Run `git add .` (stages **everything**)
4. 💬 Ask for a commit message
5. 📤 Commit, tag, and push the commit and the new tag to `origin`

> ⚠️ It stages and pushes all your changes. Run `git status` first.

## 📦 Using it in another Go project

```bash
go get github.com/GOAT-Robotics/gtstudio-proto@<version>
```

```go
import pb "github.com/GOAT-Robotics/gtstudio-proto/protos/v1"
```

---

<p align="center"><sub>Built with ❤️ by <a href="https://github.com/GOAT-Robotics">GOAT Robotics</a></sub></p>
