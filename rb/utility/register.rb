# CloudflarePages SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

CloudflarePagesUtility.registrar = ->(u) {
  u.clean = CloudflarePagesUtilities::Clean
  u.done = CloudflarePagesUtilities::Done
  u.make_error = CloudflarePagesUtilities::MakeError
  u.feature_add = CloudflarePagesUtilities::FeatureAdd
  u.feature_hook = CloudflarePagesUtilities::FeatureHook
  u.feature_init = CloudflarePagesUtilities::FeatureInit
  u.fetcher = CloudflarePagesUtilities::Fetcher
  u.make_fetch_def = CloudflarePagesUtilities::MakeFetchDef
  u.make_context = CloudflarePagesUtilities::MakeContext
  u.make_options = CloudflarePagesUtilities::MakeOptions
  u.make_request = CloudflarePagesUtilities::MakeRequest
  u.make_response = CloudflarePagesUtilities::MakeResponse
  u.make_result = CloudflarePagesUtilities::MakeResult
  u.make_point = CloudflarePagesUtilities::MakePoint
  u.make_spec = CloudflarePagesUtilities::MakeSpec
  u.make_url = CloudflarePagesUtilities::MakeUrl
  u.param = CloudflarePagesUtilities::Param
  u.prepare_auth = CloudflarePagesUtilities::PrepareAuth
  u.prepare_body = CloudflarePagesUtilities::PrepareBody
  u.prepare_headers = CloudflarePagesUtilities::PrepareHeaders
  u.prepare_method = CloudflarePagesUtilities::PrepareMethod
  u.prepare_params = CloudflarePagesUtilities::PrepareParams
  u.prepare_path = CloudflarePagesUtilities::PreparePath
  u.prepare_query = CloudflarePagesUtilities::PrepareQuery
  u.graphql_body = CloudflarePagesUtilities::GraphqlBody
  u.graphql_errors = CloudflarePagesUtilities::GraphqlErrors
  u.result_basic = CloudflarePagesUtilities::ResultBasic
  u.result_body = CloudflarePagesUtilities::ResultBody
  u.result_headers = CloudflarePagesUtilities::ResultHeaders
  u.transform_request = CloudflarePagesUtilities::TransformRequest
  u.transform_response = CloudflarePagesUtilities::TransformResponse
}
