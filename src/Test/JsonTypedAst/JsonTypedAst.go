package JsonTypedAst

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "os"
    "runtime"
    "strings"
    "time"
    "gopurs/output/gopurs_runtime"
)

var diagnosticSink gopurs_runtime.Value

func canonicalHash(text string) string {
    var value any
    if err := json.Unmarshal([]byte(text), &value); err != nil { panic(err) }
    // The JS driver uses the same HTML and line-separator escaping.
    bytes, err := json.Marshal(value); if err != nil { panic(err) }
    hash := sha256.Sum256(bytes)
    return hex.EncodeToString(hash[:])
}

func Drive(parse, decode, decodeText, encode, fingerprint gopurs_runtime.Value) func() any {
 return func() any {
    bytes,err:=os.ReadFile(os.Getenv("DIAG_CORPUS"));if err!=nil {panic(err)}
    var files []struct{Name string `json:"name"`; Contents string `json:"contents"`}
    if err=json.Unmarshal(bytes,&files);err!=nil {panic(err)}
    parsed:=make([]gopurs_runtime.Value,len(files))
    expectedJSON:=make([]string,len(files));expectedAST:=make([]string,len(files))
    for i,f:=range files {
        parsed[i]=gopurs_runtime.Apply(parse,gopurs_runtime.Str(f.Contents))
        expectedJSON[i]=canonicalHash(f.Contents)
        result:=gopurs_runtime.Apply(decode,parsed[i])
        raw:=gopurs_runtime.Apply(fingerprint,result).StrVal()
        expectedAST[i]=canonicalHash(raw)
    }
    phases:=map[string]any{}
    phaseOrder:=[]string{"parse","decode","combined"}
    if requested:=os.Getenv("DIAG_PHASES");requested!="" {phaseOrder=strings.Split(requested,",")}
    seen:=map[string]bool{}
    for _,phase:=range phaseOrder {
        if seen[phase] || (phase!="parse" && phase!="decode" && phase!="combined") {panic("invalid DIAG_PHASES")}
        seen[phase]=true
    }
    for _,phase:=range phaseOrder {
        samples:=[]map[string]float64{};best:=1e100
        for pass:=0;pass<7;pass++ {
            results:=make([]gopurs_runtime.Value,len(files))
            var before,after runtime.MemStats
            runtime.ReadMemStats(&before)
            start:=time.Now()
            for i,f:=range files {
                switch phase {
                case "parse":results[i]=gopurs_runtime.Apply(parse,gopurs_runtime.Str(f.Contents))
                case "decode":results[i]=gopurs_runtime.Apply(decode,parsed[i])
                case "combined":results[i]=gopurs_runtime.Apply(decodeText,gopurs_runtime.Str(f.Contents))
                }
                diagnosticSink=results[i]
            }
            elapsed:=float64(time.Since(start).Nanoseconds())/1000
            runtime.ReadMemStats(&after)
            for i,result:=range results {
                callback:=fingerprint;expected:=expectedAST[i]
                if phase=="parse" {callback=encode;expected=expectedJSON[i]}
                if canonicalHash(gopurs_runtime.Apply(callback,result).StrVal())!=expected {panic("Unstable "+phase+" output")}
            }
            if pass>=2 {
                if elapsed<best {best=elapsed}
                samples=append(samples,map[string]float64{"time_us":elapsed,"allocated_bytes":float64(after.TotalAlloc-before.TotalAlloc)})
            }
        }
        phases[phase]=map[string]any{"samples":samples,"time_us":best}
    }
    report:=map[string]any{"backend":"go","phase_order":phaseOrder,"modules":len(files),"fingerprints":expectedAST,"json_fingerprints":expectedJSON,"phases":phases,"go":runtime.Version(),"gomaxprocs":runtime.GOMAXPROCS(0)}
    result,err:=json.Marshal(report);if err!=nil {panic(err)};fmt.Println(string(result))
    return nil
 }
}
