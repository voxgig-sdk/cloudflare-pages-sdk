<?php
declare(strict_types=1);

// CloudflarePages SDK base feature

class CloudflarePagesBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(CloudflarePagesContext $ctx, array $options): void {}
    public function PostConstruct(CloudflarePagesContext $ctx): void {}
    public function PostConstructEntity(CloudflarePagesContext $ctx): void {}
    public function SetData(CloudflarePagesContext $ctx): void {}
    public function GetData(CloudflarePagesContext $ctx): void {}
    public function GetMatch(CloudflarePagesContext $ctx): void {}
    public function SetMatch(CloudflarePagesContext $ctx): void {}
    public function PrePoint(CloudflarePagesContext $ctx): void {}
    public function PreSpec(CloudflarePagesContext $ctx): void {}
    public function PreRequest(CloudflarePagesContext $ctx): void {}
    public function PreResponse(CloudflarePagesContext $ctx): void {}
    public function PreResult(CloudflarePagesContext $ctx): void {}
    public function PreDone(CloudflarePagesContext $ctx): void {}
    public function PreUnexpected(CloudflarePagesContext $ctx): void {}
}
