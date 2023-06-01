// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
package ent

// REUSE-IgnoreStart

//go:generate find . -not -name generate.go -and -not -name main.go -and -not -path "**/schema/*" -type f -delete
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --header "// SPDX-FileCopyrightText: The entgo authors\n// SPDX-License-Identifier: Apache-2.0\n\n// Code generated by ent, DO NOT EDIT." --feature sql/execquery,sql/versioned-migration ./schema

// REUSE-IgnoreEnd
