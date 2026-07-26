<?php
if (file_exists(__DIR__ . '/../../vendor/autoload.php')) require_once __DIR__ . '/../../vendor/autoload.php';
set_exception_handler(function($e) { echo 'FATAL: ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n"; exit(1); });
require_once __DIR__ . '/../require_all.php';
$GLOBALS['App_main']();
if (class_exists('\\Revolt\\EventLoop')) { \Revolt\EventLoop::run(); }
