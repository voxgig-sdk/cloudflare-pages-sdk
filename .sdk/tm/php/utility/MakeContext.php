<?php
declare(strict_types=1);

// CloudflarePages SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class CloudflarePagesMakeContext
{
    public static function call(array $ctxmap, ?CloudflarePagesContext $basectx): CloudflarePagesContext
    {
        return new CloudflarePagesContext($ctxmap, $basectx);
    }
}
