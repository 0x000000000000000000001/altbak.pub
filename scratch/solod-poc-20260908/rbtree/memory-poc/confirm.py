#!/usr/bin/env python3
"""Fresh confirmation of the promising region and intermediate memory budgets."""
import hashlib
import json
import random
import statistics
import sys
from pathlib import Path
import bench as b

DEST=b.ROOT/'results-confirm.json'
VARIANTS=[
    b.variant('go-original-gogc800',0,800,pool=False),
    b.variant('go-original-gogc1600',0,1600,pool=False),
    b.variant('go-pool8-gogc200',8,200),
    b.variant('go-pool16-gogc100',16,100),
    b.variant('go-pool24-gogc100',24,100),
    b.variant('go-pool32-gogc100',32,100),
    b.variant('go-pool96-gogc800',96,800),
]
REPEATS=7
MEMORY_REPEATS=3
SIZES=[10000,100000]

def save(result):
    tmp=DEST.with_suffix('.json.tmp')
    tmp.write_text(json.dumps(result,indent=2)+'\n')
    tmp.replace(DEST)

def main():
    b.VARIANTS=tuple(VARIANTS)
    result={'status':'running','metadata':b.metadata(),
        'methodology':{'repetitions':REPEATS,'memory_repetitions':MEMORY_REPEATS,
            'random_seed':20260909,'target_batch_seconds':b.TARGET_SECONDS,'minimum_trees_per_batch':3,
            'calibration_variant':b.BASELINE['name'],
            'selection':'fresh samples after the 11-variant screening; includes new intermediate pools and GC1600 control',
            'variants':VARIANTS,'sizes':SIZES,'pool_cleanup_included':True,
            'rss':'three separate fresh parents per variant/size, macOS child ru_maxrss bytes'},
        'verification':[],'workloads':{}}
    result['metadata']['confirmation_script_sha256']=hashlib.sha256(Path(__file__).read_bytes()).hexdigest()
    rng=random.Random(20260909)
    save(result)
    for n in SIZES:
        reference=None
        for item in VARIANTS:
            obs=b.invoke(b.command(item,'verify',1,n),item)
            v=b.json_output(obs)
            fields=[v[k] for k in ('count','sum','depth','black_height')]
            if reference is None: reference=fields
            assert fields==reference and v['count']==n and v['sum']==n*(n+1)//2
            if n==100000:
                assert fields==[100000,5000050000,22,17]
                if item['pool']: assert v['allocations']==2483948
            if item['pool']: assert 0<=v['pool_bytes']<=item['cap_mib']*b.MIB
            obs.update({'n':n,'result':v})
            result['verification'].append(obs)
        b.N=n
        count,calibration=b.calibrate(reference[2])
        print(f'Confirm {n} nodes: {count} trees/batch',flush=True)
        w={'n':n,'count':count,'depth':reference[2],'calibration':calibration,
            'warmups':[b.timed_batch(item,count,reference[2]) for item in VARIANTS],
            'samples':{i['name']:[] for i in VARIANTS},'order':[],
            'memory_samples':{i['name']:[] for i in VARIANTS}}
        result['workloads'][str(n)]=w
        for repetition in range(REPEATS):
            order=list(VARIANTS); rng.shuffle(order)
            w['order'].append([i['name'] for i in order])
            for item in order:
                obs=b.timed_batch(item,count,reference[2]); obs['repetition']=repetition+1
                w['samples'][item['name']].append(obs)
            save(result)
        w['summary']=b.summary(w['samples'])
        for item in VARIANTS:
            print(item['name'],round(w['summary'][item['name']]['median_ns_per_tree']/1e6,3),'ms',flush=True)
        save(result)
        for repetition in range(MEMORY_REPEATS):
            order=list(VARIANTS); rng.shuffle(order)
            for item in order:
                argv=[sys.executable,str(b.ROOT.parent/'memory.py'),*b.command(item,'bench',count,n)]
                obs=b.invoke(argv,item); b.checksum(obs,count,n,reference[2])
                match=b.RSS.search(obs['stderr']); assert match,obs
                obs['rss_bytes']=int(match.group(1)); obs['repetition']=repetition+1
                w['memory_samples'][item['name']].append(obs)
            save(result)
        w['memory_summary']={}
        for item in VARIANTS:
            values=[o['rss_bytes'] for o in w['memory_samples'][item['name']]]
            w['memory_summary'][item['name']]={'median_rss_bytes':statistics.median(values),
                'min_rss_bytes':min(values),'max_rss_bytes':max(values)}
            print(item['name'],'RSS',round(statistics.median(values)/b.MIB,2),'MiB',flush=True)
        save(result)
    assert b.binaries()==result['metadata']['binaries']
    result['status']='complete'
    result['metadata']['binary_hashes_verified_after_run']=True
    save(result)
    print('Saved',DEST,flush=True)

if __name__=='__main__': main()
