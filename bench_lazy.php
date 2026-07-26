<?php
$start = microtime(true);
$sum = 0;
for ($i = 0; $i < 1000000; $i++) {
    $f = function() use ($i) { return $i; };
    $v = null;
    $done = false;
    $lazy = function() use ($f, &$v, &$done) {
        if ($done) return $v;
        $v = $f();
        $done = true;
        return $v;
    };
    $sum += $lazy();
}
echo "Closures: " . ((microtime(true) - $start) * 1000) . " ms\n";

class Phpurs_Lazy {
    private $f, $v, $done = false;
    public function __construct($f) { $this->f = $f; }
    public function __invoke() {
        if ($this->done) return $this->v;
        $this->v = ($this->f)();
        $this->done = true;
        return $this->v;
    }
}

$start = microtime(true);
$sum = 0;
for ($i = 0; $i < 1000000; $i++) {
    $f = function() use ($i) { return $i; };
    $lazy = new Phpurs_Lazy($f);
    $sum += $lazy();
}
echo "Objects: " . ((microtime(true) - $start) * 1000) . " ms\n";
