<?php
class RBNodeCheat {
    public $color; // 0 red, 1 black
    public $left;
    public $value;
    public $right;
    public function __construct($c, $l, $v, $r) {
        $this->color = $c; $this->left = $l; $this->value = $v; $this->right = $r;
    }
}

class RBTreeNative {
    public static function balance($c, $a, $x, $b) {
        if ($c === 1) { // black
            if ($a !== null && $a->color === 0) {
                if ($a->left !== null && $a->left->color === 0) {
                    return new RBNodeCheat(0, new RBNodeCheat(1, $a->left->left, $a->left->value, $a->left->right), $a->value, new RBNodeCheat(1, $a->right, $x, $b));
                }
                if ($a->right !== null && $a->right->color === 0) {
                    return new RBNodeCheat(0, new RBNodeCheat(1, $a->left, $a->value, $a->right->left), $a->right->value, new RBNodeCheat(1, $a->right->right, $x, $b));
                }
            }
            if ($b !== null && $b->color === 0) {
                if ($b->left !== null && $b->left->color === 0) {
                    return new RBNodeCheat(0, new RBNodeCheat(1, $a, $x, $b->left->left), $b->left->value, new RBNodeCheat(1, $b->left->right, $b->value, $b->right));
                }
                if ($b->right !== null && $b->right->color === 0) {
                    return new RBNodeCheat(0, new RBNodeCheat(1, $a, $x, $b->left), $b->value, new RBNodeCheat(1, $b->right->left, $b->right->value, $b->right->right));
                }
            }
        }
        return new RBNodeCheat($c, $a, $x, $b);
    }

    public static function ins($x, $t) {
        if ($t === null) return new RBNodeCheat(0, null, $x, null);
        if ($x < $t->value) return self::balance($t->color, self::ins($x, $t->left), $t->value, $t->right);
        if ($x > $t->value) return self::balance($t->color, $t->left, $t->value, self::ins($x, $t->right));
        return $t;
    }

    public static function insert($x, $t) {
        $res = self::ins($x, $t);
        return new RBNodeCheat(1, $res->left, $res->value, $res->right);
    }

    public static function depth($t) {
        if ($t === null) return 0;
        $ld = self::depth($t->left);
        $rd = self::depth($t->right);
        return 1 + ($ld > $rd ? $ld : $rd);
    }

}

$exports['runRBTreeFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = null;
    for ($i = $n; $i > 0; $i--) {
        $acc = RBTreeNative::insert($i, $acc);
    }
    
    return RBTreeNative::depth($acc);
};
return $exports;
