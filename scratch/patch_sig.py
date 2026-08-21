import re

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "r") as f:
    content = f.read()

# Replace the closing parenthesis of balance with `, Some(v1.clone()))`
# Wait, the 4th argument is `v1.unwrap_record().vals.as_ref().unwrap()[3].clone()`
# So we can just replace `unwrap()[3].clone())` with `unwrap()[3].clone(), Some(v1.clone()))`
# Only in the lines where `Test_RBTree_ins` is calling `Test_RBTree_balance`

content = content.replace(
    ".unwrap()[3].clone())",
    ".unwrap()[3].clone(), Some(v1.clone()))"
)
# Wait, but v1 is consumed by unwrap_record().
# Let's just fix it by replacing `.unwrap()[3].clone(), Some(v1.clone()))` with `.unwrap()[3].clone(), Some(v1))`
# BUT wait! `v1.unwrap_record()` consumes `v1`!
# So `Some(v1)` is invalid if it was consumed.
# Let's change `v1.unwrap_record()` to `v1.clone().unwrap_record()` everywhere in the last argument!
content = content.replace(
    "v1.unwrap_record().vals.as_ref().unwrap()[3].clone()",
    "v1.clone().unwrap_record().vals.as_ref().unwrap()[3].clone()"
)
# Now the last argument doesn't consume `v1`. So we can pass `Some(v1)`!
content = content.replace(
    "v1.clone().unwrap_record().vals.as_ref().unwrap()[3].clone())",
    "v1.clone().unwrap_record().vals.as_ref().unwrap()[3].clone(), Some(v1))"
)

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "w") as f:
    f.write(content)
print("Patched calls!")
