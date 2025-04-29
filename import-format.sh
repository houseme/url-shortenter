#!/bin/bash
# Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
#
# This Source Code Form is subject to the terms of the MIT License.
# If a copy of the MIT was not distributed with this file,
# You can obtain one at https://github.com/houseme/url-shortenter.
#
# Unless required by applicable law or agreed to in writing, software distributed
# under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
# CONDITIONS OF ANY KIND, either express or implied. See the License for the
# specific language governing permissions and limitations under the License.


# 格式化 go.mod
go mod tidy -compat=1.18


# 处理 go imports 的格式化
rm -rf style_tool
rm -rf goimports-reviser

mkdir -p style_tool

cd style_tool

wget https://github.com/incu6us/goimports-reviser/releases/download/v3.9.1/goimports-reviser_3.8.2_linux_amd64.tar.gz
tar -zxvf goimports-reviser_3.9.1_linux_amd64.tar.gz
mv goimports-reviser ../

cd ../

ls -lstrh

find . -name "*.go" -type f | grep -v examples | grep -v .pb.go|grep -v test/tools/tools.go | grep -v ./plugin_register_generate.go | xargs -I {} goimports-reviser -rm-unused -format {} -project-name github.com/houseme/url-shortenter

# 处理 go 代码格式化
go fmt ./...