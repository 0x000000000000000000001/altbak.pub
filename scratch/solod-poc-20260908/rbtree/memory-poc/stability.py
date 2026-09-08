#!/usr/bin/env python3
"""Untimed lifetime and retention checks; no compiler or benchmark mutations."""
import json
import random
import statistics
import sys
import bench as b

DEST=b.ROOT/'results-stability.json'

def main():
    config=json.loads((b.ROOT/'results-confirm.json').read_text())
    assert config['status']=='complete','finish timing runs first'
    names=sys.argv[1:] or ['go-original-gogc800','go-pool8-gogc200','go-pool96-gogc800']
    variants={v['name']:v for v in config['methodology']['variants']}
    result={'status':'running','metadata':b.metadata(),'sequence_checks':[],
        'retention_checks':[],'retention_summary':{},'cold_start':{},
        'methodology':{'note':'separate diagnostics; forced GC only after a batch and in validation',
            'sizes_sequence':[100000,1000,0,100000], 'counts':[1,10,100],
            'rss':'fresh parent per diagnostic child; macOS bytes'}}
    def save(): DEST.write_text(json.dumps(result,indent=2)+'\n')
    # No arena reuse: one tree per process. Keep these external timings
    # separate, because process startup materially affects the 10k case.
    rng=random.Random(20260910)
    for n in [10000,100000]:
        samples={name:[] for name in names}
        for repetition in range(7):
            order=list(names);rng.shuffle(order)
            for name in order:
                item=variants[name]
                obs=b.invoke(b.command(item,'bench',1,n),item,timed=True)
                b.checksum(obs,1,n,config['workloads'][str(n)]['depth'])
                obs['repetition']=repetition+1
                samples[name].append(obs)
        summaries={name:{'median_elapsed_ns':statistics.median(o['elapsed_ns'] for o in seq),
            'min_elapsed_ns':min(o['elapsed_ns'] for o in seq),
            'max_elapsed_ns':max(o['elapsed_ns'] for o in seq)} for name,seq in samples.items()}
        result['cold_start'][str(n)]={'repetitions':7,'count':1,'samples':samples,'summary':summaries,
            'includes_process_startup':True,'arena_reused':False}
        for name,s in summaries.items(): print('Cold',n,name,round(s['median_elapsed_ns']/1e6,3),'ms',flush=True)
        save()
    for cap in [0,1,8,24,96]:
        item=b.variant(f'sequence-pool{cap}-gogc100',cap,100)
        obs=b.invoke(b.command(item,'sequence',1,100000),item)
        v=b.json_output(obs)
        assert len(v['trees'])==4
        for size,t in zip([100000,1000,0,100000],v['trees']):
            assert t['count']==size and t['sum']==size*(size+1)//2
            assert t['pool_bytes']<=cap*b.MIB
            if size==100000: assert t['depth']==22 and t['black_height']==17 and t['allocations']==2483948
        pool_bytes=v['trees'][-1]['pool_bytes']
        assert v['after_gc']['heap_alloc_bytes']<=pool_bytes+4*b.MIB
        obs['result']=v; result['sequence_checks'].append(obs)
        save()
    for name in names:
        item=variants[name]
        observations=[]
        for count in [1,10,100]:
            argv=[sys.executable,str(b.ROOT.parent/'memory.py'),*b.command(item,'stats',count,100000)]
            obs=b.invoke(argv,item)
            v=b.json_output(obs)
            assert v['checksum']==22*count
            if item['pool']: assert v['allocations']==2483948
            assert v['pool_bytes']<=item['cap_mib']*b.MIB
            assert v['after_gc']['heap_alloc_bytes']<=v['pool_bytes']+4*b.MIB
            match=b.RSS.search(obs['stderr']); assert match
            obs.update({'count':count,'result':v,'rss_bytes':int(match.group(1))})
            observations.append(obs); result['retention_checks'].append(obs)
            print(name,count,'trees; retained',round(v['after_gc']['heap_alloc_bytes']/b.MIB,2),
                'MiB; peakRSS',round(obs['rss_bytes']/b.MIB,2),'MiB',flush=True)
            save()
        heaps=[o['result']['after_gc']['heap_alloc_bytes'] for o in observations]
        assert max(heaps)-min(heaps)<4*b.MIB
        result['retention_summary'][name]={'retained_heap_min_bytes':min(heaps),
            'retained_heap_max_bytes':max(heaps),'drift_bytes':max(heaps)-min(heaps)}
    result['metadata']['binary_hashes_verified_after_run']=b.binaries()==result['metadata']['binaries']
    assert result['metadata']['binary_hashes_verified_after_run']
    result['status']='complete'; save()
    print('Saved',DEST,flush=True)

if __name__=='__main__': main()
