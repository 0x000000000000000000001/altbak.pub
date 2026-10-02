<?php
// Use the pinned Composer's own lock hash and platform detection, without
// starting its installer, solver, plugins or scripts.
require 'phar://' . $argv[1] . '/vendor/autoload.php';

$platform = [];
foreach ((new Composer\Repository\PlatformRepository())->getPackages() as $package) {
    $platform[$package->getName()] = $package->getPrettyVersion();
}
ksort($platform);
$config = Composer\Factory::createConfig(new Composer\IO\NullIO(), dirname($argv[2]));
echo json_encode([
    'content-hash' => Composer\Package\Locker::getContentHash(file_get_contents($argv[2])),
    'platform' => $platform,
    'config-sha256' => hash('sha256', json_encode($config->all(), JSON_THROW_ON_ERROR)),
], JSON_THROW_ON_ERROR);
