<?php
class AstTreeNative {
    public static function build($n) {
        if ($n === 0) return (object)['tag' => 0, 'value' => 1];
        return (object)['tag' => 1,
            'left' => (object)['tag' => 2, 'left' => (object)['tag' => 0, 'value' => $n], 'right' => self::build($n - 1)],
            'right' => (object)['tag' => 3, 'left' => self::build($n - 1), 'right' => (object)['tag' => 0, 'value' => 1]]];
    }
    public static function evaluate($tree) {
        switch ($tree->tag) {
            case 0: return $tree->value;
            case 1: return self::evaluate($tree->left) + self::evaluate($tree->right);
            case 2: return self::evaluate($tree->left) * self::evaluate($tree->right);
            case 3: return self::evaluate($tree->left) - self::evaluate($tree->right);
        }
        throw new \LogicException('Unknown AST tag');
    }
}
$exports['runAstTreeFFICheatcode'] = function($limit) {
    return AstTreeNative::evaluate(AstTreeNative::build((int)$limit));
};
return $exports;
