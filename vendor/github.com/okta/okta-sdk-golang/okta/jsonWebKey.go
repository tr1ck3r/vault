/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"time"
)

type JsonWebKey struct {
	Links       interface{} `json:"_links,omitempty"`
	Alg         string      `json:"alg,omitempty"`
	Created     *time.Time  `json:"created,omitempty"`
	E           string      `json:"e,omitempty"`
	ExpiresAt   *time.Time  `json:"expiresAt,omitempty"`
	KeyOps      []string    `json:"key_ops,omitempty"`
	Kid         string      `json:"kid,omitempty"`
	Kty         string      `json:"kty,omitempty"`
	LastUpdated *time.Time  `json:"lastUpdated,omitempty"`
	N           string      `json:"n,omitempty"`
	Status      string      `json:"status,omitempty"`
	Use         string      `json:"use,omitempty"`
	X5c         []string    `json:"x5c,omitempty"`
	X5t         string      `json:"x5t,omitempty"`
	X5tS256     string      `json:"x5t#S256,omitempty"`
	X5u         string      `json:"x5u,omitempty"`
}

func NewJsonWebKey() *JsonWebKey {
	return &JsonWebKey{}
}

func (a *JsonWebKey) IsApplicationInstance() bool {
	return true
}