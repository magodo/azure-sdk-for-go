// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package runtime

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v5.0/qnamaker/runtime"

type ErrorCodeType = original.ErrorCodeType

const (
	BadArgument       ErrorCodeType = original.BadArgument
	EndpointKeysError ErrorCodeType = original.EndpointKeysError
	ExtractionFailure ErrorCodeType = original.ExtractionFailure
	Forbidden         ErrorCodeType = original.Forbidden
	KbNotFound        ErrorCodeType = original.KbNotFound
	NotFound          ErrorCodeType = original.NotFound
	OperationNotFound ErrorCodeType = original.OperationNotFound
	QnaRuntimeError   ErrorCodeType = original.QnaRuntimeError
	QuotaExceeded     ErrorCodeType = original.QuotaExceeded
	ServiceError      ErrorCodeType = original.ServiceError
	SKULimitExceeded  ErrorCodeType = original.SKULimitExceeded
	Unauthorized      ErrorCodeType = original.Unauthorized
	Unspecified       ErrorCodeType = original.Unspecified
	ValidationFailure ErrorCodeType = original.ValidationFailure
)

type BaseClient = original.BaseClient
type Client = original.Client
type ContextDTO = original.ContextDTO
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type FeedbackRecordDTO = original.FeedbackRecordDTO
type FeedbackRecordsDTO = original.FeedbackRecordsDTO
type InnerErrorModel = original.InnerErrorModel
type MetadataDTO = original.MetadataDTO
type PromptDTO = original.PromptDTO
type PromptDTOQna = original.PromptDTOQna
type QnADTO = original.QnADTO
type QnADTOContext = original.QnADTOContext
type QnASearchResult = original.QnASearchResult
type QnASearchResultContext = original.QnASearchResultContext
type QnASearchResultList = original.QnASearchResultList
type QueryContextDTO = original.QueryContextDTO
type QueryDTO = original.QueryDTO
type QueryDTOContext = original.QueryDTOContext

func New(runtimeEndpoint string) BaseClient {
	return original.New(runtimeEndpoint)
}
func NewClient(runtimeEndpoint string) Client {
	return original.NewClient(runtimeEndpoint)
}
func NewWithoutDefaults(runtimeEndpoint string) BaseClient {
	return original.NewWithoutDefaults(runtimeEndpoint)
}
func PossibleErrorCodeTypeValues() []ErrorCodeType {
	return original.PossibleErrorCodeTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
