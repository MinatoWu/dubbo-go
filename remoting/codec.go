/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package remoting

import (
	"bytes"
)

// Codec is the interface that wrap EncodeRequest、 EncodeResponse and Decode method
// for exchangeClient.
type Codec interface {
	EncodeRequest(request *Request) (*bytes.Buffer, error)
	EncodeResponse(response *Response) (*bytes.Buffer, error)
	Decode(data []byte) (*DecodeResult, int, error)
}

type DecodeResult struct {
	IsRequest bool // indicates whether the current request is a heartbeat request
	Result    any
}

var codec = make(map[string]Codec, 2)

func RegistryCodec(protocol string, codecTmp Codec) {
	codec[protocol] = codecTmp
}

func GetCodec(protocol string) Codec {
	return codec[protocol]
}
