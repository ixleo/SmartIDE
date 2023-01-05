/*
SmartIDE - Dev Containers
Copyright (C) 2023 leansoftX.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package compose

// 暴露的映射的公共接口
type YmlSecret struct {
	File     string `yaml:"file,omitempty"`     // 文件路径
	External string `yaml:"external,omitempty"` // 是否已存在，存在不需要再创建。
	Name     string `yaml:"name,omitempty"`     // (v3.5+) 名称
}
