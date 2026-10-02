#!/usr/bin/env python3
"""Exercise locked preparation with the real pinned Composer, entirely offline."""
import contextlib
import io
import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / 'bin/php'))
import driver


class ComposerPreparation(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory(prefix='altbak-composer-')
        self.addCleanup(self.temporary.cleanup)
        self.directory = Path(self.temporary.name).resolve()
        self.php = shutil.which(os.environ.get('PHP', 'php'))
        self.env = dict(os.environ, COMPOSER_DISABLE_NETWORK='1',
                        COMPOSER_HOME=str(self.directory / 'home'),
                        COMPOSER_CACHE_DIR=str(self.directory / 'cache'))
        (self.directory / 'output').mkdir()
        shutil.copyfile(ROOT / 'run/bak/php/composer.json', self.directory / 'composer.json')
        generated = {'name': 'phpurs/lib-deps', 'require': {'php': '>=8.1'}, 'require-dev': {}}
        driver.write_json(self.directory / 'output/composer.json', generated)
        # The root manifest is the real benchmark manifest. Replace its locked
        # dependency graph by one local package; no registry/download is needed.
        lock = json.loads((ROOT / 'run/bak/php/composer.lock').read_text())
        package = lock['packages'][0]
        package['require'] = generated['require']
        lock['packages'] = [package]
        driver.write_json(self.directory / 'composer.lock', lock)
        self.lock_bytes = (self.directory / 'composer.lock').read_bytes()
        self.value = self.directory / 'output/value.php'
        self.value.write_text('<?php return 7;\n')

    def prepare(self):
        with contextlib.redirect_stdout(io.StringIO()):
            return driver.prepare_composer(self.directory, self.php, self.env,
                                           self.directory / 'composer.log')

    def state(self):
        return {str(p.relative_to(self.directory)): (p.read_bytes(), p.stat().st_mtime_ns)
                for p in (self.directory / 'vendor').rglob('*') if p.is_file()}

    def execute(self):
        return subprocess.check_output([self.php, '-r',
            'require "vendor/autoload.php"; echo require "vendor/phpurs/lib-deps/value.php";'],
            cwd=self.directory, text=True)

    def test_fresh_then_reuse_preserves_vendor_and_observes_new_output(self):
        self.assertEqual(self.prepare()['status'], 'installed')
        before = self.state()
        stamp = (self.directory / '.composer-prepared.json').stat().st_mtime_ns
        self.assertEqual(self.execute(), '7')
        # Rebuilding PHP recreates output; the installed package must still
        # refer to that new directory without an install/autoload dump.
        generated = (self.directory / 'output/composer.json').read_bytes()
        shutil.rmtree(self.directory / 'output')
        (self.directory / 'output').mkdir()
        (self.directory / 'output/composer.json').write_bytes(generated)
        self.value.write_text('<?php return 11;\n')
        with patch.object(driver, 'logged', side_effect=AssertionError('unexpected Composer install')):
            self.assertEqual(self.prepare(), {'status': 'reused', 'commands': []})
        self.assertEqual(self.state(), before)
        self.assertEqual((self.directory / '.composer-prepared.json').stat().st_mtime_ns, stamp)
        self.assertEqual(self.execute(), '11')
        self.assertEqual((self.directory / 'composer.lock').read_bytes(), self.lock_bytes)

    def test_missing_edited_extra_files_and_wrong_links_are_repaired(self):
        self.prepare()
        original = driver.vendor_artifacts(self.directory)
        target = self.directory / 'vendor/autoload.php'
        for damage in ['missing', 'edited', 'extra', 'link']:
            with self.subTest(damage=damage):
                if damage == 'missing':
                    target.unlink()
                elif damage == 'edited':
                    target.write_text('<?php throw new Exception("edited");')
                elif damage == 'extra':
                    (self.directory / 'vendor/extra.php').write_text('extra')
                else:
                    link = self.directory / 'vendor/phpurs/lib-deps'
                    link.unlink()
                    link.symlink_to(self.directory / 'absent')
                self.assertEqual(self.prepare()['status'], 'installed')
                self.assertEqual(driver.vendor_artifacts(self.directory), original)
                self.assertEqual(self.execute(), '7')

    def test_lock_is_required_and_changed_requirements_are_rejected_before_install(self):
        self.prepare()
        lock = self.directory / 'composer.lock'
        lock.unlink()
        with patch.object(driver, 'logged', side_effect=AssertionError('unexpected install')):
            with self.assertRaisesRegex(ValueError, 'Missing composer.lock'):
                self.prepare()
            lock.write_bytes(self.lock_bytes)
            generated = self.directory / 'output/composer.json'
            data = json.loads(generated.read_text())
            data['require']['ext-json'] = '*'
            driver.write_json(generated, data)
            with self.assertRaisesRegex(ValueError, 'Generated Composer require differs'):
                self.prepare()
            del data['require']['ext-json']
            driver.write_json(generated, data)
            root = self.directory / 'composer.json'
            data = json.loads(root.read_text())
            data['require']['php'] = '>=8.1'
            driver.write_json(root, data)
            with self.assertRaisesRegex(ValueError, 'composer.lock is stale'):
                self.prepare()

    def test_reviewed_dependency_change_installs_from_new_lock(self):
        self.prepare()
        generated = self.directory / 'output/composer.json'
        data = json.loads(generated.read_text())
        data['require']['ext-json'] = '*'
        driver.write_json(generated, data)
        lock = self.directory / 'composer.lock'
        data = json.loads(lock.read_text())
        data['packages'][0]['require']['ext-json'] = '*'
        driver.write_json(lock, data)
        locked_bytes = lock.read_bytes()
        self.assertEqual(self.prepare()['status'], 'installed')
        self.assertEqual(self.prepare()['status'], 'reused')
        self.assertEqual(lock.read_bytes(), locked_bytes)

    def test_corrupt_state_and_configuration_change_invalidate_reuse(self):
        self.prepare()
        (self.directory / '.composer-prepared.json').write_text('{broken')
        self.assertEqual(self.prepare()['status'], 'installed')
        # Global Composer configuration affects installation/autoloading too.
        config = self.directory / 'home/config.json'
        driver.write_json(config, {'config': {'optimize-autoloader': True}})
        self.assertEqual(self.prepare()['status'], 'installed')
        self.assertEqual(self.prepare()['status'], 'reused')

    def test_failed_install_cannot_authorize_reuse(self):
        self.prepare()
        (self.directory / 'vendor/autoload.php').unlink()
        with patch.object(driver, 'logged', side_effect=RuntimeError('interrupted install')):
            with self.assertRaisesRegex(RuntimeError, 'interrupted install'):
                self.prepare()
        self.assertFalse((self.directory / '.composer-prepared.json').exists())
        self.assertEqual(self.prepare()['status'], 'installed')


if __name__ == '__main__':
    unittest.main()
