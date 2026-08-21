with open("scratch/box_rbtree.rs", "r") as f:
    content = f.read()

content = content.replace("*llx", "llx").replace("*lx", "lx").replace("*lrx", "lrx")
content = content.replace("*rx", "rx").replace("*rlx", "rlx").replace("*rrx", "rrx")
content = content.replace("ref llx", "llx").replace("ref lx", "lx").replace("ref lrx", "lrx")
content = content.replace("ref rx", "rx").replace("ref rlx", "rlx").replace("ref rrx", "rrx")

with open("scratch/box_rbtree.rs", "w") as f:
    f.write(content)
print("Patched!")
