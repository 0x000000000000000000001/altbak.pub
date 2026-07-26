<?php
if (file_exists(__DIR__ . '/../../vendor/autoload.php')) require_once __DIR__ . '/../../vendor/autoload.php';
if (!\class_exists('\\PhpursThunks', false)) {
  class PhpursThunks {
    public static $thunks = [];
    public static $cache = [];
    public static function eval($id) {
      if (isset(self::$cache[$id]) || \array_key_exists($id, self::$cache)) return self::$cache[$id];
      if (isset(self::$thunks[$id])) {
        self::$cache[$id] = self::$thunks[$id]();
        return self::$cache[$id];
      }
      throw new \Exception("FATAL: Unknown thunk " . $id);
    }
  }
}
set_exception_handler(function($e) { echo 'FATAL: ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n"; exit(1); });
require_once __DIR__ . '/index.php';
($GLOBALS['App_main'] ?? \PhpursThunks::eval('App_main'))();
if (class_exists('\\Revolt\\EventLoop')) { \Revolt\EventLoop::run(); }
