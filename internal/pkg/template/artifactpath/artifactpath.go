// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package artifactpath holds functions to generate the S3 object path for artifacts.
package artifactpath

import (
	"crypto/sha256"
	"fmt"
	"path"
)

const (
	s3ArtifactDirName         = "manual"
	s3TemplateDirName         = "templates"
	s3ArtifactAddonsDirName   = "addons"
	s3ArtifactEnvFilesDirName = "env-files"
)

// MkdirSHA prefixes the key with the SHA256 hash of the contents of "manual/<hash>/key".
func MkdirSHA256(key string, content []byte) string {
	return path.Join(s3ArtifactDirName, fmt.Sprintf("%x", sha256.Sum256(content)), key)
}

// Addons returns the path to store addon artifact files with sha256 of the content.
// Example: manual/addons/key/sha.yml.
func Addons(key string, content []byte) string {
	return path.Join(s3ArtifactDirName, s3ArtifactAddonsDirName, key, fmt.Sprintf("%x.yml", sha256.Sum256(content)))
}

// CFNTemplate returns the path to store cloudformation templates with sha256 of the content.
// Example: manual/templates/key/sha.yml.
func CFNTemplate(key string, content []byte) string {
	return path.Join(s3ArtifactDirName, s3TemplateDirName, key, fmt.Sprintf("%x.yml", sha256.Sum256(content)))
}

// EnvFiles returns the path to store an env file artifact with sha256 of the content..
// Example: manual/env-files/key/sha.env.
func EnvFiles(key string, content []byte) string {
	return path.Join(s3ArtifactDirName, s3ArtifactEnvFilesDirName, key, fmt.Sprintf("%x.env", sha256.Sum256(content)))
}
