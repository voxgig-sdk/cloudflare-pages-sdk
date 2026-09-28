<?php
declare(strict_types=1);

// CloudflarePages SDK utility: result_headers

class CloudflarePagesResultHeaders
{
    public static function call(CloudflarePagesContext $ctx): ?CloudflarePagesResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
