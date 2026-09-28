<?php
declare(strict_types=1);

// CloudflarePages SDK utility: result_body

class CloudflarePagesResultBody
{
    public static function call(CloudflarePagesContext $ctx): ?CloudflarePagesResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
