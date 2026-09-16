#!/usr/bin/env python3
"""Isolate forced inlining from bounds-check removal; algorithm unchanged."""
from pathlib import Path
root=Path(__file__).resolve().parent/'arena'
source=(root/'kernel-unchecked.rs').read_text()
for name in ['balance','rotate_right','rotate_left']:
 old='    fn '+name+'('
 assert source.count(old)==1
 source=source.replace(old,'    #[inline(always)]\n'+old)
(root/'kernel-unchecked-inline.rs').write_text(source)
(root/'check-unchecked-inline.rs').write_text((root/'check-unchecked.rs').read_text().replace(
 'include!("kernel-unchecked.rs")','include!("kernel-unchecked-inline.rs")'))
