import os
import glob
import shutil

phpurs_dir = '/Users/0x1/Documents/htdocs/phpurs'
output_dir = '/Users/0x1/Documents/htdocs/altbak.pub/run/bak/php/output'

for root, dirs, files in os.walk(phpurs_dir):
    if 'src' in root.split(os.sep) and not 'node_modules' in root.split(os.sep) and not 'output' in root.split(os.sep):
        for file in files:
            if file.endswith('.php') and file != 'Data.Show.php' and file != 'Bench.php':
                src_path = os.path.join(root, file)
                
                # compute module name from path relative to 'src'
                parts = root.split(os.sep)
                src_idx = parts.index('src')
                mod_parts = parts[src_idx+1:] + [file[:-4]]
                mod_name = '.'.join(mod_parts)
                
                dest_path = os.path.join(output_dir, f"{mod_name}.php")
                shutil.copyfile(src_path, dest_path)
                print(f"Copied {src_path} to {dest_path}")

shutil.copyfile('/Users/0x1/Documents/htdocs/altbak.pub/src/Bench.php', os.path.join(output_dir, 'Bench.php'))
shutil.copyfile('/Users/0x1/Documents/htdocs/phpurs/phpurs-prelude/src/Data/Show.php', os.path.join(output_dir, 'Data.Show.php'))
print("Done.")
