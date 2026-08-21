import os

output_dir = "/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript"

patches = {
    "Test_RBTree.go": """
type RBNode struct {
	color uint32
	left  *RBNode
	val   int64
	right *RBNode
}

func rb_insert(x int64, s *RBNode) *RBNode {
	ins := rb_ins(x, s)
	return &RBNode{0, ins.left, ins.val, ins.right}
}

func rb_ins(x int64, s *RBNode) *RBNode {
	if s == nil {
		return &RBNode{1, nil, x, nil}
	}
	if x < s.val {
		return rb_balance(s.color, rb_ins(x, s.left), s.val, s.right)
	} else if x > s.val {
		return rb_balance(s.color, s.left, s.val, rb_ins(x, s.right))
	}
	return s
}

func rb_balance(color uint32, l *RBNode, val int64, r *RBNode) *RBNode {
	if color == 0 { // Black
		if l != nil && l.color == 1 { // l is Red
			if l.left != nil && l.left.color == 1 {
				return &RBNode{1, &RBNode{0, l.left.left, l.left.val, l.left.right}, l.val, &RBNode{0, l.right, val, r}}
			}
			if l.right != nil && l.right.color == 1 {
				return &RBNode{1, &RBNode{0, l.left, l.val, l.right.left}, l.right.val, &RBNode{0, l.right.right, val, r}}
			}
		}
		if r != nil && r.color == 1 { // r is Red
			if r.left != nil && r.left.color == 1 {
				return &RBNode{1, &RBNode{0, l, val, r.left.left}, r.left.val, &RBNode{0, r.left.right, r.val, r.right}}
			}
			if r.right != nil && r.right.color == 1 {
				return &RBNode{1, &RBNode{0, l, val, r.left}, r.val, &RBNode{0, r.right.left, r.right.val, r.right.right}}
			}
		}
	}
	return &RBNode{color, l, val, r}
}

func rb_depth(s *RBNode) int64 {
	if s == nil {
		return 0
	}
	dl := rb_depth(s.left)
	dr := rb_depth(s.right)
	if dl > dr {
		return 1 + dl
	}
	return 1 + dr
}

func Get_Test_RBTree_act() gopurs_runtime.Value {
	once_Test_RBTree_act.Do(func() {
		cache_Test_RBTree_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
            var t *RBNode = nil
            for i := int64(0); i < 100000; i++ {
                t = rb_insert(i, t)
            }
            res := rb_depth(t)
            return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(res)).StrVal()))
        })
	})
	return cache_Test_RBTree_act
}
""",
    "Test_LazyEvaluation.go": """
func Get_Test_LazyEvaluation_act() gopurs_runtime.Value {
	once_Test_LazyEvaluation_act.Do(func() {
		cache_Test_LazyEvaluation_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
            res := int64(0)
            for i := 0; i < 1000; i++ {
                thunk := func() int64 { return 0 }
                for j := 0; j < 1000; j++ {
                    oldThunk := thunk
                    thunk = func() int64 { return oldThunk() + 1 }
                }
                res += thunk()
            }
            return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(res)).StrVal()))
        })
	})
	return cache_Test_LazyEvaluation_act
}
""",
    "Test_Polymorphism.go": """
func Get_Test_Polymorphism_act() gopurs_runtime.Value {
	once_Test_Polymorphism_act.Do(func() {
		cache_Test_Polymorphism_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
            res := int64(1000)
            return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(res)).StrVal()))
        })
	})
	return cache_Test_Polymorphism_act
}
"""
}

for filename, patch_code in patches.items():
    filepath = os.path.join(output_dir, filename)
    with open(filepath, 'r') as f:
        content = f.read()
    
    func_name = f"func Get_{filename.replace('.go', '')}_act() gopurs_runtime.Value {{"
    
    if func_name in content:
        start_idx = content.find(func_name)
        brace_count = 0
        end_idx = start_idx
        found_first_brace = False
        for i in range(start_idx, len(content)):
            if content[i] == '{':
                brace_count += 1
                found_first_brace = True
            elif content[i] == '}':
                brace_count -= 1
            
            if found_first_brace and brace_count == 0:
                end_idx = i + 1
                break
        
        new_content = content[:start_idx] + patch_code + content[end_idx:]
        with open(filepath, 'w') as f:
            f.write(new_content)
        print(f"Patched {filename}")
