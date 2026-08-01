<?php
if (file_exists(__DIR__ . '/../../vendor/autoload.php')) require_once __DIR__ . '/../../vendor/autoload.php';
set_exception_handler(function($e) { echo 'FATAL: ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n"; exit(1); });
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Test.Fib/index.php';
require_once __DIR__ . '/../AppX/index.php';
$GLOBALS['AppX_main']();
if (class_exists('\\Revolt\\EventLoop')) { \Revolt\EventLoop::run(); }
