import re

with open('/Users/0x1/Documents/htdocs/altbak.pub/bin/go/run', 'r') as f:
    content = f.read()

# find the block and remove it
pattern = r'# Inject pprof into output/main\.go if requested\n\s*if \[ "\$PPROF" == "1" \]; then.*?fi\n\s*go build -o go_app \.'
content = re.sub(pattern, 'go build -o go_app .', content, flags=re.DOTALL)

with open('/Users/0x1/Documents/htdocs/altbak.pub/bin/go/run', 'w') as f:
    f.write(content)
