<?php
$exports['toCharCode'] = function($c) { return \ord($c); };
$exports['fromCharCode'] = function($c) { return \chr($c); };
