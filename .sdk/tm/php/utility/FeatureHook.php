<?php
declare(strict_types=1);

// CloudflarePages SDK utility: feature_hook

class CloudflarePagesFeatureHook
{
    public static function call(CloudflarePagesContext $ctx, string $name): void
    {
        if (!$ctx->client) {
            return;
        }
        $features = $ctx->client->features ?? null;
        if (!$features) {
            return;
        }
        foreach ($features as $f) {
            if (method_exists($f, $name)) {
                $f->$name($ctx);
            }
        }
    }
}
